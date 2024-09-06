import Calculator from "@/components/Calculator";
import CTA from "@/components/CTA";
import Section from "@/components/Section";
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
      <Section>
        <Calculator />
      </Section>
      <Spacer height={256} />
    </>
  );
};
export default Page;
