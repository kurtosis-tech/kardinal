"use client";
import Image from "next/image";
import { PropsWithChildren, ReactNode } from "react";
import { BiRightArrowAlt } from "react-icons/bi";
import styled, { CSSProperties } from "styled-components";

import Button from "@/components/Button";
import Heading from "@/components/Heading";
import Responsive from "@/components/Responsive";
import Section from "@/components/Section";
import Text from "@/components/Text";
import { mobile, tablet } from "@/constants/breakpoints";
import continuityImg from "@/public/illustrations/continuity-2.svg";

interface BaseProps {
  centered?: boolean;
  column?: boolean;
  contrast?: boolean;
  fullWidth?: boolean;
  fullWidthImageUrl?: string;
  mobileFullWidthImageUrl?: string;
  heading: ReactNode;
  imageUrl?: string;
  mobileImageUrl?: string;
  padTop?: boolean;
  padBottom?: boolean;
  reverse?: boolean;
  style?: CSSProperties;
  negativeTopOffset?: boolean;
  childrenWrapper?: "div";
}

type Props = PropsWithChildren<
  BaseProps & { buttonText?: undefined; buttonAnalyticsId?: undefined }
>;
type PropsWithButton = PropsWithChildren<
  BaseProps & { buttonText?: null | string; buttonAnalyticsId: null | string }
>;

const Content = ({
  buttonAnalyticsId,
  buttonText = "Read the docs",
  centered,
  children,
  column,
  contrast,
  fullWidth,
  fullWidthImageUrl,
  mobileFullWidthImageUrl,
  heading,
  imageUrl,
  mobileImageUrl,
  padBottom,
  padTop,
  reverse,
  style,
  negativeTopOffset,
  childrenWrapper,
}: Props | PropsWithButton) => {
  const ChildrenWrapper = childrenWrapper ?? Text.Base;

  return (
    <Section
      contrast={contrast}
      padTop={padTop}
      padBottom={padBottom}
      style={style}
    >
      <S.Content
        style={style}
        $hasImage={imageUrl != null}
        $reverse={reverse}
        $centered={centered}
        $fullWidth={fullWidth}
        $column={column}
        $negativeTopOffset={negativeTopOffset}
      >
        {fullWidth && (
          <S.ContinuityImage
            src={continuityImg}
            width={83}
            height={277}
            alt="Diagram lines"
          />
        )}
        {imageUrl && (
          <S.ContentImage
            alt="Illustration"
            src={imageUrl}
            width={856}
            height={584}
          />
        )}
        {mobileImageUrl && (
          <S.MobileContentImage
            alt="Illustration"
            src={mobileImageUrl}
            width={327}
            height={682}
          />
        )}
        <S.Container $fullWidth={fullWidth} $centered={centered}>
          {heading != null && <Heading.H2>{heading}</Heading.H2>}
          <ChildrenWrapper>{children}</ChildrenWrapper>
          {buttonText && (
            <S.ButtonWrapper>
              <Button.Secondary
                analyticsId={buttonAnalyticsId || "button_see_how"}
                href="/docs"
                iconRight={<BiRightArrowAlt size={20} />}
              >
                {buttonText}
              </Button.Secondary>
            </S.ButtonWrapper>
          )}
        </S.Container>
      </S.Content>
      {fullWidthImageUrl && !mobileFullWidthImageUrl && (
        <S.FullWidthImage
          width={1248}
          height={405}
          src={fullWidthImageUrl}
          alt="Architecture diagram"
        />
      )}
      {fullWidthImageUrl && mobileFullWidthImageUrl && (
        <>
          <Responsive.Desktop>
            <S.FullWidthImage
              width={1248}
              height={405}
              src={fullWidthImageUrl}
              alt="Architecture diagram"
              unoptimized
            />
          </Responsive.Desktop>
          <Responsive.Mobile>
            <S.FullWidthImage
              width={360}
              height={716}
              src={mobileFullWidthImageUrl}
              alt="Architecture diagram"
              style={{ maxWidth: "360px" }}
              unoptimized
            />
          </Responsive.Mobile>
        </>
      )}
    </Section>
  );
};

namespace S {
  export const Content = styled.div<{
    $hasImage?: boolean;
    $reverse?: boolean;
    $centered?: boolean;
    $fullWidth?: boolean;
    $column?: boolean;
    $negativeTopOffset?: boolean;
  }>`
    display: grid;
    grid-template-columns: ${(props) =>
      props.$column ? "minmax(auto, 1fr);" : "auto 617px"};

    gap: 0px;
    margin-top: ${(props) => (props.$negativeTopOffset ? "-300px" : "unset")};

    @media ${tablet} {
      margin-top: 0;
      display: flex;
      flex-direction: column;
      align-items: center;
    }

    @media ${mobile} {
      padding-bottom: 0px;
      grid-row-gap: 64px;
    }

    ${(props) =>
      props.$reverse &&
      `
      > *:nth-child(1) {
        order: 1;
      }

      > *:nth-child(2) {
        order: -1;
      }
      `}
  `;

  export const Container = styled.div<{
    $fullWidth?: boolean;
    $centered?: boolean;
  }>`
    display: flex;
    flex-direction: column;
    justify-content: flex-end;
    padding-bottom: 48px;
    gap: 24px;
    max-width: ${(props) => (props.$fullWidth ? "100%" : "617px")};
    text-align: ${(props) => (props.$centered ? "center" : "left")};
    padding-left: ${(props) => (props.$fullWidth ? "104px" : "0px")};

    @media ${tablet} {
      padding-bottom: 32px;
      padding-left: 0px;
      text-align: center;
      width: 100%;
      max-width: ${(props) => (props.$fullWidth ? "100%" : "420px")};
      justify-content: center;
    }

    @media ${mobile} {
      gap: 16px;
      padding-bottom: 0px;
    }
  `;

  export const ButtonWrapper = styled.div`
    display: flex;
    width: 100%;

    @media ${tablet} {
      text-align: center;
      justify-content: center;
    }
  `;

  export const ContentImage = styled(Image)`
    width: auto;
    z-index: 3;
    margin-left: -1px;
    pointer-events: none;

    @media ${tablet} {
      display: none;
    }
  `;

  export const MobileContentImage = styled(Image)`
    display: none;
    pointer-events: none;

    @media ${tablet} {
      display: block;
      flex-shrink: 0;
      height: auto;
      margin-left: -2px;
      margin-top: -120px;
      max-width: 327px;
      z-index: 3;
    }
    @media ${mobile} {
      margin-top: -64px;
    }
  `;

  export const FullWidthImage = styled(Image)`
    margin-top: 64px;
    width: 100%;
    height: auto;
    pointer-events: none;
  `;

  export const ContinuityImage = styled(Image)`
    position: absolute;
    top: 0px;
    pointer-events: none;

    @media ${tablet} {
      display: none;
    }
  `;
}

export default Content;
