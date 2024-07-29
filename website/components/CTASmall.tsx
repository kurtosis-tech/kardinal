"use client";

import { ReactNode } from "react";
import styled from "styled-components";

import Heading from "@/components/Heading";
import Section from "@/components/Section";
import { mobile, tablet } from "@/constants/breakpoints";

const CTASmall = ({
  children,
  heading,
  myPrecious,
}: {
  children?: ReactNode;
  heading: ReactNode;
  myPrecious?: boolean;
}) => {
  return (
    <Section>
      <S.CTASmall>
        {myPrecious && (
          <>
            <S.Torus1 />
            <S.Torus2 />
          </>
        )}
        <S.Content>
          <S.CTAHeading>{heading}</S.CTAHeading>
          {children}
        </S.Content>
      </S.CTASmall>
    </Section>
  );
};

namespace S {
  export const CTASmall = styled.div`
    display: flex;
    flex-direction: column;
    align-items: center;
    border-radius: 12px;
    background: rgba(252, 160, 97, 0.08);
    margin-bottom: 80px;
    padding: 64px 0;
    width: 100%;
    align-self: stretch;
    position: relative;

    @media ${mobile} {
      padding: 16px;
      margin-top: 48px;
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
}

export default CTASmall;
