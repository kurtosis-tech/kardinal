import styled from "styled-components";

import { mobile, tablet } from "@/constants/breakpoints";

const commonStyles = `
  color: var(--foreground-dark);
  font-size: 54px;
  line-height: 1.15;
  font-style: normal;
  font-weight: 300;
  letter-spacing: -1.44px;

  @media ${tablet} {
    font-size: 40px;
  }

  @media ${mobile} {
    font-size: 32px;
    line-height: normal;
  }
`;

const H1 = styled.h1`
  ${commonStyles}
`;

const H2 = styled.h2`
  ${commonStyles}
  font-size: 42px;
  line-height: normal;

  @media ${mobile} {
    font-size: 24px;
    line-height: normal;
  }
`;

const H3 = styled.h3`
  ${commonStyles}
  font-size: 21px;
  font-weight: 600;
  line-height: 1.33;

  @media ${mobile} {
    font-size: 18px;
  }
`;

const Heading = {
  H1,
  H2,
  H3,
};

export default Heading;
