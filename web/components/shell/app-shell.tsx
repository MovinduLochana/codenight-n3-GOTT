import { ChapterSidebar } from "@/components/shell/chapter-sidebar";
import { TopBar } from "@/components/shell/top-bar";
import { assessmentExercises } from "@/lib/assessment";
import { categories } from "@/lib/content";

export function AppShell({ children }: { children: React.ReactNode }) {
  return (
    <div className="flex h-dvh flex-col overflow-hidden">
      <div style={{ viewTransitionName: "site-header" }}>
        <TopBar categories={categories} />
      </div>
      <div className="flex min-h-0 flex-1">
        <div style={{ viewTransitionName: "site-sidebar" }}>
          <ChapterSidebar
            categories={categories}
            assessmentExercises={assessmentExercises}
          />
        </div>
        <div className="flex min-w-0 flex-1 flex-col">{children}</div>
      </div>
    </div>
  );
}
