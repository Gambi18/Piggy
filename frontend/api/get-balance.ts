import axios from "axios";

const BASEURL = "http://localhost:8081";

export const getBalance = async (userId: string) => {
	try {
		const response = await axios.get(`${BASEURL}/api/v1/balance?userId=${userId}`);
		return {
			data: response.data,
			success: true,
			error: undefined,
		};
	} catch (err) {
		console.error("Error fetching balance:", err);
		return {
			data: { balance: 0, totalSavings: 0, totalWithdrawals: 0 },
			success: false,
			error: err,
		};
	}
};
