import { http } from '@/utils/http'
import { Account, AccountDetail } from '@/views/accounts/list/model'

export const getAccounts = <T = Account | AccountDetail[]>(params: any): Promise<T> => {
  return http.request({
    url: `/gm/account/list`,
    method: 'get',
    params
  })
}
