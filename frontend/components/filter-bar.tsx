export function FilterBar() {
  return (
    <div className="flex flex-wrap gap-3 rounded-lg border border-slate-200 bg-white p-3">
      <input
        className="min-w-64 rounded-md border border-slate-300 px-3 py-2 text-sm"
        placeholder="Search assets, policies, owners"
      />
      <select className="rounded-md border border-slate-300 px-3 py-2 text-sm">
        <option>All risk levels</option>
        <option>High</option>
        <option>Medium</option>
        <option>Low</option>
      </select>
      <select className="rounded-md border border-slate-300 px-3 py-2 text-sm">
        <option>All classifications</option>
        <option>PII</option>
        <option>Confidential</option>
      </select>
    </div>
  );
}
