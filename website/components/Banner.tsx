"use client";

import { useEffect, useState } from "react";
import { BiX } from "react-icons/bi";
import styled from "styled-components";

import Button from "@/components/Button";
import Text from "@/components/Text";
import { mobile } from "@/constants/breakpoints";
import { useModal } from "@/context/ModalContext";

const Banner = () => {
  const { toggleModal } = useModal();
  const [bannerIsDismissed, setBannerIsDismissed] = useState(
    typeof window !== "undefined" // client-side only
      ? localStorage.getItem("bannerIsDismissed") === "true"
      : false,
  );
  useEffect(() => {
    if (bannerIsDismissed) {
      localStorage.setItem("bannerIsDismissed", "true");
    }
  }, [bannerIsDismissed]);

  if (bannerIsDismissed) return null;
  return (
    <S.Banner>
      <S.BannerText>
        Kardinal is still in build mode - please{" "}
        <S.CTAButton onClick={toggleModal} analyticsId="banner_join_waitlist">
          join the beta
        </S.CTAButton>{" "}
        to be alerted when we launch
      </S.BannerText>
      <S.CloseButton onClick={() => setBannerIsDismissed(true)}>
        <BiX size={24} />
      </S.CloseButton>
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
