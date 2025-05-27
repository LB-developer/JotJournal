import { fetchWithAuth } from "@/lib/auth";
import { API_BASE_URL } from "@/lib/config/config";
import { JotCollection } from "@/types/jotTypes";

const baseURL = `${API_BASE_URL}jots`;
export async function getJots(month: string) {
    const params = new URLSearchParams();
    params.set("month", month);

    if (month == "") {
        throw new Error("No month included in parameters");
    }

    const url = baseURL + "?" + params;

    const data: JotCollection | undefined = await fetchWithAuth<JotCollection>(
        url,
        "GET",
    );
    if (data != undefined) {
        return data;
    }
}
