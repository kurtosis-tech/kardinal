"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";
import React from "react";
import styled from "styled-components";

import Heading from "@/components/Heading";
import ResponsiveNav from "@/components/ResponsiveNav";
import { mobile } from "@/constants/breakpoints";

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
      { path: "install", title: "Installation" },
      { path: "fundamentals", title: "Fundamentals" },
      { path: "own-app", title: "Create your first flow" },
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
    path: "references",
    title: "References",
    children: [
      { path: "comparisons", title: "Comparison to alternatives" },
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
  const pathname = usePathname();

  return (
    <ResponsiveNav>
      <S.NavItemsWrapper>
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
      </S.NavItemsWrapper>
    </ResponsiveNav>
  );
};

namespace S {
  export const NavItemsWrapper = styled.div`
    padding-top: 240px;
    @media ${mobile} {
      padding-top: 0;
    }
    display: flex;
    flex-direction: column;
    gap: 16px;
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
