"use client";

import styled from "styled-components";

import { mobile } from "@/constants/breakpoints";

const commonStyles = `
  color: var(--white-60, rgba(255, 255, 255, 0.6));
  font-size: 21px;
  font-style: normal;
  font-weight: 400;
  line-height: 160%; /* 33.6px */

  @media ${mobile} {
    font-size: 16px;
  }
`;

const Ol = styled.ol`
  ${commonStyles}
`;
const Ul = styled.ul`
  ${commonStyles}
`;

const List = {
  Ol,
  Ul,
};

export default List;
