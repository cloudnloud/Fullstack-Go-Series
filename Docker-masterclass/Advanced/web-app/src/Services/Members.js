import { AxiosInstance } from "./AxiosInstance";

const config = {
    headers: {
        "Content-Type": "application/json",
    }
}

export const getAllMembersApi = async () => {
    try {
        const response = await AxiosInstance.get("/members", config)
        return response;
    } catch (error) {
        console.log(error)
        return error
    }
}

export const getMemberApi = async () => {
    try {
        const response = await AxiosInstance.get("/members", config)
        return response;
    } catch (error) {
        console.log(error)
        return error
    }
}

export const createMemberApi = async (data) => {
    try {
        const response = await AxiosInstance.post("/member", data, config)
        return response;
    } catch (error) {
        console.log(error)
        return error
    }
}


export const updateMemberApi = async (id, data) => {
    try {
        const response = await AxiosInstance.put(`/update-member/${id}`,data, config)
        return response;
    } catch (error) {
        console.log(error)
        return error
    }
}

export const deleteMemberApi = async (id) => {
    try {
        const response = await AxiosInstance.delete(`/delete-member/${id}`,config)
        return response;
    } catch (error) {
        console.log(error)
        return error
    }
}