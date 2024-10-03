import { Metadata } from "next";

import Calculator from "@/components/Calculator";
import CTA from "@/components/CTA";
import CTADemo from "@/components/CTADemo";
import SavingsGraph from "@/components/SavingsGraph";
import Spacer from "@/components/Spacer";
import { TextBase } from "@/components/Text";

export const metadata: Metadata = {
  title: "Kardinal | Calculator",
  description: "Calculate exactly how much your team could save with Kardinal",
};

const Page = () => {
  return (
    <>
      <CTA
        style={{ maxHeight: 540, minHeight: 540 }}
        imageUrl={null}
        buttonText={null}
        fullHeight
        heading={
          <>
            Save <em>~90%</em> on dev sandbox costs with <em>Kardinal</em>
          </>
        }
      >
        <TextBase>
          Check out our calculator below to see <br data-desktop />{" "}
          <em>exactly how much</em> your team could save.
        </TextBase>
      </CTA>

      <Calculator />

      <Spacer height={48} />

      <SavingsGraph />

      <CTADemo />
    </>
  );
};

export default Page;
