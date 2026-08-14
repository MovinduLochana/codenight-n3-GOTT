"use client";

import { useEffect, useState } from "react";

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
import type { PublicQuiz } from "@/lib/content";

type Result = { passed: boolean; score: number; total: number };

export function ChapterQuiz({
  categoryId,
  quiz,
  completedResult,
}: {
  categoryId: string;
  quiz: PublicQuiz;
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

  // Already completed elsewhere (e.g. another tab) — nothing left to resume.
  useEffect(() => {
    if (completedResult) window.sessionStorage.removeItem(storageKey);
  }, [completedResult, storageKey]);

  if (quiz.questions.length === 0) return null;

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
            ? "Perfect score — you've got this chapter down."
            : "You've already completed this quiz. Review the lessons above whenever you're ready."}
        </p>
      </div>
    );
  }

  return (
    <div className="border border-border bg-card p-6">
      <Questionnaire onSubmit={handleSubmit} onChange={handleChange}>
        <QuestionnaireProgress />

        {quiz.questions.map((question) => (
          <QuestionnaireItem key={question.id} name={question.id} required>
            <QuestionnaireTitle>{question.prompt}</QuestionnaireTitle>
            <QuestionnaireChoices>
              {question.choices.map((choice) => (
                <QuestionnaireChoice
                  key={choice.id}
                  value={choice.id}
                  defaultChecked={savedAnswers[question.id] === choice.id}
                >
                  {choice.text}
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
            { submitting && <Spinner className="size-4" /> }
            Submit
          </QuestionnaireSubmit>
        </QuestionnaireActions>
      </Questionnaire>

      {error ? <p className="mt-4 text-sm text-destructive">{error}</p> : null}
    </div>
  );
}
