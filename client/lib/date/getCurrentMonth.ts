export function getCurrentMonth(): string {
    const date = new Date();
    const month = date.getMonth();
    return month.toString().padStart(2, "0");
}
