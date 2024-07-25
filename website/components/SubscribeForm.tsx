import { ChangeEvent, FormEvent, useState } from "react";
import styled from "styled-components";

import Button from "@/components/Button";

// const checkboxInputs = [
//   { label: "Cloud cost reduction", value: "reduce_cloud_costs" },
//   {
//     label: "Reducing number of pre-prod environments",
//     value: "reduce_preprod_environments",
//   },
//   { label: "Maturity-based access controls", value: "maturity_controls" },
//   {
//     label: "Shorter dev and test cycles using the production environment",
//     value: "develop_on_prod",
//   },
// ];
//
// const radioInputs = [
//   { label: "Just when Kardinal is ready.", value: "when_its_ready" },
//   {
//     label:
//       "Sooner! I have some critical workflow and I want to beta test Kardinal.",
//     value: "sooner",
//   },
//   {
//     label:
//       "Right now! I want to provide feedback and help to tailor Kardinal to my needs.",
//     value: "right_now",
//   },
// ];

export interface FormData {
  email: string;
  commPreference: string;
  features: string[];
}

// Mailchimp formats fields in all caps, and can only accept strings
// These form fields are defined by Mailchimp. They can be edited here:
// https://us2.admin.mailchimp.com/lists/settings/merge-tags/

export interface MailchimpFormData {
  EMAIL: string;
  COMMS: string;
  FEATURES: string;
  UTM_SOURCE: string;
  UTM_MEDIUM: string;
  UTM_CAMP: string; // abbreviated because too long to be a field name in Mailchimp
  UTM_TERM: string;
  UTM_CONT: string; // abbreviated because too long to be a field name in Mailchimp
  PAGE_PATH: string;
}

const SubscribeForm = ({
  status,
  onValidated,
}: {
  status: "sending" | "error" | "success" | null;
  // eslint-disable-next-line no-unused-vars
  onValidated: (formData: FormData) => void;
}) => {
  const [email, setEmail] = useState("");
  const features = ["NOT_COLLECTED"];
  // const [features, setFeatures] = useState<string[]>([]);
  const commPreference = "NOT_COLLECTED";
  // const [commPreference, setCommPreference] = useState<string>("");

  // const handleFeatureChange = useCallback(
  //   (event: ChangeEvent<HTMLInputElement>) => {
  //     const feature = event.target.value;
  //     setFeatures((features) => {
  //       if (event.target.checked) {
  //         // Add the checked feature to the state
  //         return [...features, feature];
  //       } else {
  //         // Remove the unchecked feature from the state
  //         return features.filter((f) => f !== feature);
  //       }
  //     });
  //   },
  //   [],
  // );

  const handleEmailChange = (event: ChangeEvent<HTMLInputElement>) => {
    setEmail(event.target.value);
  };

  const handleSubmit = (e: FormEvent) => {
    e.preventDefault();
    email.includes("@") &&
      onValidated({
        email,
        commPreference,
        features,
      });
  };

  return (
    <S.Form onSubmit={handleSubmit}>
      <S.Fieldset>
        <S.Input
          type="email"
          value={email}
          onChange={handleEmailChange}
          placeholder="Your email"
        />
      </S.Fieldset>
      <Button.Primary
        analyticsId="button_modal_join_waitlist"
        type="submit"
        loading={status === "sending"}
      >
        Join the beta
      </Button.Primary>
    </S.Form>
  );
};
//       <Divider />
//       <S.Fieldset>
//         <Text.Base style={{ marginBottom: 8 }}>
//           What Kardinal features are most relevant to you?
//         </Text.Base>
//         {checkboxInputs.map((input) => (
//           <S.Label key={input.value}>
//             <input
//               type="checkbox"
//               value={input.value}
//               checked={features.includes(input.value)}
//               onChange={handleFeatureChange}
//             />
//             {input.label}
//           </S.Label>
//         ))}
//       </S.Fieldset>
//       <S.Fieldset>
//         <Text.Base style={{ marginBottom: 8 }}>Reach out...</Text.Base>
//         {radioInputs.map((input) => (
//           <S.Label key={input.value}>
//             <input
//               type="radio"
//               value={input.value}
//               checked={commPreference === input.value}
//               onChange={() => setCommPreference(input.value)}
//             />
//             {input.label}
//           </S.Label>
//         ))}
//       </S.Fieldset>
//
// const Divider = () => (
//   <S.Divider>
//     <S.DividerText>Optional</S.DividerText>
//   </S.Divider>
// );

namespace S {
  export const Divider = styled.div`
    height: 1px;
    background-color: var(--gray-border);
    margin-top: 12px;
    position: relative;
  `;
  export const DividerText = styled.span`
    background-color: var(--background);
    padding: 8px;
    position: absolute;
    top: -16px;
    left: calc(50% - 32px);
  `;
  export const SubscribeForm = styled.div`
    display: flex;
    flex-direction: column;
  `;

  export const Input = styled.input`
    height: 44px;
    width: 100%;
    font-size: 18px;
    background-color: var(--background);
    color: var(--foreground);
    border: 1px solid var(--gray-border);
    border-radius: 4px;
    padding: 8px;
    margin-top: 8px;

    &::placeholder {
      /* Chrome, Firefox, Opera, Safari 10.1+ */
      color: var(--foreground);
      opacity: 1; /* Firefox */
    }

    &:-ms-input-placeholder {
      /* Internet Explorer 10-11 */
      color: var(--foreground);
    }

    &::-ms-input-placeholder {
      /* Microsoft Edge */
      color: var(--foreground);
    }
  `;

  export const Select = styled.select`
    height: 44px;
    width: 100%;
    font-size: 18px;
    background-color: var(--background);
    color: var(--foreground);
    border: 1px solid var(--foreground);
    border-radius: 4px;
    padding: 8px;
    margin-top: 8px;
  `;

  export const Form = styled.form`
    display: flex;
    flex-direction: column;
    gap: 24px;
  `;

  export const Fieldset = styled.fieldset`
    text-align: left;
    border: 0;
    display: flex;
    flex-direction: column;
  `;

  export const Label = styled.label`
    cursor: pointer;
    display: flex;
    gap: 8px;
    margin-top: 8px;
  `;
}
export default SubscribeForm;
