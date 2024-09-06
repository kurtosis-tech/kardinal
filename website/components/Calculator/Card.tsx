"use client";

import styled from "styled-components";

interface Props {
  title: string;
  values: { label: string; value: string }[];
  isContrast?: boolean;
}
const Card = ({ title, values, isContrast }: Props) => {
  if (values == null || values.length !== 2) {
    throw new Error("Card must receive 2 values");
  }
  return (
    <S.Card $isContrast={isContrast}>
      <S.Title>{title}</S.Title>
      <S.Container>
        <S.ContentBox>
          <S.Value $large>{values[0].value}</S.Value>
          <S.Label>{values[0].label}</S.Label>
        </S.ContentBox>
        <S.ContentBox>
          <S.Value>{values[1].value}</S.Value>
          <S.Label>{values[1].label}</S.Label>
        </S.ContentBox>
      </S.Container>
    </S.Card>
  );
};

namespace S {
  export const Title = styled.h2`
    color: var(--foreground);
    font-size: 12px;
    letter-spacing: 0.96px;
    font-weight: 600;
    margin-bottom: 16px;
    text-transform: uppercase;
  `;

  export const Container = styled.div`
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    gap: 16px;
  `;

  export const ContentBox = styled.div`
    text-align: left;
  `;

  export const Value = styled.div<{ $large?: boolean }>`
    color: var(--gray-dark);
    font-size: ${(props) => (props.$large ? "42px" : "24px")};
  `;

  export const Label = styled.div`
    color: var(--foreground);
    font-size: 14px;
    font-weight: 400;
  `;

  export const Card = styled.div<{ $isContrast?: boolean }>`
    background: ${(props) =>
      props.$isContrast
        ? "var(--gradient-brand-reverse)"
        : "var(--gray-lightest)"};
    border-radius: 8px;
    padding: 16px;

    ${Title}, ${Value}, ${Label} {
      ${(props) => (props.$isContrast ? "color: white" : "")};
    }
  `;
}

export default Card;
