'server-only'
import { ApiError, LoginToken } from "@/types/apiTypes";
import { cookies } from "next/headers";
import { redirect } from "next/navigation";

export async function fetchWithAuth<T>(
  url: string,
  method: "GET" | "PATCH",
  headers?: Headers,
  reqBody?: any
): Promise<T | undefined> {
  const sessionToken = await getSessionToken()

  if (!sessionToken) {
    redirect("/login")
  }

  if (headers) {
    headers.set("Authorization", sessionToken)
  }

  for (let retries = 0; retries < 3; retries++) {
    try {
        const res = await fetch(url, 
        {
          method: method,
          headers: {"Authorization": sessionToken},
          body: JSON.stringify(reqBody)
        })
      

      if (res.status === 401) {
        const error = await refreshSessionToken()
        throw new Error(error)
      }

      // no content expected
      if (res.status === 204) {
        return undefined
      }

      const body = await res.json() as T
      return body

    } catch (e) {
      console.warn("Auth check did not pass")
      if (retries >= 2) {
        console.error(JSON.stringify(e))
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
      method: '',
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
