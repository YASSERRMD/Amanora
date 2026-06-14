export function LoadingState() {
  return <div className="rounded-lg border border-[#d9e0ea] bg-white p-5 text-sm text-[#657187]">Loading</div>;
}

export function ErrorState({ message }: { message: string }) {
  return <div className="rounded-lg border border-[#fecaca] bg-[#fff7f7] p-5 text-sm text-[#991b1b]">{message}</div>;
}
