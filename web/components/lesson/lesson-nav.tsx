import { ArrowLeftIcon, ArrowRightIcon, ListChecksIcon } from "lucide-react";
import Link from "next/link";

import { buttonVariants } from "@/components/ui/button";
import type { AdjacentTopic } from "@/lib/content";
import { cn } from "@/lib/utils";

export function LessonNav({
  categoryId,
  previous,
  next,
}: {
  categoryId: string;
  previous: AdjacentTopic | null;
  next: AdjacentTopic | null;
}) {
  return (
    <nav className="mt-12 flex items-stretch justify-between gap-3 border-t border-border pt-6">
      {previous ? (
        <Link
          href={`/learn/${previous.categoryId}/${previous.topicId}`}
          className={cn(
            buttonVariants({ variant: "outline" }),
            "h-auto flex-col items-start gap-1 px-4 py-3 normal-case",
          )}
        >
          <span className="flex items-center gap-1.5 text-[0.625rem] font-semibold tracking-widest text-muted-foreground uppercase">
            <ArrowLeftIcon className="size-3" />
            Previous
          </span>
          <span className="text-xs font-medium">{previous.title}</span>
        </Link>
      ) : (
        <span />
      )}

      {next ? (
        <Link
          href={`/learn/${categoryId}/quiz`}
          className={cn(
            buttonVariants({ variant: "ghost", size: "sm" }),
            "shrink-0 self-center normal-case",
          )}
        >
          <ListChecksIcon />
          Chapter Quiz
        </Link>
      ) : null}

      {next ? (
        <Link
          href={`/learn/${next.categoryId}/${next.topicId}`}
          className={cn(
            buttonVariants({ variant: "outline" }),
            "h-auto flex-col items-end gap-1 px-4 py-3 text-end normal-case",
          )}
        >
          <span className="flex items-center gap-1.5 text-[0.625rem] font-semibold tracking-widest text-muted-foreground uppercase">
            Next
            <ArrowRightIcon className="size-3" />
          </span>
          <span className="text-xs font-medium">{next.title}</span>
        </Link>
      ) : (
        <Link
          href={`/learn/${categoryId}/quiz`}
          className={cn(
            buttonVariants({ variant: "default" }),
            "h-auto flex-col items-end gap-1 px-4 py-3 text-end normal-case",
          )}
        >
          <span className="flex items-center gap-1.5 text-[0.625rem] font-semibold tracking-widest uppercase">
            Next
            <ArrowRightIcon className="size-3" />
          </span>
          <span className="text-xs font-medium">Take the Chapter Quiz</span>
        </Link>
      )}
    </nav>
  );
}
