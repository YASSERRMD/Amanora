import { ScreenPage } from "@/components/screen-page";

export default function LineagePage() {
  return <ScreenPage title="Lineage" description="Upstream and downstream impact analysis" filters={["Upstream", "Downstream", "Transformations"]} columns={["From", "Relationship", "To"]} rows={[["raw_customers", "transforms", "customers"], ["customers", "produces", "customer_report"]]} />;
}
