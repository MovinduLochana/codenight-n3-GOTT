"use client";

import { useSearchParams } from "next/navigation";

const ERROR_MESSAGES: Record<string, string> = {
  missing_code: "Authentication code is missing. Please try again.",
  expired_code: "Authentication code expired. Please log in again.",
  token_exchange_failed: "Failed to exchange authentication code. Try again.",
  session_fetch_failed: "Could not retrieve your session. Try again.",
  no_token: "Unexpected response from auth server. Try again.",
  no_user_id: "Could not identify your account. Try again.",
  network_error: "Network error reaching the auth server. Try again later.",
};

export function ErrorBanner() {
  const searchParams = useSearchParams();
  const code = searchParams.get("error");

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
