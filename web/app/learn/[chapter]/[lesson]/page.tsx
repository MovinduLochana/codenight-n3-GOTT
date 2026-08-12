import { notFound } from "next/navigation";

import { chapters, getChapter, lessonTitle, readLesson } from "@/lib/content";
import { renderMarkdown } from "@/lib/markdown";

export function generateStaticParams() {
  return chapters.flatMap((chapter) =>
    chapter.lessons.map((lesson) => ({
      chapter: chapter.slug,
      lesson: lesson.slug,
    })),
  );
}

export default async function LessonPage({
  params,
}: {
  params: Promise<{ chapter: string; lesson: string }>;
}) {
  const { chapter: chapterSlug, lesson: lessonSlug } = await params;

  const chapter = getChapter(chapterSlug);
  const markdown = readLesson(chapterSlug, lessonSlug);
  if (!chapter || markdown === null) notFound();

  const html = await renderMarkdown(markdown);

  return (
    <main className="min-h-0 flex-1 overflow-y-auto">
      <div className="mx-auto max-w-3xl px-6 py-10">
        <p className="text-[0.625rem] font-semibold tracking-widest text-primary uppercase">
          Chapter {chapter.number} · {chapter.title}
        </p>
        <h1 className="font-heading mt-2 mb-8 text-3xl font-semibold">
          {lessonTitle(lessonSlug)}
        </h1>
        <article
          className="prose prose-invert max-w-none prose-pre:border prose-pre:border-border prose-pre:bg-card"
          // biome-ignore lint/security/noDangerouslySetInnerHtml: our own markdown
          dangerouslySetInnerHTML={{ __html: html }}
        />
      </div>
    </main>
  );
}
