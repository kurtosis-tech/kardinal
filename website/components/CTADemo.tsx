"use client";

import { FiArrowRight } from "react-icons/fi";
import styled from "styled-components";

import Button from "@/components/Button";
import Heading from "@/components/Heading";
import Section from "@/components/Section";
import Text from "@/components/Text";
import { mobile } from "@/constants/breakpoints";
import { calendlyDemoUrl } from "@/constants/urls";

const CTADemo = () => {
  return (
    <Section>
      <S.CTADemo>
        <S.Content>
          <S.TextWrapper>
            <Heading.H2>Want to see Kardinal in action?</Heading.H2>
            <Text.Base>
              Use the link on the right to see how you can save your team on
              both cost and maintenance overhead:
            </Text.Base>
          </S.TextWrapper>
          <Button.Primary
            analyticsId="button_demo_cta_schedule_demo"
            iconRight={<FiArrowRight />}
            href={calendlyDemoUrl}
            rel="noopener noreferrer"
            target="_blank"
          >
            Schedule a demo
          </Button.Primary>
        </S.Content>
      </S.CTADemo>
    </Section>
  );
};

const S = {
  CTADemo: styled.div`
    display: flex;
    width: 100%;
    padding: 92px 0;
    align-items: center;
    justify-content: center;
    background: url(/bg-static.svg);
    background-position-y: -64px;
  `,

  Content: styled.div`
    border-radius: 12px;
    width: 100%;
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 48px 0;
    gap: 4px;

    @media ${mobile} {
      flex-direction: column;
      gap: 16px;
      align-items: flex-start;
    }
  `,

  TextWrapper: styled.div`
    display: flex;
    flex-direction: column;
    gap: 4px;
    max-width: 675px;
    @media ${mobile} {
      gap: 8px;
    }
  `,
};

export default CTADemo;
