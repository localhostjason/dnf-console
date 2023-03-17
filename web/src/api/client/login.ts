import { http } from '@/utils/http'
import { ClientLoginForm } from '@/views/client/login/models/login'

export const loginClient = (data: ClientLoginForm): Promise<any> => {
  return http.request({
    url: `/client/login`,
    method: 'post',
    data
  })
}
