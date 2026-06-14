import { ScreenPage } from "@/components/screen-page";

export default function RetentionPage() {
  return <ScreenPage title="Retention" description="Lifecycle rules, expiry, and violations" filters={["Review", "Archive", "Delete"]} columns={["Policy", "Days", "Action"]} rows={[["Customer PII Review", "365", "Review"], ["Finance Archive", "2555", "Archive"]]} />;
}
