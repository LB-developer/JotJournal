"server-only";
import { API_BASE_URL } from "./config/config";
import { ApiError, SessionToken } from "@/types/apiTypes";
import { cookies } from "next/headers";
import { redirect } from "next/navigation";

const baseURL = API_BASE_URL;
export async function fetchWithAuth<T>(
    url: string,
    method: "GET" | "PATCH",
    headers?: Headers,
    reqBody?: unknown,
): Promise<T | undefined> {
    const reqHeaders = new Headers();
    if (headers) {
        for (const [key, val] of headers.entries()) {
            reqHeaders.set(key, val);
        }
    }

    let sessionToken = await getSessionToken();
    for (let retries = 0; retries < 3; retries++) {
        if (!sessionToken) {
            redirect("/login");
        }

        reqHeaders.set("Authorization", sessionToken);

        try {
            const res = await fetch(url, {
                method: method,
                headers: reqHeaders,
                body: JSON.stringify(reqBody),
            });

            if (res.status === 401) {
                sessionToken = await refreshSessionToken();
                throw new Error("Unauthorized");
            }

            // no content expected
            if (res.status === 204) {
                return undefined;
            }

            const body = (await res.json()) as T;
            return body;
        } catch (e) {
            if (retries >= 2) {
                console.log(JSON.stringify(e));
                redirect("/login");
            }
            console.warn(
                `Auth check did not pass on attempt: ${retries}, retrying...`,
            );
        }
    }

    throw new Error("Unreachable: retry loop did not return or redirect");
}

export default async function refreshSessionToken(): Promise<string> {
    const sessionToken = await getSessionToken();
    if (!sessionToken) {
        redirect("/login");
    }

    const res = await fetch(baseURL + "refresh", {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
            Authorization: sessionToken,
        },
    });

    const data: SessionToken | ApiError = await res.json();
    if ("error" in data) {
        console.error(data.error);
        return `ERROR FROM SERVER`;
    }

    await setSessionToken(data.sessionToken);

    return data.sessionToken;
}

export async function setSessionToken(token: string): Promise<void> {
    const cookieStore = await cookies();
    cookieStore.set("sessionToken", token, {
        httpOnly: true,
        secure: true,
    });
    console.log("set new 'sessionToken' for user");
}

export async function getSessionToken(): Promise<string | undefined> {
    const cookieStore = await cookies();
    const sessionToken = cookieStore.get("sessionToken")?.value;
    return sessionToken;
}

export async function clearSessionToken(): Promise<void> {
    const cookieStore = await cookies();
    cookieStore.delete("sessionToken");
}
