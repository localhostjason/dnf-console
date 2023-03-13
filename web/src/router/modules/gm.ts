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
        name: 'GmAccounts',
        component: () => import('@/views/accounts/list/index.vue'),
        meta: {
          title: '账号管理',
          icon: 'person'
        }
      },
      {
        path: 'roles',
        name: 'GmRoles',
        component: () => import('@/views/accounts/roles/index.vue'),
        meta: {
          title: '角色管理',
          icon: 'person'
        }
      },
      {
        path: 'tasks',
        name: 'GmTasks',
        component: () => import('@/views/accounts/tasks/index.vue'),
        meta: {
          title: '任务清理',
          icon: 'person'
        }
      },
      {
        path: 'email',
        name: 'GmEmail',
        component: () => import('@/views/accounts/email/index.vue'),
        meta: {
          title: '邮件工具',
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
