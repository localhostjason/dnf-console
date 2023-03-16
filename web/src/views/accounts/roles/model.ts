import { AccountDetail } from '@/views/accounts/list/model/model'

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
  money: number
  QP: number
  pvp: RolePvp
}

export type RolePvp = {
  win: number
  pvp_grade: number
  pvp_point: number
}

export type RoleState = {
  loading: boolean
  data: Role[]
}
