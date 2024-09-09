"use client";

import { useState } from "react";
import styled from "styled-components";

import Section from "@/components/Section";
import { ResourceRequirement } from "@/constants/calculator";
import {
  CostInterval,
  useCalculatorContext,
} from "@/context/CalcualtorContext";
import analytics from "@/lib/analytics";

import Spacer from "../Spacer";

import CalculatorInputs from "./CalculatorInputs";
import Card from "./Card";
import CardGroup from "./CardGroup";

const WORKING_HOURS_PER_COST_INTERVAL: Record<CostInterval, number> = {
  Month: 173.2, // Average Monthly Working Hours: 40 hours/week×4.33 weeks/month=173.2 hours/month
  Year: 1920, // 48 working weeks per year
};

const HOURLY_COST_PER_RESOURCE_REQUIREMENT: Record<
  ResourceRequirement,
  number
> = {
  [ResourceRequirement.MICRO]: 0.0116,
  [ResourceRequirement.SMALL]: 0.023,
  [ResourceRequirement.MEDIUM]: 0.0464,
};

const currencyFormatter = new Intl.NumberFormat("en-US", {
  style: "currency",
  currency: "USD",
  minimumFractionDigits: 2,
  maximumFractionDigits: 2,
});

const Calculator = () => {
  const { engineers, microservices, resourceRequirement, costInterval } =
    useCalculatorContext();

  const costPerServiceHour =
    HOURLY_COST_PER_RESOURCE_REQUIREMENT[resourceRequirement];

  const costBefore = microservices * engineers * costPerServiceHour;
  const costAfter = (microservices + engineers) * costPerServiceHour;
  const savings = costBefore - costAfter;

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
              value: 100 - Math.round((costAfter / costBefore) * 100) + "%",
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
