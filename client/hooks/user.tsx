"use client";
import { logout } from "@/app/(public)/login/actions";
import { useAuth } from "@/context/AuthContext";
import { redirect } from "next/navigation";
import { useEffect } from "react";

// ensures the user is stored in context
// if unable to find force user to login again
export function useEstablishUser() {
    const context = useAuth();

    useEffect(() => {
        if (!context.user) {
            fetch("/api/me")
                .then((res) => {
                    if (!res.ok) throw new Error("Unauthorized");
                    return res.json();
                })
                .then((user) => context.setUser(user))
                .catch(() => {
                    console.error("No user found, redirecting...");
                    logout();
                    redirect("/login");
                });
        }
    }, [context.user]);
}
