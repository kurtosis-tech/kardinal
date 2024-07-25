"use client";
import { useState } from "react";
import MailchimpSubscribe from "react-mailchimp-subscribe";
import styled from "styled-components";

import EmailInput from "@/components/EmailInput";
import { MailchimpFormData } from "@/components/SubscribeForm";
import Text from "@/components/Text";
import analytics from "@/lib/analytics";

const EmailCapture = ({
  buttonAnalyticsId,
  // secondaryButtonAnalyticsId,
}: {
  buttonAnalyticsId: string;
  // secondaryButtonAnalyticsId: string;
}) => {
  const [email, setEmail] = useState("");

  const handleFormSubmit = (
    // eslint-disable-next-line no-unused-vars
    subscribe: (d: MailchimpFormData) => void,
  ) => {
    // if (!email.includes("@")) {
    //   return;
    // }
    // log some analytics in segment
    // try catch this because people use adblockers, which will cause this call to fail
    // eslint-disable-next-line no-unused-vars
    try {
      analytics.track("FORM_SUBMIT", {
        email,
        formType: "waitlist", // in case we ever have multiple forms
      });
      // we use anonymous ID and report email as an attribute. see:
      // https://segment.com/docs/connections/spec/identify/#anonymous-id
      analytics.identify({ email });
    } catch (error) {
      console.error("Segment analytics failed:", error);
    }

    const utmParams = analytics.getUtmParams();

    // do the actual mailchimp submission
    subscribe({
      EMAIL: email,
      COMMS: "NOT_COLLECTED", // this form does not have these inputs
      FEATURES: "NOT_COLLECTED", // this form does not have these inputs
      UTM_SOURCE: utmParams.utm_source || "NOT_SET",
      UTM_MEDIUM: utmParams.utm_medium || "NOT_SET",
      UTM_CAMP: utmParams.utm_campaign || "NOT_SET",
      UTM_TERM: utmParams.utm_term || "NOT_SET",
      UTM_CONT: utmParams.utm_content || "NOT_SET",
      PAGE_PATH: window.location.pathname,
    });
  };

  return (
    <S.EmailCapture>
      <MailchimpSubscribe<MailchimpFormData>
        url={process.env.NEXT_PUBLIC_MAILCHIMP_POST_URL!}
        render={({ subscribe, status, message }) => {
          return (
            <>
              <S.Form
                onSubmit={(e) => {
                  e.preventDefault();
                  handleFormSubmit(subscribe);
                }}
              >
                <EmailInput
                  value={email}
                  onChange={setEmail}
                  buttonAnalyticsId={buttonAnalyticsId}
                  isLoading={status === "sending"}
                  isSuccess={status === "success"}
                />
                {status === "error" && (
                  <S.ErrorMessage>
                    {message?.toString() === "0 - Please enter a value"
                      ? "Please enter a valid email address"
                      : message?.toString() ||
                        "An unknown error occurred. Please check that an adblocker isn't interfering and try again."}
                  </S.ErrorMessage>
                )}
                {
                  // @ts-ignore
                  status === "success" && (
                    <S.SuccessMessage>
                      Success! You’ve been added to our list of beta users. Keep
                      an eye on your inbox!
                    </S.SuccessMessage>
                  )
                }
              </S.Form>
            </>
          );
        }}
      />
    </S.EmailCapture>
  );
};

namespace S {
  export const EmailCapture = styled.div`
    display: flex;
    flex-direction: column;
    gap: 24px;
    align-items: center;
    width: 100%;
  `;

  export const Form = styled.form`
    display: flex;
    flex-direction: column;
    gap: 24px;
    align-items: center;
    width: 100%;
  `;

  export const SuccessMessage = styled(Text.Small)`
    color: #249728;
  `;

  export const ErrorMessage = styled(Text.Small)`
    color: indianred;
  `;
}

export default EmailCapture;
