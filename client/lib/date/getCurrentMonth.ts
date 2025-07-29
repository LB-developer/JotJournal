export function getCurrentMonth(): string {
    const date = new Date();
    const month = date.getMonth() + 1; // Months are 0-indexed

    return month.toString().padStart(2, "0");
}
