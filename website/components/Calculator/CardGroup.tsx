"use client";

import { PropsWithChildren } from "react";
import styled from "styled-components";

import { tablet } from "@/constants/breakpoints";

const CardGroup = ({ children }: PropsWithChildren) => {
  return <S.CardGroup>{children}</S.CardGroup>;
};

namespace S {
  export const CardGroup = styled.div`
    display: grid;
    grid-template-columns: 1fr 1fr 1fr;
    grid-gap: 16px;

    @media ${tablet} {
      grid-template-columns: 1fr;
    }
  `;
}

export default CardGroup;
