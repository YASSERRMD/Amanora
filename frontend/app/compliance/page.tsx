import { ScreenPage } from "@/components/screen-page";

export default function CompliancePage() {
  return <ScreenPage title="Compliance" description="Framework checks and evidence capture" filters={["GDPR", "UAE PDPL", "Internal"]} columns={["Framework", "Check", "Status"]} rows={[["GDPR-style", "PII owner", "Pass"], ["Internal", "Retention assigned", "Fail"]]} />;
}
