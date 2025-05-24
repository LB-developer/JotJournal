import { fetchWithAuth } from "@/lib/auth";
import { JotCollection } from "@/types/jotTypes";

const baseURL = "http://localhost:8080/api/v1/jots"
export async function getJots(month: string) {
    const params = new URLSearchParams()
    params.set("month", month)

    if (month == "") {
        throw new Error("No month included in parameters") 
    }

    const url = baseURL + "?" + params

    const data: JotCollection | undefined = await fetchWithAuth<JotCollection>(url, "GET")
    if (data != undefined) {
        return data
    }
}

