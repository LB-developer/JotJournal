'use server'

import { setSessionToken } from "@/lib/auth";
import { ApiError, ApiSuccess, LoginToken} from "@/types/apiTypes";

const baseURL = "http://localhost:8080/api/v1/"

export async function login(prevState: any, formData: FormData): Promise<ApiSuccess | ApiError> {
  const email = formData.get("email")
  const password = formData.get("password")

  try {
    const res = await fetch(baseURL + "login", {
      method: 'POST',
      body: JSON.stringify({ email, password }),
      headers: { 'Content-Type': 'application/json' }
    });

    const data: LoginToken | ApiError = await res.json()
    if ("error" in data) {
        return { type: "error", error: data.error  }
    }
    
    await setSessionToken(data)
    
    return { type: "success" }
    
  } catch (e) {
    console.error(e)
    const errorMessage = e instanceof Error ? e.message : "an unknown error occurred"
    return { type: "error", error: errorMessage  }
  }
}

export async function logout() {}
