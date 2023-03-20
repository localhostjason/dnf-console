import { http } from '@/utils/http'

export const getDashTotal = (): Promise<any> => {
  return http.request({
    url: `/dash/count/stat`,
    method: 'get'
  })
}