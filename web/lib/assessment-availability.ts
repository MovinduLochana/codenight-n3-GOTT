import "server-only";

import { eq } from "drizzle-orm";

import { db } from "@/db/drizzle";
import { assessmentAvailability } from "@/db/schema";

// Exercises are defined statically in content/final_assessment/manifest.json
// and have no DB row by default, so an exercise with no row here is
// available — the table only needs rows for exercises explicitly disabled.

// Fails open: if the availability check itself can't run (e.g. the
// assessment_availability table hasn't been migrated yet), exercises stay
// visible/runnable rather than taking down the whole Final Assessment page.
export async function isExerciseAvailable(
  exerciseId: string,
): Promise<boolean> {
  try {
    const [row] = await db
      .select({ available: assessmentAvailability.available })
      .from(assessmentAvailability)
      .where(eq(assessmentAvailability.exerciseId, exerciseId))
      .limit(1);
    return row?.available ?? true;
  } catch (error) {
    console.error("isExerciseAvailable query failed", error);
    return true;
  }
}

export async function getUnavailableExerciseIds(): Promise<Set<string>> {
  try {
    const rows = await db
      .select({ exerciseId: assessmentAvailability.exerciseId })
      .from(assessmentAvailability)
      .where(eq(assessmentAvailability.available, false));
    return new Set(rows.map((row) => row.exerciseId));
  } catch (error) {
    console.error("getUnavailableExerciseIds query failed", error);
    return new Set();
  }
}
