import { TrophyIcon } from "lucide-react";
import { Suspense, ViewTransition } from "react";

import { SuspenseLoader } from "@/components/common/suspense-loader";
import { LiveLeaderboard } from "@/components/leaderboard/live-board";
import { getLeaderboard } from "@/lib/leaderboard";
import { getSession } from "@/lib/session";

export default function LeaderboardPage() {
  return (
    <Suspense fallback={<ViewTransition exit="fade-out"><SuspenseLoader /></ViewTransition>}>
      <ViewTransition enter="fade-in" default="none">
        <LeaderboardPageContent />
      </ViewTransition>
    </Suspense>
  );
}

async function LeaderboardPageContent() {
  const [entries, session] = await Promise.all([
    getLeaderboard(),
    getSession(),
  ]);

  return (
    <main className="min-h-0 flex-1 overflow-y-auto">
      <div className="mx-auto max-w-3xl px-6 py-10">
        <div className="flex items-center gap-2">
          <TrophyIcon className="size-5 text-primary" />
          <h1 className="font-heading text-2xl font-semibold">Leaderboard</h1>
        </div>
        <p className="mt-2 text-sm text-muted-foreground">
          Ranked by chapter quiz performance and Final Assessment exercises
          passed. Updates automatically.
        </p>

        <LiveLeaderboard
          initialEntries={entries}
          currentUserId={session?.userId ?? null}
        />
      </div>
    </main>
  );
}
