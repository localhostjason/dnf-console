import { http } from '@/utils/http'
import { downloadHttp } from '@/utils/http/download'
import { User } from '@/models/user/user'

export const getUsers = (): Promise<User[]> => {
  return http.request({
    url: `/user/list`,
    method: 'get'
  })
}

export const downloadFile = <T = any>(): Promise<T> => {
  return downloadHttp.request({
    url: `/user/file`,
    method: 'get'
  })
}

export const getUsersList = async (): Promise<User[]> => {
  return [
    {
      id: 1,
      username: 'admin',
      last_login_time: '2023-02-10T10:07:56.8351023+08:00',
      time: '2023-02-09T14:41:19.3067481+08:00',
      role: 'admin',
      email: '',
      desc: '系统管理员',
      is_super_admin: true
    }
  ]
}
