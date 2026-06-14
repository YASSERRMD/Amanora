import { ScreenPage } from "@/components/screen-page";

export default function PoliciesPage() {
  return <ScreenPage title="Policies" description="Governance rules and compliance decisions" filters={["Privacy", "Retention", "Ownership"]} columns={["Policy", "Severity", "Effect"]} rows={[["PII must have owner", "High", "Warn"], ["High risk needs retention", "High", "Require review"]]} />;
}
