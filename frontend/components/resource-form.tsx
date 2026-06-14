export function ResourceForm() {
  return (
    <form className="grid gap-3 rounded-lg border border-slate-200 bg-white p-4 md:grid-cols-3">
      <input className="rounded-md border border-slate-300 px-3 py-2 text-sm" placeholder="Name" />
      <input className="rounded-md border border-slate-300 px-3 py-2 text-sm" placeholder="Type" />
      <button className="rounded-md bg-blue-600 px-4 py-2 text-sm font-semibold text-white" type="button">
        Save
      </button>
    </form>
  );
}
