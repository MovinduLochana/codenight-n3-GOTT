"use client";

import { go } from "@codemirror/lang-go";
import {
  CheckIcon,
  FileTextIcon,
  PlayIcon,
  RotateCcwIcon,
  XIcon,
} from "lucide-react";
import dynamic from "next/dynamic";
import { useState } from "react";

import { Button } from "@/components/ui/button";
import { ayuDarkTheme } from "@/lib/codemirror-theme";
import { cn } from "@/lib/utils";

const CodeEditor = dynamic(() => import("@uiw/react-codemirror"), {
  ssr: false,
  loading: () => (
    <div className="flex h-full items-center justify-center bg-background text-xs text-muted-foreground">
      Loading editor…
    </div>
  ),
});

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
    <div className="grid min-h-0 flex-1 gap-6 lg:grid-cols-[380px_1fr]">
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

        <div className="min-h-90 flex-1 overflow-hidden lg:min-h-0">
          <CodeEditor
            value={code}
            height="100%"
            className="h-full"
            extensions={[go()]}
            theme={ayuDarkTheme}
            onChange={(value) => setCode(value)}
            basicSetup={{
              lineNumbers: true,
              highlightActiveLineGutter: true,
              highlightActiveLine: true,
              foldGutter: false,
              dropCursor: false,
              allowMultipleSelections: false,
              indentOnInput: true,
              tabSize: 4,
            }}
          />
        </div>

        <div
          className={cn(
            "shrink-0 border-t border-border p-4",
            result === null
              ? "text-muted-foreground"
              : result.passed
                ? "bg-sky-500/10"
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
                  result.passed ? "text-sky-400" : "text-destructive",
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
