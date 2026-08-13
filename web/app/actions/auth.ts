"use server";

import { redirect } from "next/navigation";
import { cookies } from "next/headers";
import { z } from "zod";
import { createSession, deleteSession, getSession } from "@/lib/session";

const AUTH_API_BASE = process.env.AUTH_API_BASE ?? "https://accounts.sliitmozilla.org/api";
const APP_URL = process.env.NEXT_PUBLIC_APP_URL ?? "http://localhost:3000";

const LoginSchema = z.object({
  email: z.email({ error: "Enter a valid email address." }).trim(),
  password: z.string().min(1, { error: "Password is required." }),
});

export type AuthFormState =
  | { errors?: { email?: string[]; password?: string[] }; message?: string }
  | undefined;

export async function loginWithCredentials(
  state: AuthFormState,
  formData: FormData
): Promise<AuthFormState> {
  const validated = LoginSchema.safeParse({
    email: formData.get("email"),
    password: formData.get("password"),
  });

  if (!validated.success) {
    return { errors: validated.error.flatten().fieldErrors };
  }

  const { email, password } = validated.data;

  // Call the auth server's POST /login endpoint
  let accessToken: string;
  let refreshTokenFromServer: string | undefined;

  try {
    const res = await fetch(`${AUTH_API_BASE}/login`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ email, password }),
    });

    if (res.status === 401) {
      return { message: "Invalid email or password. Please try again." };
    }
    if (res.status === 400) {
      return { message: "Invalid request. Please check your input." };
    }
    if (!res.ok) {
      return { message: "Something went wrong. Please try again later." };
    }

    const body = await res.json() as { data?: { token?: string } };
    const token = body?.data?.token;
    if (!token) {
      return { message: "Unexpected response from auth server." };
    }

    accessToken = token;

    const setCookie = res.headers.get("set-cookie");
    if (setCookie) {
      const match = setCookie.match(/refreshToken=([^;]+)/);
      if (match) refreshTokenFromServer = match[1];
    }
  } catch {
    return { message: "Unable to reach the auth server. Please try again later." };
  }

  // Fetch the current user's ID from /session using the access token
  let userId: string;
  try {
    const sessionRes = await fetch(`${AUTH_API_BASE}/session`, {
      headers: { Authorization: `Bearer ${accessToken}` },
    });
    if (!sessionRes.ok) {
      return { message: "Failed to retrieve user session from auth server." };
    }
    const sessionBody = await sessionRes.json() as { data?: { id?: string } };
    const id = sessionBody?.data?.id;
    if (!id) {
      return { message: "User ID not found in auth server response." };
    }
    userId = id;
  } catch {
    return { message: "Unable to fetch session info. Please try again." };
  }

  await createSession(userId, accessToken, refreshTokenFromServer);

  redirect("/");
}
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
        ...(session?.accessToken ? { Authorization: `Bearer ${session.accessToken}` } : {}),
        ...(refreshToken ? { Cookie: `refreshToken=${refreshToken}` } : {}),
      },
    });
  } catch (error) {
    console.error("Error logout request:", error);
  }
  await deleteSession();
  redirect("/login");
}
