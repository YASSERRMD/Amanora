import { ResourcePage } from "@/components/resource-page";

export default function RetentionPage() {
  return (
    <ResourcePage
      title="Retention"
      description="Evaluate lifecycle rules, expiry windows, and violations."
      filters={["Review", "Archive", "Delete", "Violation"]}
      columns={["Policy", "Days", "Action"]}
      rows={[
        ["Customer PII Review", "365", "Review"],
        ["Finance Archive", "2555", "Archive"],
      ]}
    />
  );
}
