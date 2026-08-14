"use client";

import { Fragment, useEffect, useState } from "react";

import {
  Questionnaire,
  QuestionnaireActions,
  QuestionnaireChoice,
  QuestionnaireChoices,
  QuestionnaireError,
  QuestionnaireItem,
  QuestionnaireNext,
  QuestionnairePrevious,
  QuestionnaireProgress,
  QuestionnaireSubmit,
  QuestionnaireTitle,
} from "@/components/ui/questionnaire";
import { Spinner } from "@/components/ui/spinner";
import type { PublicQuizQuestion } from "@/lib/content";
import { cn } from "@/lib/utils";

export type RenderedQuizQuestion = PublicQuizQuestion & {
  codeHtml: string | null;
};
export type RenderedQuiz = { questions: RenderedQuizQuestion[] };

type Result = { passed: boolean; score: number; total: number };

function renderInlineCode(text: string): React.ReactNode {
  const parts = text.split(/(`[^`]+`)/g);
  if (parts.length === 1) return text;

  return parts.map((part, index) => {
    const key = `${index}:${part}`;
    return part.startsWith("`") && part.endsWith("`") ? (
      <code
        key={key}
        className="rounded-none bg-muted px-1 py-0.5 font-mono text-[0.85em] text-foreground"
      >
        {part.slice(1, -1)}
      </code>
    ) : (
      <Fragment key={key}>{part}</Fragment>
    );
  });
}

export function ChapterQuiz({
  categoryId,
  quiz,
  completedResult,
}: {
  categoryId: string;
  quiz: RenderedQuiz;
  completedResult: Result | null;
}) {
  const storageKey = `quiz-answers:${categoryId}`;

  const [submitting, setSubmitting] = useState(false);
  const [result, setResult] = useState<Result | null>(completedResult);
  const [error, setError] = useState<string | null>(null);
  const [savedAnswers] = useState<Record<string, string>>(() => {
    if (typeof window === "undefined") return {};
    try {
      return JSON.parse(window.sessionStorage.getItem(storageKey) ?? "{}");
    } catch {
      return {};
    }
  });

  useEffect(() => {
    if (completedResult) window.sessionStorage.removeItem(storageKey);
  }, [completedResult, storageKey]);

  if (quiz.questions.length === 0) return null;

  const resumeQuestionId =
    quiz.questions.find((question) => !(question.id in savedAnswers))?.id ??
    quiz.questions[quiz.questions.length - 1]?.id;

  function handleChange(event: React.ChangeEvent<HTMLFormElement>) {
    const formData = new FormData(event.currentTarget);
    const answers: Record<string, string> = {};
    for (const question of quiz.questions) {
      const answer = formData.get(question.id);
      if (typeof answer === "string") answers[question.id] = answer;
    }
    window.sessionStorage.setItem(storageKey, JSON.stringify(answers));
  }

  async function handleSubmit(event: React.SubmitEvent<HTMLFormElement>) {
    event.preventDefault();
    setSubmitting(true);
    setError(null);

    try {
      const formData = new FormData(event.currentTarget);
      const answers: Record<string, string> = {};
      for (const question of quiz.questions) {
        const answer = formData.get(question.id);
        if (typeof answer === "string") answers[question.id] = answer;
      }

      const response = await fetch(`/api/quiz/${categoryId}`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ answers }),
      });

      const body = (await response.json().catch(() => null)) as
        | (Result & { error?: string })
        | { error: string }
        | null;

      if (!response.ok) {
        setError(
          body?.error ?? "Something went wrong submitting your answers.",
        );
        return;
      }

      window.sessionStorage.removeItem(storageKey);
      setResult(body as Result);
    } finally {
      setSubmitting(false);
    }
  }

  if (result !== null) {
    return (
      <div className="flex flex-col items-start gap-1 border border-border bg-card p-6">
        <p className="font-heading text-2xl font-semibold">
          {result.score} / {result.total}
        </p>
        <p className="text-sm text-muted-foreground">
          {result.passed
            ? "Perfect score. You've got this chapter down."
            : "You've already completed this quiz. Review the lessons above whenever you're ready."}
        </p>
      </div>
    );
  }

  return (
    <div className="border border-border bg-card p-6">
      <Questionnaire
        onSubmit={handleSubmit}
        onChange={handleChange}
        defaultItem={resumeQuestionId}
      >
        <QuestionnaireProgress />

        {quiz.questions.map((question) => (
          <QuestionnaireItem key={question.id} name={question.id} required>
            <QuestionnaireTitle>
              {renderInlineCode(question.prompt)}
            </QuestionnaireTitle>

            {question.codeHtml ? (
              <div
                className={cn(
                  "prose prose-invert prose-sm max-w-none",
                  "prose-pre:my-0 prose-pre:border prose-pre:border-border prose-pre:bg-background",
                )}
                // biome-ignore lint/security/noDangerouslySetInnerHtml: our own rendered markdown
                dangerouslySetInnerHTML={{ __html: question.codeHtml }}
              />
            ) : null}

            <QuestionnaireChoices>
              {question.choices.map((choice) => (
                <QuestionnaireChoice
                  key={choice.id}
                  value={choice.id}
                  defaultChecked={savedAnswers[question.id] === choice.id}
                >
                  {renderInlineCode(choice.text)}
                </QuestionnaireChoice>
              ))}
            </QuestionnaireChoices>
            <QuestionnaireError />
          </QuestionnaireItem>
        ))}

        <QuestionnaireActions>
          <QuestionnairePrevious />
          <QuestionnaireNext />

          <QuestionnaireSubmit disabled={submitting}>
            {submitting && <Spinner className="size-4" />}
            Submit
          </QuestionnaireSubmit>
        </QuestionnaireActions>
      </Questionnaire>

      {error ? <p className="mt-4 text-sm text-destructive">{error}</p> : null}
    </div>
  );
}
