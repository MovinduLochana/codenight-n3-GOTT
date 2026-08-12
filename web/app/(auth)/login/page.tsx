"use client";

import { useActionState, useEffect, Suspense } from "react";
import { useSearchParams } from "next/navigation";
import {
  loginWithCredentials,
  initiateOAuthFlow,
  type AuthFormState,
} from "@/app/actions/auth";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";

function MozillaIcon() {
  return (
    <svg viewBox="0 0 100 100" className="size-5" aria-hidden="true" fill="none">
      <rect width="100" height="100" rx="18" fill="oklch(0.60 0.22 25)" />
      <text
        x="50"
        y="68"
        textAnchor="middle"
        fontSize="52"
        fontWeight="900"
        fontFamily="serif"
        fill="white"
        letterSpacing="-2"
      >
        M
      </text>
    </svg>
  );
}
function Spinner() {
  return (
    <svg
      className="size-4 animate-spin"
      xmlns="http://www.w3.org/2000/svg"
      fill="none"
      viewBox="0 0 24 24"
      aria-hidden="true"
    >
      <circle
        className="opacity-25"
        cx="12"
        cy="12"
        r="10"
        stroke="currentColor"
        strokeWidth="4"
      />
      <path
        className="opacity-75"
        fill="currentColor"
        d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"
      />
    </svg>
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
  const message = ERROR_MESSAGES[code] ?? "An unexpected error occurred.";
  return (
    <div
      role="alert"
      className="flex items-start gap-3 rounded-xl border border-red-500/30 bg-red-500/10 px-4 py-3 text-sm text-red-300"
    >
      <svg className="mt-0.5 size-4 shrink-0" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
        <path
          fillRule="evenodd"
          d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-8-5a.75.75 0 01.75.75v4.5a.75.75 0 01-1.5 0v-4.5A.75.75 0 0110 5zm0 10a1 1 0 100-2 1 1 0 000 2z"
          clipRule="evenodd"
        />
      </svg>
      {message}
    </div>
  );
}


function FieldError({ errors }: { errors?: string[] }) {
  if (!errors?.length) return null;
  return (
    <p className="mt-1.5 text-xs text-red-400" role="alert">
      {errors[0]}
    </p>
  );
}


function Divider() {
  return (
    <div className="relative my-6 flex items-center">
      <div className="flex-1 border-t border-white/10" />
      <span className="mx-4 text-xs font-medium tracking-wider text-white/30 uppercase">
        or
      </span>
      <div className="flex-1 border-t border-white/10" />
    </div>
  );
}


function LoginForm() {
  const searchParams = useSearchParams();
  const errorCode = searchParams.get("error");

  const [state, action, pending] = useActionState<AuthFormState, FormData>(
    loginWithCredentials,
    undefined
  );

  useEffect(() => {
    if (state?.message || state?.errors) {
      window.scrollTo({ top: 0, behavior: "smooth" });
    }
  }, [state]);

  return (
    <div className="mx-auto w-full max-w-md">
      {/* Card */}
      <div className="login-card relative overflow-hidden rounded-2xl border border-white/10 bg-white/5 px-8 py-10 shadow-2xl backdrop-blur-xl">
        {/* Shine overlay */}
        <div
          className="pointer-events-none absolute inset-0 rounded-2xl"
          style={{
            background:
              "linear-gradient(135deg, oklch(1 0 0 / 6%) 0%, transparent 50%)",
          }}
          aria-hidden="true"
        />

        {/* Header */}
        <div className="relative mb-8 text-center">
          <div className="mb-5 inline-flex items-center gap-2.5 rounded-full border border-white/10 bg-white/8 px-4 py-1.5 text-xs font-semibold tracking-widest text-white/60 uppercase">
            <div className="size-1.5 rounded-full bg-emerald-400 shadow-[0_0_6px_oklch(0.78_0.19_155)]" />
            SLIIT Mozilla
          </div>
          <h1 className="text-2xl font-bold tracking-tight text-white">
            Welcome back
          </h1>
          <p className="mt-2 text-sm text-white/45">
            Sign in to your SLIIT Mozilla account
          </p>
        </div>

        {/* Callback error banner — only show if there's no newer server action error */}
        {!state?.message && <ErrorBanner code={errorCode} />}

        {/* Server action global error */}
        {state?.message && (
          <div
            role="alert"
            className="mb-4 flex items-start gap-3 rounded-xl border border-red-500/30 bg-red-500/10 px-4 py-3 text-sm text-red-300"
          >
            <svg className="mt-0.5 size-4 shrink-0" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
              <path
                fillRule="evenodd"
                d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-8-5a.75.75 0 01.75.75v4.5a.75.75 0 01-1.5 0v-4.5A.75.75 0 0110 5zm0 10a1 1 0 000 2 1 1 0 000-2z"
                clipRule="evenodd"
              />
            </svg>
            {state.message}
          </div>
        )}

        {/* Email / Password Form */}
        <form action={action} className="relative space-y-5" noValidate>
          {/* Email */}
          <div>
            <label
              htmlFor="login-email"
              className="mb-2 block text-xs font-semibold tracking-wider text-white/50 uppercase"
            >
              Email address
            </label>
            <div className="login-input-wrap">
              <Input
                id="login-email"
                name="email"
                type="email"
                placeholder="you@example.com"
                autoComplete="email"
                aria-invalid={!!state?.errors?.email}
                aria-describedby={
                  state?.errors?.email ? "email-error" : undefined
                }
                className="login-input"
              />
            </div>
            <FieldError errors={state?.errors?.email} />
          </div>

          {/* Password */}
          <div>
            <div className="mb-2 flex items-center justify-between">
              <label
                htmlFor="login-password"
                className="text-xs font-semibold tracking-wider text-white/50 uppercase"
              >
                Password
              </label>
            </div>
            <div className="login-input-wrap">
              <Input
                id="login-password"
                name="password"
                type="password"
                placeholder="••••••••"
                autoComplete="current-password"
                aria-invalid={!!state?.errors?.password}
                aria-describedby={
                  state?.errors?.password ? "password-error" : undefined
                }
                className="login-input"
              />
            </div>
            <FieldError errors={state?.errors?.password} />
          </div>

          {/* Submit */}
          <Button
            type="submit"
            disabled={pending}
            className="login-btn mt-2 w-full"
            size="lg"
          >
            {pending ? (
              <span className="flex items-center gap-2">
                <Spinner />
                Signing in…
              </span>
            ) : (
              "Sign in"
            )}
          </Button>
        </form>

        <Divider />

        {/* OAuth button */}
        <form action={initiateOAuthFlow}>
          <Button
            type="submit"
            variant="outline"
            className="oauth-btn w-full"
            size="lg"
          >
            <MozillaIcon />
            Continue with SLIIT Mozilla
          </Button>
        </form>

        {/* Footer */}
        <p className="relative mt-8 text-center text-xs text-white/30">
          By signing in, you agree to the{" "}
          <a
            href="https://www.sliitmozilla.org/"
            target="_blank"
            rel="noreferrer"
            className="text-white/50 underline-offset-2 hover:text-white/70 hover:underline transition-colors"
          >
            SLIIT Mozilla Community
          </a>{" "}
          terms.
        </p>
      </div>

      {/* Styles */}
      <style>{`
        .login-card {
          box-shadow:
            0 0 0 1px oklch(1 0 0 / 6%),
            0 40px 80px oklch(0 0 0 / 60%),
            inset 0 1px 0 oklch(1 0 0 / 10%);
        }

        .login-input-wrap {
          position: relative;
        }

        .login-input {
          width: 100%;
          border-radius: 12px !important;
          background: oklch(1 0 0 / 4%) !important;
          border: 1px solid oklch(1 0 0 / 10%) !important;
          border-bottom: 1px solid oklch(1 0 0 / 10%) !important;
          padding: 0.625rem 1rem !important;
          font-size: 0.875rem !important;
          color: white !important;
          height: 44px !important;
          transition: border-color 0.2s, box-shadow 0.2s, background 0.2s;
        }
        .login-input::placeholder {
          color: oklch(1 0 0 / 25%);
        }
        .login-input:focus-visible {
          border-color: oklch(0.60 0.20 150 / 80%) !important;
          box-shadow: 0 0 0 3px oklch(0.60 0.20 150 / 20%) !important;
          background: oklch(1 0 0 / 6%) !important;
          outline: none !important;
        }
        .login-input[aria-invalid="true"] {
          border-color: oklch(0.65 0.22 25 / 60%) !important;
          box-shadow: 0 0 0 3px oklch(0.65 0.22 25 / 15%) !important;
        }

        .login-btn {
          background: linear-gradient(
            135deg,
            oklch(0.55 0.22 150) 0%,
            oklch(0.48 0.18 160) 100%
          ) !important;
          border: 1px solid oklch(0.65 0.20 150 / 40%) !important;
          border-radius: 12px !important;
          color: white !important;
          font-weight: 600 !important;
          letter-spacing: 0.04em !important;
          box-shadow:
            0 1px 0 oklch(1 0 0 / 15%) inset,
            0 4px 20px oklch(0.55 0.22 150 / 30%) !important;
          transition: opacity 0.15s, transform 0.15s, box-shadow 0.15s !important;
        }
        .login-btn:hover:not(:disabled) {
          opacity: 0.92;
          box-shadow:
            0 1px 0 oklch(1 0 0 / 15%) inset,
            0 6px 28px oklch(0.55 0.22 150 / 40%) !important;
        }
        .login-btn:active:not(:disabled) {
          transform: translateY(1px);
        }
        .login-btn:disabled {
          opacity: 0.55;
        }

        .oauth-btn {
          border-radius: 12px !important;
          border: 1px solid oklch(1 0 0 / 12%) !important;
          background: oklch(1 0 0 / 5%) !important;
          color: oklch(0.92 0 0) !important;
          font-weight: 500 !important;
          letter-spacing: 0.02em !important;
          gap: 10px !important;
          transition: background 0.15s, border-color 0.15s, transform 0.15s !important;
        }
        .oauth-btn:hover:not(:disabled) {
          background: oklch(1 0 0 / 9%) !important;
          border-color: oklch(1 0 0 / 18%) !important;
        }
        .oauth-btn:active:not(:disabled) {
          transform: translateY(1px);
        }
      `}</style>
    </div>
  );
}


export default function LoginPage() {
  return (
    <Suspense
      fallback={
        <div className="mx-auto w-full max-w-md">
          <div className="login-card relative overflow-hidden rounded-2xl border border-white/10 bg-white/5 px-8 py-10 shadow-2xl backdrop-blur-xl animate-pulse">
            <div className="h-5 w-28 mx-auto rounded-full bg-white/10 mb-6" />
            <div className="h-7 w-40 mx-auto rounded bg-white/10 mb-3" />
            <div className="h-4 w-56 mx-auto rounded bg-white/5 mb-8" />
            <div className="space-y-5">
              <div className="h-11 rounded-xl bg-white/5" />
              <div className="h-11 rounded-xl bg-white/5" />
              <div className="h-11 rounded-xl bg-white/10 mt-2" />
            </div>
            <div className="my-6 h-px bg-white/10" />
            <div className="h-11 rounded-xl bg-white/5" />
          </div>
        </div>
      }
    >
      <LoginForm />
    </Suspense>
  );
}
