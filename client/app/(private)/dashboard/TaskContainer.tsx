"use server";
import { DataTable } from "@/components/ui/table";
import { getTasks } from "../tasks/actions";
import { typeColumns } from "@/types/taskTypes";

export default async function TaskContainer({
    month,
    year,
}: {
    month: string;
    year: string;
}) {
    console.log(year);
    const tasks = await getTasks(month);
    tasks?.forEach((task) => {
        const today = new Date();
        const deadlineDate = new Date(task.deadline);
        const days = dateDiffInDays(today, deadlineDate);
        task.deadline = convertDaysToHumanReadable(days);
    });

    return <>{tasks && <DataTable columns={typeColumns} data={tasks} />}</>;
}

function dateDiffInDays(a: Date, b: Date) {
    const _MS_PER_DAY = 1000 * 60 * 60 * 24;
    // Discard the time and time-zone information.
    const utc1 = Date.UTC(a.getFullYear(), a.getMonth(), a.getDate());
    const utc2 = Date.UTC(b.getFullYear(), b.getMonth(), b.getDate());

    return Math.floor((utc2 - utc1) / _MS_PER_DAY);
}

const time = {
    past: "Days ago",
    yesterday: "Yesterday",
    today: "Today",
    tomorrow: "Tomorrow",
    future: "Days from now",
};

function convertDaysToHumanReadable(days: number): string {
    switch (true) {
        case days < -1:
            return `${String(days)} ${time.past}`;
        case days === -1:
            return time.yesterday;
        case days === 0:
            return time.today;
        case days === 1:
            return time.yesterday;
        case days === 1:
            return `${String(days)} ${time.future}`;
        default:
            return "Unknown";
    }
}
