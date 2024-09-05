import CalculatorInputs from "./CalculatorInputs";
import Card from "./Card";
import CardGroup from "./CardGroup";

const Calculator = () => {
  return (
    <>
      <CalculatorInputs />
      <CardGroup>
        <Card
          title="Your stateless costs before"
          values={[
            { label: "Monthly cost", value: "$16.10" },
            { label: "Yearly cost", value: "$30,912.00" },
          ]}
        />
        <Card
          title="Your stateless costs before"
          values={[
            { label: "Monthly cost", value: "$16.10" },
            { label: "Yearly cost", value: "$30,912.00" },
          ]}
        />
        <Card
          isContrast
          title="Your stateless costs before"
          values={[
            { label: "Monthly cost", value: "$16.10" },
            { label: "Yearly cost", value: "$30,912.00" },
          ]}
        />
      </CardGroup>
    </>
  );
};

export default Calculator;
