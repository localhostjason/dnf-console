import { AppRouteRecordRaw } from '@/router/types'
import { LAYOUT } from '@/router/constant'

const systemRoutes: Array<AppRouteRecordRaw> = [
  {
    path: '/system',
    name: 'System',
    component: LAYOUT,
    meta: {
      title: '系统管理',
      icon: 'people'
    },
    redirect: 'user_list',
    children: [
      {
        path: 'user_list',
        name: 'UserList',
        component: () => import('@/views/user/list/index.vue'),
        meta: {
          title: '用户管理',
          icon: 'person'
        }
      },
      {
        path: 'user_info',
        name: 'UserInfo',
        component: () => import('@/views/user/info/index.vue'),
        // hidden: true,
        meta: {
          title: '个人信息',
          icon: 'person'
          // hideTag: true
        }
      }
    ]
  }
]

export default systemRoutes
