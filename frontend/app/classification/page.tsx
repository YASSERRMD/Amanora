import { ResourcePage } from "@/components/resource-page";

export default function ClassificationPage() {
  return (
    <ResourcePage
      title="Classification"
      description="Review PII findings and classification jobs."
      filters={["Email", "Phone", "Passport", "False positive"]}
      columns={["Field", "PII Type", "Confidence"]}
      rows={[
        ["email", "Email", "98%"],
        ["phone", "Phone", "93%"],
      ]}
      status="High"
    />
  );
}
