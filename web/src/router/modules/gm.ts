import { AppRouteRecordRaw } from '@/router/types'
import { LAYOUT } from '@/router/constant'

const accountsRoutes: Array<AppRouteRecordRaw> = [
  {
    path: '/gm',
    name: 'GM',
    component: LAYOUT,
    meta: {
      title: 'GM管理',
      icon: 'people'
    },
    redirect: 'accounts',
    children: [
      {
        path: 'accounts',
        name: 'Accounts',
        component: () => import('@/views/accounts/list/index.vue'),
        meta: {
          title: '账号管理',
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
      }
    ]
  }
]

export default accountsRoutes
