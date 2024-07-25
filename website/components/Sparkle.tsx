"use client";
import Image from "next/image";
import styled from "styled-components";

import { mobile } from "@/constants/breakpoints";
import sparkleImg from "@/public/icons/sparkle.png";

const Sparkle = () => {
  return (
    <S.ImageElement
      src={sparkleImg}
      role="presentation"
      alt="Sparlke emoji"
      width={48}
      height={48}
      unoptimized
      style={{ marginLeft: "0.7rem" }}
    />
  );
};

export const S = {
  ImageElement: styled(Image)`
    width: 48px;
    height: 48px;

    @media ${mobile} {
      width: 32px;
      height: 32px;
    }
  `,
};

export default Sparkle;
