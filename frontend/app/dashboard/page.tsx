import { DataTable } from "@/components/data-table";
import { FilterBar } from "@/components/filter-bar";
import { StatusBadge } from "@/components/status-badge";

export default function DashboardPage() {
  return (
    <main className="space-y-6 p-6">
      <div className="flex flex-col justify-between gap-4 md:flex-row md:items-center">
        <div>
          <h1 className="text-2xl font-semibold text-[#172033]">Dashboard</h1>
          <p className="mt-1 text-sm text-[#657187]">Governance activity and risk posture</p>
        </div>
        <FilterBar filters={["All sources", "High risk", "PII", "Needs owner"]} />
      </div>
      <div className="grid gap-4 md:grid-cols-4">
        {[
          ["Sources", "18", "Running"],
          ["Assets", "12.8k", "Healthy"],
          ["Findings", "432", "High"],
          ["Policies", "91%", "Passing"],
        ].map(([label, value, status]) => (
          <section key={label} className="rounded-lg border border-[#d9e0ea] bg-white p-5">
            <p className="text-sm text-[#657187]">{label}</p>
            <p className="mt-2 text-2xl font-semibold text-[#172033]">{value}</p>
            <div className="mt-3">
              <StatusBadge status={status} />
            </div>
          </section>
        ))}
      </div>
      <DataTable
        columns={["Asset", "Workflow", "Status"]}
        rows={[
          ["customers", "Classification", "Running"],
          ["finance_transactions", "Retention", "Passing"],
          ["hr_archive", "Ownership", "High risk"],
        ]}
      />
    </main>
  );
}
