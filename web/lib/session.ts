import "server-only";
import { SignJWT, jwtVerify } from "jose";
import { cookies } from "next/headers";
import { db } from "@/db/drizzle";
import { sessions } from "@/db/schema";
import { eq, and, isNull } from "drizzle-orm";

const SESSION_COOKIE = "sliit_session";
const secretKey = process.env.SESSION_SECRET;

function getEncodedKey() {
  if (!secretKey) throw new Error("SESSION_SECRET env var is not set");
  return new TextEncoder().encode(secretKey);
}

function nowSriLanka(): Date {
  const offsetMs = 5.5 * 60 * 60 * 1000;
  return new Date(Date.now() + offsetMs);
}

async function encryptSessionId(sessionId: string): Promise<string> {
  const expiresAt = new Date(nowSriLanka().getTime() + 7 * 24 * 60 * 60 * 1000);
  return new SignJWT({ sessionId })
    .setProtectedHeader({ alg: "HS256" })
    .setIssuedAt()
    .setExpirationTime(expiresAt)
    .sign(getEncodedKey());
}

async function decryptSessionId(token: string): Promise<string | null> {
  try {
    const { payload } = await jwtVerify(token, getEncodedKey(), {
      algorithms: ["HS256"],
    });
    return (payload as { sessionId: string }).sessionId ?? null;
  } catch {
    return null;
  }
}
export async function createSession(
  userId: string,
  accessToken: string,
  refreshToken?: string
) {
  const now = nowSriLanka();
  const expiresAt = new Date(now.getTime() + 7 * 24 * 60 * 60 * 1000);

  const [existing] = await db
    .select({ id: sessions.id })
    .from(sessions)
    .where(eq(sessions.userId, userId))
    .limit(1);

  let sessionId: string;

  if (existing) {
    sessionId = existing.id;
    await db
      .update(sessions)
      .set({
        accessToken,
        refreshToken: refreshToken ?? null,
        expiresAt,
        lastLoginAt: now,
        loggedOutAt: null, 
      })
      .where(eq(sessions.userId, userId));
  } else {
    sessionId = crypto.randomUUID();
    await db.insert(sessions).values({
      id: sessionId,
      userId,
      accessToken,
      refreshToken: refreshToken ?? null,
      expiresAt,
      lastLoginAt: now,
    });
  }

  // Set encrypted session cookie (holds only the session row id)
  const encrypted = await encryptSessionId(sessionId);
  const cookieStore = await cookies();
  cookieStore.set(SESSION_COOKIE, encrypted, {
    httpOnly: true,
    secure: process.env.NODE_ENV === "production",
    sameSite: "lax",
    path: "/",
    expires: expiresAt,
  });
}

export async function getSession() {
  const cookieStore = await cookies();
  const cookie = cookieStore.get(SESSION_COOKIE)?.value;
  if (!cookie) return null;

  const sessionId = await decryptSessionId(cookie);
  if (!sessionId) return null;

  const [session] = await db
    .select()
    .from(sessions)
    .where(
      and(
        eq(sessions.id, sessionId),
        isNull(sessions.loggedOutAt) 
      )
    )
    .limit(1);

  if (!session) return null;
  if (session.expiresAt < nowSriLanka()) {
    await db
      .update(sessions)
      .set({ loggedOutAt: nowSriLanka() })
      .where(eq(sessions.id, sessionId));
    return null;
  }

  return session;
}

export async function deleteSession() {
  const cookieStore = await cookies();
  const cookie = cookieStore.get(SESSION_COOKIE)?.value;

  if (cookie) {
    const sessionId = await decryptSessionId(cookie);
    if (sessionId) {
      await db
        .update(sessions)
        .set({ loggedOutAt: nowSriLanka() })
        .where(eq(sessions.id, sessionId));
    }
  }

  cookieStore.delete(SESSION_COOKIE);
}
