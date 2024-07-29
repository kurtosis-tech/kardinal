"use client";
import Link from "next/link";
import styled from "styled-components";

import Text from "@/components/Text";
import { mobile } from "@/constants/breakpoints";

const Footer = () => {
  return (
    <S.Footer>
      <S.Content>
        <S.FooterText>
          Kardinal is built by Kurtosis Technologies.
          <br />© 2024 Kurtosis Technologies. All Rights Reserved. <br />
        </S.FooterText>
        <S.FooterText>
          <Link href="https://www.kurtosis.com/privacy-policy" target="_blank">
            Privacy policy
          </Link>
        </S.FooterText>
      </S.Content>
    </S.Footer>
  );
};

namespace S {
  export const Footer = styled.footer`
    display: flex;
    align-items: center;
    justify-content: center;
    width: 100%;
    padding: 48px 16px;

    @media ${mobile} {
      grid-template-columns: 1fr;
      gap: 24px;
      padding: 24px;
    }
  `;

  export const Content = styled.div`
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    width: 100%;
    align-items: flex-start;
    gap: 24px;
    max-width: var(--max-width);

    @media ${mobile} {
      flex-direction: column;
      gap: 8px;
    }
  `;

  export const FooterText = styled(Text.Small)`
    @media ${mobile} {
      font-size: 12px;
      line-height: 16px;
    }
  `;
}

export default Footer;
