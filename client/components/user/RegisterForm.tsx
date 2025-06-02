"use client";

import { register } from "@/app/register/actions";
import { redirect } from "next/navigation";
import { useActionState } from "react";
import { SubmitButton } from "./SubmitButton";

export default function RegisterForm() {
    const [state, registerAction] = useActionState(register, undefined);

    if (state?.type == "success") {
        redirect("/dashboard");
    }

    const handleRedirectToLoginPage = (): void => {
        redirect("/login");
    };

    return (
        <div className="min-h-screen bg-gray-100 flex items-center justify-center p-4">
            <div className="max-w-md w-full bg-white rounded-xl shadow-lg p-8">
                <h2 className="text-2xl font-bold text-gray-900 mb-6 text-center">
                    Register
                </h2>

                <form action={registerAction} className="space-y-4">
                    <div>
                        <label className="block text-sm font-medium text-gray-700 mb-1">
                            Email
                            <input
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
                            First Name
                            <input
                                id="firstName"
                                type="text"
                                name="firstName"
                                maxLength={15}
                                className="w-full px-4 py-2 text-gray-800 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 outline-none transition-all"
                                placeholder="John"
                            />
                        </label>
                    </div>
                    <div>
                        <label className="block text-sm font-medium text-gray-700 mb-1">
                            Last Name
                            <input
                                id="lastName"
                                type="text"
                                name="lastName"
                                maxLength={15}
                                className="w-full px-4 py-2 text-gray-800 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 outline-none transition-all"
                                placeholder="Smith"
                            />
                        </label>
                    </div>
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

                    <SubmitButton actionText="Register" />
                </form>

                <div className="flex gap-1 justify-center mt-6 text-center text-sm text-gray-600">
                    <p>Already have an account?</p>
                    <span
                        onClick={handleRedirectToLoginPage}
                        className="text-indigo-600 cursor-pointer hover:text-indigo-500 font-medium"
                    >
                        <p>Sign in</p>
                    </span>
                </div>
            </div>
        </div>
    );
}
