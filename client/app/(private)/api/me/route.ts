import { cookies } from "next/headers";

export async function GET() {
    const cookieStore = await cookies();
    const userCookie = cookieStore.get("user");

    if (!userCookie) {
        return new Response("Unauthorized", { status: 401 });
    }

    return Response.json(JSON.parse(userCookie.value));
}
