import Link from "next/link";

const navItems = [
  ["Dashboard", "/dashboard"],
  ["Data Sources", "/datasources"],
  ["Discovery", "/discovery"],
  ["Catalog", "/catalog"],
  ["Classification", "/classification"],
  ["Lineage", "/lineage"],
  ["Policies", "/policies"],
  ["Retention", "/retention"],
  ["Compliance", "/compliance"],
  ["Risk", "/risk"],
  ["Audit", "/audit"],
  ["Settings", "/settings"],
];

export function AppShell({ children }: { children: React.ReactNode }) {
  return (
    <div className="min-h-screen bg-slate-100 text-slate-950">
      <aside className="fixed inset-y-0 left-0 hidden w-64 border-r border-slate-200 bg-white p-5 lg:block">
        <div className="text-xl font-semibold">Amanora</div>
        <nav className="mt-8 space-y-1">
          {navItems.map(([label, href]) => (
            <Link key={href} href={href} className="block rounded-md px-3 py-2 text-sm font-medium text-slate-700 hover:bg-slate-100">
              {label}
            </Link>
          ))}
        </nav>
      </aside>
      <div className="lg:pl-64">
        <header className="border-b border-slate-200 bg-white px-6 py-4">
          <p className="text-sm text-slate-500">Autonomous data governance console</p>
        </header>
        <main className="px-6 py-6">{children}</main>
      </div>
    </div>
  );
}
