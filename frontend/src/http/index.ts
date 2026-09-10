import { getToken } from '@/storage/token';
import axios from 'axios';

import type { AxiosError, AxiosResponse, AxiosRequestConfig } from 'axios';

interface ApiResponse<T = any> {
    code: number,
    data: T,
    msg: string
}

const instance = axios.create({
    baseURL: "/api",
    timeout: 10000,
    headers: {
        "Content-Type": "application/json;charset=utf-8"
    }

})

instance.interceptors.request.use(function (config) {

    if (getToken()) {

        Object.assign(config.headers, {
            Authorization: `Bearer ${getToken()}`
        })

    }


    return config
}, (err) => {
    return Promise.reject(err)
})


instance.interceptors.response.use((response: AxiosResponse<ApiResponse>) => {


    if (response.config.responseType === "text" && response.status === 200) {
        return response.data as any
    }
    const { code, data, msg } = response.data
    if (code === 200) {


        return data as any

    }

    if (code === 401) {

    }

    return Promise.reject(msg || 'Error')


}, (error: AxiosError) => {

    // 处理 HTTP 网络错误 (非 200 状态码)
    let errorMessage = '网络请求异常，请稍后重试！';

    if (error.response) {
        switch (error.response.status) {
            case 401:
                errorMessage = '登录已过期，请重新登录';
                // 清除 token 并跳转等...
                break;
            case 403:
                errorMessage = '拒绝访问 (403)';
                break;
            case 404:
                errorMessage = '请求的资源不存在 (404)';
                break;
            case 500:
                errorMessage = '服务器内部错误 (500)';
                break;
            default:
                errorMessage = `连接错误 (${error.response.status})`;
        }
    } else if (error.message.includes('timeout')) {
        errorMessage = '网络请求超时';
    } else if (error.message.includes('Network Error')) {
        errorMessage = '网络连接断开';
    }

    console.error(errorMessage);
    // ElMessage.error(errorMessage);
    return Promise.reject(error);

})

// 4. 导出基于 Promise 的 HTTP 工具方法
const http = {
    get<T = any>(url: string, params?: object, config?: AxiosRequestConfig): Promise<T> {
        return instance.get(url, { params, ...config });
    },

    post<T = any>(url: string, data?: object, config?: AxiosRequestConfig): Promise<T> {
        return instance.post(url, data, config);
    },

    put<T = any>(url: string, data?: object, config?: AxiosRequestConfig): Promise<T> {
        return instance.put(url, data, config);
    },

    delete<T = any>(url: string, params?: object, config?: AxiosRequestConfig): Promise<T> {
        return instance.delete(url, { params, ...config });
    }
};

export default http