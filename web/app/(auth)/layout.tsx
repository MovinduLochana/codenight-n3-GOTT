export default function AuthLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <div className="relative min-h-screen flex items-center justify-center overflow-hidden bg-[oklch(0.11_0.018_240)]">
      {/* Animated gradient orbs */}
      <div
        className="pointer-events-none absolute inset-0 overflow-hidden"
        aria-hidden="true"
      >
        <div className="auth-orb auth-orb-1" />
        <div className="auth-orb auth-orb-2" />
        <div className="auth-orb auth-orb-3" />
      </div>

      {/* Grid overlay */}
      <div className="pointer-events-none absolute inset-0 auth-grid" aria-hidden="true" />

      {/* Content */}
      <div className="relative z-10 w-full px-4 py-12">{children}</div>

      <style>{`
        .auth-orb {
          position: absolute;
          border-radius: 9999px;
          filter: blur(80px);
          opacity: 0.35;
          animation: auth-float 8s ease-in-out infinite;
        }
        .auth-orb-1 {
          width: 600px;
          height: 600px;
          background: radial-gradient(circle, oklch(0.55 0.22 150) 0%, transparent 70%);
          top: -200px;
          left: -150px;
          animation-delay: 0s;
        }
        .auth-orb-2 {
          width: 500px;
          height: 500px;
          background: radial-gradient(circle, oklch(0.50 0.18 260) 0%, transparent 70%);
          bottom: -180px;
          right: -100px;
          animation-delay: -3s;
        }
        .auth-orb-3 {
          width: 350px;
          height: 350px;
          background: radial-gradient(circle, oklch(0.60 0.20 200) 0%, transparent 70%);
          top: 40%;
          left: 60%;
          animation-delay: -6s;
        }
        @keyframes auth-float {
          0%, 100% { transform: translateY(0px) scale(1); }
          33% { transform: translateY(-30px) scale(1.05); }
          66% { transform: translateY(20px) scale(0.97); }
        }
        .auth-grid {
          background-image:
            linear-gradient(oklch(1 0 0 / 3%) 1px, transparent 1px),
            linear-gradient(90deg, oklch(1 0 0 / 3%) 1px, transparent 1px);
          background-size: 48px 48px;
        }
      `}</style>
    </div>
  );
}
