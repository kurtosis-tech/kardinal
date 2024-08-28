"use client";
import styled from "styled-components";

import { ButtonTertiary } from "@/components/Button";
import EmailCapture from "@/components/EmailCapture";
import Sparkles from "@/components/icons/Sparkles";
import { mobile } from "@/constants/breakpoints";

const CTAButtons = () => {
  return (
    <S.CTAButtons>
      <ButtonTertiary
        analyticsId="button_hero_playground"
        href="https://github.com/kurtosis-tech/kardinal-playground"
        rel="noopener noreferrer"
        target="_blank"
        iconRight={<Sparkles size={16} />}
      >
        Try in Playground
      </ButtonTertiary>
      <EmailCapture buttonAnalyticsId="button_footer_join_waitlist" />
    </S.CTAButtons>
  );
};

namespace S {
  export const CTAButtons = styled.div`
    margin-top: 16px;
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 16px;
    @media ${mobile} {
      flex-direction: column;
    }
  `;
}

export default CTAButtons;
