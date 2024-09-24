"use client";
import { FiCalendar, FiTerminal } from "react-icons/fi";
import styled from "styled-components";

import { ButtonPrimary, ButtonTertiary } from "@/components/Button";
import { mobile } from "@/constants/breakpoints";
import { calendlyDemoUrl } from "@/constants/urls";
import { scrollToId } from "@/utils";

const CTAButtons = () => {
  return (
    <S.CTAButtons>
      <ButtonPrimary
        analyticsId="button_hero_get_started"
        onClick={() => scrollToId("get-started")}
        rel="noopener noreferrer"
        target="_blank"
        iconLeft={<FiTerminal size={18} />}
        size="lg"
      >
        Get Started
      </ButtonPrimary>
      <ButtonTertiary
        analyticsId="button_hero_schedule_demo"
        href={calendlyDemoUrl}
        rel="noopener noreferrer"
        target="_blank"
        iconLeft={<FiCalendar size={18} />}
      >
        Schedule a Demo
      </ButtonTertiary>
    </S.CTAButtons>
  );
};

namespace S {
  export const CTAButtons = styled.div`
    margin-top: 16px;
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 24px;
    @media ${mobile} {
      flex-direction: column;
    }
  `;
}

export default CTAButtons;
