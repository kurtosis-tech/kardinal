import { Code } from "bright";
import { IBM_Plex_Mono } from "next/font/google";
import { ComponentProps, ReactElement } from "react";

import CodeCopyButton from "@/components/CodeCopyButton";

const fontMono = IBM_Plex_Mono({
  subsets: ["latin"],
  display: "swap",
  variable: "--font-mono",
  weight: "400",
});

interface Props extends ComponentProps<typeof Code> {
  children: ReactElement;
}

const CodeBlock = (props: Props) => {
  return (
    <div style={{ position: "relative" }}>
      <CodeCopyButton text={props.children.props.children} />
      <Code
        {...props}
        style={{
          maxWidth: "calc(100vw - 32px)",
          margin: 0,
          borderRadius: "8px",
          padding: 10, // 24px - 1em = 10px
          backgroundColor: "var(--gray-bg)",
        }}
        theme={"github-dark"}
        className={fontMono.className}
      />
    </div>
  );
};

export default CodeBlock;
