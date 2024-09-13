"use client";

import Image from "next/image";
import { PropsWithChildren } from "react";
import { FiArrowRight } from "react-icons/fi";
import styled from "styled-components";

import Button from "@/components/Button";
import Heading from "@/components/Heading";
import Text from "@/components/Text";
import { tablet } from "@/constants/breakpoints";
import kardiZoomImage from "@/public/illustrations/kardi-zoom.svg";

const GetStarted = ({ children }: PropsWithChildren) => {
  return (
    <S.GetStarted id="get-started">
      <S.Content>
        <Heading.H2>
          Get started <em>now</em>
        </Heading.H2>
        <Text.Base>
          Kardinal is easy to install, and easy to uninstall: Copy the commands
          on the right to get it working on your local environment!
        </Text.Base>
        <S.ButtonWrapper>
          <Button.Secondary
            analyticsId="button_get_started_read_docs"
            iconRight={<FiArrowRight size={18} />}
            href="/docs"
          >
            Read the docs
          </Button.Secondary>
        </S.ButtonWrapper>
      </S.Content>
      <S.CodeWrapper>
        {children}
        <S.KardiZoomImage
          src={kardiZoomImage}
          width={197}
          height={103}
          alt="kardinal logo with gradient line"
        />
      </S.CodeWrapper>
    </S.GetStarted>
  );
};

const S = {
  GetStarted: styled.div`
    display: flex;
    justify-content: space-between;
    @media ${tablet} {
      flex-direction: column;
      align-items: center;
    }
  `,

  Content: styled.div`
    display: flex;
    flex-direction: column;
    gap: 16px;
    width: 50%;
    max-width: 471px;
    flex-shrink: 0;
    @media ${tablet} {
      width: 100%;
    }
  `,

  CodeWrapper: styled.div`
    max-width: 511px;
    position: relative;
    margin-top: 32px;
    text-align: left;
    @media ${tablet} {
      max-width: 100%;
    }
  `,

  KardiZoomImage: styled(Image)`
    position: absolute;
    height: 103px;
    width: 197px;
    left: -150px;
    top: 40px;

    @media ${tablet} {
      display: none;
    }
  `,

  ButtonWrapper: styled.div`
    display: flex;
    align-items: center;
    justify-content: flex-start;

    @media ${tablet} {
      justify-content: center;
    }
  `,
};

export default GetStarted;
