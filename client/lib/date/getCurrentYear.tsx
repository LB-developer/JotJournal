export function getCurrentYear(): string {
    const date = new Date();
    const year = date.getFullYear();
    return String(year);
}
