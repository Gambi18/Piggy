"use client";
import React from "react";

export interface BtnProps {
	onClick: () => void;
	text: string;
	variant?: "primary" | "secondary";
	className?: string;
}

function Button({
	onClick,
	text,
	variant = "primary",
	className,
}: BtnProps) {
	// Base Styles
	const base =
		"py-2 px-5 rounded-lg font-medium text-sm transition-colors focus:outline-none focus:ring-2 focus:ring-offset-2";
	// Variant styles
	const styles = {
		primary: `${base} bg-blue-600 hover:bg-blue-700 text-white focus:ring-blue-500`,
		secondary: `${base} bg-slate-100 hover:bg-slate-200 text-slate-700 dark:bg-slate-700 dark:hover:bg-slate-600 dark:text-slate-200 focus:ring-slate-400`,
	};

	return (
		<button
			type="button"
			className={className ? className : styles[variant]}
			onClick={onClick}
		>
			{text}
		</button>
	);
}

export default Button;
