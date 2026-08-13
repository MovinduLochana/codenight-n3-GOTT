import Image from "next/image";
import { Suspense } from "react";
import { initiateOAuthFlow } from "@/app/actions/auth";
import { Button } from "../ui/button";
import { ErrorBanner } from "./error-banner";

export function LoginCard() {
  return (
    <div className="w-full max-w-lg border border-white/10 bg-[oklch(0.21_0.028_264)] px-12 py-16 text-center">
      <Image
        src="/assets/logo.svg"
        alt="CodeNight Logo"
        width={176}
        height={87}
        priority
        className="mx-auto h-auto w-64"
      />

      <h1 className="mt-5 text-3xl font-bold text-foreground">
        Welcome to CodeNight
      </h1>
      <p className="mx-auto mt-4 max-w-88 font-mono text-sm leading-relaxed text-muted-foreground">
        Authenticate to access your workspace and terminal sessions.
      </p>

      <div className="mt-10">
        <Suspense fallback={<div>Loading...</div>}>
          <ErrorBanner />
        </Suspense>

        <form action={initiateOAuthFlow}>
          <Button
            type="submit"
            className="flex h-12 w-full items-center justify-center gap-3 font-mono text-sm font-medium  transition-opacity hover:opacity-90"
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
      </div>
    </div>
  );
}
