import { Suspense } from "react";

import { SuspenseLoader } from "@/components/common/suspense-loader";
import { AppShell } from "@/components/shell/app-shell";
import { PublicShell } from "@/components/shell/public-shell";
import { getSession } from "@/lib/session";

export default function LeaderboardLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <Suspense fallback={<SuspenseLoader />}>
      <LeaderboardLayoutContent>{children}</LeaderboardLayoutContent>
    </Suspense>
  );
}

async function LeaderboardLayoutContent({
  children,
}: {
  children: React.ReactNode;
}) {
  const session = await getSession();

  if (!session) {
    return <PublicShell>{children}</PublicShell>;
  }

  return <AppShell>{children}</AppShell>;
}
