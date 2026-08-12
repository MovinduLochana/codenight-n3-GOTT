"use client";

import { BellIcon } from "lucide-react";
import Image from "next/image";
import Link from "next/link";
import { usePathname } from "next/navigation";

import type { Category } from "@/lib/content";
import { cn } from "@/lib/utils";

const linkClass =
  "border-b-2 px-3 py-1.5 text-xs font-semibold tracking-widest uppercase transition-colors";

export function TopBar({ categories }: { categories: Category[] }) {
  const pathname = usePathname();
  const active = categories.find(
    (category) => category.id === pathname.split("/")[2],
  );

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
        <button
          type="button"
          className={cn(
            linkClass,
            "border-transparent text-muted-foreground hover:text-foreground",
          )}
        >
          Dashboard
        </button>

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

        <button
          type="button"
          className={cn(
            linkClass,
            "border-transparent text-muted-foreground hover:text-foreground",
          )}
        >
          Resources
        </button>
      </nav>

      <div className="ms-auto flex items-center gap-4">
        {active ? (
          <span className="hidden text-xs tracking-wide text-muted-foreground uppercase lg:inline">
            Chapter {active.number}: {active.title}
          </span>
        ) : null}

        <BellIcon className="size-4 text-muted-foreground" />

        <div className="flex size-8 items-center justify-center border border-border bg-muted text-[0.625rem] font-semibold tracking-wide uppercase">
          n3
        </div>
      </div>
    </header>
  );
}
