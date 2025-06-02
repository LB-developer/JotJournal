import { User } from "./userTypes";

export interface UserLoginPayload {
    email: string;
    password: string;
}

export interface SessionToken {
    sessionToken: string;
}

export type ApiError = { type: "error"; error: string };
export type ApiSuccess = { type: "success"; user: User };
export type LoginSuccess = { sessionToken: string; user: User };
