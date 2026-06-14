import { ResourcePage } from "@/components/resource-page";

export default function SettingsPage() {
  return (
    <ResourcePage
      title="Settings"
      description="Configure tenants, API keys, roles, and governance defaults."
      filters={["Tenant", "API keys", "Roles", "Headers"]}
      columns={["Setting", "Value", "Status"]}
      rows={[
        ["Tenant isolation", "Enabled", "Active"],
        ["API key auth", "Configured", "Active"],
      ]}
    />
  );
}
