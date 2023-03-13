import { http } from '@/utils/http'

export const sendEmailByRole = (characNo: number, data: any): Promise<any> => {
  return http.request({
    url: `/gm/roles/${characNo}/email`,
    method: 'post',
    data
  })
}
