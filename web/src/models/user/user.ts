export type User = {
  id: number
  username: string
  time: string
  last_login_time: string | null
  email: string
  role: string
  desc: string
}

export type UserState = {
  data: User[]
  loading: boolean
}

export type ModifyUserForm = {
  username: string
  email: string
  desc: string
  password?: string
  checkPassword?: string
}
