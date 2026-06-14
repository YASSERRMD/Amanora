import { ResourcePage } from "@/components/resource-page";

export default function CompliancePage() {
  return (
    <ResourcePage
      title="Compliance"
      description="Run framework checks and collect evidence."
      filters={["GDPR-style", "UAE PDPL", "India DPDP", "Internal"]}
      columns={["Framework", "Check", "Status"]}
      rows={[
        ["GDPR-style", "PII owner", "Pass"],
        ["Internal", "Steward assigned", "Fail"],
      ]}
      status="Fail"
    />
  );
}
