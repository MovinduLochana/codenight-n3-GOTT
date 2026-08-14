import Image from "next/image";
import Link from "next/link";

import { buttonVariants } from "@/components/ui/button";
import { cn } from "@/lib/utils";

export function PublicShell({ children }: { children: React.ReactNode }) {
  return (
    <div className="flex h-dvh flex-col overflow-hidden">
      <header className="flex h-14 shrink-0 items-center gap-3 border-b border-border bg-sidebar px-4">
        <Link href="/login" className="flex items-center">
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

        <nav className="ms-6 hidden items-center gap-6 md:flex">
          <span className="text-xs font-semibold tracking-widest text-primary uppercase">
            Leaderboard
          </span>
        </nav>

        <Link
          href="/login"
          className={cn(
            buttonVariants({ variant: "ghost", size: "sm" }),
            "ms-auto bg-transparent hover:bg-transparent",
          )}
        >
          <Image
            src="/assets/login_logo.svg"
            alt="Login"
            width={200}
            height={54}
            className="h-11 w-auto"
          />
        </Link>
      </header>

      <div className="flex min-h-0 flex-1">{children}</div>
    </div>
  );
}
