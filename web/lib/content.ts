import { readFileSync } from "node:fs";
import path from "node:path";

const repoRoot = path.join(process.cwd(), "..");

export type Level = "beginner" | "intermediate" | "advanced";

export type Exercise = {
  id: string;
  title: string;
  level: Level;
  filePath: string;
  testPath: string;
  docPath: string;
};

export type Topic = {
  id: string;
  title: string;
  contentPath: string;
  exercises: Exercise[];
};

export type Category = {
  id: string;
  number: number;
  title: string;
  topics: Topic[];
};

type Manifest = {
  title: string;
  categories: {
    id: string;
    title: string;
    topics: {
      id: string;
      title: string;
      content_path: string;
      exercises: {
        id: string;
        title: string;
        level: Level;
        file_path: string;
        test_path: string;
        doc_path: string;
      }[];
    }[];
  }[];
};

const manifest: Manifest = JSON.parse(
  readFileSync(path.join(repoRoot, "content", "manifest.json"), "utf-8"),
);

export const courseTitle = manifest.title;

export const categories: Category[] = manifest.categories.map(
  (category, index) => ({
    id: category.id,
    number: index + 1,
    title: category.title,
    topics: category.topics.map((topic) => ({
      id: topic.id,
      title: topic.title,
      contentPath: topic.content_path,
      exercises: topic.exercises.map((exercise) => ({
        id: exercise.id,
        title: exercise.title,
        level: exercise.level,
        filePath: exercise.file_path,
        testPath: exercise.test_path,
        docPath: exercise.doc_path,
      })),
    })),
  }),
);

export function getCategory(categoryId: string): Category | undefined {
  return categories.find((category) => category.id === categoryId);
}

export function getTopic(
  categoryId: string,
  topicId: string,
): Topic | undefined {
  return getCategory(categoryId)?.topics.find((topic) => topic.id === topicId);
}

export function readRepoFile(relativePath: string): string | null {
  try {
    return readFileSync(path.join(repoRoot, relativePath), "utf-8");
  } catch {
    return null;
  }
}
