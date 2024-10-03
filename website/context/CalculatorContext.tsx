"use client";

import {
  createContext,
  Dispatch,
  PropsWithChildren,
  SetStateAction,
  useContext,
  useState,
} from "react";

import { ResourceRequirement } from "@/constants/calculator";

export type CostInterval = "Year" | "Month";

interface CalculatorContextProps {
  engineers: number;
  setEngineers: Dispatch<SetStateAction<number>>;
  microservices: number;
  setMicroservices: Dispatch<SetStateAction<number>>;
  resourceRequirement: ResourceRequirement;
  setResourceRequirement: Dispatch<SetStateAction<ResourceRequirement>>;
  costInterval: CostInterval;
  setCostInterval: Dispatch<SetStateAction<CostInterval>>;
  costBefore: number;
  costAfter: number;
  savings: number;
  savingsPercent: number;
}

const CalculatorContext = createContext<CalculatorContextProps | undefined>(
  undefined,
);

const HOURLY_COST_PER_RESOURCE_REQUIREMENT: Record<
  ResourceRequirement,
  number
> = {
  [ResourceRequirement.MICRO]: 0.0116,
  [ResourceRequirement.SMALL]: 0.023,
  [ResourceRequirement.MEDIUM]: 0.0464,
};

export const CalculatorProvider = ({ children }: PropsWithChildren) => {
  const searchParams = new URLSearchParams(
    typeof window !== "undefined" ? window.location.search : "",
  );

  const initialEngineers: number =
    searchParams.get("engineers") != null
      ? parseInt(searchParams.get("engineers") || "20")
      : 20;

  const initialMicroservices: number =
    searchParams.get("services") != null
      ? parseInt(searchParams.get("services") || "60")
      : 60;

  const [engineers, setEngineers] = useState<number>(
    Math.min(Math.max(initialEngineers, 2), 100),
  ); // value from 2 to 100
  const [microservices, setMicroservices] = useState<number>(
    Math.min(Math.max(initialMicroservices, 2), 100),
  ); // max value 100
  const [resourceRequirement, setResourceRequirement] =
    useState<ResourceRequirement>(ResourceRequirement.MICRO);
  const [costInterval, setCostInterval] = useState<CostInterval>("Year");

  const costPerServiceHour =
    HOURLY_COST_PER_RESOURCE_REQUIREMENT[resourceRequirement];

  const costBefore = microservices * engineers * costPerServiceHour;
  const costAfter = (microservices + engineers) * costPerServiceHour;
  const savings = costBefore - costAfter;
  const savingsPercent = 100 - Math.round((costAfter / costBefore) * 100);

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
        costBefore,
        costAfter,
        savings,
        savingsPercent,
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
