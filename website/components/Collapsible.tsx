"use client";

import React, { useState } from 'react';
import { BiChevronDown, BiChevronRight } from 'react-icons/bi';
import styled from 'styled-components';

interface CollapsibleProps {
  title: string;
  children: React.ReactNode;
}

const Collapsible: React.FC<CollapsibleProps> = ({ title, children }) => {
  const [isOpen, setIsOpen] = useState(false);

  return (
    <S.Wrapper>
      <S.Toggle onClick={() => setIsOpen(!isOpen)}>
        {isOpen ? <BiChevronDown size={24} /> : <BiChevronRight size={24} />}
        {title}
      </S.Toggle>
      {isOpen && <S.Content>{children}</S.Content>}
    </S.Wrapper>
  );
};

namespace S {
  export const Wrapper = styled.div`
    margin-bottom: 1rem;
    border: 1px solid var(--gray-light);
    border-radius: 8px;
    overflow: hidden;
  `;

  export const Toggle = styled.button`
    background-color: var(--background);
    color: var(--foreground);
    border: none;
    padding: 1rem;
    width: 100%;
    text-align: left;
    cursor: pointer;
    font-weight: 500;
    font-size: 16px;
    display: flex;
    align-items: center;
    transition: background-color 0.2s ease;

    &:hover {
      background-color: var(--gray-light);
    }

    svg {
      margin-right: 0.5rem;
    }
  `;

  export const Content = styled.div`
    padding: 1rem;
    background-color: var(--background);
    color: var(--foreground);
    font-size: 14px;
    line-height: 1.6;
  `;
}

export default Collapsible;