import { ResourcePage } from "@/components/resource-page";

export default function AuditPage() {
  return (
    <ResourcePage
      title="Audit"
      description="Review governance actions, workers, and policy decisions."
      filters={["System", "Worker", "Policy", "Failure"]}
      columns={["Actor", "Action", "Outcome"]}
      rows={[
        ["system", "discovery.job.completed", "Success"],
        ["worker", "classification.job.completed", "Success"],
      ]}
    />
  );
}
