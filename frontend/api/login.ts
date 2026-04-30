import { LoginType } from "@/types/interfaces";
import axios from "axios";

export interface ResponseType {
	success: boolean;
	error: unknown;
}

const BASEURL = "http://localhost:8081";

export const login = async (payload: LoginType) => {
	console.log("Login function executed!");
	const response: ResponseType = await axios
		.post(BASEURL + "/api/v1/login", payload)
		.then(() => {
			return {
				error: undefined,
				success: true,
			};
		})
		.catch((err) => {
			return {
				error: err,
				success: false,
			};
		});
	return response;
};
