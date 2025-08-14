import { getCurrentMonth } from "@/lib/date/getCurrentMonth";
import JotContainer from "./JotContainer";
import TaskContainer from "./TaskContainer";
import { getCurrentYear } from "@/lib/date/getCurrentYear";
import MoveDashboardMonthButton from "@/components/jots/MoveDashboardMonthButton";

const months = [
    "January",
    "February",
    "March",
    "April",
    "May",
    "June",
    "July",
    "August",
    "September",
    "October",
    "November",
    "December",
];

export default async function DashboardPage({
    searchParams,
}: {
    searchParams: Promise<{ month?: string; year?: string }>;
}) {
    const params = await searchParams;
    const month = params.month ?? getCurrentMonth();
    const year = params.year ?? getCurrentYear();
    const activeMonth = months[Number(month) - 1];

    return (
        <div
            className={`grid grid-cols-1 md:grid-cols-10 md:grid-rows-10 gap-4 p-4 `}
        >
            <div className="flex flex-row justify-center align-center w-100 md:col-span-10 md:row-span-1 justify-self-center">
                <MoveDashboardMonthButton
                    monthNumberAsString={month}
                    yearNumberAsString={year}
                    direction="prev"
                />
                <h1 className="underline text-5xl">{`${activeMonth} ${year}`}</h1>
                <MoveDashboardMonthButton
                    monthNumberAsString={month}
                    yearNumberAsString={year}
                    direction="next"
                />
            </div>
            <div className="md:col-span-10 md:row-span-9 order-1 md:order-1">
                <JotContainer month={month} year={year} />
            </div>

            <div className="md:col-span-10 order-2 md:row-span-4 md:order-2">
                <TaskContainer month={month} />
            </div>
        </div>
    );
}
