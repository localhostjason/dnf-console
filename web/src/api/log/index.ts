import { http } from '@/utils/http'

export const getRechargeLog = (params): Promise<any> => {
  return http.request({
    url: `/log/recharge`,
    method: 'get',
    params
  })
}

export const getOperateLog = (params): Promise<any> => {
  return http.request({
    url: `/log/operate`,
    method: 'get',
    params
  })
}
