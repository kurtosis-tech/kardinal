import { PropsWithChildren } from "react";
import styled from "styled-components";

import { mobile } from "@/constants/breakpoints";
import { useModal } from "@/context/ModalContext";

const ResponsiveNav = ({ children }: PropsWithChildren) => {
  const { isNavOpen } = useModal();
  return <S.ResponsiveNav $isNavOpen={isNavOpen}>{children}</S.ResponsiveNav>;
};

namespace S {
  export const ResponsiveNav = styled.div<{ $isNavOpen: boolean }>`
    background-color: var(--background);
    transition: transform 0.3s ease-in-out;

    @media ${mobile} {
      transform: ${({ $isNavOpen }) =>
        $isNavOpen ? "translateX(0)" : "translateX(-100%)"};
      padding: 48px 24px;
      box-shadow: ${({ $isNavOpen }) =>
        $isNavOpen ? "0 0 30px 0 rgba(0, 0, 0, 1)" : "none"};
      position: fixed;
      top: 0;
      left: 0;
      width: 80%;
      height: 100%;
      z-index: 999;
    }
  `;
}
export default ResponsiveNav;
