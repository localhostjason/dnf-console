import { http } from '@/utils/http'
import { Account, AccountDetail, RechargeForm } from '@/views/accounts/list/model'

export const getAccounts = <T = Account | AccountDetail[]>(params: any): Promise<T> => {
  return http.request({
    url: `/gm/account/list`,
    method: 'get',
    params
  })
}

export const rechargeByUid = (uid: number, data: RechargeForm): Promise<any> => {
  return http.request({
    url: `/gm/account/${uid}/recharge`,
    method: 'post',
    data
  })
}
