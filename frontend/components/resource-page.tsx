import { DataTable } from "@/components/data-table";
import { FilterBar } from "@/components/filter-bar";
import { StatusBadge } from "@/components/status-badge";

type ResourcePageProps = {
  title: string;
  description: string;
  filters: string[];
  columns: string[];
  rows: Array<Array<string>>;
  status?: string;
};

export function ResourcePage({
  title,
  description,
  filters,
  columns,
  rows,
  status = "Active",
}: ResourcePageProps) {
  return (
    <main className="space-y-6 p-6">
      <div className="flex flex-col justify-between gap-4 md:flex-row md:items-center">
        <div>
          <h1 className="text-2xl font-semibold text-[#172033]">{title}</h1>
          <p className="mt-1 text-sm text-[#657187]">{description}</p>
        </div>
        <StatusBadge status={status} />
      </div>
      <FilterBar filters={filters} />
      <DataTable columns={columns} rows={rows} />
    </main>
  );
}
