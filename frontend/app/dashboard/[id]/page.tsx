"use client";
import StatsCard from "@/components/Card";
import Navbar from "@/components/navmain";
import TransactionsList from "@/components/TransactionsList";
import Button from "@/components/Button";
import { useGetTransactions } from "@/hooks/useFetchTransactions";

export default function Home() {
	const recentTransactions = useGetTransactions({ size: 5 });
	console.log("In the component: ", recentTransactions);
	return (
		<>
			<div className="flex flex-col flex-1 min-h-screen bg-slate-50 dark:bg-slate-900">
				<Navbar />
				<main className="flex-1 w-full max-w-3xl mx-auto px-6 py-10 flex flex-col gap-8">
					<section className="flex gap-4">
						<StatsCard title="Total Savings" text="53,000 CFA" />
						<StatsCard
							title="Total Withdrawals"
							text="100,000 CFA"
						/>
					</section>
					<section className="flex justify-between">
						<StatsCard title="Balance" text="53,000 CFA" />
					</section>
					<section className="flex gap-3">
						<Button
							text="Add Savings"
							onClick={() => {
								window.location.href = "./save";
							}}
						/>
						<Button
							text="Make Withdrawal"
							variant="secondary"
							onClick={() => {
								window.location.href = "./withdraw";
							}}
						/>
					</section>
					<section>
						<h2 className="text-lg font-semibold text-slate-700 dark:text-slate-300 mb-3">
							Recent Transactions
						</h2>
						<TransactionsList transactions={recentTransactions} />
					</section>
				</main>
			</div>
		</>
	);
}
