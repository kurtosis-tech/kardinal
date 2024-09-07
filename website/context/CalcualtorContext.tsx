"use client";

import { useSearchParams } from "next/navigation";
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
}

const CalculatorContext = createContext<CalculatorContextProps | undefined>(
  undefined,
);

export const CalculatorProvider = ({ children }: PropsWithChildren) => {
  const searchParams = useSearchParams();

  const initialEngineers: number =
    searchParams.get("engineers") != null
      ? parseInt(searchParams.get("engineers") || "20")
      : 20;

  const initialMicroservices: number =
    searchParams.get("services") != null
      ? parseInt(searchParams.get("services") || "60")
      : 60;

  const [engineers, setEngineers] = useState<number>(
    Math.min(initialEngineers, 100),
  ); // max value 100
  const [microservices, setMicroservices] = useState<number>(
    Math.min(initialMicroservices, 100),
  ); // max value 100
  const [resourceRequirement, setResourceRequirement] =
    useState<ResourceRequirement>(ResourceRequirement.MICRO);
  const [costInterval, setCostInterval] = useState<CostInterval>("Year");

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
