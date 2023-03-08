import { LAYOUT } from '@/router/constant'
import { AppRouteRecordRaw } from '@/router/types'

const userRoutes: Array<AppRouteRecordRaw> = [
  {
    path: '/user',
    name: 'User',
    component: LAYOUT,
    meta: {
      title: '用户管理',
      icon: 'people'
    },
    redirect: 'list',
    children: [
      {
        path: 'list',
        name: 'UserList',
        component: () => import('@/views/user/list/index.vue'),
        meta: {
          title: '用户列表',
          icon: 'person'
        }
      },
      {
        path: 'test',
        name: 'UserTest',
        component: () => import('@/views/user/test/index.vue'),
        meta: {
          title: '测试',
          icon: 'list'
        }
      },
      {
        path: 'info',
        name: 'UserInfo',
        component: () => import('@/views/user/info/index.vue'),
        hidden: true,
        meta: {
          title: '个人信息',
          // hideTag: true
        }
      }
    ]
  }
]

export default userRoutes
