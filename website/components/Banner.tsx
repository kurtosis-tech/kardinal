"use client";

import styled from "styled-components";

import Button from "@/components/Button";
import Text from "@/components/Text";
import { mobile } from "@/constants/breakpoints";

const Banner = () => {
  return (
    <S.Banner>
      <S.BannerText>
        <b>Disclaimer:</b> The Kardinal project is no longer maintained. The
        archived source code is still available on{" "}
        <a
          style={{ textDecoration: "underline" }}
          href="https://github.com/kurtosis-tech/kardinal"
        >
          Github
        </a>
        .
      </S.BannerText>
    </S.Banner>
  );
};

namespace S {
  export const Banner = styled.div`
    display: flex;
    background: var(--gradient-brand);
    width: 100vw;
    padding: 12px 24px;
    align-items: center;
    justify-content: center;

    @media ${mobile} {
      display: none;
    }
  `;

  export const BannerText = styled(Text.Small)`
    color: var(--white-100);
    font-weight: 500;
  `;

  export const CTAButton = styled(Button.Secondary)`
    display: inline;
    color: var(--white-100);
    font-weight: 500;
    text-decoration: underline;
  `;

  export const CloseButton = styled.button`
    border: none;
    background: none;
    outline: none;
    cursor: pointer;
    color: var(--white-100);
    height: 32px;
    width: 32px;
    display: flex;
    align-items: center;
    justify-content: center;
    position: absolute;
    right: 16px;
    transform: translateY(0);
    transition: all 0.2s ease-in-out;

    &:hover {
      opacity: 0.8;
      transform: translateY(-2px);
    }
  `;
}

export default Banner;
