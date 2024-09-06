"use client";

import Section from "@/components/Section";

import CalculatorInputs from "./CalculatorInputs";
import Card from "./Card";
import CardGroup from "./CardGroup";

const Calculator = () => {
  const handleCalculate = () => {
    console.log("calculate!");
  };

  return (
    <Section>
      <CalculatorInputs onCalculate={handleCalculate} />
      <CardGroup>
        <Card
          title="Your stateless costs before"
          values={[
            { label: "Services cost before (per year)", value: "$30,912.00" },
            { label: "Services cost after (per hour)", value: "$16.10" },
          ]}
        />
        <Card
          title="Your stateless costs before"
          values={[
            { label: "Services cost after (per year)", value: "$2,428.80" },
            { label: "Services cost after (per hour)", value: "$1.26" },
          ]}
        />
        <Card
          isContrast
          title="Your stateless costs before"
          values={[
            { label: "Percentage of previous cloud costs", value: "92%" },
            { label: "Cost savings per year*", value: "$28,483.20" },
          ]}
        />
      </CardGroup>
    </Section>
  );
};

export default Calculator;
