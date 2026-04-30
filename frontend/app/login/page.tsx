"use client";
import React from "react";
import { useRouter } from "next/navigation";
import Navbar from "@/components/navbar";
import Button from "@/components/Button";
import { toast } from "react-toastify";
import { LoginType } from "@/types/interfaces";
import { login } from "@/api/login";
import Link from "next/link";

function LoginPage() {
    const router = useRouter();
    const [username, setUsername] = React.useState("");
    const [password, setPassword] = React.useState("");

    // Handle login
    const handleLogin = async () => {
        console.log("Executed!");
        const payload: LoginType = {
            username,
            password,
        };
        const res = await login(payload);

        console.log("Response from fetch: ", res);

        if (res.success) {
            toast("Login successful! Redirecting...");
            router.push("/");
        } else {
            toast.error("Failed to login! Try again.");
        }
    }
    return (
        <>
            <div className="flex flex-col flex-1 min-h-screen bg-slate-50 dark:bg-slate-900">
                <Navbar />
                <main className="flex flex-1 items-center justify-center px-6 py-10">
                    <div className="w-full max-w-md bg-white dark:bg-slate-800 rounded-xl border border-slate-200 dark:border-slate-700 shadow-sm p-8">
                        <h1 className="text-xl font-semibold text-slate-800 dark:text-slate-100 mb-6">
                            Login
                        </h1>
                        <form className="flex flex-col gap-4">
                            <div className="flex flex-col gap-1">
                                <label className="text-sm font-medium text-slate-600 dark:text-slate-400">
                                    Username
                                </label>
                                <input
                                    className="border border-slate-200 dark:border-slate-600 bg-white dark:bg-slate-900 rounded-lg px-4 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 dark:text-slate-100"
                                    placeholder="Username"
                                    value={username}
                                    onChange={(e) => setUsername(e.target.value)}
                                />
                            </div>
                            <div className="flex flex-col gap-1">
                                <label className="text-sm font-medium text-slate-600 dark:text-slate-400">
                                    Password
                                </label>
                                <input
                                    className="border border-slate-200 dark:border-slate-600 bg-white dark:bg-slate-900 rounded-lg px-4 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 dark:text-slate-100"
                                    placeholder="password"
                                    type="password"
                                    value={password}
                                    onChange={(e) => setPassword(e.target.value)}
                                />
                            </div>
                            <div className="pt-2 flex justify-between">
                                <Button text="Login" onClick={handleLogin} />
                                <Link
                                    href="./signup"
                                    className="mt-3 inline-block text-center text-blue-500 hover:text-blue-600 text-sm font-medium underline"
                                >
                                    Don&apos;t have an account? Sign up here
                                </Link>
                            </div>
                        </form>
                    </div>
                </main>
            </div>
        </>
    );
}

export default LoginPage;
