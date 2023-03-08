import AxiosHttp from '@/utils/http/core'
import { HttpTimeOut, HttpUrl } from '@/utils/http/const'
import { AxiosRequestConfig } from 'axios'

const downloadConfig: AxiosRequestConfig = {
  baseURL: HttpUrl, // api的base_url
  timeout: HttpTimeOut, // request timeout
  responseType: 'blob'
}

// 下载 返回字节 blob
export const downloadHttp = new AxiosHttp(downloadConfig, 'download')
