import { NextResponse } from "next/server";

import { getLeaderboard } from "@/lib/leaderboard";
import { getSession } from "@/lib/session";

export async function GET() {
  const session = await getSession();
  if (!session) {
    return NextResponse.json({ error: "Unauthorized" }, { status: 401 });
  }

  const entries = await getLeaderboard();
  return NextResponse.json({ entries });
}
