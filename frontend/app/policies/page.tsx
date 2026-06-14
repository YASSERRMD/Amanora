import { ResourcePage } from "@/components/resource-page";

export default function PoliciesPage() {
  return (
    <ResourcePage
      title="Policies"
      description="Manage policy documents and governance decisions."
      filters={["Privacy", "Ownership", "Retention", "Failed"]}
      columns={["Policy", "Category", "Severity"]}
      rows={[
        ["PII data must have owner", "Ownership", "High"],
        ["High-risk data needs retention", "Retention", "High"],
      ]}
    />
  );
}
