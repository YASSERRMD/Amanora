import { ResourcePage } from "@/components/resource-page";

export default function CatalogPage() {
  return (
    <ResourcePage
      title="Catalog"
      description="Search governed data assets, fields, ownership, tags, and risk."
      filters={["PII", "High risk", "Owned", "Tagged"]}
      columns={["Asset", "Owner", "Risk"]}
      rows={[
        ["customers", "Privacy", "High"],
        ["finance_transactions", "Finance", "Medium"],
      ]}
    />
  );
}
