"use client";

import {
  createContext,
  Dispatch,
  PropsWithChildren,
  SetStateAction,
  useContext,
  useState,
} from "react";

export type ResourceRequirement =
  | "1vCPU, 2GB RAM (t2 small)"
  | "2vCPU, 4GB RAM (t2 medium)"
  | "4vCPU, 8GB RAM (t2 large)";

export type CostInterval = "Yearly" | "Monthly" | "Weekly";

interface CalculatorContextProps {
  engineers: number;
  setEngineers: Dispatch<SetStateAction<number>>;
  microservices: number;
  setMicroservices: Dispatch<SetStateAction<number>>;
  resourceRequirement: ResourceRequirement;
  setResourceRequirement: Dispatch<SetStateAction<ResourceRequirement>>;
  costInterval: CostInterval;
  setCostInterval: Dispatch<SetStateAction<CostInterval>>;
}

const CalculatorContext = createContext<CalculatorContextProps | undefined>(
  undefined,
);

export const CalculatorProvider = ({ children }: PropsWithChildren) => {
  const [engineers, setEngineers] = useState<number>(35);
  const [microservices, setMicroservices] = useState<number>(20);
  const [resourceRequirement, setResourceRequirement] =
    useState<ResourceRequirement>("1vCPU, 2GB RAM (t2 small)");
  const [costInterval, setCostInterval] = useState<CostInterval>("Yearly");

  return (
    <CalculatorContext.Provider
      value={{
        engineers,
        setEngineers,
        microservices,
        setMicroservices,
        resourceRequirement,
        setResourceRequirement,
        costInterval,
        setCostInterval,
      }}
    >
      {children}
    </CalculatorContext.Provider>
  );
};

export const useCalculatorContext = () => {
  const context = useContext(CalculatorContext);
  if (!context) {
    throw new Error(
      "useCalculatorContext must be used within a CalculatorProvider",
    );
  }
  return context;
};
