"use client";

import { FiArrowRight, FiChevronDown } from "react-icons/fi";
import styled from "styled-components";

import { ButtonPrimary } from "@/components/Button";
import {
  CostInterval,
  ResourceRequirement,
  useCalculatorContext,
} from "@/context/CalcualtorContext";

interface Props {
  onCalculate: () => void;
}

const CostSavingsCalculator = ({ onCalculate }: Props) => {
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

  const resourceRequirementsOptions: ResourceRequirement[] = [
    "1vCPU, 2GB RAM (t2 small)",
    "2vCPU, 4GB RAM (t2 medium)",
    "4vCPU, 8GB RAM (t2 large)",
  ];

  const costIntervalOptions: CostInterval[] = ["Yearly", "Monthly", "Weekly"];

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
          />
          <S.SliderValue $value={engineers}>{engineers}</S.SliderValue>
        </S.SliderContainer>

        <S.SliderContainer>
          <S.SliderLabel>
            Number of stateless microservices in your architecture:
          </S.SliderLabel>
          <S.Slider
            value={microservices}
            onChange={(e) => setMicroservices(parseInt(e.target.value))}
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
                {o}
              </option>
            ))}
          </S.Select>
          <S.Chevron size={20} role="presentation" />
        </S.SelectContainer>
      </S.Columns>

      <S.ButtonWrapper>
        <ButtonPrimary
          analyticsId={"calculator_calculate"}
          iconRight={<FiArrowRight />}
          onClick={onCalculate}
        >
          Calculate!
        </ButtonPrimary>
      </S.ButtonWrapper>
    </S.Wrapper>
  );
};

namespace S {
  export const Columns = styled.div`
    display: grid;
    grid-template-columns: 1fr 1fr;
    grid-gap: 16px;
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
    gap: 24px;
  `;

  export const SliderContainer = styled.div`
    margin-bottom: 20px;
    position: relative;
  `;

  export const SliderLabel = styled.label`
    display: block;
    margin-bottom: 8px;
    font-size: 14px;
    color: #555;
  `;

  export const Slider = styled.input.attrs({
    type: "range",
    min: "0",
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
    left: calc(${(props) => props.$value * 0.962}% - 5px);
    bottom: -44px;
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
    margin-bottom: 20px;
    position: relative;
  `;

  export const Select = styled.select`
    width: 100%;
    border: 1px solid #ddd;
    border-radius: 8px;
    font-size: 14px;
    display: flex;
    height: 48px;
    padding: 12px 16px;
    justify-content: center;
    align-items: flex-start;
    align-self: stretch;
    border: 1px solid var(--gray-200, #e5e7eb);
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

  export const ButtonWrapper = styled.div`
    display: flex;
    align-items: center;
    justify-content: flex-end;
  `;
}

export default CostSavingsCalculator;
