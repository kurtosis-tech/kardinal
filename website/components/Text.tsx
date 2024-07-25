"use client";

import styled from "styled-components";

import { mobile } from "@/constants/breakpoints";

export const TextBase = styled.p`
  color: var(--foreground);
  font-size: 21px;
  font-style: normal;
  font-weight: 400;
  line-height: 32px;

  @media ${mobile} {
    font-size: 18px;
    line-height: 28px;
  }
`;

export const TextSmall = styled(TextBase)`
  font-size: 16px;
  line-height: 24px; /* 150% */
`;

const Text = {
  Base: TextBase,
  Small: TextSmall,
};

export default Text;
