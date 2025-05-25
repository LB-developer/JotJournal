import LogoutButton from "@/components/user/LogoutButton";
import JotContainer from "./JotContainer";
import getCurrentMonth from "@/lib/date/getCurrentMonth";

export default async function DashboardPage({ searchParams }: { searchParams: { month?: string } }) {
  const params = await searchParams
  const month = params.month ?? getCurrentMonth()

  return (
    <>
      <JotContainer month={month} />
      <LogoutButton />
    </>
  );
}
