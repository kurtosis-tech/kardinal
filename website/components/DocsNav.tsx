"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";
import React from "react";
import styled from "styled-components";

import Heading from "@/components/Heading";
import { mobile } from "@/constants/breakpoints";
import { useModal } from "@/context/ModalContext";

interface NavItem {
  path: string;
  title: string;
  children?: NavItem[];
}

const navItems: NavItem[] = [
  {
    path: "getting-started",
    title: "Getting Started",
    children: [
      { path: "demo", title: "Run on a demo app" },
      { path: "install", title: "Run on your own app" },
    ],
  },
  {
    path: "architecture",
    title: "Architecture",
    children: [{ path: "overview", title: "Overview" }],
  },
  {
    path: "concepts",
    title: "Concepts",
    children: [
      { path: "flows", title: "Flows" },
      { path: "plugins", title: "Stateful Service Plugins" },
    ],
  },
  {
    path: "use-cases",
    title: "Use Cases",
    children: [
      {
        path: "isolated-dev-sandbox-flows",
        title: "Isolated Dev Sandbox Flows",
      },
      {
        path: "preview-and-feature-branch-flows",
        title: "Preview and Feature Branch Flows",
      },
      {
        path: "qa-flows",
        title: "QA Flows",
      },
    ],
  },
];

const NavItemImpl = ({
  item,
  parentBaseUrl,
}: {
  item: NavItem;
  parentBaseUrl: string;
}) => {
  const baseUrl = [parentBaseUrl, item.path].join("/");
  const pathname = usePathname();
  return item.children ? (
    <div key={baseUrl}>
      <S.ItemGroupHeading>{item.title}</S.ItemGroupHeading>
      <S.ItemGroup>
        {item.children.map((item) => (
          <NavItemImpl
            key={baseUrl + item.path}
            item={item}
            parentBaseUrl={baseUrl}
          />
        ))}
      </S.ItemGroup>
    </div>
  ) : (
    <S.Item key={baseUrl} $isActive={baseUrl === pathname}>
      <S.NavLink href={baseUrl} $isActive={baseUrl === pathname}>
        {item.title}
      </S.NavLink>
    </S.Item>
  );
};

const DocsNav = () => {
  const { isNavOpen } = useModal();
  const pathname = usePathname();

  return (
    <S.DocsNav $isNavOpen={isNavOpen}>
      <S.ItemGroup>
        <S.Item $isActive={pathname === "/docs"} style={{ marginLeft: 0 }}>
          <S.NavLink href={"/docs"} $isActive={pathname === "/docs"}>
            Introduction
          </S.NavLink>
        </S.Item>
      </S.ItemGroup>

      {navItems.map((item) => (
        <NavItemImpl item={item} parentBaseUrl="/docs" key={item.path} />
      ))}
    </S.DocsNav>
  );
};

namespace S {
  export const DocsNav = styled.div<{ $isNavOpen: boolean }>`
    background-color: var(--background);
    padding-top: 240px;
    display: flex;
    flex-direction: column;
    gap: 16px;
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

  export const ItemGroupHeading = styled(Heading.H3)`
    font-size: 16px;
    line-height: normal;
    font-weight: 400;
    margin-bottom: 0;
    padding-left: 12px;
    display: flex;
    align-items: center;
    height: 32px;
    color: var(--gray-dark);
    letter-spacing: unset;
  `;

  export const ItemGroup = styled.ul`
    list-style: none;
    display: flex;
    flex-direction: column;
    margin-left: 0px;
  `;

  export const Item = styled.li<{ $isActive?: boolean }>`
    min-height: 32px;
    padding: 8px 12px;
    display: flex;
    align-items: center;
    margin-left: 16px;
    border-left: 2px solid
      ${({ $isActive }) => ($isActive ? "var(--brand-primary)" : "transparent")};
  `;

  export const NavLink = styled(Link)<{ $isActive?: boolean }>`
    font-size: 14px;
    font-weight: 400;
    line-height: normal;
    text-decoration: none;
    transition: color 0.2s ease-in-out;
    leading-trim: both;
    text-edge: cap;
    color: ${({ $isActive }) =>
      $isActive ? "var(--brand-primary)" : "var(--gray)"};

    &:hover {
      color: var(--brand-primary) !important;
    }
  `;
}
export default DocsNav;
