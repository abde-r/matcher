import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  /* config options here */
  env: {
    API_URL: process.env.API_URL,
  },
  images: {
    domains: ['cdn.intra.42.fr', 'media.licdn.com'],
  },
};

export default nextConfig;
