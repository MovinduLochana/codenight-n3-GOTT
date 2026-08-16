import { readdirSync, readFileSync } from "node:fs";
import path from "node:path";

import type { Level } from "@/lib/content";

const repoRoot = path.join(process.cwd(), ".data");

export type AssessmentExercise = {
  id: string;
  number: number;
  title: string;
  level: Level;
  dirPath: string;
  taskPath: string;
  starterPath: string;
  testPath: string;
};

type AssessmentManifest = {
  exercises: {
    id: string;
    number: number;
    title: string;
    level: Level;
    dir: string;
    task_path: string;
    starter_path: string;
    test_path: string;
  }[];
};

const manifest: AssessmentManifest = JSON.parse(
  readFileSync(
    path.join(repoRoot, "content", "final_assessment", "manifest.json"),
    "utf-8",
  ),
);

export const assessmentExercises: AssessmentExercise[] = manifest.exercises.map(
  (exercise) => ({
    id: exercise.id,
    number: exercise.number,
    title: exercise.title,
    level: exercise.level,
    dirPath: exercise.dir,
    taskPath: exercise.task_path,
    starterPath: exercise.starter_path,
    testPath: exercise.test_path,
  }),
);

export function getAssessmentExercise(
  id: string,
): AssessmentExercise | undefined {
  return assessmentExercises.find((exercise) => exercise.id === id);
}

const NON_FIXTURE_FILES = new Set(["main.go", "main_test.go", "task.md"]);

export function getFixtureFiles(
  exercise: AssessmentExercise,
): { name: string; content: string }[] {
  const dir = path.join(repoRoot, exercise.dirPath);
  return readdirSync(dir, { withFileTypes: true })
    .filter((entry) => entry.isFile() && !NON_FIXTURE_FILES.has(entry.name))
    .map((entry) => ({
      name: entry.name,
      content: readFileSync(path.join(dir, entry.name), "utf-8"),
    }));
}
