import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  reactCompiler: true,

  cacheComponents: true,
  // partialPrefetching: true,

  experimental: {
    turbopackRustReactCompiler: true,
    useOffline: true,
    optimizePackageImports: ["@uiw/react-codemirror", "lucide-react"],
  },
};

export default nextConfig;
