import { FiArrowRight } from "react-icons/fi";

import { ButtonPrimary } from "@/components/Button";
import Calculator from "@/components/Calculator";
import CTA from "@/components/CTA";
import CTASmall from "@/components/CTASmall";
import SavingsGraph from "@/components/SavingsGraph";
import Spacer from "@/components/Spacer";
import { TextBase } from "@/components/Text";

const Page = () => {
  return (
    <>
      <CTA
        style={{ maxHeight: 640 }}
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

      <CTASmall heading={"Want a demo?"} hasBackground>
        <TextBase>
          Use the link below to book <br data-desktop /> a personalized demo of
          Kardinal.
        </TextBase>
        <ButtonPrimary
          analyticsId="button_calculator_cta_get_demo"
          href="https://calendly.com/d/cqhd-tgj-vmc/45-minute-meeting"
          rel="noopener noreferrer"
          target="_blank"
          size="lg"
        >
          Get a Demo
        </ButtonPrimary>
      </CTASmall>
    </>
  );
};

export default Page;
