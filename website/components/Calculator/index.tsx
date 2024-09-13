"use client";

import styled from "styled-components";

import Section from "@/components/Section";
import {
  CostInterval,
  useCalculatorContext,
} from "@/context/CalculatorContext";

import Spacer from "../Spacer";

import CalculatorInputs from "./CalculatorInputs";
import Card from "./Card";
import CardGroup from "./CardGroup";

const WORKING_HOURS_PER_COST_INTERVAL: Record<CostInterval, number> = {
  Month: 173.2, // Average Monthly Working Hours: 40 hours/week×4.33 weeks/month=173.2 hours/month
  Year: 1920, // 48 working weeks per year
};

const currencyFormatter = new Intl.NumberFormat("en-US", {
  style: "currency",
  currency: "USD",
  minimumFractionDigits: 2,
  maximumFractionDigits: 2,
});

const Calculator = () => {
  const { costInterval, costBefore, costAfter, savings, savingsPercent } =
    useCalculatorContext();

  return (
    <Section>
      <S.Title>
        {"put in your organization numbers to see cost savings 👇🏻"}
      </S.Title>
      <CalculatorInputs />
      <CardGroup>
        <Card
          title="Your costs before"
          values={[
            {
              label: `Services cost before (per ${costInterval.toLowerCase()})`,
              value: currencyFormatter.format(
                costBefore * WORKING_HOURS_PER_COST_INTERVAL[costInterval],
              ),
            },
            {
              label: "Services cost before (per hour)",
              value: currencyFormatter.format(costBefore),
            },
          ]}
        />
        <Card
          title="Your costs after"
          values={[
            {
              label: `Services cost after (per ${costInterval.toLowerCase()})`,
              value: currencyFormatter.format(
                costAfter * WORKING_HOURS_PER_COST_INTERVAL[costInterval],
              ),
            },
            {
              label: "Services cost after (per hour)",
              value: currencyFormatter.format(costAfter),
            },
          ]}
        />
        <Card
          isContrast
          title="Your cost savings"
          values={[
            {
              label: "Percentage of previous cloud costs saved",
              value: `${savingsPercent}%`,
            },
            {
              label: `Cost savings per ${costInterval.toLowerCase()}*`,
              value: currencyFormatter.format(
                savings * WORKING_HOURS_PER_COST_INTERVAL[costInterval],
              ),
            },
          ]}
        />
      </CardGroup>
      <Spacer height={24} />
      <small>
        * Assumes dev sandboxes are up 8hrs/day Mon-Fri, 48 weeks/year
      </small>
    </Section>
  );
};

namespace S {
  export const Title = styled.h2`
    color: var(--foreground);
    font-size: 12px;
    font-style: normal;
    font-weight: 600;
    line-height: normal;
    letter-spacing: 0.96px;
    text-transform: uppercase;
    margin-bottom: 24px;
  `;
}

export default Calculator;
