import { http } from '@/utils/http'
import { Account, AccountDetail, RechargeForm } from '@/views/gm/account/model/model'
import { ChangePasswordForm, CreateAccountForm, EditAccountForm } from '@/views/gm/account/model/info'

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

export const resetCreateCharac = (uid: number): Promise<any> => {
  return http.request({
    url: `/gm/account/${uid}/reset_create_charac`,
    method: 'post'
  })
}

export const deleteAccount = (uid: number): Promise<any> => {
  return http.request({
    url: `/gm/account/${uid}`,
    method: 'delete'
  })
}

export const updateAccountPasswordByUid = (uid: number, data: ChangePasswordForm): Promise<any> => {
  return http.request({
    url: `/gm/account/${uid}/change_password`,
    method: 'post',
    data
  })
}

export const updateAccountInfo = (uid: number, data: EditAccountForm): Promise<any> => {
  return http.request({
    url: `/gm/account/${uid}`,
    method: 'put',
    data
  })
}

export const createAccountInfo = (data: CreateAccountForm): Promise<any> => {
  return http.request({
    url: `/gm/account`,
    method: 'post',
    data
  })
}
