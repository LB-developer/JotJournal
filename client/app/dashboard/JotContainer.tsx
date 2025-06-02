'use server'
import { getJots } from "../jots/actions";
import dynamic from "next/dynamic";

const JotDisplay = dynamic(() => import('@/components/jots/JotDisplay'));

export default async function JotContainer({ month }: { month: string }) {
  const jots = await getJots(month)


  return (
    <>
      {jots && <JotDisplay jotCollection={jots} />}
    </>
  );
}
