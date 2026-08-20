"use client";

import { ChevronLeftIcon, ChevronRightIcon, CrownIcon } from "lucide-react";
import { useEffect, useState } from "react";

import type { LeaderboardEntry } from "@/lib/leaderboard";
import { cn } from "@/lib/utils";

const POLL_INTERVAL_MS = 30_000; // 30 seconds
const PAGE_SIZE = 10;

const PODIUM_HEIGHT: Record<1 | 2 | 3, string> = {
  1: "min-h-[232px] pt-10",
  2: "min-h-[192px] pt-7",
  3: "min-h-[160px] pt-6",
};

export function LiveLeaderboard({
  initialEntries,
  currentUserId,
}: {
  initialEntries: LeaderboardEntry[];
  currentUserId: string | null;
}) {
  const [entries, setEntries] = useState(initialEntries);
  const [page, setPage] = useState(1);

  useEffect(() => {
    let cancelled = false;
    let intervalId: ReturnType<typeof setInterval> | null = null;

    async function poll() {
      try {
        const response = await fetch("/api/leaderboard", { cache: "no-store" });
        if (!response.ok || cancelled) return;
        const body = (await response.json()) as { entries: LeaderboardEntry[] };
        if (!cancelled) setEntries(body.entries);
      } catch {}
    }

    function start() {
      if (!intervalId) intervalId = setInterval(poll, POLL_INTERVAL_MS);
    }

    function stop() {
      if (intervalId) {
        clearInterval(intervalId);
        intervalId = null;
      }
    }

    function handleVisibility() {
      if (document.hidden) {
        stop();
      } else {
        poll();
        start();
      }
    }

    start();
    document.addEventListener("visibilitychange", handleVisibility);

    return () => {
      cancelled = true;
      stop();
      document.removeEventListener("visibilitychange", handleVisibility);
    };
  }, []);

  if (entries.length === 0) {
    return (
      <div className="mt-8 border border-border bg-card p-8 text-center text-sm text-muted-foreground">
        No one's on the board yet. Complete a chapter quiz or an exercise to
        appear here.
      </div>
    );
  }

  const [first, second, third, ...rest] = entries;
  const podiumSlots: { entry: LeaderboardEntry | null; rank: 1 | 2 | 3 }[] = [
    { entry: second ?? null, rank: 2 },
    { entry: first ?? null, rank: 1 },
    { entry: third ?? null, rank: 3 },
  ];

  const totalPages = Math.max(1, Math.ceil(rest.length / PAGE_SIZE));
  const currentPage = Math.min(page, totalPages);
  const pageStart = (currentPage - 1) * PAGE_SIZE;
  const pageEntries = rest.slice(pageStart, pageStart + PAGE_SIZE);

  return (
    <>
      <div className="mt-10 grid grid-cols-3 items-end gap-3 sm:gap-4">
        {podiumSlots.map((slot) => (
          <PodiumCard
            key={slot.rank}
            entry={slot.entry}
            rank={slot.rank}
            isCurrentUser={slot.entry?.userId === currentUserId}
          />
        ))}
      </div>

      {rest.length > 0 ? (
        <>
          <div className="mt-8 divide-y divide-border border border-border bg-card">
            {pageEntries.map((entry, index) => (
              <RankRow
                key={entry.userId}
                rank={pageStart + index + 4}
                entry={entry}
                isCurrentUser={entry.userId === currentUserId}
              />
            ))}
          </div>

          {totalPages > 1 ? (
            <Pagination
              page={currentPage}
              totalPages={totalPages}
              onPageChange={setPage}
            />
          ) : null}
        </>
      ) : null}
    </>
  );
}

function PodiumCard({
  entry,
  rank,
  isCurrentUser,
}: {
  entry: LeaderboardEntry | null;
  rank: 1 | 2 | 3;
  isCurrentUser: boolean;
}) {
  return (
    <div
      className={cn(
        "relative flex flex-col items-center overflow-hidden border bg-card px-3 pb-5 text-center",
        PODIUM_HEIGHT[rank],
        entry
          ? isCurrentUser
            ? "border-primary bg-primary/5"
            : "border-primary/50"
          : "border-dashed border-border opacity-50",
      )}
    >
      {rank === 1 && entry ? (
        <>
          <div className="pointer-events-none absolute inset-0 bg-gradient-to-b from-primary/10 to-transparent" />
          <CrownIcon
            className="relative mb-1.5 size-4 text-primary"
            fill="currentColor"
          />
        </>
      ) : null}

      <div
        className={cn(
          "relative flex shrink-0 items-center justify-center border font-heading font-semibold",
          rank === 1 ? "size-10 text-lg" : "size-8 text-sm",
          entry
            ? "border-primary text-primary"
            : "border-border text-muted-foreground",
        )}
      >
        {rank}
      </div>
      <p className="relative mt-3 w-full truncate text-sm font-medium">
        {entry ? entry.displayName : "—"}
      </p>
      <p className="relative mt-1 font-heading text-xl font-semibold text-primary">
        {entry ? entry.score : "—"}
      </p>
      <p className="relative text-[0.625rem] tracking-widest text-muted-foreground uppercase">
        points
      </p>
    </div>
  );
}

function RankRow({
  rank,
  entry,
  isCurrentUser,
}: {
  rank: number;
  entry: LeaderboardEntry;
  isCurrentUser: boolean;
}) {
  return (
    <div
      className={cn(
        "flex items-center gap-4 px-4 py-3 transition-colors",
        isCurrentUser && "border-l-2 border-l-primary bg-primary/5",
      )}
    >
      <span className="w-6 shrink-0 font-mono text-xs text-muted-foreground tabular-nums">
        {String(rank).padStart(2, "0")}
      </span>
      <span className="min-w-0 flex-1 truncate text-sm font-medium">
        {entry.displayName}
      </span>
      <span className="shrink-0 font-heading text-sm font-semibold tabular-nums">
        {entry.score}
        <span className="ms-1 text-[0.625rem] font-normal tracking-widest text-muted-foreground uppercase">
          pts
        </span>
      </span>
    </div>
  );
}

function Pagination({
  page,
  totalPages,
  onPageChange,
}: {
  page: number;
  totalPages: number;
  onPageChange: (page: number) => void;
}) {
  const pages = Array.from({ length: totalPages }, (_, i) => i + 1);

  return (
    <div className="mt-6 flex items-center justify-center gap-2">
      <button
        type="button"
        onClick={() => onPageChange(Math.max(1, page - 1))}
        disabled={page === 1}
        className="flex size-7 items-center justify-center border border-border text-muted-foreground transition-colors hover:border-primary hover:bg-primary hover:text-primary-foreground disabled:pointer-events-none disabled:opacity-30"
      >
        <ChevronLeftIcon className="size-3.5" />
      </button>

      {pages.map((p) => (
        <button
          key={p}
          type="button"
          onClick={() => onPageChange(p)}
          className={cn(
            "flex size-7 items-center justify-center border text-xs font-semibold transition-colors",
            p === page
              ? "border-primary bg-primary text-primary-foreground"
              : "border-border text-muted-foreground hover:border-primary hover:bg-primary hover:text-primary-foreground",
          )}
        >
          {p}
        </button>
      ))}

      <button
        type="button"
        onClick={() => onPageChange(Math.min(totalPages, page + 1))}
        disabled={page === totalPages}
        className="flex size-7 items-center justify-center border border-border text-muted-foreground transition-colors hover:border-primary hover:bg-primary hover:text-primary-foreground disabled:pointer-events-none disabled:opacity-30"
      >
        <ChevronRightIcon className="size-3.5" />
      </button>
    </div>
  );
}
