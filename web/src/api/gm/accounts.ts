import { http } from '@/utils/http'
import { Account } from '@/views/accounts/list/model'

export const getAccounts = (params: any): Promise<Account> => {
  return http.request({
    url: `/gm/account/list`,
    method: 'get',
    params
  })
}
