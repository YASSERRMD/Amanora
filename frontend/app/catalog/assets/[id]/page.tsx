import { ScreenPage } from "@/components/screen-page";

export default function AssetDetailPage() {
  return <ScreenPage title="Asset Detail" description="Field metadata, tags, ownership, and risk" filters={["Fields", "Tags", "Lineage"]} columns={["Field", "Type", "Classification"]} rows={[["email", "text", "email"], ["phone", "text", "phone_number"]]} />;
}
