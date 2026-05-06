"use client";
import Navbar from "@/components/navbar";
import Button from "@/components/Button";

export default function Home() {
	return (
		<>
			<div className="flex flex-col flex-1 min-h-screen bg-slate-50 dark:bg-slate-900">
				<Navbar />
				<main className="flex-1 w-full max-w-3xl mx-auto px-6 py-10 flex flex-col gap-8">
					<section className="flex gap-4">
						{/* landing page */}
						<div className="flex flex-col items-center justify-center gap-4 w-full h-[80vh]">
							<h1 className="text-6xl font-bold text-slate-700 dark:text-slate-300 mb-3">Welcome to Piggy</h1>
							<p className="text-2xl text-slate-600 dark:text-slate-400 mb-6">A simple app for tracking your savings and withdrawals.</p>
							<div className="flex gap-4 ">
								<Button text="Login" onClick={() => {
									window.location.href = "./login";
								}} />
								<Button text="Sign Up" variant="secondary" onClick={() => {
									window.location.href = "./signup";
								}} />
							</div>
						</div>
					</section>
				</main>
			</div>
		</>
	);
}
