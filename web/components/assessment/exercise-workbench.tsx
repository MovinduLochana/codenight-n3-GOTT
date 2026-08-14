"use client";

import Editor from "@monaco-editor/react";
import {
  CheckIcon,
  FileTextIcon,
  PlayIcon,
  RotateCcwIcon,
  XIcon,
} from "lucide-react";
import { useState } from "react";

import { Button } from "@/components/ui/button";
import { cn } from "@/lib/utils";

type Result = { passed: boolean; output: string };

export function ExerciseWorkbench({
  exerciseId,
  taskHtml,
  starterCode,
  savedCode,
  initialResult,
}: {
  exerciseId: string;
  taskHtml: string;
  starterCode: string;
  savedCode: string | null;
  initialResult: Result | null;
}) {
  const [code, setCode] = useState(savedCode ?? starterCode);
  const [running, setRunning] = useState(false);
  const [result, setResult] = useState<Result | null>(initialResult);

  async function runTests() {
    setRunning(true);
    try {
      const response = await fetch(`/api/assessment/${exerciseId}`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ code }),
      });

      if (!response.ok) {
        const body = (await response.json().catch(() => null)) as {
          error?: string;
        } | null;
        setResult({
          passed: false,
          output: body?.error ?? "Something went wrong running your code.",
        });
        return;
      }

      const body = (await response.json()) as Result;
      setResult(body);
    } finally {
      setRunning(false);
    }
  }

  function reset() {
    setCode(starterCode);
    setResult(null);
  }

  return (
    <div className="grid min-h-0 flex-1 gap-6 lg:grid-cols-[300px_1fr]">
      <div className="min-w-0 lg:max-h-full lg:overflow-y-auto lg:pe-2">
        <div className="border border-border bg-card p-5">
          <p className="flex items-center gap-1.5 text-[0.625rem] font-semibold tracking-widest text-primary uppercase">
            <FileTextIcon className="size-3" />
            Task
          </p>
          <article
            className={cn(
              "prose prose-invert prose-sm mt-4 max-w-none",
              "prose-p:leading-relaxed prose-p:text-foreground",
              "prose-strong:text-foreground",
              "prose-blockquote:border-primary/40 prose-blockquote:font-normal prose-blockquote:text-foreground prose-blockquote:not-italic",
              "prose-code:rounded-none prose-code:text-primary prose-code:before:content-none prose-code:after:content-none",
              "prose-pre:border prose-pre:border-border prose-pre:bg-background",
            )}
            // biome-ignore lint/security/noDangerouslySetInnerHtml: our own markdown
            dangerouslySetInnerHTML={{ __html: taskHtml }}
          />
        </div>
      </div>

      <div className="flex min-h-0 min-w-0 flex-col border border-border">
        <div className="flex shrink-0 items-center justify-between gap-2 border-b border-border bg-card px-4 py-2">
          <span className="font-mono text-xs text-muted-foreground">
            main.go
          </span>
          <div className="flex items-center gap-2">
            <Button
              variant="outline"
              size="xs"
              onClick={reset}
              disabled={running}
            >
              <RotateCcwIcon />
              Reset
            </Button>
            <Button size="xs" onClick={runTests} disabled={running}>
              <PlayIcon />
              {running ? "Running…" : "Run Tests"}
            </Button>
          </div>
        </div>

        <div className="min-h-[360px] flex-1 lg:min-h-0">
          <Editor
            height="100%"
            language="go"
            theme="vs-dark"
            value={code}
            onChange={(value) => setCode(value ?? "")}
            options={{
              minimap: { enabled: false },
              fontSize: 13,
              scrollBeyondLastLine: false,
              padding: { top: 12 },
            }}
          />
        </div>

        <div
          className={cn(
            "shrink-0 border-t border-border p-4",
            result === null
              ? "text-muted-foreground"
              : result.passed
                ? "bg-primary/5"
                : "bg-destructive/5",
          )}
        >
          {result === null ? (
            <p className="text-xs text-muted-foreground">
              Run the tests to see your result here.
            </p>
          ) : (
            <>
              <p
                className={cn(
                  "flex items-center gap-1.5 text-xs font-semibold tracking-widest uppercase",
                  result.passed ? "text-primary" : "text-destructive",
                )}
              >
                {result.passed ? (
                  <CheckIcon className="size-3.5" />
                ) : (
                  <XIcon className="size-3.5" />
                )}
                {result.passed ? "Tests passed" : "Tests failed"}
              </p>
              <pre className="mt-3 max-h-40 overflow-auto font-mono text-xs whitespace-pre-wrap text-muted-foreground">
                {result.output}
              </pre>
            </>
          )}
        </div>
      </div>
    </div>
  );
}
