"use client";

import styled from "styled-components";

import { mobile } from "@/constants/breakpoints";

const ResponsiveMobile = styled.div`
  display: none;
  @media ${mobile} {
    display: flex;
    justify-content: center;
  }
`;

const ResponsiveDesktop = styled.div`
  display: flex;
  align-items: center;

  @media ${mobile} {
    display: none;
  }
`;

const Responsive = {
  Mobile: ResponsiveMobile,
  Desktop: ResponsiveDesktop,
};

export default Responsive;
