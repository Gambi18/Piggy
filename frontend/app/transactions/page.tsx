"use client";
import { Suspense } from "react";
import TransactionsList from "@/components/TransactionsList";
import { useGetTransactions } from "@/hooks/useFetchTransactions";
import { GetTransactionsParamsType } from "@/types/interfaces";
import { useSearchParams } from "next/navigation";
import React from "react";

export default function TransactionsPage() {
	return (
		<Suspense fallback={<div>Loading transactions...</div>}>
			<TransactionsPageContent />
		</Suspense>
	);
}

function TransactionsPageContent() {
	const query = useSearchParams();
	const type = query.get("type");
	const filtered = useGetTransactions({
		type: type,
	} as GetTransactionsParamsType);

	console.log("IN Transactions: ", filtered);

	return (
		<div className="flex flex-col flex-1 min-h-screen bg-slate-50 dark:bg-slate-900">
			<main className="w-full max-w-3xl mx-auto px-6 py-10">
				<h1 className="text-xl font-semibold text-slate-700 dark:text-slate-300 mb-4">
					All Transactions
				</h1>
				<TransactionsList transactions={filtered} />
			</main>
		</div>
	);
}

