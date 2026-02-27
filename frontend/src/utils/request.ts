import axios, { type AxiosResponse } from 'axios'

const instance = axios.create({
    baseURL: '/api',
})

// 请求拦截器
instance.interceptors.request.use(config => { return config }, error => Promise.reject(error))

// 响应拦截器
instance.interceptors.response.use(
    (response: AxiosResponse) => {
        const res = response.data
        if (res.code === 500) {
            alert(res.message)
            return Promise.reject(res.message);
        }
        return res
    },
    error => { return Promise.reject(error) }
)
export default instance
