import { Metadata } from "next";

import LandingPageTemplate from "@/components/LandingPageTemplate";
import trafficFlowImgMobile from "@/public/illustrations/traffic-flow-mobile-optimize.svg";
import trafficFlowImg from "@/public/illustrations/traffic-flow-optimize.svg";

export const metadata: Metadata = {
  title: "Kardinal",
  description: "The lightest-weight k8s dev environments in the world.",
};

const Page = () => {
  return (
    <LandingPageTemplate
      heading={
        <>
          The <em>lightest-weight</em> Kubernetes dev environments in the world.
          <br />
        </>
      }
      trafficFlowImg={trafficFlowImg}
      trafficFlowImgMobile={trafficFlowImgMobile}
      trafficFlowHeading={
        <>
          Don&apos;t duplicate - <em>consolidate</em> your pre-production
          clusters.
        </>
      }
      trafficFlowContent={
        <>
          Run it all in a single cluster with Kardinal&apos;s{" "}
          <em>open-source</em>, hyper-lightweight multitenancy framework.
        </>
      }
      videoStepperHeading={
        <>
          Isolate development and QA workflows within <em>one cluster</em>
        </>
      }
      videoStepperContent={
        'Kardinal creates logical "environments" within one cluster to isolate traffic and data access between dev, test, and staging workflows.'
      }
    >
      Stop duplicating everything in your Kubernetes clusters.{" "}
      <br data-desktop />
      Deploy the <em>minimum resources necessary</em> to dev and test
      <br data-desktop />
      directly in one production-like environment.
    </LandingPageTemplate>
  );
};

export default Page;
