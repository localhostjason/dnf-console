export type ChangePasswordForm = {
  password: string
  check_password: string
}

export type EditAccountForm = {
  qq: string
}

export interface CreateAccountForm extends ChangePasswordForm, EditAccountForm {
  account_name: string
}
