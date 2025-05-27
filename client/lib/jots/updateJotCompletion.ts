import { Jot } from "@/types/jotTypes";

// Takes in a Jot[] and finds the target using binary search
// > needle should always be derived from an id in the haystack (Jot[])
//
export function updateJotCompletion(
    haystack: Jot[],
    left: number,
    right: number,
    needle: number,
) {
    const mid = left + Math.floor((right - left) / 2);

    if (haystack[mid].id === needle) {
        // change the current jot
        haystack[mid].isCompleted = !haystack[mid].isCompleted;
        return;
    } else if (haystack[mid].id < needle) {
        left = mid + 1;
    } else if (haystack[mid].id > needle) {
        right = mid;
    }

    updateJotCompletion(haystack, left, right, needle);
}
