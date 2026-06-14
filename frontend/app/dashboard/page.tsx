import { AppShell } from "@/components/app-shell";
import { DataTable } from "@/components/data-table";
import { FilterBar } from "@/components/filter-bar";
import { StatusBadge } from "@/components/status-badge";

const metrics = [
  ["Sources", "18", "success"],
  ["Catalog assets", "12.8k", "neutral"],
  ["PII findings", "432", "warning"],
  ["High risk", "37", "danger"],
] as const;

const rows = [
  { asset: "customers", owner: "Privacy", risk: "High", status: "Review" },
  { asset: "finance_transactions", owner: "Finance", risk: "Medium", status: "Compliant" },
  { asset: "hr_archive", owner: "People", risk: "High", status: "Retention" },
];

export default function DashboardPage() {
  return (
    <AppShell>
      <div className="space-y-6">
        <div>
          <h1 className="text-3xl font-semibold">Dashboard</h1>
          <p className="mt-1 text-sm text-slate-600">Operational overview across discovery, policy, compliance, and risk.</p>
        </div>
        <div className="grid gap-4 md:grid-cols-4">
          {metrics.map(([label, value, tone]) => (
            <div key={label} className="rounded-lg border border-slate-200 bg-white p-4">
              <p className="text-sm text-slate-500">{label}</p>
              <div className="mt-3 flex items-center justify-between">
                <p className="text-2xl font-semibold">{value}</p>
                <StatusBadge tone={tone}>{tone}</StatusBadge>
              </div>
            </div>
          ))}
        </div>
        <FilterBar />
        <DataTable
          rows={rows}
          columns={[
            { key: "asset", header: "Asset", render: (row) => row.asset },
            { key: "owner", header: "Owner", render: (row) => row.owner },
            { key: "risk", header: "Risk", render: (row) => row.risk },
            { key: "status", header: "Status", render: (row) => <StatusBadge>{row.status}</StatusBadge> },
          ]}
        />
      </div>
    </AppShell>
  );
}
