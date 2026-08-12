"use client";

import Image from "next/image";
import { Suspense } from "react";
import { useSearchParams } from "next/navigation";
import { initiateOAuthFlow } from "@/app/actions/auth";

function MozillaIcon() {
  return (
    <Image
      src="/assets/mozilla-logo.png"
      alt=""
      width={28}
      height={28}
      className="size-7"
    />
  );
}

const ERROR_MESSAGES: Record<string, string> = {
  missing_code: "Authentication code is missing. Please try again.",
  expired_code: "Authentication code expired. Please log in again.",
  token_exchange_failed: "Failed to exchange authentication code. Try again.",
  session_fetch_failed: "Could not retrieve your session. Try again.",
  no_token: "Unexpected response from auth server. Try again.",
  no_user_id: "Could not identify your account. Try again.",
  network_error: "Network error reaching the auth server. Try again later.",
};

function ErrorBanner({ code }: { code: string | null }) {
  if (!code) return null;

  return (
    <p
      role="alert"
      className="mb-6 border border-destructive/30 bg-destructive/10 px-4 py-3 text-center font-mono text-xs leading-relaxed text-destructive"
    >
      {ERROR_MESSAGES[code] ?? "An unexpected error occurred."}
    </p>
  );
}

function Wordmark() {
  return (
    <Image
      src="/assets/logo.svg"
      alt="CodeNight"
      width={176}
      height={87}
      priority
      unoptimized
      className="mx-auto h-auto w-64"
    />
  );
}

function LoginCard() {
  const searchParams = useSearchParams();

  return (
    <div className="w-full max-w-lg border border-white/10 bg-[oklch(0.21_0.028_264)] px-12 py-16 text-center">
      <Wordmark />

      <h1 className="mt-5 text-3xl font-bold text-foreground">
        Welcome to CodeNight
      </h1>
      <p className="mx-auto mt-4 max-w-[22rem] font-mono text-sm leading-relaxed text-muted-foreground">
        Authenticate to access your workspace and terminal sessions.
      </p>

      <div className="mt-10">
        <ErrorBanner code={searchParams.get("error")} />

        <form action={initiateOAuthFlow}>
          <button
            type="submit"
            className="flex h-12 w-full items-center justify-center gap-3 bg-foreground font-mono text-sm font-medium text-background transition-opacity hover:opacity-90"
          >
            <MozillaIcon />
            Sign in with SLIIT Mozilla
          </button>
        </form>
      </div>
    </div>
  );
}

export default function LoginPage() {
  return (
    <Suspense
      fallback={
        <div className="w-full max-w-lg animate-pulse border border-white/10 bg-[oklch(0.21_0.028_264)] px-12 py-16">
          <div className="mx-auto h-16 w-64 bg-white/10" />
          <div className="mx-auto mt-5 h-9 w-72 bg-white/10" />
          <div className="mx-auto mt-4 h-10 w-full bg-white/5" />
          <div className="mt-10 h-12 w-full bg-white/10" />
        </div>
      }
    >
      <LoginCard />
    </Suspense>
  );
}
