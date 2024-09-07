"use client";

import { useState } from "react";
import { FiArrowDown, FiArrowRight } from "react-icons/fi";
import styled from "styled-components";

import Heading from "@/components/Heading";
import Section from "@/components/Section";
import Text from "@/components/Text";
import { tablet } from "@/constants/breakpoints";
import analytics from "@/lib/analytics";

const SavingsSection = () => {
  const [engineers, setEngineers] = useState(20);
  const [microservices, setMicroservices] = useState(60);

  return (
    <Section>
      <S.SavingsSection>
        <Heading.H2>
          See how much <em>you can save</em> <br data-desktop /> with Kardinal
        </Heading.H2>
        <Text.Base>
          Replace your dev sandboxes with Kardinal,
          <br data-desktop /> and see how much money you could save.
        </Text.Base>
        <S.CalculatorPlaceholder>
          <S.Columns>
            <S.InputWrapper>
              <label>
                Number of engineers using dev sandboxes in your organization:
              </label>
              <S.Input
                value={engineers}
                onChange={(e) => setEngineers(parseInt(e.target.value))}
              />
            </S.InputWrapper>
            <S.InputWrapper>
              <label>Number of microservices in your architecture:</label>
              <S.Input
                value={microservices}
                onChange={(e) => setMicroservices(parseInt(e.target.value))}
              />
            </S.InputWrapper>

            <S.ButtonWrapper>
              <S.CalculateButton
                href={`/calculator?engineers=${engineers}&services=${microservices}`}
                onClick={() => {
                  analytics.track("BUTTON_CLICK", {
                    analyticsId: "button_landingpage_calculate",
                  });
                  analytics.track("CALCULATE", {
                    numEngineers: engineers,
                    numServices: microservices,
                  });
                }}
              >
                Calculate!
                <S.CalculateButtonIcon role="presentation">
                  <FiArrowRight />
                </S.CalculateButtonIcon>
              </S.CalculateButton>
            </S.ButtonWrapper>
          </S.Columns>
          <S.PlaceholderFooter>
            Potential savings:
            <S.SavingsAmount>{"$26,726.40"}</S.SavingsAmount>
            <S.SavingsPercentage>
              <FiArrowDown size={16} />
              ~93%
            </S.SavingsPercentage>
          </S.PlaceholderFooter>
        </S.CalculatorPlaceholder>
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
    max-width: 827px;
    margin: 0 auto;
  `;

  export const CalculatorPlaceholder = styled.div`
    background: var(--gradient-brand-reverse);
    width: 100%;
    color: white;
    border-radius: 21px;
    margin-top: 32px;
  `;

  export const PlaceholderFooter = styled.div`
    background: rgba(168, 50, 5, 0.4);
    margin-top: -16px;
    width: 100%;
    border-radius: 0px 0px 21px 21px;
    padding: 8px 32px;
    display: flex;
    gap: 12px;
    align-items: center;
    color: rgba(255, 255, 255, 0.8);
    font-size: 12px;
    font-style: normal;
    font-weight: 600;
    line-height: normal;
    letter-spacing: 0.96px;
    text-transform: uppercase;
    @media ${tablet} {
      flex-direction: column;
    }
  `;

  export const SavingsAmount = styled.span`
    color: var(--white);
    font-size: 32px;
    font-style: normal;
    font-weight: 500;
    line-height: normal;
  `;

  export const SavingsPercentage = styled.span`
    display: flex;
    align-items: center;
    gap: 4px;
  `;

  export const Columns = styled.div`
    display: grid;
    grid-template-columns: 2fr 2fr 1fr;
    grid-column-gap: 24px;
    grid-row-gap: 16px;
    text-align: left;
    padding: 32px;
    @media ${tablet} {
      grid-template-columns: 1fr;
    }
  `;

  export const InputWrapper = styled.div`
    display: flex;
    flex-direction: column;
    gap: 8px;
  `;

  export const Input = styled.input.attrs({ type: "number" })`
    display: flex;
    padding: 12px 16px;
    justify-content: center;
    align-items: flex-start;
    align-self: stretch;
    border-radius: 8px;
    background: rgba(168, 50, 5, 0.3);
    border: none;
    -moz-appearance: textfield;
    outline: 2px solid transparent;
    transition: outline 0.1s ease-in-out;
    color: var(--white);
    font-size: 16px;
    font-style: normal;
    font-weight: 400;
    height: 48px;
    line-height: 24px; /* 171.429% */

    &::-webkit-outer-spin-button,
    &::-webkit-inner-spin-button {
      -webkit-appearance: none;
      margin: 0;
    }

    &:focus {
      outline: 2px solid var(--white);
    }
  `;

  export const CalculateButton = styled.a`
    height: 48px;
    background: var(--white);
    color: var(--brand-primary);
    display: flex;
    height: 47px;
    justify-content: center;
    align-items: center;
    gap: 12px;
    border: none;
    border-radius: 8px;
    text-align: right;
    leading-trim: both;
    text-edge: cap;
    font-size: 16px;
    font-style: normal;
    font-weight: 500;
    line-height: 28px; /* 175% */
    padding: 8px;
    display: flex;
    gap: 12px;
    width: 100%;
    tranform: translateY(0);
    transition: transform 0.2s ease-in-out;

    &:hover {
      transform: translateY(-2px);
      cursor: pointer;
    }
  `;

  export const CalculateButtonIcon = styled.span`
    height: 24px;
    width: 24px;
    background: var(--brand-primary);
    color: white;
    border-radius: 50%;
    display: flex;
    align-items: center;
    justify-content: center;
  `;

  export const ButtonWrapper = styled.div`
    width: 100%;
    height: 100%;
    display: flex;
    align-items: flex-end;
  `;
}

export default SavingsSection;
