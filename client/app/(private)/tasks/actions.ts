import { fetchWithAuth } from "@/lib/auth";
import { API_BASE_URL } from "@/lib/config/config";
import { Task } from "@/types/taskTypes";

const baseURL = `${API_BASE_URL}tasks`;
export async function getTasks(month: string) {
    const params = new URLSearchParams();
    params.set("month", month);

    if (month == "") {
        throw new Error("No month included in parameters");
    }

    const url = baseURL + "?" + params;

    const data: Task[] | undefined = await fetchWithAuth<Task[]>(url, "GET");
    if (data != undefined) {
        return data;
    }
}
