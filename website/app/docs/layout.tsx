import { Metadata } from "next";
import { PropsWithChildren } from "react";
import Head from "next/head";

import DocsLayout from "@/components/DocsLayout";

export const metadata: Metadata = {
  title: "Kardinal Docs",
  description: "The lightest-weight Kubernetes dev environments in the world",
};

const Layout = ({ children }: PropsWithChildren) => {
  return (
    <>
      <Head>
        <link rel="canonical" href="https://kardinal.dev/docs" />
      </Head>
      <DocsLayout>{children}</DocsLayout>
    </>
  );
};

export default Layout;
