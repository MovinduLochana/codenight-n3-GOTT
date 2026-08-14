"use client";

import { BellIcon } from "lucide-react";
import Image from "next/image";
import Link from "next/link";
import { usePathname } from "next/navigation";

import { Avatar, AvatarFallback } from "@/components/ui/avatar";
import { Badge } from "@/components/ui/badge";
import type { Category } from "@/lib/content";
import { cn } from "@/lib/utils";

const linkClass =
  "border-b-2 px-3 py-1.5 text-xs font-semibold tracking-widest uppercase transition-colors";

export function TopBar({ categories }: { categories: Category[] }) {
  const pathname = usePathname();
  const pathParts = pathname.split("/");
  const activeChapterId = pathParts[1] === "learn" ? pathParts[2] : undefined;
  const active = categories.find((category) => category.id === activeChapterId);

  return (
    <header className="flex h-14 shrink-0 items-center gap-3 border-b border-border bg-sidebar px-4">
      <Link href="/learn" className="flex items-center">
        <Image
          src="/assets/logo.svg"
          alt="CodeNight"
          width={89}
          height={44}
          priority
          unoptimized
          className="h-11 w-auto"
        />
      </Link>

      <nav className="hidden items-center gap-1 md:flex">
        <Link
          href="/learn"
          className={cn(
            linkClass,
            pathname.startsWith("/learn")
              ? "border-primary text-foreground"
              : "border-transparent text-muted-foreground hover:text-foreground",
          )}
        >
          Curriculum
        </Link>
      </nav>

      <div className="ms-auto flex items-center gap-4">
        {active ? (
          <Badge variant="secondary" className="hidden lg:inline-flex">
            Chapter {active.number}: {active.title}
          </Badge>
        ) : null}

        <BellIcon className="size-4 text-muted-foreground transition-colors hover:text-foreground cursor-pointer" />

        <Avatar size="sm">
          <AvatarFallback className="font-mono text-xs font-bold text-primary">
            N3
          </AvatarFallback>
        </Avatar>
      </div>
    </header>
  );
}
