import { ScreenPage } from "@/components/screen-page";

export default function ClassificationPage() {
  return <ScreenPage title="PII Findings" description="Sensitive data detections and review status" filters={["Open", "Confirmed", "False positive"]} columns={["Asset", "Field", "PII Type"]} rows={[["customers", "email", "email"], ["customers", "phone", "phone_number"]]} />;
}
