"use client";
import Navbar from "@/components/navmain";
import TransactionsList from "@/components/TransactionsList";
import Button from "@/components/Button";
import { useRouter, useParams } from "next/navigation";
import { useGetTransactions } from "@/hooks/useFetchTransactions";

export async function generateStaticParams() {
	return [
		{ id: '1' },
		{ id: '2' },
		{ id: '3' }
	];
}

export default function WithdrawalsPage() {
	const params = useParams();
	const router = useRouter();
	const withdrawalTransactions = useGetTransactions({
		userId: params.id as string,
		type: "withdrawal",
		size: 100 // Get all withdrawal transactions
	});

	const handleBack = () => {
		router.push(`/dashboard/${params.id}`);
	};

	return (
		<div className="flex flex-col flex-1 min-h-screen bg-slate-50 dark:bg-slate-900">
			<Navbar />
			<main className="flex-1 w-full max-w-4xl mx-auto px-6 py-10 flex flex-col gap-6">
				<section className="flex items-center gap-4">
					<Button
						text="← Back"
						variant="secondary"
						onClick={handleBack}
					/>
					<h1 className="text-2xl font-bold text-slate-800 dark:text-slate-100">
						Withdrawal Transactions
					</h1>
				</section>

				<section>
					<div className="bg-white dark:bg-slate-800 rounded-xl border border-slate-200 dark:border-slate-700 shadow-sm p-6">
						<h2 className="text-lg font-semibold text-slate-700 dark:text-slate-300 mb-4">
							All Withdrawals
						</h2>
						{Array.isArray(withdrawalTransactions) && withdrawalTransactions.length > 0 ? (
							<TransactionsList transactions={withdrawalTransactions} />
						) : (
							<div className="text-center py-8 text-slate-500 dark:text-slate-400">
								<p>No withdrawal transactions found.</p>
								<Button
									text="Make Withdrawal"
									onClick={() => router.push(`/dashboard/${params.id}/withdraw`)}
									className="mt-4"
								/>
							</div>
						)}
					</div>
				</section>
			</main>
		</div>
	);
}
