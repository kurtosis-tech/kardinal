"use client";
import styled from "styled-components";

import { ButtonPrimary, ButtonTertiary } from "@/components/Button";
import { mobile } from "@/constants/breakpoints";

const CTAButtons = () => {
  return (
    <S.CTAButtons>
      <ButtonPrimary
        analyticsId="button_hero_github"
        href="https://github.com/kurtosis-tech/kardinal"
        rel="noopener noreferrer"
        target="_blank"
      >
        View on GitHub
      </ButtonPrimary>
      <ButtonTertiary
        analyticsId="button_hero_playground"
        href="https://github.com/kurtosis-tech/kardinal-playground"
        rel="noopener noreferrer"
        target="_blank"
      >
        Try in playground
      </ButtonTertiary>
    </S.CTAButtons>
  );
};

namespace S {
  export const CTAButtons = styled.div`
    display: flex;
    flex-direction: row;
    gap: 16px;
    @media ${mobile} {
      flex-direction: column;
    }
  `;
}

export default CTAButtons;
