"use client";
import Image from "next/image";
import { PropsWithChildren, ReactNode } from "react";
import styled from "styled-components";

import Heading from "@/components/Heading";
import Section from "@/components/Section";
import Spacer from "@/components/Spacer";
import Text from "@/components/Text";
import { mobile } from "@/constants/breakpoints";

const DefaultHeading = () => {
  return (
    <>
      Develop in prod
      <br />
      <em>Fearlessly</em>
    </>
  );
};

const CTA = ({
  children,
  heading,
  imageUrl,
  invert,
  fullHeight,
  padTop,
  padBottom,
}: PropsWithChildren<{
  heading?: ReactNode;
  buttonText?: null | string;
  imageUrl?: null | string;
  invert?: boolean;
  fullHeight?: boolean;
  padTop?: boolean;
  padBottom?: boolean;
}>) => {
  return (
    <S.CTA
      $invert={invert}
      $fullHeight={fullHeight}
      $padTop={padTop}
      $padBottom={padBottom}
    >
      <Section>
        <S.Content>
          {imageUrl !== null && (
            <Image
              src={imageUrl || "/bg.png"}
              alt="Get started now"
              width={280}
              height={138}
              unoptimized
            />
          )}
          <Heading.H1>{heading || <DefaultHeading />}</Heading.H1>

          {children || (
            <Text.Base>
              {
                'Kardinal introduces "maturity-based access controls”, allowing you to progressively grant access to sensitive production data as service versions move from “dev” to “QA” to “production” quality.'
              }
            </Text.Base>
          )}
          <Spacer height={0} />
          {/*
          <Button.Codespaces
            href="https://github.com/kurtosis-tech/kardinal-playground/"
            target="_blank"
            analyticsId="button_codespaces_demo"
          >
            Run the demo in Github Codespaces
          </Button.Codespaces>
          */}
        </S.Content>
      </Section>
    </S.CTA>
  );
};

namespace S {
  export const CTA = styled.div<{
    $invert?: boolean;
    $fullHeight?: boolean;
    $padTop?: boolean;
    $padBottom?: boolean;
  }>`
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    flex-shrink: 0;
    background-color: ${(props) =>
      props.$invert ? "var(--gray-bg)" : "transparent"};
    background-image: ${(props) => (props.$invert ? "unset" : "url(/bg.svg)")};
    background-repeat: ${(props) => (props.$invert ? "repeat-x" : "no-repeat")};
    background-size: contain;
    background-position: center;
    width: 100%;
    height: ${(props) => (props.$fullHeight ? "90vh" : "unset")};
    max-height: 1024px;
    min-height: 720px;
    padding-top: ${(props) => (props.$padTop ? "192px" : "0px")};
    padding-bottom: ${(props) => (props.$padBottom ? "128px" : "0px")};

    @media ${mobile} {
      max-height: unset;
      margin-top: 0;
      background-size: cover;
    }

    ${(props) =>
      props.$invert &&
      `
      --foreground: var(--foreground-inverted);
      --foreground-dark: var(--foreground-dark-inverted);
      `}
  `;

  export const Content = styled.div`
    display: flex;
    flex-direction: column;
    gap: 24px;
    align-items: center;
    max-width: 816px;
    margin: 0 auto;
    text-align: center;
    z-index: 1;

    @media ${mobile} {
      max-width: 100%;
      gap: 16px;
    }
  `;
}

export default CTA;
