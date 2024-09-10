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
            <Heading.H2>Fancy a demo?</Heading.H2>
            <Text.Base>
              Schedule some time for a personalized demo of Kardinal.
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
    padding: 64px 0;
    align-items: center;
    justify-content: center;
  `,

  Content: styled.div`
    max-width: 1038px;
    border-radius: 12px;
    background-color: rgba(252, 160, 97, 0.08);
    width: 100%;
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 48px;
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
    @media ${mobile} {
      gap: 8px;
    }
  `,
};

export default CTADemo;
