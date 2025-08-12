export type JotCollection = Record<string, Jot[]>;

export type Jot = {
    date: string;
    habit: string;
    id: number;
    isCompleted: boolean;
};

export type UpdateJotPayload = {
    jotID: number;
    isCompleted: boolean;
};

export type CreateJotPayload = {
    name: string;
    month: string;
    year: string;
};

export type DeleteJotPayload = {
    habit: string;
    month: number;
    year: number;
};

export type JotActionBody =
    | { jotID: number; isCompleted?: boolean }
    | { habit: string; month: number; year: number };
