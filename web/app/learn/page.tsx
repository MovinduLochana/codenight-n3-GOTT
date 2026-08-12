import Link from "next/link";

import { chapters } from "@/lib/content";

export default function CurriculumPage() {
  return (
    <main className="min-h-0 flex-1 overflow-y-auto">
      <div className="mx-auto max-w-3xl px-6 py-10">
        <h1 className="font-heading text-2xl font-semibold">Curriculum</h1>

        <ul className="mt-8 flex flex-col gap-2">
          {chapters.map((chapter) => (
            <li key={chapter.slug}>
              <Link
                href={`/learn/${chapter.slug}/${chapter.lessons[0]?.slug}`}
                className="flex items-baseline gap-3 border border-border bg-card p-4 transition-colors hover:border-primary/40"
              >
                <span className="font-mono text-xs text-primary">
                  {String(chapter.number).padStart(2, "0")}
                </span>
                <span className="font-medium">{chapter.title}</span>
                <span className="ms-auto text-xs text-muted-foreground">
                  {chapter.lessons.length} lessons
                </span>
              </Link>
            </li>
          ))}
        </ul>
      </div>
    </main>
  );
}
