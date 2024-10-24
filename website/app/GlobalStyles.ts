"use client";

import { createGlobalStyle, css } from "styled-components";

const styles = css`
  :root {
    color-scheme: only light;

    --max-width: 1248px;
    --border-radius: 4px;

    /* --font-mono: ui-monospace, Menlo, Monaco, "Cascadia Mono", "Segoe UI Mono", */
    /*   "Roboto Mono", "Oxygen Mono", "Ubuntu Monospace", "Source Code Pro", */
    /*   "Fira Mono", "Droid Sans Mono", "Courier New", monospace; */

    /* color vars */
    --brand-primary: #ff602c;
    --brand-primary-dark: rgba(168, 50, 5, 0.4);
    --brand-secondary: #fca061;
    --gradient-from-bg: linear-gradient(84deg, transparent 0%, white 100%);
    --gradient-to-bg: linear-gradient(180deg, transparent 0%, white 100%);

    --gradient-brand: linear-gradient(
      to right,
      var(--brand-primary),
      var(--brand-secondary)
    );
    --gradient-brand-reverse: linear-gradient(
      to right,
      var(--brand-secondary),
      var(--brand-primary)
    );

    --gradient-brand-dark: linear-gradient(
      90deg,
      #fca061,
      rgba(168, 50, 5, 0.4)
    );
    --white-10: rgba(255, 255, 255, 0.1);
    --white-20: rgba(255, 255, 255, 0.2);
    --white-40: rgba(255, 255, 255, 0.4);
    --white-60: rgba(255, 255, 255, 0.6);
    --white-70: rgba(255, 255, 255, 0.7);
    --white-90: rgba(255, 255, 255, 0.9);
    --white-100: #ffffff;

    /* TODO: too many grays... consolidate */
    --gray-100: #f3f4f6;
    --gray-200: #e5e7eb;
    --gray-400: #9ca3af;
    --gray-600: #4b5563;
    --gray-bg-light: #212121;
    --gray-bg: #1f232a;
    --gray-dark: #21262d;
    --gray-diagramlines: #666a6f;
    --gray-icon: #8b949e;
    --gray-light: #aaa;
    --gray-lighter: #f2f2f2;
    --gray-lightest: #fafafa;
    --gray: #525157;

    --blue: #1f56a5;
    --blue-light: #58a6ff;

    --black: #000330;
    --white: #ffffff;

    /* semantic vars */
    --foreground: var(--gray);
    --foreground-dark: var(--black);
    --foreground-light: var(--gray-light);

    --foreground-inverted: var(--gray-light);
    --foreground-dark-inverted: var(--gray-lighter);

    --background: var(--white);
    --background-light: var(--gray-bg-light, #212121);
    --background-dark: var(--gray-bg-dark, #21262d);

    --background-contrast: #fff9f2;
    --background-inverted: var(--gray-bg);

    --gray-border: #ddd;
  }

  * {
    -moz-font-feature-settings: "liga" on;
    -moz-osx-font-smoothing: grayscale;
    -webkit-font-smoothing: antialiased;
    box-sizing: border-box;
    font-smooth: always;
    margin: 0;
    padding: 0;
    text-rendering: optimizeLegibility;
  }

  html,
  body {
    background: var(--background);
    overflow-x: hidden;
    scroll-behavior: smooth;
  }

  body {
    color: var(--foreground);
  }

  a {
    color: inherit;
    text-decoration: none;
  }

  em {
    font-weight: 500;
    background: var(--gradient-brand-reverse);
    background-size: auto;
    background-clip: border-box;
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
    font-style: normal;
  }

  s {
    text-decoration: none;
    position: relative;
    display: inline-block;
  }

  s:after {
    content: "";
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background-image: url("/scribble-animated.svg");
    background-size: 100%;
    background-repeat: no-repeat;
  }

  hr {
    border: 0;
    border-bottom: 1px solid var(--gray-border);
  }

  br[data-mobile="true"] {
    display: none;
  }

  @media (max-width: 768px) {
    br[data-desktop="true"] {
      display: none;
    }
    br[data-mobile="true"] {
      display: block;
    }
  }

  html {
    color-scheme: light;
  }

  pre {
    background-color: var(--gray-bg) !important;
    padding: 12px !important;
  }

  code {
    line-height: 1.6;
    font-size: 14px;
    font-style: normal;
    font-weight: 400;
  }

  @media (max-width: 768px) {
    code {
      font-size: 14px;
    }
  }
`;

const GlobalStyles = createGlobalStyle`${styles}`;

export default GlobalStyles;
