"use client";

import styled from "styled-components";

import Heading from "@/components/Heading";
import Section from "@/components/Section";
import Text from "@/components/Text";

const SavingsSection = () => {
  return (
    <Section>
      <S.SavingsSection>
        <Heading.H2>
          See how much <em>you can save</em> <br data-desktop /> with Kardinal
          💸
        </Heading.H2>
        <Text.Base>
          Replace your dev sandboxes with Kardinal, 
          <br data-desktop /> and see how much money you could save.
        </Text.Base>
        <S.CalculatorPlaceholder>
          <S.Columns>
            <div>Number of engineers</div>
            <div>numver of services</div>
            <div>calculate</div>
          </S.Columns>
        </S.CalculatorPlaceholder>
        <S.PlaceholderFooter>Potential savings:</S.PlaceholderFooter>
      </S.SavingsSection>
    </Section>
  );
};

namespace S {
  export const SavingsSection = styled.div`
    display: flex;
    flex-direction: column;
    gap: 16px;
    width: 100%;
    align-items: center;
    justify-content: center;
    text-align: center;
  `;

  export const CalculatorPlaceholder = styled.div`
    background: var(--gradient-brand);
    width: 100%;
    padding: 32px;
    color: white;
    border-radius: 21px 21px 0px 0px;
  `;

  export const PlaceholderFooter = styled.div`
    border-radius: 0px 0px 21px 21px;
    background: rgba(168, 50, 5, 0.4);
  `;

  export const Columns = styled.div`
    display: grid;
    grid-template-columns: 2fr 2fr 1fr;
    text-align: left;
  `;
}

export default SavingsSection;
