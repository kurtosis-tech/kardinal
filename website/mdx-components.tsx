import type { MDXComponents } from "mdx/types";
import Image from "next/image";
import Link from "next/link";

import CodeBlock from "@/components/CodeBlock";
import DocsHeading from "@/components/DocsHeading";

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
    h1: (props) => <DocsHeading as="h1" {...props} />,
    h2: (props) => <DocsHeading as="h2" {...props} />,
    h3: (props) => <DocsHeading as="h3" {...props} />,
    h4: (props) => <DocsHeading as="h4" {...props} />,
    h5: (props) => <DocsHeading as="h5" {...props} />,
    h6: (props) => <DocsHeading as="h6" {...props} />,
    // @ts-expect-error children will not be undefined
    pre: CodeBlock,
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
