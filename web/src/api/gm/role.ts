import { http } from '@/utils/http'

export const getRoles = (id: number): Promise<any[]> => {
  return http.request({
    url: `/gm/account/${id}/roles`,
    method: 'get'
  })
}
