export type ChangePassword = {
  password: string
  check_password: string
}

export type EditAccountForm = {
  qq: string
}

export interface CreateAccountForm extends ChangePassword, EditAccountForm {
  account_name: string
}
