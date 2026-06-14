import { ScreenPage } from "@/components/screen-page";

export default function RiskPage() {
  return <ScreenPage title="Risk" description="Risk scores by governed asset" filters={["Critical", "High", "Medium"]} columns={["Asset", "Score", "Level"]} rows={[["customers", "72", "High"], ["hr_archive", "88", "Critical"]]} />;
}
