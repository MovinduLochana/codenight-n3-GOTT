import { redirect } from "next/navigation";
import { getSession } from "@/lib/session";

export const instant = false;

export default async function Home() {
  const session = await getSession();

  redirect(session ? "/learn" : "/login");
}
