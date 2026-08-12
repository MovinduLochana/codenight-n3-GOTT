import { getSession } from "@/lib/session";
import { logout } from "@/app/actions/auth";
import { redirect } from "next/navigation";

export default async function Home() {
  const session = await getSession();
  
  if (!session) redirect("/login");
  return (
    <div className="flex min-h-screen flex-col items-center justify-center gap-6 bg-[oklch(0.11_0.018_240)] p-8">
      <div className="w-full max-w-md rounded-2xl border border-white/10 bg-white/5 p-8 text-center shadow-2xl backdrop-blur-xl">
        <div className="mb-2 inline-flex size-14 items-center justify-center rounded-full bg-emerald-500/20 text-emerald-400 text-2xl">
          ✓
        </div>
        <h1 className="mt-4 text-2xl font-bold text-white">You're signed in!</h1>
        <p className="mt-2 text-sm text-white/50">
          User ID: <span className="font-mono text-white/80">{session.userId}</span>
        </p>
        <form action={logout} className="mt-8">
          <button
            type="submit"
            className="w-full rounded-xl border border-white/10 bg-white/5 px-6 py-2.5 text-sm font-semibold text-white/70 transition-colors hover:bg-white/10 hover:text-white"
          >
            Sign out
          </button>
        </form>
      </div>
    </div>
  );
import { redirect } from "next/navigation";

export default function Home() {
  redirect("/learn");
}

