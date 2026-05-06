import { SignupType } from "@/types/interfaces";
import axios from "axios";

export interface ResponseType {
	success: boolean;
	error: unknown;
}

const BASEURL = "http://localhost:8081";

export const signup = async (payload: SignupType) => {
	console.log("Signup function executed!");
	const response: ResponseType = await axios
		.post(BASEURL + "/api/v1/signup", payload)
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
