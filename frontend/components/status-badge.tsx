type StatusBadgeProps = {
  status: string;
};

export function StatusBadge({ status }: StatusBadgeProps) {
  const normalized = status.toLowerCase();
  const tone =
    normalized.includes("fail") || normalized.includes("high")
      ? "bg-[#fff1e7] text-[#a54712]"
      : normalized.includes("running") || normalized.includes("medium")
        ? "bg-[#eef4ff] text-[#1d4ed8]"
        : "bg-[#e6f4f1] text-[#0f766e]";

  return (
    <span className={`inline-flex rounded-md px-2 py-1 text-xs font-semibold ${tone}`}>
      {status}
    </span>
  );
}
