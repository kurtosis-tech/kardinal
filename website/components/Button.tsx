"use client";
import Link from "next/link";
import { ButtonHTMLAttributes, ReactElement, useEffect, useState } from "react";
import { BiCheck, BiLogoGithub } from "react-icons/bi";
import styled, { css, keyframes } from "styled-components";

import analytics from "@/lib/analytics";

interface StyledProps {
  $loading?: boolean;
  $size?: "md" | "lg";
  $gradientDirection?: "left" | "right";
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

const ButtonIcon = styled.span<{
  $size?: "md" | "lg";
  $variant?: ButtonVariant;
}>`
  pointer-events: none;
  user-select: none;
  display: flex;
  align-items: center;
  justify-content: center;
  height: ${({ $size }) => ($size === "lg" ? "32px" : "24px")};
  width: ${({ $size }) => ($size === "lg" ? "32px" : "24px")};
  background-color: ${({ $variant }) =>
    $variant === "primary" ? "rgba(168, 50, 5, 0.4)" : "transparent"};
  border-radius: 50%;
`;

const primaryButtonStyles = css<StyledProps>`
  height: ${({ $size }) => ($size === "lg" ? "48px" : "40px")};
  align-items: center;
  background: ${({ $gradientDirection }) =>
    $gradientDirection === "left"
      ? "var(--gradient-brand-reverse)"
      : "var(--gradient-brand)"};
  border-radius: 54px;
  border: none;
  cursor: pointer;
  color: var(--white-100);
  display: inline-flex;
  gap: 8px;
  font-size: 16px;
  font-weight: 500;
  line-height: 28px;
  justify-content: center;
  padding: 8px;
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

const TextSpacer = styled.span``;

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
  color: var(--gray-dark);

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

type ButtonVariant = "primary" | "secondary" | "tertiary" | "codespaces";

interface Props extends ButtonHTMLAttributes<HTMLButtonElement> {
  onClick?: () => void;
  href?: string;
  iconLeft?: ReactElement<{ size: number }>;
  iconRight?: ReactElement<{ size: number }>;
  analyticsId: string;
  loading?: boolean;
  isSuccess?: boolean;
  variant?: ButtonVariant;
  Component?: any;
  target?: string;
  size?: "md" | "lg";
}

const ButtonIconImpl = ({
  success,
  loading,
  icon,
  size,
  variant,
}: {
  success?: boolean;
  loading?: boolean;
  icon: ReactElement<{ size: number }>;
  size?: "md" | "lg";
  variant?: ButtonVariant;
}) => {
  return (
    <ButtonIcon $size={size} $variant={variant} role="presentation">
      {success ? <BiCheck size={20} /> : loading ? <Spinner /> : icon}
    </ButtonIcon>
  );
};

const ButtonImpl = ({
  analyticsId,
  onClick,
  iconLeft,
  iconRight,
  loading,
  isSuccess,
  children,
  variant,
  Component,
  size,
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
      $size={size}
      $loading={loading}
      $gradientDirection={iconLeft != null ? "left" : "right"}
      disabled={loading}
      onClick={() => {
        analytics.track("BUTTON_CLICK", { analyticsId });
        onClick && onClick();
      }}
    >
      {variant === "codespaces" && <BiLogoGithub size={24} />}
      {iconLeft == null && <TextSpacer />}
      {iconLeft != null && (
        <ButtonIconImpl
          size={size}
          icon={iconLeft}
          loading={loading}
          success={success}
          variant={variant}
        />
      )}
      {children}
      {iconRight != null && (
        <ButtonIconImpl
          size={size}
          icon={iconRight}
          loading={loading}
          success={success}
          variant={variant}
        />
      )}
      {iconRight == null && <TextSpacer />}
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
