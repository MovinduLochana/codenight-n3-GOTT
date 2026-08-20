import { neon, neonConfig } from "@neondatabase/serverless";
import { drizzle } from "drizzle-orm/neon-http";
import { Agent, fetch as undiciFetch, type RequestInfo, type RequestInit } from "undici";

// Some networks (mobile hotspots, WSL2, certain VPNs/campus wifi) advertise
// IPv6 routes to Neon's endpoint that don't actually work. Node's fetch races
// the IPv4 and IPv6 addresses together, so one dead IPv6 leg is enough to
// fail the whole request even though IPv4 alone connects fine. Force IPv4.
const ipv4Agent = new Agent({ connect: { family: 4 } });
neonConfig.fetchFunction = (url: RequestInfo, init?: RequestInit) =>
  undiciFetch(url, { ...init, dispatcher: ipv4Agent });

const dbUrl = process.env.DATABASE_URL ?? "";
const sql = neon(dbUrl);
export const db = drizzle({ client: sql });
