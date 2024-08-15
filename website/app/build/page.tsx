import { Metadata } from "next";

import Sparkles from "@/components/icons/Sparkles";
import LandingPageTemplate from "@/components/LandingPageTemplate";

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
            <Sparkles />
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
