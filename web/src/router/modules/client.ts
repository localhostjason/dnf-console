import { AppRouteRecordRaw } from '@/router/types'
import { LAYOUT } from '@/router/constant'

const ClientRoutes: Array<AppRouteRecordRaw> = [
  {
    path: '/client',
    name: 'Client',
    component: LAYOUT,
    meta: {
      title: '客户端',
      icon: 'bell'
    },
    redirect: 'accounts',
    children: [
      {
        path: 'login',
        name: 'ClientLogin',
        component: () => import('@/views/client/login/index.vue'),
        meta: {
          title: '客户端登录',
          icon: 'key'
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
      }
    ]
  }
]

export default ClientRoutes
