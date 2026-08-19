import { HighlightStyle, syntaxHighlighting } from "@codemirror/language";
import type { Extension } from "@codemirror/state";
import { EditorView } from "@codemirror/view";
import { tags as t } from "@lezer/highlight";

export const ayuDarkEditorTheme = EditorView.theme(
  {
    "&": {
      color: "#ededed",
      backgroundColor: "#0d1017",
      height: "100%",
      fontSize: "15px",
      fontFamily: "var(--font-fira-code), monospace",
      fontWeight: "600",
    },
    ".cm-scroller": {
      fontFamily: "var(--font-fira-code), monospace",
      lineHeight: "1.6",
      padding: "8px 0",
      overflow: "auto",
    },
    ".cm-content": {
      caretColor: "#e6b450",
      padding: "0 12px",
    },
    ".cm-cursor, .cm-dropCursor": {
      borderLeftColor: "#e6b450",
      borderLeftWidth: "2px",
    },
    "&.cm-focused .cm-selectionBackground, .cm-selectionBackground, .cm-content ::selection":
      {
        backgroundColor: "#3388ff40 !important",
      },
    ".cm-activeLine": {
      backgroundColor: "#161a2480",
    },
    ".cm-gutters": {
      backgroundColor: "#0d1017",
      color: "#5a6378a6",
      borderRight: "1px solid rgba(255, 255, 255, 0.06)",
    },
    ".cm-activeLineGutter": {
      backgroundColor: "#161a2480",
      color: "#5a6378",
    },
    ".cm-lineNumbers .cm-gutterElement": {
      padding: "0 12px 0 16px",
      minWidth: "40px",
      textAlign: "right",
    },
    ".cm-foldPlaceholder": {
      backgroundColor: "#1b1f29",
      border: "none",
      color: "#5a6673",
    },
    ".cm-tooltip": {
      backgroundColor: "#141821",
      border: "1px solid #1b1f29",
      color: "#ededed",
    },
    ".cm-tooltip .cm-tooltip-arrow:before": {
      borderTopColor: "#1b1f29",
      borderBottomColor: "#1b1f29",
    },
    ".cm-tooltip .cm-tooltip-arrow:after": {
      borderTopColor: "#141821",
      borderBottomColor: "#141821",
    },
    ".cm-tooltip-autocomplete": {
      "& > ul > li[aria-selected]": {
        backgroundColor: "#47526640",
        color: "#ededed",
      },
    },
  },
  { dark: true },
);

// Syntax highlighting styles matching AyuDark
export const ayuDarkHighlightStyle = HighlightStyle.define([
  // Comments
  { tag: [t.comment, t.lineComment, t.blockComment, t.docComment], color: "#5a6673", fontStyle: "italic" },

  // Strings & Escapes
  { tag: [t.string, t.special(t.string)], color: "#aad94c" },
  { tag: [t.escape, t.regexp], color: "#95e6cb" },

  // Numbers & Constants
  { tag: [t.number, t.integer, t.float, t.bool], color: "#d2a6ff" },
  { tag: [t.constant(t.variableName), t.atom], color: "#d2a6ff" },

  // Keywords & Control
  { tag: [t.keyword, t.controlKeyword, t.definitionKeyword, t.moduleKeyword], color: "#ff8f40" },
  { tag: [t.operatorKeyword, t.modifier], color: "#ff8f40" },

  // Operators & Delimiters
  { tag: [t.operator, t.compareOperator, t.arithmeticOperator, t.logicOperator], color: "#f29668" },
  { tag: [t.punctuation, t.separator, t.bracket, t.paren, t.brace], color: "#bfbdb6" },

  // Variables
  { tag: [t.variableName, t.local(t.variableName)], color: "#bfbdb6" },
  { tag: [t.special(t.variableName), t.self], color: "#39bae6", fontStyle: "italic" },
  { tag: [t.propertyName, t.attributeName], color: "#f07178" },

  // Functions & Methods
  { tag: [t.function(t.variableName), t.function(t.propertyName)], color: "#ffb454" },
  { tag: [t.definition(t.function(t.variableName))], color: "#ffb454" },
  { tag: [t.standard(t.function(t.variableName))], color: "#f07178" },

  // Types & Structs
  { tag: [t.typeName, t.standard(t.typeName), t.className, t.namespace], color: "#59c2ff" },
  { tag: [t.definition(t.typeName)], color: "#59c2ff" },

  // Errors / Invalid
  { tag: [t.invalid], color: "#d95757" },
]);

export const ayuDarkTheme: Extension = [
  ayuDarkEditorTheme,
  syntaxHighlighting(ayuDarkHighlightStyle),
];
