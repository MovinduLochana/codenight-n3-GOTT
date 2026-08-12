import { marked } from "marked";
import { cacheLife } from "next/cache";
import { createHighlighter } from "shiki";

const theme = "github-dark";

let configured: Promise<void> | null = null;

function configure(): Promise<void> {
  if (configured === null) {
    configured = createHighlighter({ themes: [theme], langs: ["go"] }).then(
      (highlighter) => {
        marked.use({
          renderer: {
            code({ text, lang }) {
              const language = lang === "go" ? "go" : "text";
              return highlighter.codeToHtml(text, { lang: language, theme });
            },
          },
        });
      },
    );
  }

  return configured;
}

export async function renderMarkdown(markdown: string): Promise<string> {
  "use cache";
  cacheLife("max");

  await configure();
  return marked.parse(markdown, { async: false });
}
