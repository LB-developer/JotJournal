import JotDisplay from "@/components/jots/JotDisplay";
import LogoutButton from "@/components/user/LogoutButton";
import { getJots } from "../jots/actions";

export default async function JotContainer() {
  const jots = await getJots("4")

  return (
    <>
      {jots && <JotDisplay jotCollection={jots} />}
      <LogoutButton />
    </>
  );
}
