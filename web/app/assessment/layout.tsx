import { AppShell } from "@/components/shell/app-shell";

export default function AssessmentLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return <AppShell>{children}</AppShell>;
}
