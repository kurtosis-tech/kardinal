import withMDX from "@next/mdx";
import remarkGfm from "remark-gfm"; // github flavored markdown

/** @type {import('next').NextConfig} */
const config = {
  distDir: "out",
  pageExtensions: ["mdx", "ts", "tsx"],
  compiler: {
    styledComponents: {
      ssr: true,
      displayName: true,
    },
  },
  output: "export",

  // If localhost proxy is required (e.g. for CORS), uncomment this
  // async rewrites() {
  //   return [
  //     {
  //       source: "/api/voting-api-proxy/:path*",
  //       destination: "https://voting.kardinal.dev:9111/:path*", // Proxy to API
  //     },
  //   ];
  // },
  webpack(config, { isServer }) {
    const prefix = config.assetPrefix ?? config.basePath ?? "";
    config.module.rules.push({
      test: /\.mp4$/,
      use: [
        {
          loader: "file-loader",
          options: {
            publicPath: `${prefix}/_next/static/media/`,
            outputPath: `${isServer ? "../" : ""}static/media/`,
            name: "[name].[hash].[ext]",
          },
        },
      ],
    });

    return config;
  },
};

export default withMDX({
  // Add markdown plugins here, as desired
  options: {
    remarkPlugins: [remarkGfm],
    rehypePlugins: [],
  },
})(config);
