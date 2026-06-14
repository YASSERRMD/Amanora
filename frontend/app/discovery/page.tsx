import { ResourcePage } from "@/components/resource-page";

export default function DiscoveryPage() {
  return (
    <ResourcePage
      title="Discovery"
      description="Track schema, table, field, and sampling jobs."
      filters={["Running", "Completed", "Failed", "Queued"]}
      columns={["Job", "Source", "Status"]}
      rows={[
        ["disc_000001", "Customer Warehouse", "Completed"],
        ["disc_000002", "Customer Exports", "Running"],
      ]}
      status="Running"
    />
  );
}
