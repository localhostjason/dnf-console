export type AccountDetail = {
  uid: number
  account_name: string
  qqq: string
  roles: number
  money: any
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

export type ModifyAccountForm = {}

export type FilterAccountForm = {
  uid: string
}
