"use server";
import { setSessionToken } from "@/lib/auth";
import { API_BASE_URL } from "@/lib/config/config";
import { ApiError, ApiSuccess, LoginSuccess } from "@/types/apiTypes";

const baseURL = API_BASE_URL;

export async function register(
    prevState: ApiSuccess | ApiError | undefined,
    formData: FormData,
): Promise<ApiSuccess | ApiError> {
    const email = formData.get("email");
    const password = formData.get("password");
    const firstName = formData.get("firstName");
    const lastName = formData.get("lastName");

    try {
        const res = await fetch(baseURL + "register", {
            method: "POST",
            body: JSON.stringify({ email, password, firstName, lastName }),
            headers: { "Content-Type": "application/json" },
        });

        const data: LoginSuccess | ApiError = await res.json();
        if ("error" in data) {
            return { type: "error", error: data.error };
        }

        await setSessionToken(data.sessionToken);

        return { type: "success", user: data.user };
    } catch (e) {
        console.error(e);
        const errorMessage =
            e instanceof Error ? e.message : "an unknown error occurred";
        return { type: "error", error: errorMessage };
    }
}
