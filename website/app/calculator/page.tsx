import Card from "@/components/Card";
import CardGroup from "@/components/CardGroup";
import CTA from "@/components/CTA";
import Section from "@/components/Section";
import { TextBase } from "@/components/Text";

const Page = () => {
  return (
    <>
      <CTA
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
        <CardGroup>
          <Card
            title="Your stateless costs before"
            values={[
              { label: "Monthly cost", value: "$16.10" },
              { label: "Yearly cost", value: "$30,912.00" },
            ]}
          />
          <Card
            title="Your stateless costs before"
            values={[
              { label: "Monthly cost", value: "$16.10" },
              { label: "Yearly cost", value: "$30,912.00" },
            ]}
          />
          <Card
            isContrast
            title="Your stateless costs before"
            values={[
              { label: "Monthly cost", value: "$16.10" },
              { label: "Yearly cost", value: "$30,912.00" },
            ]}
          />
        </CardGroup>
      </Section>
    </>
  );
};
export default Page;
