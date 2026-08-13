"use server";

import { cookies } from "next/headers";
import { redirect } from "next/navigation";

import { deleteSession, getSession } from "@/lib/session";

const AUTH_API_BASE = process.env.AUTH_API_BASE ?? "https://accounts.sliitmozilla.org/api";
const APP_URL = process.env.NEXT_PUBLIC_APP_URL ?? "http://localhost:3000";

export async function initiateOAuthFlow(): Promise<never> {
  const callbackUrl = `${APP_URL}/api/auth/callback`;
  const authorizeUrl = `${AUTH_API_BASE}/authorize?redirect=${encodeURIComponent(callbackUrl)}`;
  redirect(authorizeUrl);
}

export async function logout(): Promise<never> {
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
    console.error("Error logout request:", error);
  }
  await deleteSession();
  redirect("/login");
}
