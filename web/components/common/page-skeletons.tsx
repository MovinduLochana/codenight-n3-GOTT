import { Skeleton } from "@/components/ui/skeleton";

// Shape-matched loading states for each route's Suspense boundary. Each one
// mirrors the real layout it's standing in for, so streaming in the actual
// content doesn't cause a visible reflow/pop.

export function AssessmentListSkeleton() {
  return (
    <main className="min-h-0 flex-1 overflow-y-auto">
      <div className="mx-auto max-w-3xl px-6 py-10">
        <Skeleton className="h-8 w-64 rounded" />
        <Skeleton className="mt-3 h-4 w-full max-w-md rounded" />

        <ul className="mt-8 flex flex-col gap-2">
          {Array.from({ length: 8 }, (_, index) => (
            <li
              // biome-ignore lint/suspicious/noArrayIndexKey: static placeholder count, never reordered
              key={index}
              className="flex items-center gap-3 border border-border bg-card p-4"
            >
              <Skeleton className="h-3 w-4 shrink-0 rounded" />
              <Skeleton className="h-4 w-1/3 min-w-0 flex-1 rounded" />
              <Skeleton className="h-5 w-20 shrink-0 rounded-full" />
              <Skeleton className="h-3 w-10 shrink-0 rounded" />
              <Skeleton className="h-3 w-24 shrink-0 rounded" />
            </li>
          ))}
        </ul>
      </div>
    </main>
  );
}

export function LeaderboardSkeleton() {
  return (
    <main className="min-h-0 flex-1 overflow-y-auto">
      <div className="mx-auto max-w-3xl px-6 py-10">
        <div className="flex items-center gap-2">
          <Skeleton className="size-5 rounded" />
          <Skeleton className="h-7 w-40 rounded" />
        </div>
        <Skeleton className="mt-3 h-4 w-full max-w-lg rounded" />
        <Skeleton className="mt-1.5 h-4 w-2/3 max-w-md rounded" />

        <ul className="mt-8 flex flex-col gap-1.5">
          {Array.from({ length: 10 }, (_, index) => (
            <li
              // biome-ignore lint/suspicious/noArrayIndexKey: static placeholder count, never reordered
              key={index}
              className="flex items-center gap-3 border border-border bg-card p-4"
            >
              <Skeleton className="h-3 w-4 shrink-0 rounded" />
              <Skeleton
                className="h-4 rounded"
                style={{ width: `${45 - index}%` }}
              />
              <Skeleton className="ms-auto h-4 w-10 shrink-0 rounded" />
            </li>
          ))}
        </ul>
      </div>
    </main>
  );
}

export function ExerciseHeaderSkeleton() {
  return (
    <>
      <Skeleton className="h-2.5 w-56 rounded" />
      <Skeleton className="mt-3 mb-6 h-8 w-80 rounded" />
    </>
  );
}

export function ExerciseWorkbenchSkeleton() {
  return (
    <div className="grid min-h-0 flex-1 gap-6 lg:grid-cols-[380px_1fr]">
      <div className="flex min-w-0 flex-col gap-3 lg:max-h-full lg:overflow-y-auto lg:pe-2">
        <Skeleton className="h-5 w-2/3 rounded" />
        <Skeleton className="h-4 w-full rounded" />
        <Skeleton className="h-4 w-11/12 rounded" />
        <Skeleton className="h-4 w-4/5 rounded" />
        <Skeleton className="mt-4 h-24 w-full rounded" />
        <Skeleton className="mt-2 h-4 w-full rounded" />
        <Skeleton className="h-4 w-3/4 rounded" />
      </div>

      <div className="flex min-h-0 min-w-0 flex-col border border-border">
        <div className="flex items-center justify-between border-b border-border bg-card px-3 py-2">
          <Skeleton className="h-3 w-24 rounded" />
          <div className="flex gap-2">
            <Skeleton className="h-7 w-16 rounded-md" />
            <Skeleton className="h-7 w-16 rounded-md" />
          </div>
        </div>
        <div className="flex-1 space-y-2 bg-background p-4">
          {Array.from({ length: 10 }, (_, index) => (
            <Skeleton
              // biome-ignore lint/suspicious/noArrayIndexKey: static placeholder count, never reordered
              key={index}
              className="h-3.5 rounded"
              style={{ width: `${85 - ((index * 13) % 55)}%` }}
            />
          ))}
        </div>
      </div>
    </div>
  );
}

export function ChapterQuizSkeleton() {
  return (
    <main className="min-h-0 flex-1 overflow-y-auto">
      <div className="mx-auto max-w-3xl px-6 py-10">
        <Skeleton className="mb-8 h-3.5 w-28 rounded" />
        <Skeleton className="h-2.5 w-40 rounded" />
        <Skeleton className="mt-2 mb-2 h-8 w-48 rounded" />
        <Skeleton className="mb-8 h-4 w-72 rounded" />

        <div className="flex flex-col gap-6 border border-border bg-card p-6">
          <Skeleton className="h-3 w-20 rounded" />
          <Skeleton className="h-6 w-full rounded" />
          <Skeleton className="h-6 w-4/5 rounded" />
          <div className="mt-2 flex flex-col gap-2">
            {Array.from({ length: 4 }, (_, index) => (
              <Skeleton
                // biome-ignore lint/suspicious/noArrayIndexKey: static placeholder count, never reordered
                key={index}
                className="h-11 w-full rounded-md"
              />
            ))}
          </div>
        </div>

        <div className="mt-8 flex items-center justify-between gap-3">
          <Skeleton className="h-9 w-32 rounded-md" />
          <Skeleton className="h-9 w-32 rounded-md" />
        </div>
      </div>
    </main>
  );
}
