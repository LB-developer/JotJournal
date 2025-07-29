"use client";

import * as React from "react";
import { Home } from "lucide-react";

import { NavUser } from "@/components/nav-user";
import {
    Sidebar,
    SidebarContent,
    SidebarFooter,
    SidebarGroup,
    SidebarGroupContent,
    SidebarGroupLabel,
    SidebarMenu,
    SidebarMenuButton,
    SidebarMenuItem,
} from "@/components/ui/sidebar";
import { useAuth } from "@/context/AuthContext";

// This is sample data.
const data = {
    defaultUser: {
        name: "shadcn",
        email: "m@example.com",
        avatar: "",
    },
    items: [
        {
            title: "Dashboard",
            url: "/dashboard",
            icon: Home,
        },
    ],
};

export function AppSidebar() {
    const { user } = useAuth();

    let sideBarInfoUser;
    if (!user) {
        sideBarInfoUser = data.defaultUser;
    } else {
        sideBarInfoUser = {
            name: user?.firstName,
            email: user?.email,
            avatar: "",
        };
    }

    return (
        <Sidebar>
            <SidebarContent>
                <SidebarGroup>
                    <SidebarGroupLabel>Jot Journal</SidebarGroupLabel>
                    <SidebarGroupContent>
                        <SidebarMenu>
                            {data.items.map((item) => (
                                <SidebarMenuItem key={item.title}>
                                    <SidebarMenuButton asChild>
                                        <a href={item.url}>
                                            <item.icon />
                                            <span>{item.title}</span>
                                        </a>
                                    </SidebarMenuButton>
                                </SidebarMenuItem>
                            ))}
                        </SidebarMenu>
                    </SidebarGroupContent>
                </SidebarGroup>
            </SidebarContent>
            <SidebarFooter>
                <NavUser user={sideBarInfoUser} />
            </SidebarFooter>
        </Sidebar>
    );
}
