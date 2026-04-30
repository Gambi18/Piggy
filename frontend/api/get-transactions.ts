import { GetTransactionsParamsType, TransactionType } from "@/types/interfaces";
import axios from "axios";


export interface GetTransactionsRes {
	transactions: TransactionType[];
	error: unknown;
}

const BASEURL = "http://localhost:8081";
export const getAllTransactions = async (
	query: GetTransactionsParamsType,
): Promise<GetTransactionsRes> => {
	const queries = [];
	if (query.size) {
		queries.push("size=" + query.size);
	}

	if (query.type) {
		queries.push("type=" + query.type);
	}

	const url =
		BASEURL +
		"/api/v1/transactions" +
		(queries.length > 0 ? "?" : "") +
		queries.join("&");

	try {
		const res = await axios.get(url);

		return {
			transactions: res.data as TransactionType[],
			error: null,
		};
	} catch (error) {
		return {
			transactions: [],
			error,
		};
	}
};

