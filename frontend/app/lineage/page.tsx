import { ResourcePage } from "@/components/resource-page";

export default function LineagePage() {
  return (
    <ResourcePage
      title="Lineage"
      description="Inspect upstream dependencies and downstream impact."
      filters={["Upstream", "Downstream", "Transformations"]}
      columns={["From", "Relationship", "To"]}
      rows={[
        ["Customer Warehouse", "ingests", "customers"],
        ["customers", "transforms", "customer_mart"],
      ]}
    />
  );
}
