import Link from "next/link";
import { Suspense } from "react";

import { SuspenseLoader } from "@/components/common/suspense-loader";
import { LoginCard } from "@/components/login/login-card";

export default function LoginPage() {
  return (
    <Suspense fallback={<SuspenseLoader />}>
      <div className="flex w-full max-w-lg flex-col items-center gap-6">
        <div className="flex border border-border">
          <span className="bg-primary px-6 py-2 text-xs font-semibold tracking-widest text-primary-foreground uppercase">
            Login
          </span>
          <Link
            href="/leaderboard"
            className="px-6 py-2 text-xs font-semibold tracking-widest text-muted-foreground uppercase transition-colors hover:text-foreground"
          >
            Leaderboard
          </Link>
        </div>

        <LoginCard />
      </div>
    </Suspense>
  );
}
