import { fetchWithAuth } from "@/lib/auth";
import { JotCollection, UpdateJotPayload } from "@/types/jotTypes";
import { cookies } from "next/headers";
import { NextRequest } from "next/server";

const baseURL = "http://localhost:8080/api/v1/jots"

export async function getJots(month: string) {
    const params = new URLSearchParams()
    params.set("month", month)

    if (month == "") {
        throw new Error("No month included in parameters") 
    }

    // fetch jots for user by month from db
    const req = new NextRequest(baseURL + "?" + params, {
        method: 'GET'
    })

    const data: JotCollection = await fetchWithAuth(req)
    if (data != undefined) {
        return data
    }
}



export async function updateJot(jot: UpdateJotPayload) {
    const cookieStore = await cookies()

    const token = cookieStore.get("refreshToken")?.value ?? ""

    try {
        // update the specified jot
        const res = await fetch(baseURL, {
            method: 'PATCH',
            headers: { 
                'Content-Type': 'application/json',
                'Authorization': token },
            body: JSON.stringify(jot)
    
        })

        if (!res.ok) {
          const { error } = await res.json();
          throw new Error(error ?? 'Unknown error');
        }

    } catch (e) {
        console.error(e)
    }
}
