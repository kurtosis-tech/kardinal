"use client";

import { useState } from "react";
import { FiArrowRight, FiChevronDown } from "react-icons/fi";
import styled from "styled-components";

import { ButtonPrimary } from "@/components/Button";

type CostInterval = "Yearly" | "Monthly" | "Weekly";
type ResourceRequirement =
  | "1vCPU, 2GB RAM (t2 small)"
  | "2vCPU, 4GB RAM (t2 medium)"
  | "4vCPU, 8GB RAM (t2 large)";

const CostSavingsCalculator = () => {
  const [engineers, setEngineers] = useState<number>(35);
  const [microservices, setMicroservices] = useState<number>(20);
  const [resourceRequirement, setResourceRequirement] =
    useState<ResourceRequirement>("1vCPU, 2GB RAM (t2 small)");
  const [costType, setCostType] = useState<CostInterval>("Yearly");

  return (
    <S.Card>
      <S.Columns>
        <S.SliderContainer>
          <S.SliderLabel>
            Number of engineers on your organization using dev sandboxes:
          </S.SliderLabel>
          <S.Slider
            min="0"
            max="100"
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
            min="0"
            max="100"
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
              setResourceRequirement(
                e.target.value as unknown as ResourceRequirement,
              )
            }
          >
            <option>1vCPU, 2GB RAM (t2 small)</option>
            <option>2vCPU, 4GB RAM (t2 medium)</option>
            <option>4vCPU, 8GB RAM (t2 large)</option>
          </S.Select>
          <S.Chevron size={20} role="presentation" />
        </S.SelectContainer>

        <S.SelectContainer>
          <S.SliderLabel>Show costs:</S.SliderLabel>
          <S.Select
            value={costType}
            onChange={(e) =>
              setCostType(e.target.value as unknown as CostInterval)
            }
          >
            <option>Yearly</option>
            <option>Monthly</option>
            <option>Weekly</option>
          </S.Select>
          <S.Chevron size={20} role="presentation" />
        </S.SelectContainer>
      </S.Columns>

      <S.ButtonWrapper>
        <ButtonPrimary
          analyticsId={"calculator_calculate"}
          iconRight={<FiArrowRight />}
        >
          Calculate!
        </ButtonPrimary>
      </S.ButtonWrapper>
    </S.Card>
  );
};

namespace S {
  export const Columns = styled.div`
    display: grid;
    grid-template-columns: 1fr 1fr;
    grid-gap: 16px;
  `;

  export const Card = styled.div`
    background-color: white;
    border-radius: 8px;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    padding: 24px;
    width: 100%;
    margin: 20px auto;
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

  export const Slider = styled.input.attrs({ type: "range" })`
    width: 100%;
    -webkit-appearance: none;
    height: 4px;
    border-radius: 2px;
    background: #f0f0f0;
    outline: none;
    &::-webkit-slider-thumb {
      -webkit-appearance: none;
      appearance: none;
      width: 20px;
      height: 20px;
      border-radius: 50%;
      background: #ff7f50;
      cursor: pointer;
    }
  `;

  export const SliderValue = styled.span<{ $value: number }>`
    display: inline-flex;
    align-items: center;
    justify-content: center;
    position: absolute;
    left: calc(${(props) => props.$value}% - 14px);
    bottom: -40px;
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
