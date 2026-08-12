import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  reactCompiler: true,

  cacheComponents: true,
  partialPrefetching: true,

  experimental: {
    turbopackRustReactCompiler: true,
    useOffline: true,
  },
  allowedDevOrigins:['192.168.1.104']
};

export default nextConfig;
