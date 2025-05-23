import LogoutButton from "@/components/user/LogoutButton";
import { getJots } from "../jots/actions";

export default async function DashboardPage() {
  // TODO: make page ("4") dynamic
  await getJots("4")

  return (
    <>
      <p>User is logged in and at the dashboard</p>
      <LogoutButton />
    </>
  );
}
