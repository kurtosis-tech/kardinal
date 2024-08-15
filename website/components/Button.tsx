"use client";
import Link from "next/link";
import { ButtonHTMLAttributes, useEffect, useState } from "react";
import { BiCheck, BiLogoGithub, BiRightArrowAlt } from "react-icons/bi";
import styled, { css, keyframes } from "styled-components";

import analytics from "@/lib/analytics";

interface StyledProps {
  $loading?: boolean;
}

const spin = keyframes`
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
`;

const Spinner = styled.div`
  border: 2px solid var(--gray-border);
  border-top: 2px solid transparent;
  border-radius: 50%;
  width: 24px;
  height: 24px;
  animation: ${spin} 0.8s ease-in-out infinite;
  position: absolute;
`;

const ButtonIcon = styled.span`
  pointer-events: none;
  user-select: none;
  display: flex;
  align-items: center;
  justify-content: center;
  height: 24px;
  width: 24px;
  background-color: rgba(168, 50, 5, 0.4);
  border-radius: 50%;
`;

const primaryButtonStyles = css<StyledProps>`
  height: 40px;
  align-items: center;
  background: var(--gradient-brand);
  border-radius: 54px;
  border: none;
  cursor: pointer;
  color: var(--white-100);
  display: inline-flex;
  gap: 12px;
  font-size: 16px;
  font-weight: 500;
  line-height: 28px;
  justify-content: center;
  padding: 8px 8px 8px 16px;
  position: relative;
  z-index: 1;
  text-decoration: none;
  transform: translateY(0);
  background-size: 100%;
  transition: all 0.2s ease-in-out;

  &:hover:not(:disabled) {
    transform: translateY(-2px);
    border: none;
    color: var(--white-100);
    background-size: 200%;
  }

  &:disabled {
    cursor: not-allowed;
  }
`;

const PrimaryButton = styled.button<StyledProps>`
  ${primaryButtonStyles}
`;

const secondaryButtonStyles = css<StyledProps>`
  color: var(--brand-primary);
  text-align: right;
  leading-trim: both;
  text-edge: cap;
  background: transparent;
  border: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 4px;
  font-size: 16px;
  font-style: normal;
  font-weight: 500;
  line-height: 28px; /* 175% */
  transform: translateY(0);
  transition: all 0.2s ease-in-out;
  font-family: inherit;

  &:hover:not(:disabled) {
    opacity: 0.8;
    transform: translateY(-2px);
    cursor: pointer;
  }
`;

const tertiaryButtonStyles = css<StyledProps>`
  ${secondaryButtonStyles}
  color: var(--foreground);

  &:hover {
    color: var(--brand-primary);
  }
`;

const SecondaryButton = styled.button<StyledProps>`
  ${secondaryButtonStyles}
`;

export const PrimaryLink = styled(Link)<StyledProps>`
  ${primaryButtonStyles}
`;

const SecondaryLink = styled(Link)<StyledProps>`
  ${secondaryButtonStyles}
`;

const TertiaryButton = styled.button<StyledProps>`
  ${tertiaryButtonStyles}
`;

const TertiaryLink = styled(Link)<StyledProps>`
  ${tertiaryButtonStyles}
`;

const CodespacesLink = styled(Link)<StyledProps>`
  color: var(--white-100);
  display: inline-flex;
  padding: 8px 12px;
  align-items: center;
  gap: 8px;
  border-radius: 12px;
  background: linear-gradient(90deg, #21262d 0%, #3d3d3d 100%);
  text-align: center;
  font-size: 16px;
  font-style: normal;
  font-weight: 400;
  line-height: 28px; /* 175% */
  transfrom: translateY(0);
  transition: all 0.2s ease-in-out;

  &:hover {
    opacity: 0.8;
    transform: translateY(-2px);
  }
`;

interface Props extends ButtonHTMLAttributes<HTMLButtonElement> {
  onClick?: () => void;
  href?: string;
  analyticsId: string;
  loading?: boolean;
  isSuccess?: boolean;
  variant?: "primary" | "secondary" | "tertiary" | "codespaces";
  Component?: any;
  target?: string;
}

const ButtonImpl = ({
  analyticsId,
  onClick,
  loading,
  isSuccess,
  children,
  variant,
  Component,
  ...buttonProps
}: Props) => {
  const [success, setSuccess] = useState(false);

  // show a checkmark for 5 seconds after success, then reset
  useEffect(() => {
    if (isSuccess) {
      setSuccess(true);
    }

    const delay = setTimeout(() => {
      setSuccess(false);
    }, 5000);

    return () => {
      clearTimeout(delay);
    };
  }, [isSuccess]);
  return (
    <Component
      {...buttonProps}
      $loading={loading}
      disabled={loading}
      onClick={() => {
        analytics.track("BUTTON_CLICK", { analyticsId });
        onClick && onClick();
      }}
    >
      {variant === "codespaces" && <BiLogoGithub size={24} />}
      {children}
      {variant === "primary" && (
        <ButtonIcon role="presentation">
          {success ? (
            <BiCheck size={20} />
          ) : loading ? (
            <Spinner />
          ) : (
            <BiRightArrowAlt size={20} />
          )}
        </ButtonIcon>
      )}
      {variant === "tertiary" && <BiRightArrowAlt size={24} />}
    </Component>
  );
};

export const ButtonPrimary = (props: Props) => (
  <ButtonImpl
    {...props}
    variant="primary"
    Component={props.href != null ? PrimaryLink : PrimaryButton}
  />
);

export const ButtonSecondary = (props: Props) => (
  <ButtonImpl
    {...props}
    variant="secondary"
    Component={props.href != null ? SecondaryLink : SecondaryButton}
  />
);

export const ButtonTertiary = (props: Props) => (
  <ButtonImpl
    {...props}
    variant="tertiary"
    Component={props.href != null ? TertiaryLink : TertiaryButton}
  />
);

export const ButtonCodespaces = (props: Props) => (
  <ButtonImpl {...props} variant="codespaces" Component={CodespacesLink} />
);

const Button = {
  Primary: ButtonPrimary,
  Secondary: ButtonSecondary,
  Tertiary: ButtonTertiary,
  Codespaces: ButtonCodespaces,
};

export default Button;
