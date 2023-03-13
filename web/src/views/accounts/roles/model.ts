import { AccountDetail } from '@/views/accounts/list/model'

export type AccountOptions = {
  data: AccountDetail[]
}

export type SelectForm = {
  uid: number
}

export type Role = {
  mid: number
  charac_no: number
  charac_name: string
  lev: string
  create_time: string
}

export type RoleState = {
  loading: boolean
  data: Role[]
}
