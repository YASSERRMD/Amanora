export default function Home() {
  return (
    <main className="min-h-screen bg-slate-950 px-6 py-8 text-slate-50">
      <section className="mx-auto flex min-h-[calc(100vh-4rem)] max-w-6xl flex-col justify-center">
        <p className="text-sm font-semibold uppercase tracking-[0.18em] text-cyan-300">
          Amanora
        </p>
        <h1 className="mt-4 max-w-4xl text-5xl font-semibold tracking-tight text-white md:text-7xl">
          Autonomous data governance for enterprise metadata, risk, and policy.
        </h1>
        <p className="mt-6 max-w-2xl text-lg leading-8 text-slate-300">
          Discover sources, classify sensitive data, map lineage, apply retention
          policies, and audit governance decisions from one platform.
        </p>
      </section>
    </main>
  );
}
