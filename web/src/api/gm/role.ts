import { http } from '@/utils/http'
import { QpForm } from '@/views/accounts/roles/model/qp'

export const getRoles = (id: number): Promise<any[]> => {
  return http.request({
    url: `/gm/account/${id}/roles`,
    method: 'get'
  })
}

export const updateQPbyRole = (characNo: number, data: QpForm): Promise<any> => {
  return http.request({
    url: `/gm/roles/${characNo}/qp`,
    method: 'put',
    data
  })
}
