"use client";

import { BiTimeFive, BiX } from "react-icons/bi";
import styled from "styled-components";
import { default as ReactModal } from "styled-react-modal";
import { ModalProvider } from "styled-react-modal";

import EmailCapture from "@/components/EmailCapture";
import Heading from "@/components/Heading";
import Text from "@/components/Text";
import { mobile } from "@/constants/breakpoints";
import { useModal } from "@/context/ModalContext";

const Modal = () => {
  const { isOpen, toggleModal } = useModal();

  return (
    <ModalProvider backgroundComponent={S.Background}>
      <S.Modal
        isOpen={isOpen}
        onBackgroundClick={toggleModal}
        onEscapeKeydown={toggleModal}
      >
        <S.CloseButton onClick={toggleModal}>
          <BiX size={32} />
        </S.CloseButton>
        <S.BorderTop />
        <S.WaitlistIcon>
          <BiTimeFive size={24} />
        </S.WaitlistIcon>
        <S.Content>
          <S.ModalHeading>
            Join the <em>beta</em>
          </S.ModalHeading>
          <Text.Base>
            We&apos;ll let you know whe the Kardinal beta is available. We
            won&apos;t use your email for anything else.
          </Text.Base>
          <EmailCapture buttonAnalyticsId="button_modal_join_waitlist" />
        </S.Content>
      </S.Modal>
    </ModalProvider>
  );
};

namespace S {
  export const Modal = ReactModal.styled`
    width: 100%;
    margin: 24px;
    max-width: 826px;
    display: flex;
    justify-content: center;
    align-items: center;
    background-color: var(--background);
    display: flex;
    flex-direction: column;
    gap: 24px;
    border-radius: 24px;
    padding: 104px 48px;
    position: relative;
    text-align: center;

    @media ${mobile} {
      padding: 64px 16px 32px 16px;
    }
  `;

  export const Content = styled.div`
    max-width: 484px;
    display: flex;
    flex-direction: column;
    gap: 24px;
  `;

  export const CloseButton = styled.button`
    border: 0;
    background: transparent;
    border: 1px solid transparent;
    border-radius: 4px;
    display: flex;
    align-items: center;
    justify-content: center;
    transition: border 0.1s ease-in-out;
    position: fixed;
    color: var(--white);
    top: 56px;
    right: 96px;

    &:hover {
      border: 1px solid var(--brand-primary);
      cursor: pointer;
    }

    @media ${mobile} {
      top: 24px;
      right: 24px;
    }
  `;

  export const Background = styled.div`
    display: flex;
    align-items: center;
    justify-content: center;
    position: fixed;
    top: 0;
    left: 0;
    width: 100vw;
    height: 100vh;
    z-index: 30;
    background: rgba(32, 36, 42, 0.92);
  `;

  export const Success = styled.div`
    display: flex;
    align-items: center;
    justify-content: center;
    flex-direction: column;
    text-align: center;
  `;

  export const WaitlistIcon = styled.div`
    position: absolute;
    top: -36px;
    left: calc(50% - 36px);
    width: 72px;
    height: 72px;
    border-radius: 50%;
    background-color: var(--background);
    display: flex;
    align-items: center;
    justify-content: center;
    color: var(--brand-secondary);
    filter: drop-shadow(0px 1px 2px rgba(252, 129, 73, 0.1))
      drop-shadow(0px 4px 4px rgba(252, 129, 73, 0.09))
      drop-shadow(0px 8px 5px rgba(252, 129, 73, 0.05))
      drop-shadow(0px 15px 6px rgba(252, 129, 73, 0.01))
      drop-shadow(0px 23px 7px rgba(252, 129, 73, 0));
  `;

  export const BorderTop = styled.div`
    height: 2px;
    flex-shink: 0;
    width: 100%;
    background: linear-gradient(
      90deg,
      rgba(252, 160, 97, 0) 0%,
      #fca061 50%,
      rgba(252, 160, 97, 0) 100%
    );
    position: absolute;
    top: 0;
  `;

  export const ModalHeading = styled(Heading.H2)`
    font-size: 52px;
    letter-spacing: -1.04px;

    @media ${mobile} {
      font-size: 24px;
      line-height: normal;
    }
  `;
}

export default Modal;
