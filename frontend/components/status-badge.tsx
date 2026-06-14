type StatusBadgeProps = {
  tone?: "success" | "warning" | "danger" | "neutral";
  children: React.ReactNode;
};

const tones = {
  success: "bg-emerald-50 text-emerald-700 border-emerald-200",
  warning: "bg-amber-50 text-amber-700 border-amber-200",
  danger: "bg-rose-50 text-rose-700 border-rose-200",
  neutral: "bg-slate-50 text-slate-700 border-slate-200",
};

export function StatusBadge({ tone = "neutral", children }: StatusBadgeProps) {
  return (
    <span className={`inline-flex rounded-md border px-2 py-1 text-xs font-semibold ${tones[tone]}`}>
      {children}
    </span>
  );
}
