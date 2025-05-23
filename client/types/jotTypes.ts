export type JotCollection = Record<string, Jot[]>;

export interface Jot {
  date: string
  habit: string
  id: number
  isCompleted: boolean
}

export interface UpdateJotPayload {
  id: number
  isCompleted: boolean
}
