type LoginType = 'uid' | 'password'

export type ClientLoginForm = {
  loginType: LoginType
  uid: number
  username: string
  password: string
  public_pem: string
  game_dir: string
}
