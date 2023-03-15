export type AccountDetail = {
  uid: number
  account_name: string
  qqq: string
  roles: number
  money: any
  capacity: number

  cera_point: number // d点
  cera: number // d币
}

export type Account = {
  items: AccountDetail[]
  total: number
  page: number
  per_page: number
  pages: number
}

export type AccountState = {
  data: AccountDetail[]
  loading: boolean
  total: number
}
export type FilterAccountForm = {
  uid: string
}

export type RechargeForm = {
  cera_point: number // d点
  cera: number // d币

  cera_option: string
}
