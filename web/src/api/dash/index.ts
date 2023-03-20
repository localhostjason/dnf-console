import { http } from '@/utils/http'

export const getDashTotal = (): Promise<any> => {
  return http.request({
    url: `/dash/stat/count`,
    method: 'get'
  })
}

export const getDashTop5 = (): Promise<any> => {
  return http.request({
    url: `/dash/stat/top5`,
    method: 'get'
  })
}
