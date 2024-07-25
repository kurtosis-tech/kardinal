import { Metadata } from "next";

import LandingPageTemplate from "@/components/LandingPageTemplate";
import Sparkle from "@/components/Sparkle";

export const metadata: Metadata = {
  title: "Kardinal",
  description: "Develop in prod, fearlessly.",
};

const Page = () => {
  return (
    <LandingPageTemplate
      heading={
        <>
          Develop in prod
          <br />
          <em>
            Fearlessly
            <Sparkle />
          </em>
        </>
      }
    >
      Develop with production data, services,
      <br data-desktop /> or traffic without the risk.
    </LandingPageTemplate>
  );
};

export default Page;
