import { ChapterSidebar } from "@/components/shell/chapter-sidebar";
import { TopBar } from "@/components/shell/top-bar";
import { chapters } from "@/lib/content";

export default function LearnLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <div className="flex h-dvh flex-col overflow-hidden">
      <TopBar chapters={chapters} />
      <div className="flex min-h-0 flex-1">
        <ChapterSidebar chapters={chapters} />
        <div className="flex min-w-0 flex-1 flex-col">{children}</div>
      </div>
    </div>
  );
}
