import "server-only";

import { eq, sql } from "drizzle-orm";

import { db } from "@/db/drizzle";
import { assessmentProgress, quizProgress, sessions } from "@/db/schema";

const POINTS_PER_QUIZ_QUESTION = 10;
const POINTS_PER_PASSED_EXERCISE = 50;

export type LeaderboardEntry = {
  userId: string;
  displayName: string;
  score: number;
};

function fallbackName(userId: string): string {
  return `Learner #${userId.slice(0, 6)}`;
}

export async function getLeaderboard(limit = 100): Promise<LeaderboardEntry[]> {
  const [quizTotals, exerciseTotals, allSessions] = await Promise.all([
    db
      .select({
        userId: quizProgress.userId,
        totalScore: sql<number>`sum(${quizProgress.score})`,
      })
      .from(quizProgress)
      .groupBy(quizProgress.userId),
    db
      .select({
        userId: assessmentProgress.userId,
        passedCount: sql<number>`count(*)`,
      })
      .from(assessmentProgress)
      .where(eq(assessmentProgress.passed, true))
      .groupBy(assessmentProgress.userId),
    db
      .select({ userId: sessions.userId, displayName: sessions.displayName })
      .from(sessions),
  ]);

  const nameByUserId = new Map(
    allSessions.map((row) => [row.userId, row.displayName]),
  );
  const scoreByUserId = new Map<string, number>();

  for (const row of quizTotals) {
    const points = Number(row.totalScore ?? 0) * POINTS_PER_QUIZ_QUESTION;
    scoreByUserId.set(
      row.userId,
      (scoreByUserId.get(row.userId) ?? 0) + points,
    );
  }
  for (const row of exerciseTotals) {
    const points = Number(row.passedCount ?? 0) * POINTS_PER_PASSED_EXERCISE;
    scoreByUserId.set(
      row.userId,
      (scoreByUserId.get(row.userId) ?? 0) + points,
    );
  }

  return Array.from(scoreByUserId.entries())
    .filter(([, score]) => score > 0)
    .map(([userId, score]) => ({
      userId,
      displayName: nameByUserId.get(userId) || fallbackName(userId),
      score,
    }))
    .sort((a, b) => b.score - a.score)
    .slice(0, limit);
}
