export const AyuDark = {
  base: "vs-dark",
  inherit: true,
  rules: [
    // Default
    { token: "", foreground: "ededed", background: "0d1017" },

    // Comments
    { token: "comment", foreground: "5a6673", fontStyle: "italic" },
    { token: "comment.doc", foreground: "5a6673", fontStyle: "italic" },

    // Strings & regex
    { token: "string", foreground: "aad94c" },
    { token: "string.escape", foreground: "95e6cb" },
    { token: "regexp", foreground: "95e6cb" },

    // Numbers & constants
    { token: "number", foreground: "d2a6ff" },
    { token: "number.hex", foreground: "d2a6ff" },
    { token: "number.float", foreground: "d2a6ff" },
    { token: "constant", foreground: "d2a6ff" },
    { token: "constant.language", foreground: "d2a6ff" },
    { token: "constant.numeric", foreground: "d2a6ff" },

    // Keywords & storage
    { token: "keyword", foreground: "ff8f40" },
    { token: "keyword.control", foreground: "ff8f40" },
    { token: "keyword.operator", foreground: "f29668" },
    { token: "storage", foreground: "ff8f40" },
    { token: "storage.type", foreground: "ff8f40" },
    { token: "storage.modifier", foreground: "ff8f40" },

    // Operators & punctuation
    { token: "operator", foreground: "f29668" },
    { token: "delimiter", foreground: "bfbdb6" },
    { token: "delimiter.bracket", foreground: "bfbdb6" },
    { token: "delimiter.parenthesis", foreground: "bfbdb6" },
    { token: "delimiter.square", foreground: "bfbdb6" },
    { token: "delimiter.angle", foreground: "bfbdb6" },
    { token: "punctuation", foreground: "bfbdb6" },

    // Variables
    { token: "variable", foreground: "bfbdb6" },
    { token: "variable.name", foreground: "bfbdb6" },
    { token: "variable.parameter", foreground: "d2a6ff" },
    { token: "variable.language", foreground: "39bae6", fontStyle: "italic" },
    { token: "variable.other", foreground: "bfbdb6" },
    { token: "variable.member", foreground: "f07178" },

    // Functions
    { token: "function", foreground: "ffb454" },
    { token: "function.name", foreground: "ffb454" },
    { token: "support.function", foreground: "f07178" },
    { token: "meta.function-call", foreground: "ffb454" },

    // Types / classes
    { token: "type", foreground: "59c2ff" },
    { token: "type.identifier", foreground: "59c2ff" },
    { token: "class", foreground: "59c2ff" },
    { token: "class.name", foreground: "59c2ff" },
    { token: "interface", foreground: "39bae6" },
    { token: "namespace", foreground: "59c2ff" },
    { token: "struct", foreground: "59c2ff" },
    { token: "enum", foreground: "59c2ff" },
    { token: "type.primitive", foreground: "39bae6" },

    // Tags (HTML/JSX)
    { token: "tag", foreground: "39bae6" },
    { token: "tag.id", foreground: "59c2ff" },
    { token: "tag.class", foreground: "59c2ff" },
    { token: "attribute.name", foreground: "ffb454" },
    { token: "attribute.value", foreground: "aad94c" },
    { token: "metatag", foreground: "39bae6" },
    { token: "metatag.content", foreground: "bfbdb6" },

    // Support / library
    { token: "support", foreground: "59c2ff" },
    { token: "support.class", foreground: "59c2ff" },
    { token: "support.type", foreground: "59c2ff" },
    { token: "support.constant", foreground: "d2a6ff" },
    { token: "support.variable", foreground: "f07178" },

    // Invalid / errors
    { token: "invalid", foreground: "d95757" },
    { token: "invalid.illegal", foreground: "d95757" },

    // Markup / Markdown extras
    { token: "markup.heading", foreground: "ffb454", fontStyle: "bold" },
    { token: "markup.bold", fontStyle: "bold" },
    { token: "markup.italic", fontStyle: "italic" },
    { token: "markup.underline", fontStyle: "underline" },
    { token: "markup.quote", foreground: "95e6cb", fontStyle: "italic" },
    { token: "markup.inline.raw", foreground: "f29668" },
    { token: "markup.inserted", foreground: "70bf56" },
    { token: "markup.deleted", foreground: "f26d78" },
    { token: "markup.changed", foreground: "73b8ff" },
  ],
  colors: {
    // Editor core
    "editor.background": "#0d1017",
    "editor.foreground": "#ededed",
    "editorCursor.foreground": "#e6b450",
    "editor.lineHighlightBackground": "#161a24",
    "editor.selectionBackground": "#3388ff40",
    "editor.inactiveSelectionBackground": "#80b5ff26",
    "editor.selectionHighlightBackground": "#70bf5626",
    "editor.findMatchBackground": "#4c4126",
    "editor.findMatchHighlightBackground": "#4c412680",
    "editor.wordHighlightBackground": "#73b8ff14",
    "editor.wordHighlightStrongBackground": "#70bf5614",
    "editor.rangeHighlightBackground": "#4c412633",

    // Line numbers & guides
    "editorLineNumber.foreground": "#5a6378a6",
    "editorLineNumber.activeForeground": "#5a6378",
    "editorIndentGuide.background": "#5a637842",
    "editorIndentGuide.activeBackground": "#5a6378a1",
    "editorWhitespace.foreground": "#5a6378a6",
    "editorRuler.foreground": "#5a637842",

    // Gutter
    "editorGutter.background": "#0d1017",
    "editorGutter.addedBackground": "#70bf56",
    "editorGutter.modifiedBackground": "#73b8ff",
    "editorGutter.deletedBackground": "#f26d78",

    // Bracket matching
    "editorBracketMatch.background": "#5a63784d",
    "editorBracketMatch.border": "#5a63784d",

    // Widgets
    "editorWidget.background": "#141821",
    "editorWidget.border": "#1b1f29",
    "editorHoverWidget.background": "#141821",
    "editorHoverWidget.border": "#1b1f29",
    "editorSuggestWidget.background": "#141821",
    "editorSuggestWidget.border": "#1b1f29",
    "editorSuggestWidget.selectedBackground": "#47526640",
    "editorSuggestWidget.highlightForeground": "#e6b450",

    // Scrollbar
    "scrollbarSlider.background": "#5a637866",
    "scrollbarSlider.hoverBackground": "#5a637899",
    "scrollbarSlider.activeBackground": "#5a6378b3",

    // Diff
    "diffEditor.insertedTextBackground": "#70bf561f",
    "diffEditor.removedTextBackground": "#f26d781f",
    "diffEditor.diagonalFill": "#1b1f29",

    // Overview ruler
    "editorOverviewRuler.border": "#1b1f29",
    "editorOverviewRuler.errorForeground": "#d95757",
    "editorOverviewRuler.warningForeground": "#e6b450",
    "editorOverviewRuler.addedForeground": "#70bf56",
    "editorOverviewRuler.modifiedForeground": "#73b8ff",
    "editorOverviewRuler.deletedForeground": "#f26d78",
    "editorOverviewRuler.findMatchForeground": "#4c4126",

    // Errors / warnings
    "editorError.foreground": "#d95757",
    "editorWarning.foreground": "#e6b450",
    "editorInfo.foreground": "#39bae6",

    // Inlay hints
    "editorInlayHint.foreground": "#bfbdb680",
    "editorInlayHint.background": "#00000000",
  },

  // old values
  //   {
  //   base: "vs-dark",
  //   inherit: true,
  //   rules: [
  //   { token: 'comment', foreground: '5c6773', fontStyle: 'italic' },
  //   { token: 'keyword', foreground: 'ff7733' },
  //   { token: 'number', foreground: 'f29718' },
  //   { token: 'string', foreground: 'bae67e' },
  //   { token: 'variable', foreground: 'f07178' },
  //   { token: 'type', foreground: 'ffa759' },
  //   { token: 'function', foreground: 'ffd580' },
  // ],
  //   colors: {
  //     "editor.background": "#0c0e19",
  //     "editor.foreground": "#ededed",
  //     "editorCursor.foreground": "#40fd51",
  //     "editor.lineHighlightBackground": "#11141f",
  //     "editor.selectionBackground": "#40fd5133",
  //     "editorLineNumber.foreground": "#4b5468",
  //     "editorLineNumber.activeForeground": "#40fd51",
  //     "editorIndentGuide.background": "#1a1e2e",
  //     "editorIndentGuide.activeBackground": "#2a2f45",
  //     "editorWidget.background": "#0c0e19",
  //     "editorWidget.border": "#1a1e2e",
  //     "editorSuggestWidget.background": "#0c0e19",
  //     "editorSuggestWidget.border": "#1a1e2e",
  //     "editorSuggestWidget.selectedBackground": "#40fd511a",
  //     "scrollbarSlider.background": "#40fd5122",
  //     "scrollbarSlider.hoverBackground": "#40fd5140",
  //   },
  // }
};
