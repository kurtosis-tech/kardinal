"use client";
import styled from "styled-components";

import { ButtonPrimary, ButtonTertiary } from "@/components/Button";

const CTAButtons = () => {
  return (
    <S.CTAButtons>
      <ButtonPrimary analyticsId="button_hero_github">
        View on GitHub
      </ButtonPrimary>
      <ButtonTertiary analyticsId="button_hero_github">
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
  `;
}

export default CTAButtons;
