"use client";
import { PropsWithChildren } from "react";
import styled from "styled-components";

import DocsNav from "@/components/DocsNav";
import VotingWidget from "@/components/VotingWidget";
import { mobile } from "@/constants/breakpoints";

const DocsLayout = ({ children }: PropsWithChildren) => {
  return (
    <S.Wrapper>
      <DocsNav />
      <div>
        <S.DocsMain id="mdx-content">
          {children}
          <VotingWidget />
        </S.DocsMain>
      </div>
    </S.Wrapper>
  );
};

namespace S {
  export const Wrapper = styled.div`
    display: grid;
    grid-template-columns: 195px 1fr;
    padding-bottom: 128px;
    width: 100%;
    max-width: var(--max-width);

    @media ${mobile} {
      grid-template-columns: 1fr;
    }
  `;

  export const DocsMain = styled.main`
    display: flex;
    flex-direction: column;
    gap: 24px;
    padding: 160px 0 48px 0;
    width: 100%;
    max-width: 827px;
    margin: 0 auto;
    min-height: calc(100vh - 272px);
    position: relative;

    @media ${mobile} {
      padding: 140px 24px 32px 24px;
    }

    h1,
    h2,
    h3,
    h4,
    h5,
    h6 {
      margin-bottom: 0;
      margin-top: 0;
      color: var(--foreground-dark);
      font-size: 42px;
      font-style: normal;
      font-weight: 500;
      line-height: normal;
      letter-spacing: unset;

      @media (max-width: 768px) {
        font-size: 24px;
      }
    }

    h3 {
      font-size: 32px;
      font-weight: 600;
      line-height: 1.33;

      @media (max-width: 768px) {
        font-size: 18px;
      }
    }

    ol {
      margin-top: 1rem;
    }

    ul {
      margin-bottom: 1rem;
    }

    p,
    ol,
    ul {
      color: var(--foreground);
      font-size: 18px;
      font-style: normal;
      font-weight: 400;
      line-height: 28px;

      line-height: 160%; /* 33.6px */
      margin-bottom: 0;

      @media (max-width: 768px) {
        font-size: 16px;
        margin-bottom: 1rem;
      }
    }

    ol,
    ul {
      margin-left: 2rem;
    }

    small {
      font-size: 16px;
      line-height: 24px; /* 150% */

      @media (max-width: 768px) {
        font-size: 14px;
      }
    }

    em {
      animation-name: none;
      animation-duration: unset;
      animation-name: unset;
      animation-iteration-count: unset;
      background: unset;
      -webkit-background-clip: unset;
      -webkit-text-fill-color: unset;
      color: var(--brand-primary);
    }

    a {
      color: var(--foreground);
      text-decoration: underline;
      transition: color 0.1s ease;
    }

    a:hover {
      color: var(--brand-primary);
    }

    video {
      max-width: 100%;
      height: auto;
    }

    table {
      border-spacing: 0;
      border-collapse: collapse;
    }

    td,
    th {
      border-bottom: 1px solid var(--gray-light);
      margin-right: -1px;
      margin-left: -1px;
      padding: 12px 8px;
      max-width: 256px;
    }

    th {
      font-weight: 600;
      color: var(--foreground-dark);
    }
  `;
}

export default DocsLayout;
