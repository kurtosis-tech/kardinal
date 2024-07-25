"use client";
import Image from "next/image";
import Link from "next/link";
import { usePathname } from "next/navigation";
import { BiMenu, BiX } from "react-icons/bi";
import { FiGithub } from "react-icons/fi";
import { TbSparkles } from "react-icons/tb";
import styled from "styled-components";

import Banner from "@/components/Banner";
import Button from "@/components/Button";
import Responsive from "@/components/Responsive";
import { mobile, tablet } from "@/constants/breakpoints";
import { useModal } from "@/context/ModalContext";
import analytics from "@/lib/analytics";
import kardinalLogo from "@/public/kardinal-orange.png";

const WaitlistButton = () => {
  const { toggleModal } = useModal();
  return (
    <Button.Secondary
      analyticsId="button_nav_join_waitlist"
      onClick={toggleModal}
    >
      <TbSparkles size={20} />
      Join the beta
    </Button.Secondary>
  );
};

const NavLinksAndButton = () => {
  return (
    <>
      <S.NavLink
        href={"/docs"}
        onClick={() =>
          analytics.track("BUTTON_CLICK", { analyticsId: "link_nav_docs" })
        }
      >
        Docs
      </S.NavLink>
      {/*
      <Responsive.Desktop>
        <S.NavLink
          href={"https://github.com/kurtosis-tech/kardinal-playground/"}
          target="_blank"
          onClick={() =>
            analytics.track("BUTTON_CLICK", {
              analyticsId: "link_nav_playground",
            })
          }
        >
          <TbPrompt size={22} />
          Playground
        </S.NavLink>
      </Responsive.Desktop>
      */}
      <S.NavLink
        href={"https://github.com/kurtosis-tech/kardinal"}
        target="_blank"
        onClick={() =>
          analytics.track("BUTTON_CLICK", { analyticsId: "link_nav_github" })
        }
      >
        <FiGithub size={16} />
        Github
      </S.NavLink>
      <WaitlistButton />
    </>
  );
};

const Nav = () => {
  const { toggleNav, isNavOpen } = useModal();
  const pathname = usePathname();
  return (
    <S.Nav>
      {pathname.includes("/docs") && <Banner />}
      <S.Container>
        <S.Wordmark href={"/"}>
          <S.LogoImage
            src={kardinalLogo}
            alt="Kardinal logo"
            width={32}
            height={32}
            unoptimized
          />
          <S.LogoText>Kardinal</S.LogoText>
        </S.Wordmark>
        <S.NavSpacer />

        {pathname.includes("/docs") ? (
          <>
            <Responsive.Mobile>
              <WaitlistButton />
            </Responsive.Mobile>
            <S.MobileDocsMenu onClick={toggleNav}>
              {isNavOpen ? <BiX size={24} /> : <BiMenu size={24} />}
            </S.MobileDocsMenu>
            <Responsive.Desktop>
              <NavLinksAndButton />
            </Responsive.Desktop>
          </>
        ) : (
          <NavLinksAndButton />
        )}
      </S.Container>
    </S.Nav>
  );
};

namespace S {
  export const Nav = styled.nav`
    align-items: center;
    background-color: var(--background);
    display: flex;
    flex-direction: column;
    justify-content: center;
    left: 0;
    margin: 0 auto;
    margin: 0 auto;
    padding: 0 24px 24px 24px;
    position: fixed;
    right: 0;
    top: 0;
    width: 100vw;
    z-index: 10;
  `;

  export const Container = styled.div`
    align-items: center;
    display: flex;
    justify-content: center;
    margin: 0 auto;
    padding-top: 24px;
    max-width: var(--max-width);
    width: 100%;

    @media ${tablet} {
      padding-right: 16px;
    }

    @media ${mobile} {
      padding-right: 0;
    }
  `;

  export const NavSpacer = styled.div`
    flex: 1;
  `;

  export const Wordmark = styled(Link)`
    align-items: center;
    color: var(--black);
    display: flex;
    font-size: 21px;
    font-weight: 600;
    gap: 8px;
    justify-content: center;
    line-height: 28px;
    text-decoration: none;
    transform: translateY(0);
    transition: all 0.2s ease-in-out;
    user-select: none;

    &:hover {
      color: var(--black);
      cursor: pointer;
      transform: translateY(-2px);
      opacity: 0.8;
    }
  `;

  export const LogoText = styled.span`
    @media ${mobile} {
      display: none;
    }
  `;

  export const NavLink = styled(Link)`
    align-items: center;
    display: flex;
    gap: 4px;
    margin-right: 20px;
    text-decoration: none;
    transform: translateY(0);
    transition: all 0.2s ease-in-out;
    transition: all 0.2s ease-in-out;
    user-select: none;

    &:hover {
      transform: translateY(-2px);
      color: var(--brand-primary);
    }
  `;

  export const MobileDocsMenu = styled.button`
    display: none;

    @media ${mobile} {
      background: transparent;
      border-radius: 4px;
      border: 0;
      cursor: pointer;
      display: block;
      height: 24px;
      width: 24px;
      margin-left: 20px;
      color: var(--gray);
    }
  `;

  export const LogoImage = styled(Image)`
    height: 32px;
    width: 32px;
  `;
}
export default Nav;
