import { ChapterSidebar } from "@/components/shell/chapter-sidebar";
import { TopBar } from "@/components/shell/top-bar";
import { assessmentExercises } from "@/lib/assessment";
import { categories } from "@/lib/content";

export function AppShell({ children }: { children: React.ReactNode }) {
  return (
    <div className="flex h-dvh flex-col overflow-hidden">
      <TopBar categories={categories} />
      <div className="flex min-h-0 flex-1">
        <ChapterSidebar
          categories={categories}
          assessmentExercises={assessmentExercises}
        />
        <div className="flex min-w-0 flex-1 flex-col">{children}</div>
      </div>
    </div>
  );
}
