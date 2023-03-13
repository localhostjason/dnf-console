import { http } from '@/utils/http'
import { Task } from '@/views/accounts/tasks/model'

export const getTasks = (characNo: number): Promise<Task[]> => {
  return http.request({
    url: `/gm/roles/${characNo}/tasks`,
    method: 'get'
  })
}

export const updateTasks = (characNo: number, data: any): Promise<any> => {
  return http.request({
    url: `/gm/roles/${characNo}/tasks`,
    method: 'put',
    data
  })
}
