"use client";
import { PropsWithChildren } from "react";
import styled, { CSSProperties } from "styled-components";

import { mobile } from "@/constants/breakpoints";

const Section = ({
  children,
  contrast,
  padTop,
  padBottom,
  flexDirection,
  style,
  noPadBottomMobile,
}: PropsWithChildren<{
  flexDirection?: string;
  id?: string;
  padTop?: boolean;
  padBottom?: boolean;
  contrast?: boolean;
  style?: CSSProperties;
  noPadBottomMobile?: boolean;
}>) => {
  return (
    <S.Section $contrast={contrast} style={style}>
      <S.Content
        $flexDirection={flexDirection}
        $padTop={padTop}
        $padBottom={padBottom}
        $noPadBottomMobile={noPadBottomMobile}
      >
        {children}
      </S.Content>
    </S.Section>
  );
};

namespace S {
  export const Content = styled.div<{
    $flexDirection?: string;
    $padTop?: boolean;
    $padBottom?: boolean;
    $noPadBottomMobile?: boolean;
  }>`
    display: flex;
    flex-direction: ${(props) => props.$flexDirection ?? "column"};
    padding: 0 16px;
    padding-top: ${(props) => (props.$padTop ? "192px" : "0px")};
    padding-bottom: ${(props) => (props.$padBottom ? "128px" : "0px")};
    width: 100%;
    max-width: var(--max-width);

    @media ${mobile} {
      padding: ${(props) =>
        props.$padTop
          ? "48px 16px 32px 16px"
          : props.$noPadBottomMobile
            ? "0px 16px 0px 16px"
            : "0px 16px 32px 16px"};
    }
  `;

  export const Section = styled.section<{
    $contrast?: boolean;
  }>`
    width: 100%;
    display: flex;
    justify-content: center;
    position: relative;
    min-height: ${(props) => (props.$contrast ? "593px" : "unset")};

    background-image: ${(props) =>
      props.$contrast ? "url('/dotted-bg.svg')" : "transparent"};
    background-repeat: no-repeat;

    @media ${mobile} {
      min-height: unset;
      background: transparent;
    }
  `;
}

export default Section;
