import { ReactNode } from "react";
import { FiCalendar } from "react-icons/fi";

import { ButtonPrimary } from "@/components/Button";
import CodeBlock from "@/components/CodeBlock";
import Content from "@/components/Content";
import CTA from "@/components/CTA";
import CTAButtons from "@/components/CTAButtons";
import CTASmall from "@/components/CTASmall";
import GetStarted from "@/components/GetStarted";
import SavingsSection from "@/components/SavingsSection";
import Spacer from "@/components/Spacer";
import { TextBase } from "@/components/Text";
import VideoStepper from "@/components/VideoStepper";
// assets
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
}) => {
  return (
    <>
      <CTA imageUrl={null} buttonText={null} fullHeight heading={heading}>
        <TextBase>{children}</TextBase>
        <CTAButtons />
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

      <div id="get-started" />
      <Content
        contrast
        column
        padTop
        padBottom
        fullWidth
        heading={null}
        buttonText={null}
        buttonAnalyticsId={null}
        childrenWrapper="div"
      >
        <GetStarted />
      </Content>

      <SavingsSection />

      <CTASmall heading={"Want a demo?"} hasBackground>
        <TextBase>
          Use the link below to book <br data-desktop /> a personalized demo of
          Kardinal.
        </TextBase>
        <ButtonPrimary
          analyticsId="button_cta_get_demo"
          href=""
          rel="noopener noreferrer"
          target="_blank"
          iconLeft={<FiCalendar size={18} />}
          size="lg"
        >
          Get a Demo
        </ButtonPrimary>
      </CTASmall>
    </>
  );
};

export default LandingPageTemplate;
