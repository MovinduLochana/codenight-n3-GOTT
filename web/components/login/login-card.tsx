import Image from "next/image";
import { Suspense } from "react";
import { initiateOAuthFlow } from "@/actions/auth";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { ErrorBanner } from "./error-banner";

export function LoginCard() {
  return (
    <Card className="w-full max-w-lg border-white/10 bg-[oklch(0.21_0.028_264)] py-12 text-center shadow-xl">
      <CardHeader className="items-center">
        <Image
          src="/assets/logo.svg"
          alt="CodeNight Logo"
          width={176}
          height={87}
          priority
          className="mx-auto h-auto w-64"
        />
        <CardTitle className="mt-5 text-2xl font-bold text-foreground">
          Welcome to CodeNight
        </CardTitle>
        <CardDescription className="mx-auto mt-2 max-w-sm font-mono text-xs leading-relaxed text-muted-foreground">
          Authenticate to access your workspace and terminal sessions.
        </CardDescription>
      </CardHeader>

      <CardContent className="mt-6">
        <Suspense fallback={<div>Loading...</div>}>
          <ErrorBanner />
        </Suspense>

        <form action={initiateOAuthFlow}>
          <Button
            type="submit"
            size="lg"
            className="flex h-12 w-full items-center justify-center gap-3 font-mono text-xs font-semibold tracking-wide uppercase transition-opacity hover:opacity-95"
          >
            <Image
              src="/assets/mozilla-logo.png"
              alt="MOZILLA Logo"
              width={28}
              height={28}
              className="size-7"
            />
            Sign in with SLIIT Mozilla
          </Button>
        </form>
      </CardContent>
    </Card>
  );
}
