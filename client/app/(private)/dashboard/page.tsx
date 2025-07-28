import { getCurrentMonth } from "@/lib/date/getCurrentMonth";
import JotContainer from "./JotContainer";
import TaskContainer from "./TaskContainer";
import { getCurrentYear } from "@/lib/date/getCurrentYear";

const days = [
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

    return (
        <div
            className={`grid grid-cols-1 md:grid-cols-10 md:grid-rows-10 gap-4 p-4`}
        >
            <h1 className="text-center underline md:col-span-10 md:row-span-1 text-5xl">
                {`${days[Number(month)]}${year}`}
            </h1>
            <div className="md:col-span-10 md:row-span-9 order-1 md:order-1">
                <JotContainer month={month} year={year} />
            </div>

            <div className="md:col-span-10 order-2 md:row-span-4 md:order-2">
                <TaskContainer month={month} year={year} />
            </div>
        </div>
    );
}
