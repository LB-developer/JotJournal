export class ApiError extends Error {
    constructor(
        message: string,
        public status?: number,
        public details?: unknown,
    ) {
        super(message);
        this.name = "ApiError";
    }
}
