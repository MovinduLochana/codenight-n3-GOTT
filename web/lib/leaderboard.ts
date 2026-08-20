import "server-only";

import { eq, sql } from "drizzle-orm";
import { cacheLife, cacheTag } from "next/cache";
import { db } from "@/db/drizzle";
import { assessmentProgress, quizProgress, sessions } from "@/db/schema";
import { assessmentExercises } from "@/lib/assessment";

export const POINTS_PER_QUIZ_QUESTION = 10;
export const POINTS_PER_PASSED_EXERCISE: Record<string, number> = {
  beginner: 50,
  intermediate: 100,
  advanced: 150,
};

export type LeaderboardEntry = {
  userId: string;
  displayName: string;
  score: number;
};

function fallbackName(userId: string): string {
  return `Learner #${userId.slice(0, 6)}`;
}

export async function getLeaderboard(limit = 100): Promise<LeaderboardEntry[]> {
  "use cache";
  cacheTag("leaderboard");
  cacheLife("minutes");

  const passedAssessments = await db
    .select({
      userId: assessmentProgress.userId,
      exerciseId: assessmentProgress.exerciseId,
    })
    .from(assessmentProgress)
    .where(eq(assessmentProgress.passed, true));

  const passedScoreByUser = new Map<string, number>();
  for (const row of passedAssessments) {
    const exercise = assessmentExercises.find((e) => e.id === row.exerciseId);
    if (exercise) {
      const points = POINTS_PER_PASSED_EXERCISE[exercise.level] || 0;
      passedScoreByUser.set(
        row.userId,
        (passedScoreByUser.get(row.userId) || 0) + points,
      );
    }
  }

  const quizScores = await db
    .select({
      userId: quizProgress.userId,
      totalScore: sql<number>`coalesce(sum(${quizProgress.score}), 0)`,
    })
    .from(quizProgress)
    .groupBy(quizProgress.userId);

  const quizScoreByUser = new Map<string, number>();
  for (const row of quizScores) {
    quizScoreByUser.set(row.userId, Number(row.totalScore));
  }

  const rows = await db
    .select({
      userId: sessions.userId,
      displayName: sessions.displayName,
    })
    .from(sessions);

  return rows
    .map((row) => {
      const quizScore = quizScoreByUser.get(row.userId) || 0;
      const assessmentScore = passedScoreByUser.get(row.userId) || 0;
      const score =
        quizScore * POINTS_PER_QUIZ_QUESTION + assessmentScore;
      return {
        userId: row.userId,
        displayName: row.displayName || fallbackName(row.userId),
        score,
      };
    })
    .filter((entry) => entry.score > 0)
    .sort((a, b) => b.score - a.score)
    .slice(0, limit);
}
