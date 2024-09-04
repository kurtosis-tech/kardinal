"use client";
import { FiCalendar, FiGithub } from "react-icons/fi";
import styled from "styled-components";

import { ButtonPrimary, ButtonTertiary } from "@/components/Button";
import { mobile } from "@/constants/breakpoints";

const CTAButtons = () => {
  return (
    <S.CTAButtons>
      <ButtonPrimary
        analyticsId="button_hero_get_demo"
        href="https://calendly.com/d/cqhd-tgj-vmc/45-minute-meeting"
        rel="noopener noreferrer"
        target="_blank"
        iconLeft={<FiCalendar size={18} />}
        size="lg"
      >
        Get a Demo
      </ButtonPrimary>
      <ButtonTertiary
        analyticsId="button_hero_github"
        href="https://github.com/kurtosis-tech/kardinal"
        rel="noopener noreferrer"
        target="_blank"
        iconLeft={<FiGithub size={18} />}
        size="lg"
      >
        View on GitHub
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
    gap: 16px;
    @media ${mobile} {
      flex-direction: column;
    }
  `;
}

export default CTAButtons;
