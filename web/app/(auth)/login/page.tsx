import { Suspense } from "react";
import { SuspenseLoader } from "@/components/common/suspense-loader";
import { LoginCard } from "@/components/login/login-card";

export default function LoginPage() {
  return (
    <Suspense fallback={<SuspenseLoader />}>
      <LoginCard />
    </Suspense>
  );
}
