import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  /* config options here */
  reactCompiler: true,
  images: {
    remotePatterns: [
      new URL("https://images.metmuseum.org/CRDImages/*/original/*"),
      new URL("https://images.metmuseum.org/CRDImages/*/web-large/*"),
    ],
  },
};

export default nextConfig;
