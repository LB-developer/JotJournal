"use server";
import { getJots } from "../jots/actions";
import dynamic from "next/dynamic";

const JotDisplay = dynamic(() => import("@/components/jots/JotDisplay"));

export default async function JotContainer({
    month,
    year,
}: {
    month: string;
    year: string;
}) {
    const jots = await getJots(month, year);

    return (
        <section className="flex flex-row">
            {jots && (
                <JotDisplay jotCollection={jots} month={month} year={year} />
            )}
        </section>
    );
}
