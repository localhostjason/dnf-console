type LoginType = 'uid' | 'password'

export type ClientLoginForm = {
  login_type: LoginType
  uid: number
  username: string
  password: string
  public_pem: string
  game_dir: string
}
