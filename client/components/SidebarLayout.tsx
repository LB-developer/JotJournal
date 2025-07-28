"use client";

import { useSidebar } from "@/components/ui/sidebar";

export default function SidebarLayout({
    children,
}: {
    children: React.ReactNode;
}) {
    const { open } = useSidebar();

    return (
        <div className="flex w-full overflow-hidden ">
            {/* Main content fills remaining space */}
            <div
                className={`transition-all duration-300 ease-in-out 
                            w-screen
                            ${
                                open
                                    ? "md:w-[calc(100vw-16rem)]"
                                    : "md:w-[calc(100vw-3rem)]"
                            }`}
            >
                {children}
            </div>
        </div>
    );
}
