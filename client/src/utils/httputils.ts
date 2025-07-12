import axios from "axios";

const api = axios.create({
  baseURL: "http://localhost:8080",
  headers: {
    "Content-Type": "application/json",
  },
});

export const postAPI = async <T = unknown, R = unknown>(
  url: string,
  payload: T
): Promise<R> => {
  const response = await api.post(url, payload);
  if (response.status !== 200) {
    throw new Error("Failed API call");
  }
  return response.data;
};
