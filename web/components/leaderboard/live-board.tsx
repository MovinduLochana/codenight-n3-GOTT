"use client";

import { useEffect, useState } from "react";

import type { LeaderboardEntry } from "@/lib/leaderboard";
import { cn } from "@/lib/utils";

const POLL_INTERVAL_MS = 3000;

export function LiveLeaderboard({
  initialEntries,
  currentUserId,
}: {
  initialEntries: LeaderboardEntry[];
  currentUserId: string | null;
}) {
  const [entries, setEntries] = useState(initialEntries);

  useEffect(() => {
    let cancelled = false;

    async function poll() {
      try {
        const response = await fetch("/api/leaderboard", { cache: "no-store" });
        if (!response.ok || cancelled) return;
        const body = (await response.json()) as { entries: LeaderboardEntry[] };
        if (!cancelled) setEntries(body.entries);
      } catch {
        // Transient network hiccup — the next tick will retry.
      }
    }

    const id = setInterval(poll, POLL_INTERVAL_MS);
    return () => {
      cancelled = true;
      clearInterval(id);
    };
  }, []);

  if (entries.length === 0) {
    return (
      <div className="mt-8 border border-border bg-card p-8 text-center text-sm text-muted-foreground">
        No one's on the board yet — complete a chapter quiz or an exercise to
        appear here.
      </div>
    );
  }

  const [first, second, third, ...rest] = entries;
  const podium: { entry: LeaderboardEntry; rank: 1 | 2 | 3 }[] = [
    second && { entry: second, rank: 2 as const },
    first && { entry: first, rank: 1 as const },
    third && { entry: third, rank: 3 as const },
  ].filter((slot): slot is { entry: LeaderboardEntry; rank: 1 | 2 | 3 } =>
    Boolean(slot),
  );

  return (
    <>
      {podium.length > 0 ? (
        <div className="mt-10 grid grid-cols-3 items-end gap-3">
          {[0, 1, 2].map((slot) => {
            const item = podium[slot];
            return item ? (
              <PodiumCard
                key={item.entry.userId}
                entry={item.entry}
                rank={item.rank}
                isCurrentUser={item.entry.userId === currentUserId}
              />
            ) : (
              <div key={slot} />
            );
          })}
        </div>
      ) : null}

      {rest.length > 0 ? (
        <ul className="mt-8 flex flex-col gap-2">
          {rest.map((entry, index) => (
            <li key={entry.userId}>
              <RankRow
                rank={index + 4}
                entry={entry}
                isCurrentUser={entry.userId === currentUserId}
              />
            </li>
          ))}
        </ul>
      ) : null}
    </>
  );
}

function PodiumCard({
  entry,
  rank,
  isCurrentUser,
}: {
  entry: LeaderboardEntry;
  rank: 1 | 2 | 3;
  isCurrentUser: boolean;
}) {
  return (
    <div
      className={cn(
        "flex flex-col items-center border bg-card px-3 pb-5 text-center",
        rank === 1 ? "min-h-[220px] pt-10" : "min-h-[176px] pt-6",
        rank === 1 ? "border-primary/50" : "border-border",
        isCurrentUser &&
          "ring-2 ring-primary/40 ring-offset-2 ring-offset-background",
      )}
    >
      <div
        className={cn(
          "flex shrink-0 items-center justify-center border font-heading font-semibold",
          rank === 1
            ? "size-10 border-primary text-lg text-primary"
            : "size-8 border-border text-sm text-muted-foreground",
        )}
      >
        {rank}
      </div>
      <p className="mt-3 w-full truncate text-sm font-medium">
        {entry.displayName}
      </p>
      <p className="mt-1 font-heading text-xl font-semibold text-primary">
        {entry.score}
      </p>
      <p className="text-[0.625rem] tracking-widest text-muted-foreground uppercase">
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
        "flex items-center gap-3 border border-border bg-card p-4 transition-colors",
        isCurrentUser && "border-primary/40 bg-primary/5",
      )}
    >
      <span className="font-mono text-xs text-primary">
        {String(rank).padStart(2, "0")}
      </span>
      <span className="truncate font-medium">{entry.displayName}</span>
      <span className="ms-auto font-heading text-sm font-semibold">
        {entry.score}
        <span className="ms-1 text-[0.625rem] font-normal tracking-widest text-muted-foreground uppercase">
          pts
        </span>
      </span>
    </div>
  );
}
