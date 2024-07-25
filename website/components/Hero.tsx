"use client";
import Image from "next/image";
import { PropsWithChildren, ReactNode } from "react";
import styled from "styled-components";

import Button from "@/components/Button";
import Heading from "@/components/Heading";
import Section from "@/components/Section";
import Text from "@/components/Text";
import { mobile } from "@/constants/breakpoints";
import { useModal } from "@/context/ModalContext";
import { scrollToId } from "@/utils";

interface Props extends PropsWithChildren {
  heading: ReactNode;
  image: string;
  buttonText?: string;
}
const Hero = ({ children, heading, image, buttonText = "See how" }: Props) => {
  const { toggleModal } = useModal();
  return (
    <Section>
      <S.Hero>
        <S.Content>
          <Heading.H1>{heading}</Heading.H1>
          <Text.Base>{children}</Text.Base>
          {buttonText && (
            <S.ButtonWrapper>
              <Button.Primary
                analyticsId="button_hero_get_started"
                onClick={() => {
                  toggleModal();
                }}
              >
                Get started
              </Button.Primary>
              <Button.Secondary
                analyticsId="button_hero_see_how"
                onClick={() => {
                  scrollToId("see_how");
                }}
              >
                {buttonText}
              </Button.Secondary>
            </S.ButtonWrapper>
          )}
        </S.Content>
        <S.Image>
          <Image src={image} alt="Hero image" width={727} height={588} />
        </S.Image>
      </S.Hero>
    </Section>
  );
};

namespace S {
  export const Hero = styled.div`
    display: grid;
    grid-template-columns: 1fr 1fr;

    @media ${mobile} {
      grid-template-columns: 1fr;
    }
  `;

  export const Content = styled.div`
    padding: 32px 0;
    display: flex;
    flex-direction: column;
    gap: 16px;
  `;

  export const Image = styled.div`
    position: relative;

    img {
      position: absolute;
      left: 0;
      bottom: -72px;
      max-width: 100%;
      z-index: 2;

      @media ${mobile} {
        position: relative;
        height: auto;
        bottom: 0;
      }
    }
  `;

  export const ButtonWrapper = styled.div`
    display: flex;
    gap: 16px;
  `;
}

export default Hero;
