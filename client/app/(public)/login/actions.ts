"use server";

import {
    clearSessionToken,
    clearUserFromCookies,
    getSessionToken,
    setSessionToken,
    setUserInCookies,
} from "@/lib/auth";
import { API_BASE_URL } from "@/lib/config/config";
import { ApiError, ApiSuccess, LoginSuccess } from "@/types/apiTypes";
import { redirect } from "next/navigation";

const baseURL = API_BASE_URL;

export async function login(
    _: ApiSuccess | ApiError | undefined,
    formData: FormData,
): Promise<ApiSuccess | ApiError> {
    const email = formData.get("email");
    const password = formData.get("password");

    try {
        const res = await fetch(baseURL + "login", {
            method: "POST",
            body: JSON.stringify({ email, password }),
            headers: { "Content-Type": "application/json" },
        });

        const json: LoginSuccess | ApiError = await res.json();
        if ("error" in json) {
            throw new Error(json.error);
        }

        await setSessionToken(json.sessionToken);
        await setUserInCookies(json.user);

        return { type: "success", user: json.user };
    } catch (e) {
        console.error(e);
        const errorMessage =
            e instanceof Error ? e.message : "an unknown error occurred";
        return { type: "error", error: errorMessage };
    }
}

export async function logout(): Promise<void> {
    const sessionToken = await getSessionToken();

    try {
        const res = await fetch(baseURL + "logout", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: `Bearer ${sessionToken}`,
            },
        });

        if (!res.ok) {
            const error: ApiError = await res.json();
            console.error(error);
            throw new Error(error.error);
        }
    } catch (e) {
        console.error(e);
    }

    clearUserFromCookies();
    clearSessionToken();
    redirect("/login");
}
