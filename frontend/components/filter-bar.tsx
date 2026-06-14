type FilterBarProps = {
  filters: string[];
};

export function FilterBar({ filters }: FilterBarProps) {
  return (
    <div className="flex flex-wrap items-center gap-2">
      {filters.map((filter) => (
        <button
          key={filter}
          className="rounded-md border border-[#cfd8e6] bg-white px-3 py-2 text-sm font-medium text-[#4f5e75]"
        >
          {filter}
        </button>
      ))}
    </div>
  );
}
