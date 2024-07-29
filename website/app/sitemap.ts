import type { MetadataRoute } from "next";

export default function sitemap(): MetadataRoute.Sitemap {
  return [
    {
      url: "https://kardinal.dev",
      lastModified: new Date(),
      changeFrequency: "weekly",
      priority: 1,
    },
    {
      url: "https://kardinal.dev/build",
      lastModified: new Date(),
      changeFrequency: "weekly",
      priority: 0.1,
    },
  ];
}
