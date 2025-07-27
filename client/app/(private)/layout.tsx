import { AppSidebar } from "@/components/app-sidebar";
import SidebarLayout from "@/components/SidebarLayout";
import { SidebarProvider, SidebarTrigger } from "@/components/ui/sidebar";

export default function ProtectedLayout({
    children,
}: Readonly<{
    children: React.ReactNode;
}>) {
    return (
        <SidebarProvider>
            <AppSidebar />
            <main>
                <SidebarTrigger />
                <SidebarLayout>{children}</SidebarLayout>
            </main>
        </SidebarProvider>
    );
}
