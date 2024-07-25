"use client";

import { PropsWithChildren, ReactNode } from "react";
import { BiLogoKubernetes, BiShield, BiSliderAlt } from "react-icons/bi";
import styled from "styled-components";

import Heading from "@/components/Heading";
import Section from "@/components/Section";
import Text from "@/components/Text";
import { mobile } from "@/constants/breakpoints";

export const DefaultContent = () => {
  return (
    <>
      <FeatureCard title="Seamless Integration" icon={<BiLogoKubernetes />}>
        Works out of the box with your Kubernetes cluster inside of your
        favorite cloud provider (AWS, GCP, Azure).
      </FeatureCard>
      <FeatureCard title="Safely and Rapidly Deploy" icon={<BiShield />}>
        Each engineer on your team can safely deploy to prod as fast as they can
        complete their work - speeding up feature releases immensely.
      </FeatureCard>
      <FeatureCard title="Configurable Stability" icon={<BiSliderAlt />}>
        You can control stability by configuring the amount of production
        traffic that new software versions get as they mature, as well as their
        access to production data.
      </FeatureCard>
    </>
  );
};

const Features = ({ children }: PropsWithChildren) => {
  return (
    <Section>
      <S.Features>{children || <DefaultContent />}</S.Features>
    </Section>
  );
};

export const FeatureCard = ({
  children,
  title,
  icon,
}: PropsWithChildren<{ title: string; icon: ReactNode }>) => {
  return (
    <S.Card>
      <S.Icon>{icon}</S.Icon>
      <Heading.H3>{title}</Heading.H3>
      <Text.Small>{children}</Text.Small>
    </S.Card>
  );
};

namespace S {
  export const Features = styled.div`
    display: flex;
    gap: 16px;
    display: grid;
    grid-template-columns: repeat(3, minmax(400px, 1fr));
    grid-column-gap: 16px;

    @media ${mobile} {
      grid-template-columns: 1fr;
    }
  `;

  export const Card = styled.div`
    display: flex;
    padding: 16px 16px 32px 16px;
    flex-direction: column;
    align-items: flex-start;
    gap: 8px;
    align-self: stretch;
    border-radius: 6px;
    border: 1px solid var(--gray-border, rgba(48, 54, 61, 0.4));
    background: var(--gradient-from-bg);
  `;

  export const Icon = styled.div`
    width: 52px;
    height: 52px;
    background: var(--gray-card, #21262d);
    border-radius: 50%;
    display: flex;
    align-items: center;
    justify-content: center;

    svg {
      width: 24px;
      height: 24px;
      flex-shrink: 0;
    }
  `;
}

export default Features;
