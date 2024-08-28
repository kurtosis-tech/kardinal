"use client";
import { FiGithub } from "react-icons/fi";
import styled from "styled-components";

import { ButtonTertiary } from "@/components/Button";
import EmailCapture from "@/components/EmailCapture";
import { mobile } from "@/constants/breakpoints";

const CTAButtons = () => {
  return (
    <S.CTAButtons>
      <ButtonTertiary 
        analyticsId="button_hero_github"
        href="https://github.com/kurtosis-tech/kardinal"
        rel="noopener noreferrer"
        target="_blank"
        iconLeft={<FiGithub size={18} />}
        size="lg"
      >
        View on GitHub
      </ButtonTertiary >
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
