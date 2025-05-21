'use client';

import { login } from "@/app/login/actions";
import { redirect } from "next/navigation";
import React, { useActionState} from "react";
import { useFormStatus } from "react-dom";

export default function LoginForm() {
  const [state, loginAction] = useActionState(login, undefined);

  if (state?.type == "success") {
    redirect("/")
  }

  return (
  <div className="min-h-screen bg-gray-100 flex items-center justify-center p-4">
    <div className="max-w-md w-full bg-white rounded-xl shadow-lg p-8">
      <h2 className="text-2xl font-bold text-gray-900 mb-6 text-center">Sign In</h2>
      
      <form action={loginAction} className="space-y-4">
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Email</label>
          <input 
            type="email" 
            name="email"
            maxLength={50}
            className="w-full px-4 py-2 text-gray-800 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 outline-none transition-all"
            placeholder="your@email.com"
          />
        </div>
        {state?.type == "error" && (
          <p className="text-red-500">{state?.error}</p>
        )}

        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Password</label>
          <input 
            type="password" 
            name="password"
            minLength={4}
            className="w-full px-4 py-2 border text-gray-900 border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 outline-none transition-all"
            placeholder="••••••••"
          />
        </div>

        <div className="flex items-center justify-between">
          <label className="flex items-center">
            <input type="checkbox" className="rounded border-gray-300 text-indigo-600 focus:ring-indigo-500"/>
            <span className="ml-2 text-sm text-gray-600">Remember me</span>
          </label>
          <a href="#" className="text-sm text-indigo-600 hover:text-indigo-500">Forgot password?</a>
        </div>

        <SubmitButton />
      </form>

      <div className="flex gap-1 justify-center mt-6 text-center text-sm text-gray-600">
        <p>Don't have an account?</p>
        <a href="#" className="text-indigo-600 hover:text-indigo-500 font-medium"><p>Sign up</p></a>
      </div>
    </div>
  </div>
  )
}

function SubmitButton() {
    const { pending } = useFormStatus()

    return (
        <button 
          className="w-full bg-indigo-600 hover:bg-indigo-700 text-white font-medium py-2.5 rounded-lg transition-colors"
          disabled={pending} 
          type="submit"
        >
          {pending ? "Submitting" : "Sign in"}
        </button>
    )
}
