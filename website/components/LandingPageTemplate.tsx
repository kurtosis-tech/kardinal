import { ReactNode } from "react";

import Content from "@/components/Content";
import CTA from "@/components/CTA";
import CTASmall from "@/components/CTASmall";
import EmailCapture from "@/components/EmailCapture";
import { TextBase } from "@/components/Text";
import VideoStepper from "@/components/VideoStepper";
// assets
import architectureDiagram from "@/public/illustrations/architecture-diagram.svg";
import architectureDiagramMobile from "@/public/illustrations/architecture-diagram-mobile.svg";
import trafficFlow from "@/public/illustrations/traffic-flow.svg";
import trafficFlowMobile from "@/public/illustrations/traffic-flow-mobile.svg";

const LandingPageTemplate = ({
  heading,
  children,
  videoStepperVariant,
  trafficFlowImg,
  trafficFlowImgMobile,
  trafficFlowHeading,
  trafficFlowContent,
  videoStepperHeading,
  videoStepperContent,
  iAmTheLordOfTheRings,
}: {
  heading: ReactNode;
  children?: ReactNode;
  videoStepperVariant?: "v1" | "v2" | "v5";
  trafficFlowImg?: string;
  trafficFlowImgMobile?: string;
  trafficFlowHeading?: ReactNode;
  trafficFlowContent?: ReactNode;
  videoStepperHeading?: ReactNode;
  videoStepperContent?: ReactNode;
  iAmTheLordOfTheRings?: boolean;
}) => {
  return (
    <>
      <CTA imageUrl={null} buttonText={null} fullHeight heading={heading}>
        <TextBase>{children}</TextBase>
        <EmailCapture buttonAnalyticsId="button_hero_join_waitlist" />
      </CTA>
      <Content
        negativeTopOffset
        imageUrl={trafficFlowImg || trafficFlow}
        mobileImageUrl={trafficFlowImgMobile || trafficFlowMobile}
        heading={
          trafficFlowHeading || (
            <>
              Dev <em>on prod</em>
            </>
          )
        }
      >
        {trafficFlowContent ||
          "Kardinal injects production data and service dependencies into your dev and test workflows safely and securely."}
      </Content>

      <VideoStepper
        variation={videoStepperVariant || "v2"}
        heading={
          videoStepperHeading || (
            <>
              Learn about our <em>isolation layer</em>
            </>
          )
        }
      >
        {videoStepperContent ||
          "Kardinal uses traffic flow controls and a data isolation layer to protect production while you're developing:"}
      </VideoStepper>

      <Content
        contrast
        column
        padTop
        padBottom
        fullWidth
        heading={
          <>
            Easy to install,{" "}
            <em>
              easy to uninstall <br data-desktop />
            </em>
          </>
        }
        buttonText={null}
        buttonAnalyticsId={null}
        fullWidthImageUrl={architectureDiagram}
        mobileFullWidthImageUrl={architectureDiagramMobile}
      >
        Kardinal uses a sidecar - control plane architecture. Just drop sidecars
        next to your services in your environment and configure your traffic
        flows in the control plane.
      </Content>

      <CTASmall heading={heading} myPrecious={iAmTheLordOfTheRings}>
        <TextBase>{children}</TextBase>
        <EmailCapture buttonAnalyticsId="button_footer_join_waitlist" />
      </CTASmall>
    </>
  );
};

export default LandingPageTemplate;
