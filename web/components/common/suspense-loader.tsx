
export function SuspenseLoader() {
    return (
        <div className="w-full max-w-lg animate-pulse border border-white/10 bg-[oklch(0.21_0.028_264)] px-12 py-16">
          <div className="mx-auto h-16 w-64 bg-white/10" />
          <div className="mx-auto mt-5 h-9 w-72 bg-white/10" />
          <div className="mx-auto mt-4 h-10 w-full bg-white/5" />
          <div className="mt-10 h-12 w-full bg-white/10" />
        </div>
    )
}