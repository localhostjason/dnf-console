import { http } from '@/utils/http'

export const sendEmailByRole = (characNo: number, data: any): Promise<any> => {
  return http.request({
    url: `/gm/roles/${characNo}/email`,
    method: 'post',
    data
  })
}

export const getGolds = (params: object): Promise<any> => {
  return http.request({
    url: `/gm/gold/list`,
    method: 'get',
    params
  })
}
