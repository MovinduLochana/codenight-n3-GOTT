import { revalidateTag } from "next/cache";
import { type NextRequest, NextResponse } from "next/server";

import { db } from "@/db/drizzle";
import { quizProgress } from "@/db/schema";
import { getCategory, getChapterQuiz } from "@/lib/content";
import { getSession } from "@/lib/session";

export async function POST(
  request: NextRequest,
  { params }: { params: Promise<{ categoryId: string }> },
) {
  const session = await getSession();
  if (!session) {
    return NextResponse.json({ error: "Unauthorized" }, { status: 401 });
  }

  const { categoryId } = await params;
  const category = getCategory(categoryId);
  if (!category) {
    return NextResponse.json({ error: "Chapter not found" }, { status: 404 });
  }

  const quiz = await getChapterQuiz(categoryId);
  if (!quiz) {
    return NextResponse.json({ error: "Quiz not found" }, { status: 404 });
  }

  const body = (await request.json().catch(() => null)) as {
    answers?: unknown;
  } | null;
  const answers = body?.answers;
  if (!answers || typeof answers !== "object" || Array.isArray(answers)) {
    return NextResponse.json({ error: "Invalid submission" }, { status: 400 });
  }
  const answerMap = answers as Record<string, unknown>;

  const total = quiz.questions.length;
  const score = quiz.questions.filter(
    (question) => answerMap[question.id] === question.correct,
  ).length;
  const passed = score === total;

  // Atomic insert-if-absent: a plain check-then-insert would race if the
  // same user double-submits (double click, retry), tripping the
  // quiz_progress_user_category_unique constraint. onConflictDoNothing
  // makes the check and the write a single statement.
  const [inserted] = await db
    .insert(quizProgress)
    .values({
      id: crypto.randomUUID(),
      userId: session.userId,
      categoryId,
      passed,
      score,
      total,
    })
    .onConflictDoNothing({
      target: [quizProgress.userId, quizProgress.categoryId],
    })
    .returning({ id: quizProgress.id });

  if (!inserted) {
    return NextResponse.json(
      { error: "You've already completed this quiz." },
      { status: 409 },
    );
  }

  revalidateTag("leaderboard", "max");

  return NextResponse.json({ passed, score, total });
}
