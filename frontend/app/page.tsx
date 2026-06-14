export default function Home() {
  const metrics = [
    { label: "Data sources", value: "18", detail: "4 scanning" },
    { label: "Cataloged assets", value: "12.8k", detail: "96% owned" },
    { label: "PII findings", value: "432", detail: "37 high risk" },
    { label: "Policy pass rate", value: "91%", detail: "+6% this week" },
  ];

  const activities = [
    ["PostgreSQL warehouse", "Schema discovery", "Running"],
    ["Customer exports", "PII classification", "Queued"],
    ["Finance mart", "Retention check", "Passed"],
    ["HR archive", "Owner validation", "Needs review"],
  ];

  return (
    <main className="min-h-screen bg-[#f5f7fb] text-[#172033]">
      <header className="border-b border-[#d9e0ea] bg-white">
        <div className="mx-auto flex max-w-7xl items-center justify-between px-6 py-4">
          <div>
            <p className="text-xl font-semibold">Amanora</p>
            <p className="text-sm text-[#657187]">Autonomous data governance</p>
          </div>
          <nav className="hidden items-center gap-6 text-sm font-medium text-[#4f5e75] md:flex">
            <a href="#dashboard">Dashboard</a>
            <a href="#catalog">Catalog</a>
            <a href="#policies">Policies</a>
            <a href="#audit">Audit</a>
          </nav>
        </div>
      </header>

      <section
        id="dashboard"
        className="mx-auto grid max-w-7xl gap-6 px-6 py-8 lg:grid-cols-[1.35fr_0.65fr]"
      >
        <div className="space-y-6">
          <div className="rounded-lg border border-[#d9e0ea] bg-white p-6 shadow-sm">
            <div className="flex flex-col justify-between gap-4 md:flex-row md:items-start">
              <div>
                <p className="text-sm font-semibold uppercase text-[#0f766e]">
                  Governance overview
                </p>
                <h1 className="mt-3 max-w-3xl text-3xl font-semibold tracking-normal text-[#172033] md:text-5xl">
                  Discover, classify, govern, and audit enterprise data.
                </h1>
              </div>
              <div className="rounded-md border border-[#cfd8e6] bg-[#f8fafc] px-4 py-3 text-sm">
                <p className="font-semibold text-[#172033]">API status</p>
                <p className="mt-1 text-[#657187]">Ready for `/healthz`</p>
              </div>
            </div>

            <div className="mt-8 grid gap-4 md:grid-cols-4">
              {metrics.map((metric) => (
                <div
                  key={metric.label}
                  className="rounded-lg border border-[#d9e0ea] bg-[#fbfcfe] p-4"
                >
                  <p className="text-sm text-[#657187]">{metric.label}</p>
                  <p className="mt-2 text-2xl font-semibold text-[#172033]">
                    {metric.value}
                  </p>
                  <p className="mt-1 text-sm text-[#0f766e]">{metric.detail}</p>
                </div>
              ))}
            </div>
          </div>

          <div className="grid gap-6 lg:grid-cols-2">
            <section
              id="catalog"
              className="rounded-lg border border-[#d9e0ea] bg-white p-5 shadow-sm"
            >
              <div className="flex items-center justify-between">
                <h2 className="text-lg font-semibold">Discovery activity</h2>
                <span className="rounded-md bg-[#e6f4f1] px-3 py-1 text-sm font-medium text-[#0f766e]">
                  Live
                </span>
              </div>
              <div className="mt-5 divide-y divide-[#e6ebf2]">
                {activities.map(([source, task, status]) => (
                  <div key={source} className="grid grid-cols-3 gap-3 py-3 text-sm">
                    <span className="font-medium text-[#172033]">{source}</span>
                    <span className="text-[#657187]">{task}</span>
                    <span className="text-right font-medium text-[#334155]">
                      {status}
                    </span>
                  </div>
                ))}
              </div>
            </section>

            <section
              id="policies"
              className="rounded-lg border border-[#d9e0ea] bg-white p-5 shadow-sm"
            >
              <h2 className="text-lg font-semibold">Policy coverage</h2>
              <div className="mt-5 space-y-4">
                {[
                  ["Ownership", "96%"],
                  ["Retention", "84%"],
                  ["PII controls", "91%"],
                  ["Lineage mapped", "72%"],
                ].map(([label, value]) => (
                  <div key={label}>
                    <div className="flex justify-between text-sm">
                      <span className="text-[#4f5e75]">{label}</span>
                      <span className="font-medium">{value}</span>
                    </div>
                    <div className="mt-2 h-2 rounded-full bg-[#e7edf5]">
                      <div
                        className="h-2 rounded-full bg-[#2563eb]"
                        style={{ width: value }}
                      />
                    </div>
                  </div>
                ))}
              </div>
            </section>
          </div>
        </div>

        <aside
          id="audit"
          className="rounded-lg border border-[#d9e0ea] bg-white p-5 shadow-sm"
        >
          <h2 className="text-lg font-semibold">Risk queue</h2>
          <div className="mt-5 space-y-4">
            {[
              ["Customer email exports", "High", "Missing steward"],
              ["Passport sample columns", "High", "Policy violation"],
              ["Legacy HR archive", "Medium", "Retention review"],
              ["Analytics sandbox", "Low", "Owner confirmed"],
            ].map(([asset, level, reason]) => (
              <div
                key={asset}
                className="rounded-lg border border-[#e0e7f0] bg-[#fbfcfe] p-4"
              >
                <div className="flex items-center justify-between gap-3">
                  <p className="font-medium text-[#172033]">{asset}</p>
                  <span className="rounded-md bg-[#fff1e7] px-2 py-1 text-xs font-semibold text-[#a54712]">
                    {level}
                  </span>
                </div>
                <p className="mt-2 text-sm text-[#657187]">{reason}</p>
              </div>
            ))}
          </div>
        </aside>
      </section>
    </main>
  );
}
