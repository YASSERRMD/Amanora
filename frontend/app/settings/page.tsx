import { ScreenPage } from "@/components/screen-page";

export default function SettingsPage() {
  return <ScreenPage title="Settings" description="Tenant, API, and governance configuration" filters={["Tenant", "API Keys", "Roles"]} columns={["Setting", "Value", "Status"]} rows={[["Tenant", "Amanora Demo", "Active"], ["API access", "Enabled", "Healthy"]]} />;
}
