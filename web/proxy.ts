import { type NextRequest, NextResponse } from "next/server";

// Routes that require authentication — protect everything except these public paths
const PUBLIC_PATHS = new Set(["/login", "/api/auth/callback"]);

// Paths that auth users should be redirected away from (e.g., login)
const AUTH_ONLY_PATHS = new Set(["/login"]);

export function proxy(request: NextRequest) {
  const { pathname } = request.nextUrl;

  // Skip Next.js internals and static assets
  if (
    pathname.startsWith("/_next") ||
    pathname.startsWith("/favicon") ||
    pathname.match(/\.(png|jpg|jpeg|svg|ico|webp|css|js|woff2?)$/)
  ) {
    return NextResponse.next();
  }

  const sessionCookie = request.cookies.get("sliit_session")?.value;
  const isAuthenticated = !!sessionCookie; // middleware does a presence check; full verification happens in server components

  // Redirect unauthenticated users trying to access protected routes
  if (!isAuthenticated && !PUBLIC_PATHS.has(pathname)) {
    const url = request.nextUrl.clone();
    url.pathname = "/login";
    return NextResponse.redirect(url);
  }

  // Redirect authenticated users away from the login page
  if (isAuthenticated && AUTH_ONLY_PATHS.has(pathname)) {
    const url = request.nextUrl.clone();
    url.pathname = "/";
    return NextResponse.redirect(url);
  }

  return NextResponse.next();
}

export const config = {
  matcher: [
    // Match all paths except Next.js internals
    "/((?!_next/static|_next/image|favicon.ico).*)",
  ],
};
