import { useState, useEffect } from "react";
import { getBalance } from "@/api/get-balance";

export const useGetBalance = (params: { id: string | string[] | undefined }) => {
	const [user, setUser] = useState({ balance: 0, totalSavings: 0, totalWithdrawals: 0 });

	useEffect(() => {
		if (!params.id) return;
		
		// Simple UUID format check (8-4-4-4-12 hex chars)
		const uuidRegex = /^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/i;
		if (!uuidRegex.test(params.id as string)) {
			console.warn(`Skipping balance fetch for invalid ID: ${params.id}`);
			return;
		}

		const fetchBalance = async () => {
			const res = await getBalance(params.id as string);
			if (res.success) {
				setUser(res.data);
			}
		};

		fetchBalance();
	}, [params.id]);

	return user;
};
