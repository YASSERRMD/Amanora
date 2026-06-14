import { DataTable } from "./data-table";
import { FilterBar } from "./filter-bar";
import { StatusBadge } from "./status-badge";

type ScreenPageProps = {
  title: string;
  description: string;
  filters: string[];
  columns: string[];
  rows: string[][];
};

export function ScreenPage({ title, description, filters, columns, rows }: ScreenPageProps) {
  return (
    <main className="space-y-6 p-6">
      <div className="flex flex-col justify-between gap-4 md:flex-row md:items-center">
        <div>
          <h1 className="text-2xl font-semibold text-[#172033]">{title}</h1>
          <p className="mt-1 text-sm text-[#657187]">{description}</p>
        </div>
        <FilterBar filters={filters} />
      </div>
      <DataTable columns={columns} rows={rows} />
      <section className="rounded-lg border border-[#d9e0ea] bg-white p-5">
        <div className="flex items-center justify-between">
          <p className="font-medium text-[#172033]">Workflow status</p>
          <StatusBadge status="Running" />
        </div>
      </section>
    </main>
  );
}
