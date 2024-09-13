import CodeSnippet from "./CodeSnippet.mdx";
import GetStarted from "./GetStarted";

// This wrapper is required because the MDX file cannot be rendered in a client
// component, it must be SSR
const GetStartedWrapper = () => {
  return (
    <GetStarted>
      <CodeSnippet />
    </GetStarted>
  );
};

export default GetStartedWrapper;
