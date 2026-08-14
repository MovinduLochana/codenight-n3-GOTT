import { and, eq } from "drizzle-orm";
import { ArrowLeftIcon, ArrowRightIcon } from "lucide-react";
import Link from "next/link";
import { notFound } from "next/navigation";
import { Suspense } from "react";
import { SuspenseLoader } from "@/components/common/suspense-loader";
import { ChapterQuiz } from "@/components/lesson/chapter-quiz";
import { buttonVariants } from "@/components/ui/button";
import { db } from "@/db/drizzle";
import { quizProgress } from "@/db/schema";
import {
  categories,
  getCategory,
  getChapterQuiz,
  getNextChapterFirstTopic,
  toPublicQuiz,
} from "@/lib/content";
import { getSession } from "@/lib/session";

export function generateStaticParams() {
  return categories.map((category) => ({ chapter: category.id }));
}

export default function ChapterQuizPage({
  params,
}: {
  params: Promise<{ chapter: string }>;
}) {
  return (
    <Suspense fallback={<SuspenseLoader />}>
      <ChapterQuizPageContent params={params} />
    </Suspense>
  );
}

async function ChapterQuizPageContent({
  params,
}: {
  params: Promise<{ chapter: string }>;
}) {
  const { chapter: categoryId } = await params;

  const category = getCategory(categoryId);
  if (!category) notFound();

  const quiz = getChapterQuiz(categoryId);
  if (!quiz) notFound();

  const firstTopic = category.topics[0];
  const nextChapterTopic = getNextChapterFirstTopic(categoryId);

  const session = await getSession();
  const [completed] = session
    ? await db
        .select({
          passed: quizProgress.passed,
          score: quizProgress.score,
          total: quizProgress.total,
        })
        .from(quizProgress)
        .where(
          and(
            eq(quizProgress.userId, session.userId),
            eq(quizProgress.categoryId, categoryId),
          ),
        )
        .limit(1)
    : [];

  return (
    <main className="min-h-0 flex-1 overflow-y-auto">
      <div className="mx-auto max-w-3xl px-6 py-10">
        {firstTopic ? (
          <Link
            href={`/learn/${category.id}/${firstTopic.id}`}
            className="mb-8 flex items-center gap-1.5 text-xs font-medium text-muted-foreground transition-colors hover:text-foreground"
          >
            <ArrowLeftIcon className="size-3.5" />
            Back to lessons
          </Link>
        ) : null}

        <p className="text-[0.625rem] font-semibold tracking-widest text-primary uppercase">
          Chapter {category.number} · {category.title}
        </p>
        <h1 className="font-heading mt-2 mb-2 text-3xl font-semibold">
          Chapter Quiz
        </h1>
        <p className="mb-8 text-sm text-muted-foreground">
          {quiz.questions.length} questions covering every lesson in this
          chapter.
        </p>

        <ChapterQuiz
          categoryId={categoryId}
          quiz={toPublicQuiz(quiz)}
          completedResult={completed ?? null}
        />

        <div className="mt-8 flex items-center justify-between gap-3">
          {firstTopic ? (
            <Link
              href={`/learn/${category.id}/${firstTopic.id}`}
              className={buttonVariants({ variant: "outline" })}
            >
              Back to lessons
            </Link>
          ) : (
            <span />
          )}

          {nextChapterTopic ? (
            <Link
              href={`/learn/${nextChapterTopic.categoryId}/${nextChapterTopic.topicId}`}
              className={buttonVariants({ variant: "default" })}
            >
              Next chapter
              <ArrowRightIcon />
            </Link>
          ) : null}
        </div>
      </div>
    </main>
  );
}
