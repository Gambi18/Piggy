"use client";
import { saveTransaction } from "@/api/create-transaction";
import Button from "@/components/Button";
import Navbar from "@/components/navmain";
import { TransactionType } from "@/types/interfaces";
import { useRouter, useParams } from "next/navigation";
import React, { useState } from "react";
import { toast } from "react-toastify";

function WithdrawPage() {
	const router = useRouter();
	const params = useParams();
	const [amount, setAmount] = useState("");
	const [reason, setReason] = useState("");
	// Handle withdrawal
	const handleWithdraw = async () => {
		console.log("Executed Withdrawal action!");
		const payload: TransactionType = {
			userId: params.id as string,
			amount: Number(amount),
			reason,
			type: "withdrawal",
		};
		const res = await saveTransaction(payload);

		console.log("Response from fetch: ", res);

		if (res.success) {
			toast("Withdrawal successful! Redirecting...");
			router.push(`/dashboard/${params.id}`);
		} else {
			toast.error("Failed to withdraw money! Try again.");
		}
	};

	return (
		<div className="flex flex-col flex-1 min-h-screen bg-slate-50 dark:bg-slate-900">
			<Navbar />
			<main className="flex flex-1 items-center justify-center px-6 py-10">
				<div className="w-full max-w-md bg-white dark:bg-slate-800 rounded-xl border border-slate-200 dark:border-slate-700 shadow-sm p-8">
					<h1 className="text-xl font-semibold text-slate-800 dark:text-slate-100 mb-6">
						Make a Withdrawal
					</h1>
					<form className="flex flex-col gap-4">
						<div className="flex flex-col gap-1">
							<label className="text-sm font-medium text-slate-600 dark:text-slate-400">
								Amount
							</label>
							<input
								className="border border-slate-200 dark:border-slate-600 bg-white dark:bg-slate-900 rounded-lg px-4 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 dark:text-slate-100"
								placeholder="e.g. 1000"
								value={amount}
								onChange={(v) => setAmount(v.target.value)}
							/>
						</div>
						<div className="flex flex-col gap-1">
							<label className="text-sm font-medium text-slate-600 dark:text-slate-400">
								Reason
							</label>
							<input
								className="border border-slate-200 dark:border-slate-600 bg-white dark:bg-slate-900 rounded-lg px-4 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 dark:text-slate-100"
								placeholder="e.g. Groceries"
								value={reason}
								onChange={(v) => setReason(v.target.value)}
							/>
						</div>
						<div className="pt-2">
							<Button text="Withdraw" onClick={handleWithdraw} />
						</div>
					</form>
				</div>
			</main>
		</div>
	);
}

export default WithdrawPage;
