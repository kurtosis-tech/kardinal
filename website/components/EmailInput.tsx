import { useEffect } from "react";
import styled from "styled-components";

import Button from "@/components/Button";
import { mobile } from "@/constants/breakpoints";

const EmailInput = ({
  value,
  onChange,
  buttonAnalyticsId,
  isLoading,
  isSuccess,
}: {
  value: string;
  // eslint-disable-next-line no-unused-vars
  onChange: (v: string) => void;
  buttonAnalyticsId: string;
  isLoading?: boolean;
  isSuccess?: boolean;
}) => {
  // clean value on success
  useEffect(() => {
    if (isSuccess) {
      onChange("");
    }
  }, [isSuccess, onChange]);

  return (
    <S.Fieldset>
      <S.Input
        type="email"
        value={value}
        onChange={(e) => onChange(e.target.value)}
        placeholder="Your email"
      />
      <S.SubmitButton
        analyticsId={buttonAnalyticsId}
        type="submit"
        loading={isLoading}
        isSuccess={isSuccess}
      >
        Get a demo
      </S.SubmitButton>
    </S.Fieldset>
  );
};

namespace S {
  export const Fieldset = styled.fieldset`
    align-items: center;
    background: var(--background);
    border-radius: 56px;
    border: 0;
    border: 1px solid var(--gray-border);
    display: flex;
    flex-direction: row;
    gap: 8px;
    text-align: left;
    width: 100%;
    max-width: 402px;
    padding: 8px 10px;

    @media ${mobile} {
      flex-direction: column;
      border: 0;
      background: transparent;
    }
  `;

  export const Input = styled.input`
    height: 40px;
    width: 100%;
    font-size: 16px;
    background-color: var(--background);
    border: 0;
    border-radius: 56px;
    color: var(--foreground);
    padding-left: 8px;
    outline: 2px solid transparent;
    transition: outline 0.1s ease-in-out;

    &:focus {
      outline: 2px solid var(--brand-primary);
    }

    &::placeholder {
      /* Chrome, Firefox, Opera, Safari 10.1+ */
      color: var(--foreground-light);
      opacity: 1; /* Firefox */
    }

    &:-ms-input-placeholder {
      /* Internet Explorer 10-11 */
      color: var(--foreground-light);
    }

    &::-ms-input-placeholder {
      /* Microsoft Edge */
      color: var(--foreground-light);
    }

    @media ${mobile} {
      padding-left: 16px;
      border: 1px solid var(--gray-border);
      width: 100%;
    }
  `;

  export const SubmitButton = styled(Button.Primary)`
    flex-shrink: 0;
    @media ${mobile} {
      width: 100%;
    }
  `;
}

export default EmailInput;
