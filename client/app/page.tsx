import { Button } from "@/components/ui/button";
import Link from "next/link";

export default function HomePage() {
    return (
        <Button>
            <Link href="/register">Register</Link>
        </Button>
    );
}
