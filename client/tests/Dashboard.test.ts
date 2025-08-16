import test, { expect } from "@playwright/test";
import { randomUUID } from "crypto";

test.describe("Dashboard", () => {
    test("adding a jot", async ({ page }) => {
        const randomJotName = randomUUID().slice(0, 10);

        await page.goto("/dashboard");

        const createJotButton = page.getByTestId("create-jot-button");
        await expect(createJotButton).toBeVisible();
        await createJotButton.click();

        const jotNameInput = page.getByTestId("jot-name-input");
        await expect(jotNameInput).toBeVisible();
        await jotNameInput.fill(randomJotName);

        await page.getByRole("button", { name: "Add" }).click();

        await expect(page.getByText(randomJotName)).toBeVisible();
    });

    test("deleting a jot", async ({ page }) => {
        await page.goto("/dashboard");

        const randomJotName = randomUUID().slice(0, 10);

        const createJotButton = page.getByTestId("create-jot-button");
        await createJotButton.click();

        await page.getByTestId("jot-name-input").fill(randomJotName);
        await page.getByRole("button", { name: "Add" }).click();

        await expect(page.getByText(randomJotName)).toBeVisible();

        await page.getByTestId(`delete-${randomJotName}-button`).click();
        const deleteButton = page.getByRole("button", { name: "Delete" });

        await deleteButton.click();

        await expect(deleteButton).not.toBeVisible();

        await expect(page.getByText(randomJotName)).not.toBeVisible();
    });
});
