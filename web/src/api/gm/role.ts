import { http } from '@/utils/http'
import { QpForm } from '@/views/gm/roles/model/qp'
import { PvpForm } from '@/views/gm/roles/model/pvp'

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

export const updatePvpByRole = (characNo: number, data: PvpForm): Promise<any> => {
  return http.request({
    url: `/gm/roles/${characNo}/pvp`,
    method: 'put',
    data
  })
}
