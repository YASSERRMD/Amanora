import { ScreenPage } from "@/components/screen-page";

export default function DiscoveryPage() {
  return <ScreenPage title="Discovery Jobs" description="Schema, table, field, and sampling runs" filters={["Queued", "Running", "Completed"]} columns={["Job", "Source", "Status"]} rows={[["disc_000001", "Customer Warehouse", "Completed"], ["disc_000002", "HR Exports", "Running"]]} />;
}
