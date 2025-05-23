'server-only'
import { ApiError, LoginToken } from "@/types/apiTypes";
import { cookies } from "next/headers";
import { redirect } from "next/navigation";
import { NextRequest } from "next/server";

export async function fetchWithAuth<T>(req: NextRequest): Promise<T> {
  const sessionToken = await getSessionToken()
  if (!sessionToken) {
    redirect("/login")
  }
  req.headers.set("Authorization", sessionToken)

  for (let retries = 0; retries < 3; retries++) {
    try {
      const res = await fetch(req)

      if (res.status === 401) {
        const error = await refreshSessionToken()
        throw new Error(error)
      }

      if (!res.ok) {
        const error = await res.json()
        throw new Error(error)
      }

      const json = await res.json()
      return json as T

    } catch (e) {
      console.warn("Auth check did not pass")
      if (retries >= 2) {
        console.error(e)
        redirect("/login")
      }

    }
  }

  throw new Error("Unreachable: retry loop did not return or redirect")
}

const baseURL = "http://localhost:8080/api/v1/"

export default async function refreshSessionToken(): Promise<string> {
  const sessionToken = await getSessionToken()
  if (!sessionToken) {
    redirect("/login")
  }

  const res = await fetch(baseURL + "refresh", {
      method: 'POST',
      headers: {
          'Content-Type': 'application/json',
          "Authorization": sessionToken
      }
  })

  const data: LoginToken | ApiError = await res.json()
  if ("error" in data) {
      return data.error
  }
  
  await setSessionToken(data)

  return "No error, new session token retrieved and set..."
}

export async function setSessionToken(token: LoginToken): Promise<void> {
  const cookieStore = await cookies()
  cookieStore.set('sessionToken', token.sessionToken, { httpOnly: true, secure: true })
}

export async function getSessionToken(): Promise<string | undefined> {
  const cookieStore = await cookies()
  const sessionToken = cookieStore.get("sessionToken")?.value
  return sessionToken
}

export async function clearSessionToken(): Promise<void> {
  const cookieStore = await cookies()
  cookieStore.delete("sessionToken")
}
