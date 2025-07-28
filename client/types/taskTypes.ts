import { ColumnDef } from "@tanstack/react-table";

export interface Task {
    daily: boolean;
    deadline: string;
    description: string;
    id: number;
    isCompleted: boolean;
    monthly: boolean;
    userId: number;
    weekly: boolean;
}

export const typeColumns: ColumnDef<Task>[] = [
    {
        accessorKey: "id",
        header: "ID",
    },
    {
        accessorKey: "deadline",
        header: "Deadline",
    },
    {
        accessorKey: "isCompleted",
        header: "Completed",
    },
    {
        accessorKey: "description",
        header: "Description",
    },
];
