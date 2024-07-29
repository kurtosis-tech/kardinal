import Content from "@/components/Content";

const CanaryServices = () => {
  return (
    <Content
      heading={
        <>
          Progressively deploy with <s>a canary environment</s>
          <br />
          <em>canary services</em>
        </>
      }
    >
      Canary environments are all or nothing - as software is promoted it’s
      either in canary, or in production. With Kardinal, you can progressively
      promote new versions of a microservice from a small percentage of traffic
      to a large percentage of traffic as they prove their stability.
    </Content>
  );
};

export default CanaryServices;
