import { ResourceForm } from "@/components/resource-form";
import { ScreenPage } from "@/components/screen-page";

export default function DataSourcesPage() {
  return (
    <>
      <ScreenPage title="Data Sources" description="Registered enterprise systems and connectors" filters={["All", "Healthy", "Degraded"]} columns={["Name", "Type", "Status"]} rows={[["Customer Warehouse", "PostgreSQL", "Healthy"], ["HR Exports", "CSV", "Registered"]]} />
      <div className="px-6 pb-6"><ResourceForm /></div>
    </>
  );
}
