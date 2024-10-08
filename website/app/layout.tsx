import { GoogleTagManager } from "@next/third-parties/google";
import type { Metadata } from "next";
import { Inter } from "next/font/google";
import { ReactNode } from "react";

import Footer from "@/components/Footer";
import Main from "@/components/Main";
import Modal from "@/components/Modal";
import Nav from "@/components/Nav";
import SegmentAnalytics from "@/components/SegmentAnalytics";

import GlobalStyles from "./GlobalStyles";
import Providers from "./Providers";

const inter = Inter({
  subsets: ["latin"],
  display: "block",
  variable: "--font-sans",
  weight: ["300", "400", "500", "600", "700"],
  preload: true,
});

export const metadata: Metadata = {
  title: "Kardinal",
  description: "The lightest-weight k8s dev environments in the world",
  metadataBase: new URL("https://kardinal.dev"),
  alternates: {
    canonical: "./",
  },
};

export default function RootLayout({
  children,
}: Readonly<{
  children: ReactNode;
}>) {
  return (
    <html lang="en" style={{ fontFamily: "sans-serif" }}>
      <head>
        <link
          rel="icon"
          href="/icon?<generated>"
          type="image/<generated>"
          sizes="<generated>"
        />
        <link
          rel="apple-touch-icon"
          href="/apple-icon?<generated>"
          type="image/<generated>"
          sizes="<generated>"
        />
        <link rel="icon" href="/favicon.ico" sizes="any" />
      </head>
      <SegmentAnalytics />
      <GoogleTagManager gtmId={process.env.NEXT_PUBLIC_GTM_ID!} />
      <body className={inter.className}>
        <Providers>
          <GlobalStyles />
          <Nav />
          <Main>{children}</Main>
          <Footer />
          <Modal />
        </Providers>
      </body>
    </html>
  );
}
