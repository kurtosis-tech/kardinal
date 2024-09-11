"use client";
import { ReactNode } from "react";

function textToKebabCase(text: string): string {
  // Convert to lowercase
  let kebabCase = text.toLowerCase();

  // Replace non-alphanumeric characters with hyphens
  kebabCase = kebabCase.replace(/[^a-z0-9]+/g, "-");

  // Remove leading and trailing hyphens
  kebabCase = kebabCase.replace(/^-+|-+$/g, "");

  return kebabCase;
}

const DocsHeading = ({
  as,
  children,
  ...restProps
}: {
  as: "h1" | "h2" | "h3" | "h4" | "h5" | "h6";
  children?: ReactNode;
}) => {
  const Tag = as;
  const id = textToKebabCase(children?.toString() || "");
  return (
    <Tag {...restProps} style={{ position: "relative" }}>
      <span
        style={{ position: "absolute", top: -80 }}
        id={id}
        role="presentation"
      ></span>
      <a href={`#${id}`} style={{ textDecoration: "none", color: "inherit" }}>
        {children}
      </a>
    </Tag>
  );
};

export default DocsHeading;
