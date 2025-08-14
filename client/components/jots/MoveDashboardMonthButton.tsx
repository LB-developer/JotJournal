"use client";

import Link from "next/link";

interface Props {
    monthNumberAsString: string;
    yearNumberAsString: string;
    direction: "prev" | "next";
}

export default function MoveDashboardMonthButton({
    monthNumberAsString,
    yearNumberAsString,
    direction,
}: Props) {
    const handleMoveToNextJot = (month: string, year: string) => {
        const currentMonth = Number(month);
        const currentYear = Number(year);

        let newMonth: number;
        let newYear: number;

        if (direction === "prev") {
            newMonth = currentMonth - 1;
            if (newMonth < 1) {
                newMonth = 12;
                newYear = currentYear - 1;
            } else {
                newYear = currentYear;
            }
        } else {
            newMonth = currentMonth + 1;
            if (newMonth > 12) {
                newMonth = 1;
                newYear = currentYear + 1;
            } else {
                newYear = currentYear;
            }
        }

        const newMonthStringWithPadded0 = newMonth.toString().padStart(2, "0");
        const newYearString = newYear.toString();

        const url = `/dashboard?month=${newMonthStringWithPadded0}&year=${newYearString}`;

        return (
            <Link href={url}>
                <p>{direction === "prev" ? "<" : ">"}</p>
            </Link>
        );
    };

    return (
        <div>
            {handleMoveToNextJot(monthNumberAsString, yearNumberAsString)}
        </div>
    );
}
