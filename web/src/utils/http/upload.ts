import { AxiosRequestConfig } from 'axios'
import { HttpTimeOut, HttpUrl } from '@/utils/http/const'
import AxiosHttp from '@/utils/http/core'

const uploadConfig: AxiosRequestConfig = {
  baseURL: HttpUrl, // api的base_url
  timeout: HttpTimeOut, // request timeout
  headers: {
    'Content-Type': 'multipart/form-data'
  }
}

export const uploadHttp = new AxiosHttp(uploadConfig)
