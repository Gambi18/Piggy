"use client";
import StatsCard from "@/components/Card";
import Navbar from "@/components/navmain";
import TransactionsList from "@/components/TransactionsList";
import Button from "@/components/Button";
import { useRouter, useParams } from "next/navigation";
import { useGetTransactions } from "@/hooks/useFetchTransactions";
import { useGetBalance } from "@/hooks/useFetchBalance";

export async function generateStaticParams() {
	return [
		{ id: '1' },
		{ id: '2' },
		{ id: '3' }
	];
}

export default function Home() {
	const params = useParams();
	const router = useRouter();
	const recentTransactions = useGetTransactions({ userId: params.id as string, size: 5 });
	const user = useGetBalance({ id: params.id });
	console.log("In the component: ", recentTransactions);
	return (
		<>
			<div className="flex flex-col flex-1 min-h-screen bg-slate-50 dark:bg-slate-900">
				<Navbar />
				<main className="flex-1 w-full max-w-3xl mx-auto px-6 py-10 flex flex-col gap-8">
					<section className="flex gap-4">
						<StatsCard title="Total Savings" text={`${user.totalSavings.toLocaleString()} CFA`} />
						<StatsCard
							title="Total Withdrawals"
							text={`${user.totalWithdrawals.toLocaleString()} CFA`}
						/>
					</section>
					<section className="flex justify-between">
						<StatsCard title="Balance" text={`${user.balance.toLocaleString()} CFA`} />
					</section>
					<section className="flex gap-3">
						<Button
							text="Add Savings"
							onClick={() => {
								router.push(`/dashboard/${params.id}/save`);
							}}
						/>
						<Button
							text="Make Withdrawal"
							variant="secondary"
							onClick={() => {
								router.push(`/dashboard/${params.id}/withdraw`);
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
