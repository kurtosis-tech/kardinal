import { FiCalendar } from "react-icons/fi";

import { ButtonPrimary } from "@/components/Button";
import Calculator from "@/components/Calculator";
import CTA from "@/components/CTA";
import CTASmall from "@/components/CTASmall";
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
            Kardinal helps you <em>save money</em> on your infrastructure
          </>
        }
      >
        <TextBase>
          Replace your dev sandboxes with Kardinal <br data-desktop /> and see
          how much money you could save.
        </TextBase>
      </CTA>

      <Calculator />

      <Spacer height={48} />

      <CTASmall heading={"Want a demo?"}>
        <TextBase>
          Use the link below to book <br data-desktop /> a personalized demo of
          Kardinal.
        </TextBase>
        <ButtonPrimary
          analyticsId="button_calculator_cta_get_demo"
          href="https://calendly.com/d/cqhd-tgj-vmc/45-minute-meeting"
          rel="noopener noreferrer"
          target="_blank"
          iconLeft={<FiCalendar size={18} />}
          size="lg"
        >
          Get a Demo
        </ButtonPrimary>
      </CTASmall>

      <Spacer height={256} />
    </>
  );
};
export default Page;
