export default function getDaysInMonth(year: number, month: number): number {
    // months are 0 indexed
    month -= 1;
    return new Date(year, month, 0).getDate();
}
