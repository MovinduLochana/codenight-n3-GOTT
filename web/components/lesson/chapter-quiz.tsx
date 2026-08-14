"use client";

import { RotateCcwIcon } from "lucide-react";
import { useState } from "react";

import { Button } from "@/components/ui/button";
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
import type { Quiz } from "@/lib/content";

export function ChapterQuiz({ quiz }: { quiz: Quiz }) {
  const [attempt, setAttempt] = useState(0);
  const [score, setScore] = useState<number | null>(null);

  if (quiz.questions.length === 0) return null;

  const total = quiz.questions.length;

  function handleSubmit(event: React.SubmitEvent<HTMLFormElement>) {
    event.preventDefault();
    const formData = new FormData(event.currentTarget);
    const correct = quiz.questions.filter(
      (question) => formData.get(question.id) === question.correct,
    ).length;
    setScore(correct);
  }

  function retake() {
    setScore(null);
    setAttempt((value) => value + 1);
  }

  if (score !== null) {
    return (
      <div className="flex flex-col items-start gap-4 border border-border bg-card p-6">
        <div>
          <p className="font-heading text-2xl font-semibold">
            {score} / {total}
          </p>
          <p className="mt-1 text-sm text-muted-foreground">
            {score === total
              ? "Perfect score — you've got this chapter down."
              : "Review the lessons in this chapter and give it another shot."}
          </p>
        </div>
        <Button variant="outline" size="sm" onClick={retake}>
          <RotateCcwIcon />
          Retake
        </Button>
      </div>
    );
  }

  return (
    <div className="border border-border bg-card p-6">
      <Questionnaire key={attempt} onSubmit={handleSubmit}>
        <QuestionnaireProgress />

        {quiz.questions.map((question) => (
          <QuestionnaireItem key={question.id} name={question.id} required>
            <QuestionnaireTitle>{question.prompt}</QuestionnaireTitle>
            <QuestionnaireChoices>
              {question.choices.map((choice) => (
                <QuestionnaireChoice key={choice.id} value={choice.id}>
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
          <QuestionnaireSubmit />
        </QuestionnaireActions>
      </Questionnaire>
    </div>
  );
}
