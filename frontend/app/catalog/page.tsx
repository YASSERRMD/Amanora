import { ScreenPage } from "@/components/screen-page";

export default function CatalogPage() {
  return <ScreenPage title="Catalog" description="Searchable governed data assets" filters={["PII", "Owned", "High risk"]} columns={["Asset", "Owner", "Classification"]} rows={[["customers", "Privacy", "email"], ["transactions", "Finance", "confidential"]]} />;
}
