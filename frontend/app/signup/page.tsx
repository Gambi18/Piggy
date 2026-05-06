"use client";
import React, { useState } from "react";
import { useRouter } from "next/navigation";
import Navbar from "@/components/navbar";
import Button from "@/components/Button";
import { toast } from "react-toastify";
import { signup } from "@/api/signup";
import { SignupType } from "@/types/interfaces";
import Link from "next/link";

function SignupPage() {
	const router = useRouter();
	const [username, setUsername] = useState("");
	const [name, setName] = useState("");
	const [email, setEmail] = useState("");
	const [password, setPassword] = useState("");

	// Handle signup
	const handleSignup = async () => {
		console.log("Executed!");
		const payload: SignupType = {
			username,
			name,
			email,
			password,
		};
		const res = await signup(payload);

		console.log("Response from fetch: ", res);

		if (res.success) {
			toast("Signup successful! Redirecting to login...");
			router.push("/login");
		} else {
			toast.error("Failed to signup! Try again.");
		}
	};

	return (
		<div className="flex flex-col flex-1 min-h-screen bg-slate-50 dark:bg-slate-900">
			<Navbar />
			<main className="flex flex-1 items-center justify-center px-6 py-10">
				<div className="w-full max-w-md bg-white dark:bg-slate-800 rounded-xl border border-slate-200 dark:border-slate-700 shadow-sm p-8">
					<h1 className="text-xl font-semibold text-slate-800 dark:text-slate-100 mb-6">
						Create an Account
					</h1>
					<form className="flex flex-col gap-4">
						<div className="flex flex-col gap-1">
							<label className="text-sm font-medium text-slate-600 dark:text-slate-400">
								Full Name
							</label>
							<input
								className="border border-slate-200 dark:border-slate-600 bg-white dark:bg-slate-900 rounded-lg px-4 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 dark:text-slate-100"
								placeholder="John Doe"
								value={name}
								onChange={(e) => setName(e.target.value)}
							/>
						</div>
						<div className="flex flex-col gap-1">
							<label className="text-sm font-medium text-slate-600 dark:text-slate-400">
								Username
							</label>
							<input
								className="border border-slate-200 dark:border-slate-600 bg-white dark:bg-slate-900 rounded-lg px-4 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 dark:text-slate-100"
								placeholder="johndoe123"
								value={username}
								onChange={(e) => setUsername(e.target.value)}
							/>
						</div>
						<div className="flex flex-col gap-1">
							<label className="text-sm font-medium text-slate-600 dark:text-slate-400">
								Email
							</label>
							<input
								className="border border-slate-200 dark:border-slate-600 bg-white dark:bg-slate-900 rounded-lg px-4 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 dark:text-slate-100"
								placeholder="john@example.com"
								type="email"
								value={email}
								onChange={(e) => setEmail(e.target.value)}
							/>
						</div>
						<div className="flex flex-col gap-1">
							<label className="text-sm font-medium text-slate-600 dark:text-slate-400">
								Password
							</label>
							<input
								className="border border-slate-200 dark:border-slate-600 bg-white dark:bg-slate-900 rounded-lg px-4 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 dark:text-slate-100"
								placeholder="••••••••"
								type="password"
								value={password}
								onChange={(e) => setPassword(e.target.value)}
							/>
						</div>
						<div className="pt-2 flex flex-col gap-3">
							<Button text="Sign Up" onClick={handleSignup} />
							<Link
								href="/login"
								className="text-center text-blue-500 hover:text-blue-600 text-sm font-medium underline"
							>
								Already have an account? Login here
							</Link>
						</div>
					</form>
				</div>
			</main>
		</div>
	);
}

export default SignupPage;
