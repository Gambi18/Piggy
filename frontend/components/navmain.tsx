import Link from "next/link";
import React from "react";

function Navbar() {
    return (
        <header className="w-full border-b border-slate-200 dark:border-slate-700 bg-white dark:bg-slate-800 shadow-sm">
            <div className="max-w-3xl mx-auto flex items-center justify-between px-6 py-4">
                <Link
                    href="/"
                    className="text-2xl font-bold text-blue-600 dark:text-blue-400 tracking-tight"
                >
                    🐷 Piggy
                </Link>
                <nav className="flex items-center gap-6">
                    <ul className="flex gap-6 text-sm font-medium text-slate-600 dark:text-slate-300">
                        <li>
                            <Link
                                href="/transactions?type=withdrawal"
                                className="hover:text-blue-600 dark:hover:text-blue-400 transition-colors"
                            >
                                Withdrawals
                            </Link>
                        </li>
                        <li>
                            <Link
                                href="/transactions?type=saving"
                                className="hover:text-blue-600 dark:hover:text-blue-400 transition-colors"
                            >
                                Savings
                            </Link>
                        </li>
                    </ul>
                    <div className="flex items-center gap-4 text-sm font-medium">
                        <Link
                            href="/"
                            className="text-slate-500 hover:text-red-500 dark:text-slate-400 dark:hover:text-red-400 transition-colors"
                        >
                            Logout
                        </Link>
                    </div>
                </nav>
            </div>
        </header>
    );
}

export default Navbar;
