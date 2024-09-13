"use client";

import { useEffect } from "react";
import { FiChevronDown } from "react-icons/fi";
import styled from "styled-components";

import { tablet } from "@/constants/breakpoints";
import { ResourceRequirement } from "@/constants/calculator";
import {
  CostInterval,
  useCalculatorContext,
} from "@/context/CalculatorContext";
import analytics from "@/lib/analytics";

const resourceRequirementsOptions: ResourceRequirement[] = [
  ResourceRequirement.MICRO,
  ResourceRequirement.SMALL,
  ResourceRequirement.MEDIUM,
];

const costIntervalOptions: CostInterval[] = ["Year", "Month"];

const CostSavingsCalculator = () => {
  const {
    engineers,
    setEngineers,
    microservices,
    setMicroservices,
    resourceRequirement,
    setResourceRequirement,
    costInterval,
    setCostInterval,
  } = useCalculatorContext();

  const trackCalculate = () => {
    analytics.track("CALCULATE", {
      numEngineers: engineers,
      numServices: microservices,
      costInterval,
      resourceRequirement,
    });
  };

  useEffect(() => {
    trackCalculate();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [costInterval, resourceRequirement]);

  return (
    <S.Wrapper>
      <S.Columns>
        <S.SliderContainer>
          <S.SliderLabel>
            Number of engineers on your organization using dev sandboxes:
          </S.SliderLabel>
          <S.Slider
            value={engineers}
            onChange={(e) => setEngineers(parseInt(e.target.value))}
            onMouseUp={trackCalculate}
            onTouchEnd={trackCalculate}
          />
          <S.SliderValue $value={engineers}>{engineers}</S.SliderValue>
        </S.SliderContainer>

        <S.SliderContainer>
          <S.SliderLabel>
            Number of microservices in your architecture:
          </S.SliderLabel>
          <S.Slider
            value={microservices}
            onChange={(e) => setMicroservices(parseInt(e.target.value))}
            onMouseUp={trackCalculate}
            onTouchEnd={trackCalculate}
          />
          <S.SliderValue role="presentation" $value={microservices}>
            {microservices}
          </S.SliderValue>
        </S.SliderContainer>
      </S.Columns>

      <S.Columns>
        <S.SelectContainer>
          <S.SliderLabel>
            Average resource requirement per service:
          </S.SliderLabel>
          <S.Select
            value={resourceRequirement}
            onChange={(e) =>
              setResourceRequirement(e.target.value as ResourceRequirement)
            }
          >
            {resourceRequirementsOptions.map((o) => (
              <option key={o} value={o}>
                {o}
              </option>
            ))}
          </S.Select>
          <S.Chevron size={20} role="presentation" />
        </S.SelectContainer>

        <S.SelectContainer>
          <S.SliderLabel>Show costs:</S.SliderLabel>
          <S.Select
            value={costInterval}
            onChange={(e) => setCostInterval(e.target.value as CostInterval)}
          >
            {costIntervalOptions.map((o) => (
              <option key={o} value={o}>
                {o}ly
              </option>
            ))}
          </S.Select>
          <S.Chevron size={20} role="presentation" />
        </S.SelectContainer>
      </S.Columns>
    </S.Wrapper>
  );
};

namespace S {
  export const Columns = styled.div`
    display: grid;
    grid-template-columns: 1fr 1fr;
    grid-gap: 32px;

    @media ${tablet} {
      grid-template-columns: 1fr;
    }
  `;

  export const Wrapper = styled.div`
    background-color: white;
    border-radius: 8px;
    border: 1px solid var(--gray-100);
    box-shadow: 0px -4px 22px -12px rgba(33, 38, 45, 0.12);
    padding: 24px;
    width: 100%;
    margin-bottom: 24px;
    display: flex;
    flex-direction: column;
    gap: 52px;
  `;

  export const SliderContainer = styled.div`
    position: relative;
    display: flex;
    flex-direction: column;
    gap: 16px;

    @media ${tablet} {
      margin-bottom: 32px;
    }
  `;

  export const SliderLabel = styled.label`
    display: block;
    font-size: 14px;
  `;

  export const Slider = styled.input.attrs({
    type: "range",
    min: "2",
    max: "100",
  })`
    appearance: none;
    -webkit-appearance: none;
    height: 4px;
    width: 100%;
    border-radius: 2px;
    background: var(--brand-secondary);
    outline: none;
    &::-webkit-slider-thumb,
    &::-moz-range-thumb {
      -webkit-appearance: none;
      appearance: none;
      width: 20px;
      height: 20px;
      border-radius: 50%;
      background: var(--white);
      border: 1px solid var(--brand-secondary);
      cursor: pointer;
    }
  `;

  export const SliderValue = styled.span<{ $value: number }>`
    display: inline-flex;
    align-items: center;
    justify-content: center;
    position: absolute;
    /* This math strange but exists to approximately map the percentage width
     * of the slider to the position of the slider control, which otherwise
     * over-extends the ends of the slider (so its total range is over 100% of
     * the width).
     */
    left: calc(${(props) => props.$value * 0.98}% - 16px);
    bottom: -48px;
    width: 32px;
    height: 32px;
    line-height: 30px;
    text-align: center;
    background-color: var(--gray-600);
    color: white;
    border-radius: 12px;

    &:before {
      content: "";
      display: block;
      width: 0px;
      height: 0px;
      border-style: solid;
      border-width: 0 6px 6px 6px;
      border-color: transparent transparent var(--gray-600) transparent;
      transform: rotate(0deg);
      position: absolute;
      top: -5px;
    }
  `;

  export const SelectContainer = styled.div`
    position: relative;
    display: flex;
    flex-direction: column;
    gap: 8px;
  `;

  export const Select = styled.select`
    width: 100%;
    border: 1px solid var(--gray-400);
    border-radius: 8px;
    font-size: 14px;
    display: flex;
    height: 48px;
    padding: 12px 16px;
    justify-content: center;
    align-items: flex-start;
    align-self: stretch;
    border: 1px solid var(--gray-200);
    background-color: var(--white);
    -moz-appearance: none; /* Firefox */
    -webkit-appearance: none; /* Safari and Chrome */
    appearance: none;
  `;

  export const Chevron = styled(FiChevronDown)`
    position: absolute;
    right: 16px;
    top: calc(50% + 4px);
    color: var(--gray-400);
  `;
}

export default CostSavingsCalculator;
