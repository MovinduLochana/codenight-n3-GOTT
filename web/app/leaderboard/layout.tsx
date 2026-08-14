import { AppShell } from "@/components/shell/app-shell";
import { PublicShell } from "@/components/shell/public-shell";
import { getSession } from "@/lib/session";

export default async function LeaderboardLayout({
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
