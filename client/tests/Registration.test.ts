import { test, expect } from "@playwright/test";
import { randomUUID } from "crypto";

test.describe("Registration", () => {
    test("navigates from login to registration, submits new user data, and reaches dashboard", async ({
        page,
    }) => {
        await page.goto("http://localhost:3000/login");

        const loginTitle = page.getByText("Sign in");
        const loginEmailField = page.getByLabel("Email");
        const loginPasswordField = page.getByLabel("Password");

        // Should render 'login' page
        await expect(loginTitle).toBeVisible();
        await expect(loginEmailField).toBeVisible();
        await expect(loginPasswordField).toBeVisible();

        // Move to register page
        const signUpButton = page.getByText("Sign up");
        await signUpButton.click();
        await expect(page).toHaveURL("http://localhost:3000/register");

        // Assign new user inputs
        const newEmail = randomUUID() + "@gmail.com";
        const newPassword = randomUUID();
        const newFirstName = "John";
        const newLastName = "Smith";

        // Fill new user
        const emailInputBox = page.getByLabel("Email");
        const firstNameInputBox = page.getByLabel("First Name");
        const lastNameInputBox = page.getByLabel("Last Name");
        const passwordInputBox = page.getByLabel("Password");

        await emailInputBox.fill(newEmail);
        await firstNameInputBox.fill(newFirstName);
        await lastNameInputBox.fill(newLastName);
        await passwordInputBox.fill(newPassword);

        await expect(emailInputBox).toHaveValue(newEmail);
        await expect(firstNameInputBox).toHaveValue(newFirstName);
        await expect(lastNameInputBox).toHaveValue(newLastName);
        await expect(passwordInputBox).toHaveValue(newPassword);

        // Submit registration for new user
        const registerButton = page.getByRole("button", { name: "Register" });
        await Promise.all([
            page.waitForURL("**/dashboard*"),
            registerButton.click(),
        ]);

        await expect(page).toHaveURL(/http:\/\/localhost:3000\/dashboard*/i);
    });
});
