import { ResourcePage } from "@/components/resource-page";

export default function DataSourcesPage() {
  return (
    <ResourcePage
      title="Data Sources"
      description="Register, test, and monitor governed source connections."
      filters={["PostgreSQL", "CSV", "Healthy", "Needs review"]}
      columns={["Source", "Type", "Status"]}
      rows={[
        ["Customer Warehouse", "PostgreSQL", "Healthy"],
        ["Customer Exports", "CSV", "Registered"],
        ["REST CRM", "REST", "Placeholder"],
      ]}
    />
  );
}
