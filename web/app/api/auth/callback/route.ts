import { cookies } from "next/headers";
import { type NextRequest, NextResponse } from "next/server";
import { createSession, deleteSession, getSession } from "@/lib/session";

const AUTH_API_BASE =
  process.env.AUTH_API_BASE ?? "https://accounts.sliitmozilla.org/api";
const APP_URL = process.env.NEXT_PUBLIC_APP_URL ?? "http://localhost:3000";

export async function GET(request: NextRequest) {
  const { searchParams } = new URL(request.url);
  const code = searchParams.get("code");

  if (!code) {
    return NextResponse.redirect(`${APP_URL}/login?error=missing_code`);
  }

  // Exchange code for tokens
  let accessToken: string;
  let refreshTokenFromServer: string | undefined;

  try {
    const tokenRes = await fetch(
      `${AUTH_API_BASE}/token?code=${encodeURIComponent(code)}`,
      { method: "POST" },
    );

    if (tokenRes.status === 401) {
      return NextResponse.redirect(`${APP_URL}/login?error=expired_code`);
    }
    if (!tokenRes.ok) {
      return NextResponse.redirect(
        `${APP_URL}/login?error=token_exchange_failed`,
      );
    }

    const tokenBody = (await tokenRes.json()) as { data?: { token?: string } };
    const token = tokenBody?.data?.token;
    if (!token) {
      return NextResponse.redirect(`${APP_URL}/login?error=no_token`);
    }

    accessToken = token;

    const setCookie = tokenRes.headers.get("set-cookie");
    if (setCookie) {
      const match = setCookie.match(/refreshToken=([^;]+)/);
      if (match) refreshTokenFromServer = match[1];
    }
  } catch {
    return NextResponse.redirect(`${APP_URL}/login?error=network_error`);
  }

  let userId: string;
  let displayName: string | undefined;
  try {
    const [sessionRes, meRes] = await Promise.all([
      fetch(`${AUTH_API_BASE}/session`, {
        headers: { Authorization: `Bearer ${accessToken}` },
      }),
      fetch(`${AUTH_API_BASE}/users/me`, {
        headers: { Authorization: `Bearer ${accessToken}` },
      }),
    ]);

    if (!sessionRes.ok) {
      return NextResponse.redirect(
        `${APP_URL}/login?error=session_fetch_failed`,
      );
    }

    const sessionBody = (await sessionRes.json()) as {
      data?: { id?: string; roles?: string[] };
    };
    const id = sessionBody?.data?.id;
    if (!id) {
      return NextResponse.redirect(`${APP_URL}/login?error=no_user_id`);
    }

    userId = id;

    if (meRes.ok) {
      const meBody = (await meRes.json()) as {
        data?: { name?: string; email?: string };
      };
      displayName = meBody.data?.name ?? meBody.data?.email ?? undefined;
    }
  } catch {
    return NextResponse.redirect(`${APP_URL}/login?error=network_error`);
  }

  await createSession(userId, accessToken, refreshTokenFromServer, displayName);

  return NextResponse.redirect(`${APP_URL}/learn`);
}

export async function POST() {
  const cookieStore = await cookies();
  const refreshToken = cookieStore.get("refreshToken")?.value;
  const session = await getSession();

  try {
    await fetch(`${AUTH_API_BASE}/logout`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        ...(session?.accessToken
          ? { Authorization: `Bearer ${session.accessToken}` }
          : {}),
        ...(refreshToken ? { Cookie: `refreshToken=${refreshToken}` } : {}),
      },
    });
  } catch (error) {
    console.error("Error during logout request:", error);
  }

  await deleteSession();
  const response = NextResponse.redirect(`${APP_URL}/login`);
  response.cookies.delete("refreshToken");
  return response;
}
