"use client";

import Image from "next/image";
import { FiArrowRight } from "react-icons/fi";
import styled from "styled-components";

import { ButtonPrimary } from "@/components/Button";
import Heading from "@/components/Heading";
import Section from "@/components/Section";
import Text from "@/components/Text";
import { tablet } from "@/constants/breakpoints";
import savingsGraphImg from "@/public/illustrations/savings-graph.svg";

const SavingsGraph = () => {
  return (
    <Section padTop padBottom>
      <S.SavingsGraph>
        <Image
          src={savingsGraphImg}
          alt="Savings graph"
          style={{ maxWidth: "100%", height: "auto" }}
        />
        <S.Content>
          <Heading.H2>
            Don&apos;t duplicate - <em>consolidate</em> your pre-production
            clusters.
          </Heading.H2>
          <Text.Base>
            Replace your dev sandboxes with Kardinal and get easier DX and lower
            costs.
          </Text.Base>
          <div>
            <ButtonPrimary
              analyticsId={"button_calculator_get_started"}
              href="/docs"
              iconRight={<FiArrowRight />}
            >
              Get started
            </ButtonPrimary>
          </div>
        </S.Content>
      </S.SavingsGraph>
      <small>
        * Graph values are approximate. Based on use case with 20 microservices.
      </small>
    </Section>
  );
};

namespace S {
  export const SavingsGraph = styled.div`
    display: grid;
    grid-template-columns: 1fr 1fr;
    grid-gap: 64px;
    margin-bottom: 24px;
    @media ${tablet} {
      grid-template-columns: 1fr;
    }
  `;

  export const Content = styled.div`
    display: flex;
    gap: 24px;
    flex-direction: column;
  `;
}

export default SavingsGraph;
