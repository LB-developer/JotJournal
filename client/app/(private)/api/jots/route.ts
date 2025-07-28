import {
    CreateJotPayload,
    Jot,
    JotCollection,
    UpdateJotPayload,
} from "@/types/jotTypes";
import { NextRequest, NextResponse } from "next/server";
import { fetchWithAuth } from "@/lib/auth";
import { API_BASE_URL } from "@/lib/config/config";
import { ApiError } from "@/types/apiTypes";

const baseURL = `${API_BASE_URL}jots`;
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

    const tags: string[] = [];
    for (const key of Object.keys(tags)) {
        tags.push(`userid-${key}`);
    }

    return NextResponse.json(res);
}

export async function PATCH(req: NextRequest) {
    const method = "PATCH";
    const headers = new Headers();
    headers.set("Content-Type", "application/json");

    // TODO: revalidateTag("tags here");

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
    } catch (e) {
        console.error(e);
        return new NextResponse(undefined, { status: 500 });
    }
}

export async function POST(
    req: NextRequest,
): Promise<NextResponse<Jot[] | ApiError | undefined>> {
    const method = "POST";
    const headers = new Headers();
    headers.set("Content-Type", "application/json");

    // TODO: revalidateTag("tags here");

    let body: CreateJotPayload;
    try {
        body = await req.json();
    } catch {
        return NextResponse.json(
            { type: "error", error: "Invalid JSON body" },
            { status: 400 },
        );
    }
    const { name, month, year } = body;
    if (!name || !month || !year) {
        return NextResponse.json(
            { type: "error", error: "Missing or invalid fields" },
            { status: 400 },
        );
    }

    try {
        // create the specified jot
        const res = await fetchWithAuth<Jot[] | undefined>(
            baseURL,
            method,
            headers,
            body,
        );
        return NextResponse.json(res, { status: 201 });
    } catch (e) {
        console.error(e);
        return new NextResponse(undefined, { status: 500 });
    }
}
