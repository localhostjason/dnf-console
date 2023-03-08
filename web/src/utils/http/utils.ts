import isString from 'lodash/isString'
import isPlainObject from 'lodash/isPlainObject'
import { errorMessage } from '@/utils/element/message'
import router from '@/router'
import { useUserStoreWithOut } from '@/store/modules/user'
import { AxiosRequestConfig } from 'axios'

const trim = (data: any): any => {
  let newData = {}
  for (const [k, v] of Object.entries(data)) {
    newData[k] = isString(v) ? v.trim() : v
  }
  return newData
}

const trimArgs = (config: AxiosRequestConfig): AxiosRequestConfig => {
  // trim 参数
  let { method, data, params } = config
  try {
    if (method !== 'get') {
      if (isPlainObject(data)) {
        data = trim(data)
      }
      config['data'] = data
    } else {
      if (isPlainObject(params)) {
        params = trim(params)
      }
      config['params'] = params
    }
  } catch (e) {
    console.log('trim error:', e)
    return config
  }
  return config
}

export { trimArgs }

const httpResponseMessageByCode = (code: number, errMsg: string): string => {
  // 根据自己业务 拦截error
  let message = ''
  switch (code) {
    case 403:
      message = '权限不足'
      break

    case 404:
      message = '相关的资源不存在'
      break

    case 401:
      message = errMsg
      break

    case 422:
      message = errMsg
      break

    case 500:
      message = '服务器错误'
      break
    default:
      message = '未知错误'
  }
  return message
}

// 拦截器 response err
export const httpResponseErr = async <T = never>(error: any): Promise<T> => {
  let status = 0
  const userStore = useUserStoreWithOut()
  try {
    status = error.response.status
  } catch (e) {
    const msg = '连接不上后台，已超时'
    errorMessage(msg)
    return Promise.reject(msg)
  }

  const { msg, code } = error.response.data
  const message = httpResponseMessageByCode(status, msg)
  errorMessage(message)

  if (status === 401) {
    await userStore.removeUserStore()
    await router.push({ path: '/login' })
  }
  return Promise.reject(error)
}

// 拦截器 download response err
export const downloadHttpResponseErr = <T = never>(error: any): Promise<T> => {
  let status = 0
  try {
    status = error.response.status
  } catch (e) {
    const msg = '连接不上后台，已超时'
    errorMessage(msg)
    return Promise.reject(msg)
  }

  let message
  switch (status) {
    case 403:
      message = '权限不足'
      break
    case 404:
      message = '相关的资源不存在'
      break
    default:
      message = ''
  }
  if (message) {
    errorMessage(message)
    return Promise.reject({ code: status, msg: '' })
  }

  const fileReader = new FileReader()
  fileReader.onload = function () {
    if (typeof fileReader.result === 'string') {
      const resData = JSON.parse(fileReader.result)
      errorMessage(resData.msg)
    }
  }

  const data = error.response.data
  fileReader.readAsText(data) // FileReader的API
  return Promise.reject({ code: status, msg: '' })
}
