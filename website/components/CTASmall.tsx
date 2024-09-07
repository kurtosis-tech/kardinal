"use client";

import Image from "next/image";
import { ReactNode } from "react";
import styled from "styled-components";

import Heading from "@/components/Heading";
import Section from "@/components/Section";
import { mobile, tablet } from "@/constants/breakpoints";
import calendarLineImg from "@/public/illustrations/calendar-line.svg";

const CTASmall = ({
  children,
  heading,
  hasBackground,
}: {
  children?: ReactNode;
  heading: ReactNode;
  hasBackground?: boolean;
}) => {
  return (
    <Section>
      <S.CTASmall $hasBackground={hasBackground}>
        <Image
          src={calendarLineImg}
          width={83}
          height={336}
          alt="Image of a line with a calendar icon"
        />
        <S.Content>
          <S.CTAHeading>{heading}</S.CTAHeading>
          <S.ChildrenWrapper>{children}</S.ChildrenWrapper>
        </S.Content>
      </S.CTASmall>
    </Section>
  );
};

namespace S {
  export const CTASmall = styled.div<{ $hasBackground?: boolean }>`
    display: flex;
    flex-direction: column;
    align-items: center;
    border-radius: 12px;
    margin-bottom: 80px;
    padding-bottom: 64px;
    width: 100%;
    align-self: stretch;
    position: relative;
    background: ${({ $hasBackground }) =>
      $hasBackground ? "url(/bg-static.svg)" : "none"};

    @media ${mobile} {
      padding: 16px;
    }
  `;
  export const Content = styled.div`
    max-width: 616px;
    text-align: center;
    display: flex;
    flex-direction: column;
    gap: 24px;
    z-index: 1;
  `;
  export const Torus = styled.div`
    position: absolute;
    opacity: 0.1;
    background-repeat: no-repeat;
    background-blend-mode: normal;
  `;
  export const Torus1 = styled(Torus)`
    width: 554px;
    height: 302px;
    background-image: url("/illustrations/torus-1.svg");
    left: -30%;
    @media ${tablet} {
      left: -50%;
    }
    @media ${mobile} {
      display: none;
    }
  `;
  export const Torus2 = styled(Torus)`
    width: 515px;
    height: 362px;
    background-image: url("/illustrations/torus-2.svg");
    right: -30%;
    @media ${tablet} {
      right: -50%;
    }
    @media ${mobile} {
      display: none;
    }
  `;
  export const CTAHeading = styled(Heading.H2)`
    img {
      width: 32px;
      height: 32px;
    }
    @media ${mobile} {
      img {
        width: 24px;
        height: 24px;
      }
    }
  `;
  export const ChildrenWrapper = styled.div`
    display: flex;
    flex-direction: column;
    gap: 24px;
    align-items: center;
    justify-content: center;
  `;
}

export default CTASmall;
