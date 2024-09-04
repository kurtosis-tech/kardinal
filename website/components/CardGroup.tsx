"use client";

import { PropsWithChildren } from "react";
import styled from "styled-components";

const CardGroup = ({ children }: PropsWithChildren) => {
  return <S.CardGroup>{children}</S.CardGroup>;
};

namespace S {
  export const CardGroup = styled.div`
    display: grid;
    grid-template-columns: 1fr 1fr 1fr;
    grid-gap: 16px;
  `;
}

export default CardGroup;
