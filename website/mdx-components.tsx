import { Code } from "bright";
import type { MDXComponents } from "mdx/types";
import { IBM_Plex_Mono } from "next/font/google";
import Image from "next/image";
import Link from "next/link";

import { titleBar } from "./lib/bright-extensions";

const fontMono = IBM_Plex_Mono({
  subsets: ["latin"],
  display: "swap",
  variable: "--font-mono",
  weight: "400",
});

export function useMDXComponents(components: MDXComponents): MDXComponents {
  return {
    ...components,
    a: (props) => {
      const isExternal = props.href?.startsWith("http");
      return (
        <Link
          {...props}
          href={isExternal ? props.href || "" : `/docs/${props.href}`}
          target={isExternal ? "_blank" : undefined}
          rel={isExternal ? "noopener noreferrer" : undefined}
        />
      );
    },
    img: (props) => {
      if (props.src?.includes(".mp4")) {
        return (
          <video playsInline loop muted autoPlay preload="auto">
            <source src={props.src} type="video/mp4" />
          </video>
        );
      }

      return (
        // @ts-ignore
        <Image
          {...props}
          width={684}
          height={0}
          style={{ width: "100%", height: "auto" }}
          alt="Documentation image"
          unoptimized
        />
      );
    },
    pre: (props) => {
      return (
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
          extensions={[titleBar]}
        />
      );
    },
    Vimeo: ({ id }: { id: string }) => {
      return (
        <div style={{ padding: "56% 0 0 0 ", position: "relative" }}>
          <iframe
            src={`https://player.vimeo.com/video/${id}?badge=0&amp;autopause=0&amp;player_id=0&amp;app_id=58479`}
            frameBorder="0"
            allow="autoplay; fullscreen; picture-in-picture; clipboard-write"
            style={{
              position: "absolute",
              top: 0,
              left: 0,
              width: "100%",
              height: "100%",
            }}
            title="test video"
          ></iframe>
        </div>
      );
    },
  };
}
