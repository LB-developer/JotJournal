import { chromium } from "@playwright/test";
import dotenv from "dotenv";
import path from "path";

dotenv.config({ path: path.resolve(__dirname, "..", ".env") });

export default async () => {
    const b = await chromium.launch();
    const p = await b.newPage();
    await p.goto(process.env.BASE_URL ?? "http://localhost:3000/login");
    await p.getByLabel("Email").fill(process.env.E2E_EMAIL!);
    await p.getByLabel("Password").fill(process.env.E2E_PASSWORD!);
    await p.getByRole("button", { name: "Login" }).click();
    await p.waitForURL("**/dashboard*");
    await p.context().storageState({ path: "storageState.json" });
    console.log("User logged in and storage state saved");
    await b.close();
};
