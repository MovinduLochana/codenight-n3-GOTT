import { readFileSync } from "node:fs";
import path from "node:path";

const contentRoot = path.join(process.cwd(), "..", "content");

const manifest: Record<string, string[]> = JSON.parse(
  readFileSync(path.join(contentRoot, "manifest.json"), "utf-8"),
);

const chapterNames: { slug: string; title: string }[] = [
  { slug: "fundamentals", title: "Go Fundamentals" },
  { slug: "control_flow", title: "Control Flow" },
  { slug: "collections", title: "Collections" },
  { slug: "function", title: "Functions" },
  { slug: "struct_and_interfaces", title: "Structs & Interfaces" },
  { slug: "errors_and_concurrency", title: "Errors & Concurrency" },
];

export type Lesson = {
  slug: string;
  title: string;
};

export type Chapter = {
  slug: string;
  number: number;
  title: string;
  lessons: Lesson[];
};

export const chapters: Chapter[] = chapterNames.map((chapter, index) => ({
  ...chapter,
  number: index + 1,
  lessons: (manifest[chapter.slug] ?? []).map((slug) => ({
    slug,
    title: lessonTitle(slug),
  })),
}));

export function getChapter(slug: string): Chapter | undefined {
  return chapters.find((chapter) => chapter.slug === slug);
}

export function lessonTitle(slug: string): string {
  return slug
    .split("_")
    .map((word) => word.charAt(0).toUpperCase() + word.slice(1))
    .join(" ");
}

export function readLesson(
  chapterSlug: string,
  lessonSlug: string,
): string | null {
  try {
    return readFileSync(
      path.join(contentRoot, chapterSlug, lessonSlug, "content.md"),
      "utf-8",
    );
  } catch {
    return null;
  }
}
