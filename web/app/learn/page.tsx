import Link from "next/link";
import { ViewTransition } from "react";

import { categories, courseTitle } from "@/lib/content";

export default function CurriculumPage() {
  return (
    <main className="min-h-0 flex-1 overflow-y-auto">
      <div className="mx-auto max-w-3xl px-6 py-10">
        <ViewTransition name="curriculumTitle" share="auto" default="none">
          <h1 className="font-heading text-2xl font-semibold">{courseTitle}</h1>
        </ViewTransition>

        <ul className="mt-8 flex flex-col gap-2">
          {categories.map((category) => (
            <li key={category.id}>
              <Link
                href={`/learn/${category.id}/${category.topics[0]?.id}`}
                className="flex items-baseline gap-3 border border-border bg-card p-4 transition-colors hover:border-primary/40"
              >
                <span className="font-mono text-xs text-primary">
                  {String(category.number).padStart(2, "0")}
                </span>
                <span className="font-medium">{category.title}</span>
                <span className="ms-auto text-xs text-muted-foreground">
                  {category.topics.length} topics
                </span>
              </Link>
            </li>
          ))}
        </ul>
      </div>
    </main>
  );
}
