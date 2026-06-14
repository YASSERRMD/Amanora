import { ScreenPage } from "@/components/screen-page";

export default function AuditPage() {
  return <ScreenPage title="Audit" description="Governance actions and system decisions" filters={["System", "Worker", "User"]} columns={["Actor", "Action", "Outcome"]} rows={[["system", "discovery.job.completed", "Success"], ["worker", "classification.job.completed", "Success"]]} />;
}
