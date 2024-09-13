"use client";

import Image from "next/image";
import { useEffect, useRef, useState } from "react";
import { FiCopy, FiSend } from "react-icons/fi";
import styled from "styled-components";

import Button from "@/components/Button";
import Heading from "@/components/Heading";
import Section from "@/components/Section";
import Text from "@/components/Text";
import { mobile, tablet } from "@/constants/breakpoints";
import { useCalculatorContext } from "@/context/CalculatorContext";

const ShareKardinal = () => {
  const { engineers, microservices, savingsPercent } = useCalculatorContext();
  const [isCopied, setIsCopied] = useState(false);
  const messageRef = useRef<HTMLDivElement>(null);

  const handleCopy = () => {
    if (messageRef.current == null) {
      throw new Error("messageRef is null, expected a value");
    }
    const html = messageRef.current.innerHTML;

    const clipboardItem = new ClipboardItem({
      "text/html": new Blob([html], { type: "text/html" }),
      "text/plain": new Blob([html], { type: "text/plain" }),
    });

    navigator.clipboard.write([clipboardItem]);
    setIsCopied(true);
  };

  useEffect(() => {
    if (isCopied) {
      const timeout = setTimeout(() => {
        setIsCopied(false);
      }, 2000);
      return () => clearTimeout(timeout);
    }
  }, [isCopied]);

  const plainTextBody = `Hey,
I came across a dev tool called Kardinal that could help reduce maintenance overhead and cost with our current dev and test environment setup. Their savings calculator shows ${savingsPercent}% savings with ${engineers} engineers and ${microservices} microservices.

A few interesting callouts:
- You can spin up multiple ephemeral environments within a single cluster.
- It only deploys services you’re working on, and reuses the rest.
- Works with stateless and stateful services.

Might be worth checking out to see if it aligns with our needs: https://github.com/kurtosis-tech/kardinal
`;
  const mailtoLink = `mailto:?subject=Check out Kardinal&body=${encodeURI(plainTextBody)}`;

  return (
    <>
      <Section noPadBottomMobile>
        <S.ShareKardinal>
          <S.Content>
            <Heading.H2>
              Want to share <em>Kardinal</em>
              <br data-desktop /> with your team?
            </Heading.H2>
            <Text.Base>
              Copy the following message to send it on slack,
              <br data-desktop /> teams, or paste it into an email!
            </Text.Base>
            <S.ButtonWrapper>
              <Button.Primary
                analyticsId="button_calculator_share_email"
                iconLeft={<FiSend />}
                size="lg"
                href={mailtoLink}
              >
                Send via email
              </Button.Primary>
              <Button.Tertiary
                analyticsId="button_calculator_share_copy"
                iconRight={
                  isCopied ? undefined : (
                    <FiCopy color="var(--brand-secondary)" size={18} />
                  )
                }
                onClick={handleCopy}
              >
                {isCopied ? "Copied!" : "Copy message content"}
              </Button.Tertiary>
            </S.ButtonWrapper>
          </S.Content>
          <S.MessageWrapper>
            <S.Message>
              <S.MessageContent ref={messageRef}>
                Hey, <br />I came across a dev tool called Kardinal that could
                help reduce maintenance overhead and cost with our current dev
                and test environment setup. Their savings calculator shows{" "}
                <em>{savingsPercent}% savings</em> with{" "}
                <b>{engineers} engineers</b> and{" "}
                <b>{microservices} microservices</b>
                .
                <br />A few interesting callouts:
                <ul>
                  <li>
                    You can spin up multiple ephemeral environments within a
                    single cluster.
                  </li>
                  <li>
                    It only deploys services you’re working on, and reuses the
                    rest.
                  </li>
                  <li>Works with stateless and stateful services.</li>
                </ul>
                Might be worth checking out to see if it aligns with our needs:{" "}
                <a href="https://github.com/kurtosis-tech/kardinal">
                  https://github.com/kurtosis-tech/kardinal
                </a>
              </S.MessageContent>
              <S.BottomGradient />
              <S.WindowButtons>
                <S.WindowButton />
                <S.WindowButton />
                <S.WindowButton />
              </S.WindowButtons>
              <S.ShareArrowImage
                src="/illustrations/share-arrow.svg"
                alt="Arrow pointing to the message with a slack and microsoft teams logo"
                width={235}
                height={228}
              />
            </S.Message>
          </S.MessageWrapper>
        </S.ShareKardinal>
      </Section>
      <S.SectionDivider />
    </>
  );
};

const S = {
  ShareKardinal: styled.div`
    display: grid;
    grid-template-columns: 1fr 1fr;
    width: 100%;
    margin: 0 auto;
    padding-top: 64px;
    position: relative;

    @media ${tablet} {
      grid-template-columns: 1fr;
      grid-row-gap: 24px;
    }
  `,

  Content: styled.div`
    max-width: 616px;
    display: flex;
    flex-direction: column;
    gap: 24px;
  `,

  MessageWrapper: styled.div`
    background: linear-gradient(90deg, #fca061 2.75%, #febd3a 70.63%);
    border-top-left-radius: 24px;
    border-top-right-radius: 24px;
    padding: 24px 24px 0 24px;
    margin-right: -72px;
    width: 688px;

    @media ${mobile} {
      width: 100%;
      padding: 16px 16px 0 16px;
    }
  `,

  Message: styled.div`
    background: var(--white);
    border-top-left-radius: 12px;
    border-top-right-radius: 12px;
    padding: 24px 32px 24px 24px;
    font-size: 16px;
    font-style: normal;
    font-weight: 400;
    line-height: 26px; /* 162.5% */
    width: 100%;
    position: relative;

    @media ${mobile} {
      line-height: 20px;
    }

    ul {
      margin-left: 24px;
    }

    a {
      background: var(--gradient-brand-reverse);
      background-size: auto;
      background-clip: border-box;
      -webkit-background-clip: text;
      -webkit-text-fill-color: transparent;
      font-style: normal;
    }
  `,

  BottomGradient: styled.div`
    background: linear-gradient(
      180deg,
      rgba(255, 255, 255, 0) -7.69%,
      #fff 71.79%
    );
    height: 78px;
    width: 100%;
    position: absolute;
    bottom: 0;
    left: 0;
  `,

  WindowButtons: styled.div`
    display: flex;
    gap: 8px;
    position: absolute;
    top: 24px;
    right: 24px;
  `,

  WindowButton: styled.div`
    height: 10px;
    width: 10px;
    background: var(--gray-200);
    border-radius: 50%;
  `,

  ShareArrowImage: styled(Image)`
    position: absolute;
    top: -132px;
    left: -208px;

    @media ${tablet} {
      display: none;
    }
  `,

  ButtonWrapper: styled.div`
    display: flex;
    gap: 24px;
    @media ${mobile} {
      flex-direction: column;
      align-items: flex-start;
    }
  `,

  // dummy wrapper for copying content
  MessageContent: styled.div`
    @media ${mobile} {
      font-size: 14px;
    }
  `,

  SectionDivider: styled.div`
    height: 1px;
    width: 100%;
    max-width: 1440px;
    margin: 0 auto;
    background: linear-gradient(
      90deg,
      rgba(254, 189, 58, 0) 0.88%,
      rgba(254, 189, 58, 0) 6.77%,
      #febd3a 42.33%,
      #fca061 65.81%,
      rgba(252, 160, 97, 0) 100%
    );
  `,
};

export default ShareKardinal;
