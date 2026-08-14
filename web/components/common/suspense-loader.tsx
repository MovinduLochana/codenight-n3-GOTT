import { Loader2Icon } from "lucide-react";

export function SuspenseLoader() {
  return (
    <div className="flex min-h-75 w-full max-w-lg flex-col items-center justify-center rounded-xl border border-white/10 bg-card p-12 text-center shadow-lg">
      <Loader2Icon className="size-10 animate-spin text-primary" />
      <p className="mt-4 font-mono text-xs text-muted-foreground tracking-wide uppercase animate-pulse">
        Loading session...
      </p>
    </div>
  );
}
