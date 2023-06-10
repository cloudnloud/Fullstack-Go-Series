import axios from "axios";

export const AxiosInstance = axios.create({
    baseURL: process.env.REACT_APP_API_URL
});