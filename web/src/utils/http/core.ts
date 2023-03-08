import Axios, { AxiosRequestConfig, Method, CancelTokenStatic, AxiosInstance, Canceler, AxiosResponse } from 'axios'

import { getRequestConfig } from './config'
import { trimArgs, downloadHttpResponseErr, httpResponseErr } from '@/utils/http/utils'
import { useUserStoreWithOut } from '@/store/modules/user'

export type RequestMethods = Extract<Method, 'get' | 'post' | 'put' | 'delete' | 'patch' | 'option' | 'head'>
type RequestType = 'request' | 'download' | 'upload'

const userStore = useUserStoreWithOut()

class AxiosHttp {
  private axiosInstance: AxiosInstance
  private readonly requestType: RequestType

  constructor(config?: AxiosRequestConfig, requestType?: RequestType) {
    this.axiosInstance = Axios.create(getRequestConfig(config))
    this.requestType = requestType

    // 拦截器配置
    this.httpInterceptorsRequest()
    this.httpInterceptorsResponse()
  }

  private httpInterceptorsRequest(): void {
    this.axiosInstance.interceptors.request.use(
      (config: AxiosRequestConfig): any => {
        const token = userStore.getToken
        if (token) {
          config.headers['Authorization'] = `Bearer ${token}`
        }
        // trim 参数
        return trimArgs(config)
      },
      error => {
        return Promise.reject(error)
      }
    )
  }

  /*
   * 1. 下载 http {data:xxx，filename: xxx}
   * 2. 其他 data
   * */
  private httpInterceptorsResponse(): void {
    this.axiosInstance.interceptors.response.use(
      (response: AxiosResponse): any => {
        if (this.requestType === 'download') {
          const content = response.headers['content-disposition']
          let filename = content.split('filename=')[1]

          filename = decodeURIComponent(filename)
          return { data: response.data, filename: filename }
        }

        return response.data
      },

      async error => {
        if (this.requestType === 'download') {
          return downloadHttpResponseErr(error)
        }

        return httpResponseErr(error)
      }
    )
  }

  public request<T = any>(data: AxiosRequestConfig): Promise<T> {
    return new Promise((resolve, reject) => {
      this.axiosInstance
        .request(data)
        .then((response: any) => {
          return resolve(response)
        })
        .catch((error: any) => {
          reject(error)
        })
    })
  }
}

export default AxiosHttp
