"use client";

import { useEffect, useState } from "react";
import { BiCheckDouble, BiCopy } from "react-icons/bi";
import styled from "styled-components";

interface Props {
  text: string;
}

const CodeCopyButton = ({ text }: Props) => {
  const [isCopied, setIsCopied] = useState(false);
  useEffect(() => {
    if (isCopied) {
      const timeout = setTimeout(() => {
        setIsCopied(false);
      }, 2000);
      return () => clearTimeout(timeout);
    }
  }, [isCopied]);

  const handleCopy = () => {
    if (navigator.clipboard == null) {
      console.error("Clipboard API not available");
      return;
    }
    setIsCopied(true);
    navigator.clipboard
      .writeText(text.trim())
      .then(() => {
        console.log("Text copied to clipboard");
      })
      .catch((err) => {
        console.error("Failed to copy text: ", err);
      });
  };

  return (
    <S.CodeCopyButton
      aria-label="Copy to clipboard"
      onClick={handleCopy}
      $isCopied={isCopied}
    >
      {isCopied ? (
        <BiCheckDouble
          size={24}
          style={{ pointerEvents: "none" }}
          role="presentation"
        />
      ) : (
        <BiCopy
          size={24}
          style={{ pointerEvents: "none" }}
          role="presentation"
        />
      )}
    </S.CodeCopyButton>
  );
};

namespace S {
  export const CodeCopyButton = styled.button<{ $isCopied: boolean }>`
    display: flex;
    position: absolute;
    right: 8px;
    top: 8px;
    color: white;
    background-color: ${(props) =>
      props.$isCopied
        ? "rgba(200, 255, 200, 0.5)"
        : "rgba(255, 255, 255, 0.1)"};
    border: none;
    outline: none;
    cursor: pointer;
    border-radius: 4px;
    padding: 4px;
    transition: background-color 0.2s ease-in-out;

    &:hover {
      background-color: ${(props) =>
        props.$isCopied
          ? "rgba(200, 255, 200, 0.5)"
          : "rgba(255, 255, 255, 0.2)"};
    }
  `;
}

export default CodeCopyButton;
