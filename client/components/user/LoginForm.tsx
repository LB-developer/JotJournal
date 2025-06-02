"use client";

import { login } from "@/app/login/actions";
import { useActionState, useEffect, useTransition } from "react";
import { SubmitButton } from "./SubmitButton";
import getCurrentMonth from "@/lib/date/getCurrentMonth";
import { useAuth } from "@/context/AuthContext";
import { useRouter } from "next/navigation";

export default function LoginForm() {
    const router = useRouter();
    const { setUser } = useAuth();
    const [state, loginAction] = useActionState(login, undefined);
    const [, startTransition] = useTransition();

    useEffect(() => {
        if (state?.type === "success") {
            startTransition(() => {
                setUser(state.user);
                router.push(`/dashboard?month=${getCurrentMonth()}`);
            });
        }
    }, [state, setUser, router]);

    const handleRedirectToRegisterPage = (): void => {
        router.push("/register");
    };

    return (
        <div className="min-h-screen bg-gray-100 flex items-center justify-center p-4">
            <div className="max-w-md w-full bg-white rounded-xl shadow-lg p-8">
                <h2 className="text-2xl font-bold text-gray-900 mb-6 text-center">
                    Sign In
                </h2>

                <form action={loginAction} className="space-y-4">
                    <div>
                        <label className="block text-sm font-medium text-gray-700 mb-1">
                            Email
                            <input
                                id="email-input"
                                type="email"
                                name="email"
                                maxLength={50}
                                className="w-full px-4 py-2 text-gray-800 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 outline-none transition-all"
                                placeholder="your@email.com"
                            />
                        </label>
                    </div>
                    {state?.type == "error" && (
                        <p className="text-red-500">{state?.error}</p>
                    )}

                    <div>
                        <label className="block text-sm font-medium text-gray-700 mb-1">
                            Password
                            <input
                                id="password-input"
                                type="password"
                                name="password"
                                minLength={4}
                                className="w-full px-4 py-2 border text-gray-900 border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 outline-none transition-all"
                                placeholder="••••••••"
                            />
                        </label>
                    </div>

                    <div className="flex items-center justify-between">
                        <label className="flex items-center">
                            <input
                                type="checkbox"
                                className="rounded border-gray-300 text-indigo-600 focus:ring-indigo-500"
                            />
                            <span className="ml-2 text-sm text-gray-600">
                                Remember me
                            </span>
                        </label>
                        <a
                            href="#"
                            className="text-sm text-indigo-600 hover:text-indigo-500"
                        >
                            Forgot password?
                        </a>
                    </div>

                    <SubmitButton actionText="Login" />
                </form>

                <div className="flex gap-1 justify-center mt-6 text-center text-sm text-gray-600">
                    <p>{"Don't have an account?"}</p>
                    <span
                        onClick={handleRedirectToRegisterPage}
                        className="text-indigo-600 cursor-pointer hover:text-indigo-500 font-medium"
                    >
                        <p>Sign Up</p>
                    </span>
                </div>
            </div>
        </div>
    );
}
