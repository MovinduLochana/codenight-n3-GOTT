import { Skeleton } from "@/components/ui/skeleton";


export function SuspenseLoader() {
  return (
    <div className="flex min-h-screen w-full items-center justify-center p-4">
      <div className="flex w-full max-w-lg flex-col items-center justify-center rounded-xl border border-white/10 bg-card p-12 text-center shadow-lg">
        <Skeleton className="size-12 rounded-full" />
        <Skeleton className="mt-5 h-4 w-40 rounded" />
        <Skeleton className="mt-3 h-4 w-56 rounded" />
        <Skeleton className="mt-6 h-10 w-full max-w-xs rounded-md" />
        <Skeleton className="mt-3 h-10 w-full max-w-xs rounded-md" />
      </div>
    </div>
  );
}
