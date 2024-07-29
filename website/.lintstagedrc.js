module.exports = {
  // Type check TypeScript files

  "**/*.(ts|tsx)": () => "bun tsc --noEmit",

  // Lint & Prettify TS and JS files

  "**/*.(ts|tsx|js)": (filenames) => [
    `bun eslint ${filenames.join(" ")}`,

    `bun prettier --write ${filenames.join(" ")}`,
  ],

  // Prettify only Markdown and JSON files

  "**/*.(md|json)": (filenames) =>
    `bun prettier --write ${filenames.join(" ")}`,
};
