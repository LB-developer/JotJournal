export interface UserLoginPayload {
  email : string
  password : string
}

export type ApiError = { type: "error"; error: string }
export type ApiSuccess = { type: "success"}
export type LoginToken = { sessionToken: string }
