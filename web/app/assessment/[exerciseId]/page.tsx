import { and, eq } from "drizzle-orm";
import { LockIcon } from "lucide-react";
import { cacheLife } from "next/cache";
import Link from "next/link";
import { notFound } from "next/navigation";
import { Suspense, ViewTransition } from "react";

import { ExerciseWorkbench } from "@/components/assessment/exercise-workbench";
import {
  ExerciseHeaderSkeleton,
  ExerciseWorkbenchSkeleton,
} from "@/components/common/page-skeletons";
import { db } from "@/db/drizzle";
import { assessmentProgress } from "@/db/schema";
import {
  type AssessmentExercise,
  assessmentExercises,
  getAssessmentExercise,
} from "@/lib/assessment";
import { isExerciseAvailable } from "@/lib/assessment-availability";
import { readRepoFile } from "@/lib/content";
import { renderMarkdown } from "@/lib/markdown";
import { getSession } from "@/lib/session";

export function generateStaticParams() {
  return assessmentExercises.map((exercise) => ({ exerciseId: exercise.id }));
}

export default function AssessmentExercisePage({
  params,
}: {
  params: Promise<{ exerciseId: string }>;
}) {
  return (
    <Suspense
      fallback={
        <ViewTransition exit="fade-out">
          <main className="flex min-h-0 flex-1 flex-col overflow-y-auto lg:overflow-hidden">
            <div className="flex min-h-0 flex-1 flex-col px-6 py-6">
              <ExerciseHeaderSkeleton />
              <ExerciseWorkbenchSkeleton />
            </div>
          </main>
        </ViewTransition>
      }
    >
      <ViewTransition enter="fade-in" default="none">
        <AssessmentExerciseContent params={params} />
      </ViewTransition>
    </Suspense>
  );
}

async function AssessmentExerciseContent({
  params,
}: {
  params: Promise<{ exerciseId: string }>;
}) {
  const { exerciseId } = await params;

  const exercise = getAssessmentExercise(exerciseId);
  if (!exercise) notFound();
  if (!(await isExerciseAvailable(exerciseId))) {
    return <LockedExercise exercise={exercise} />;
  }

  const taskMarkdown = await readRepoFile(exercise.taskPath);
  const starterCode = await readRepoFile(exercise.starterPath);
  if (taskMarkdown === null || starterCode === null) notFound();

  const taskHtml = await renderMarkdown(taskMarkdown);

  return (
    <main className="flex min-h-0 flex-1 flex-col overflow-y-auto lg:overflow-hidden">
      <div className="flex min-h-0 flex-1 flex-col px-6 py-6">
        <ViewTransition
          key={exerciseId}
          name="exercise"
          share="auto"
          default="none"
        >
          <p className="text-[0.625rem] font-semibold tracking-widest text-primary uppercase">
            Final Assessment · {exercise.level}
          </p>
          <h1 className="font-heading mt-1 mb-6 text-2xl font-semibold">
            {String(exercise.number).padStart(2, "0")} · {exercise.title}
          </h1>

          <Suspense
            fallback={
              <ViewTransition exit="fade-out">
                <ExerciseWorkbenchSkeleton />
              </ViewTransition>
            }
          >
            <ViewTransition enter="fade-in" default="none">
              <Editor
                exerciseId={exercise.id}
                taskHtml={taskHtml}
                starterCode={starterCode}
              />
            </ViewTransition>
          </Suspense>
        </ViewTransition>
      </div>
    </main>
  );
}

async function getSavedProgress(userId: string, exerciseId: string) {
  "use cache";
  cacheLife("seconds");

  const [saved] = await db
    .select({
      code: assessmentProgress.code,
      passed: assessmentProgress.passed,
      output: assessmentProgress.output,
    })
    .from(assessmentProgress)
    .where(
      and(
        eq(assessmentProgress.userId, userId),
        eq(assessmentProgress.exerciseId, exerciseId),
      ),
    )
    .limit(1);

  return saved ?? null;
}

async function Editor({
  exerciseId,
  taskHtml,
  starterCode,
}: {
  exerciseId: string;
  taskHtml: string;
  starterCode: string;
}) {
  const session = await getSession();
  const saved = session
    ? await getSavedProgress(session.userId, exerciseId)
    : null;

  return (
    <ExerciseWorkbench
      exerciseId={exerciseId}
      taskHtml={taskHtml}
      starterCode={starterCode}
      savedCode={saved?.code ?? null}
      initialResult={
        saved ? { passed: saved.passed, output: saved.output } : null
      }
    />
  );
}

function LockedExercise({ exercise }: { exercise: AssessmentExercise }) {
  return (
    <main className="flex min-h-0 flex-1 items-center justify-center overflow-y-auto px-6 py-6">
      <div className="flex max-w-sm flex-col items-center gap-4 text-center">
        <div className="flex size-14 items-center justify-center rounded-full border border-dashed border-border bg-card">
          <LockIcon className="size-6 text-muted-foreground" />
        </div>
        <div className="space-y-1.5">
          <p className="text-[0.625rem] font-semibold tracking-widest text-muted-foreground uppercase">
            Final Assessment · {exercise.level}
          </p>
          <h1 className="font-heading text-xl font-semibold">
            {String(exercise.number).padStart(2, "0")} · {exercise.title}
          </h1>
        </div>
        <p className="text-sm text-muted-foreground">
          This exercise isn't open yet. Check back once it's unlocked, or head
          back to see what's available now.
        </p>
        <Link
          href="/assessment"
          className="mt-2 inline-flex items-center justify-center border border-border bg-card px-4 py-2 text-sm font-medium transition-colors hover:border-primary/40"
        >
          Back to Final Assessment
        </Link>
      </div>
    </main>
  );
}
