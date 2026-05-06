import { getAllTransactions } from "@/api/get-transactions";
import { GetTransactionsParamsType, TransactionType } from "@/types/interfaces";
import { useEffect, useState } from "react";

export const useGetTransactions = (query: GetTransactionsParamsType) => {
	const [transactions, setTransactions] = useState<TransactionType[]>([]);
	useEffect(() => {
		const fetchTransactions = async () => {
			const res = await getAllTransactions(query);
			console.log(res);
			setTransactions(res.transactions);
		};

		fetchTransactions();
	}, [query.size, query.type, query.userId]);

	return transactions;
};
