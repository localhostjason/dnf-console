import { AxiosRequestConfig } from 'axios'
import { HttpTimeOut, HttpUrl } from '@/utils/http/const'

const defaultConfig: AxiosRequestConfig = {
  baseURL: HttpUrl, // api的base_url
  timeout: HttpTimeOut, // request timeout
  headers: {
    Accept: 'application/json, text/plain, */*; charset=utf-8',
    'Content-Type': 'application/json; charset=utf-8',
    Pragma: 'no-cache',
    'Cache-Control': 'no-cache'
  }
}

const getRequestConfig = (config?: AxiosRequestConfig): AxiosRequestConfig => {
  if (!config) {
    return defaultConfig
  }
  const { headers } = config
  if (headers && typeof headers === 'object') {
    defaultConfig.headers = {
      ...defaultConfig.headers,
      ...headers
    }
  }

  return { ...excludeProps(config!, 'headers'), ...defaultConfig }
}

export { getRequestConfig }

function excludeProps<T extends { [key: string]: any }>(origin: T, prop: string): { [key: string]: T } {
  return Object.keys(origin)
    .filter(key => !prop.includes(key))
    .reduce((res, key) => {
      res[key] = origin[key]
      return res
    }, {} as { [key: string]: T })
}
