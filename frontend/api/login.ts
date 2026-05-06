import { LoginType } from "@/types/interfaces";
import axios from "axios";

export interface ResponseType {
	success: boolean;
	error: unknown;
	data?: any; // eslint-disable-line @typescript-eslint/no-explicit-any
}

const BASEURL = "http://localhost:8081";

export const login = async (payload: LoginType) => {
	console.log("Login function executed!");
	const response: ResponseType = await axios
		.post(BASEURL + "/api/v1/login", payload)
		.then((res) => {
			return {
				error: undefined,
				success: true,
				data: res.data,
			};
		})
		.catch((err) => {
			console.log(err);
			return {
				error: err,
				success: false,
				data: undefined,
			};
		});
	return response;
};
