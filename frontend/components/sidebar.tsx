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

export function Sidebar() {
  return (
    <aside className="flex min-h-screen w-64 shrink-0 flex-col border-r border-[#d9e0ea] bg-white">
      <div className="border-b border-[#d9e0ea] px-5 py-4">
        <p className="text-lg font-semibold text-[#172033]">Amanora</p>
        <p className="text-sm text-[#657187]">Governance console</p>
      </div>
      <nav className="flex-1 space-y-1 px-3 py-4">
        {navItems.map(([label, href]) => (
          <a
            key={href}
            href={href}
            className="block rounded-md px-3 py-2 text-sm font-medium text-[#4f5e75] hover:bg-[#f1f5f9] hover:text-[#172033]"
          >
            {label}
          </a>
        ))}
      </nav>
    </aside>
  );
}
