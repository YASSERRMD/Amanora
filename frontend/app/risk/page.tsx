import { ResourcePage } from "@/components/resource-page";

export default function RiskPage() {
  return (
    <ResourcePage
      title="Risk"
      description="Prioritize assets using PII, ownership, retention, and compliance signals."
      filters={["Critical", "High", "Medium", "Low"]}
      columns={["Asset", "Score", "Level"]}
      rows={[
        ["customers", "72", "High"],
        ["analytics_sandbox", "22", "Low"],
      ]}
      status="High"
    />
  );
}
