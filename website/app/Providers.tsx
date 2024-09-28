import { ReactNode } from "react";

import { CalculatorProvider } from "@/context/CalculatorContext";
import { ModalProvider } from "@/context/ModalContext";
import { VotingProvider } from "@/context/VotingContext";

import StyledComponentsRegistry from "../lib/registry";

interface Props {
  children: ReactNode;
}

const Providers = ({ children }: Props) => {
  return (
    <ModalProvider>
      <VotingProvider>
        <CalculatorProvider>
          <StyledComponentsRegistry>{children}</StyledComponentsRegistry>
        </CalculatorProvider>
      </VotingProvider>
    </ModalProvider>
  );
};

export default Providers;
