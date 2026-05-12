import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  // distDir: 'build',
  images: {
    unoptimized: true
  },
  output: 'standalone'
};

export default nextConfig;
