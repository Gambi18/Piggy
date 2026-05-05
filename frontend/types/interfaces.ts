export interface TransactionType {
	userId: string;
	amount: number;
	reason: string;
	type: "saving" | "withdrawal";
	id?: string;
	createdAt?: string;
}

export interface GetTransactionsParamsType {
	userId: string;
	type?: "saving" | "withdrawal";
	size?: number;
}

export interface LoginType {
    username: string;
    password: string;
}

export interface SignupType {
    username: string;
    name: string;
    email: string;
    password: string;
}
