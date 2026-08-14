"use client";

import {
  BookOpenIcon,
  CodeIcon,
  LayoutGridIcon,
  ListChecksIcon,
  LogOutIcon,
  UsersIcon,
} from "lucide-react";
import Link from "next/link";
import { usePathname } from "next/navigation";

import { logout } from "@/app/actions/auth";
import type { Category } from "@/lib/content";
import { cn } from "@/lib/utils";

const navItemClass =
  "flex w-full items-center gap-2.5 px-3 py-2 text-xs font-semibold tracking-widest uppercase transition-colors";

const comingSoon = [
  { label: "Editor", icon: CodeIcon },
  { label: "Community", icon: UsersIcon },
];

export function ChapterSidebar({ categories }: { categories: Category[] }) {
  const pathname = usePathname();
  const activeCategory = pathname.split("/")[2];

  return (
    <aside className="hidden w-64 shrink-0 flex-col border-e border-border bg-sidebar md:flex">
      <nav className="flex flex-col gap-0.5 p-2">
        <button
          type="button"
          className={cn(
            navItemClass,
            "text-muted-foreground hover:bg-sidebar-accent hover:text-sidebar-accent-foreground",
          )}
        >
          <LayoutGridIcon className="size-3.5" />
          Overview
        </button>

        <Link
          href="/learn"
          className={cn(
            navItemClass,
            pathname.startsWith("/learn")
              ? "bg-primary text-primary-foreground"
              : "text-muted-foreground hover:bg-sidebar-accent hover:text-sidebar-accent-foreground",
          )}
        >
          <BookOpenIcon className="size-3.5" />
          Lessons
        </Link>

        {comingSoon.map((item) => (
          <button
            key={item.label}
            type="button"
            className={cn(
              navItemClass,
              "text-muted-foreground hover:bg-sidebar-accent hover:text-sidebar-accent-foreground",
            )}
          >
            <item.icon className="size-3.5" />
            {item.label}
          </button>
        ))}
      </nav>

      <div className="min-h-0 flex-1 overflow-y-auto border-t border-border p-2">
        <p className="px-3 py-2 text-[0.625rem] font-semibold tracking-widest text-muted-foreground uppercase">
          Chapters
        </p>

        {categories.map((category) => {
          const open = activeCategory === category.id;

          return (
            <div key={category.id} className="mb-1">
              <Link
                href={`/learn/${category.id}/${category.topics[0]?.id}`}
                className={cn(
                  "flex items-baseline gap-2 px-3 py-1.5 text-xs transition-colors",
                  open
                    ? "text-foreground"
                    : "text-muted-foreground hover:text-foreground",
                )}
              >
                <span className="font-mono text-[0.625rem] text-primary">
                  {String(category.number).padStart(2, "0")}
                </span>
                <span className="truncate font-medium">{category.title}</span>
              </Link>

              {open ? (
                <ul className="ms-[1.375rem] border-s border-border">
                  {category.topics.map((topic) => {
                    const href = `/learn/${category.id}/${topic.id}`;

                    return (
                      <li key={topic.id}>
                        <Link
                          href={href}
                          className={cn(
                            "-ms-px flex border-s-2 py-1.5 ps-3 pe-2 text-xs transition-colors",
                            pathname === href
                              ? "border-primary text-foreground"
                              : "border-transparent text-muted-foreground hover:border-border hover:text-foreground",
                          )}
                        >
                          <span className="truncate">{topic.title}</span>
                        </Link>
                      </li>
                    );
                  })}

                  <li>
                    <Link
                      href={`/learn/${category.id}/quiz`}
                      className={cn(
                        "-ms-px flex items-center gap-1.5 border-s-2 py-1.5 ps-3 pe-2 text-xs transition-colors",
                        pathname === `/learn/${category.id}/quiz`
                          ? "border-primary text-foreground"
                          : "border-transparent text-muted-foreground hover:border-border hover:text-foreground",
                      )}
                    >
                      <ListChecksIcon className="size-3" />
                      <span className="truncate">Chapter Quiz</span>
                    </Link>
                  </li>
                </ul>
              ) : null}
            </div>
          );
        })}
      </div>

      <div className="border-t border-border p-2">
        <form action={logout}>
          <button
            type="submit"
            className={cn(
              navItemClass,
              "justify-center bg-red-600 text-white transition-colors hover:bg-red-700",
            )}
          >
            <LogOutIcon className="size-3.5" />
            Logout
          </button>
        </form>
      </div>
    </aside>
  );
}
