import { JotCollection, UpdateJotPayload } from "@/types/jotTypes";
import { NextRequest, NextResponse } from "next/server";
import { fetchWithAuth } from "@/lib/auth";

const baseURL = "http://localhost:8080/api/v1/jots";
export async function GET(req: NextRequest) {
    const method = "GET";

    const month = req.nextUrl.searchParams.get("month");
    if (!month) {
        throw new Error("No month in params");
    }

    const params = new URLSearchParams();
    params.set("month", month!);

    const url = baseURL + "?" + params;

    // fetch jots for user by month
    const res = await fetchWithAuth<JotCollection>(url, method);

    return NextResponse.json(res);
}

export async function PATCH(req: NextRequest) {
    const method = "PATCH";
    const headers = new Headers();
    headers.set("Content-Type", "application/json");

    // revalidateTag("tags here")

    let body: UpdateJotPayload;
    try {
        body = await req.json();
    } catch {
        return NextResponse.json(
            { error: "Invalid JSON body" },
            { status: 400 },
        );
    }
    const { jotID, isCompleted } = body;
    if (!jotID || typeof isCompleted !== "boolean") {
        return NextResponse.json(
            { error: "Missing or invalid fields" },
            { status: 400 },
        );
    }

    try {
        // update the specified jot
        await fetchWithAuth<undefined>(baseURL, method, headers, body);
        return new NextResponse(undefined, { status: 204 });
    } catch {
        return new NextResponse(undefined, { status: 500 });
    }
}
