import { AppRouteRecordRaw } from '@/router/types'
import { LAYOUT } from '@/router/constant'

const accountsRoutes: Array<AppRouteRecordRaw> = [
  {
    path: '/gm',
    name: 'GM',
    component: LAYOUT,
    meta: {
      title: 'GM管理',
      icon: 'cog'
    },
    redirect: 'accounts',
    children: [
      {
        path: 'accounts',
        name: 'GmAccounts',
        component: () => import('@/views/gm/account/index.vue'),
        meta: {
          title: '账号管理',
          icon: 'list'
        }
      },
      {
        path: 'roles',
        name: 'GmRoles',
        component: () => import('@/views/gm/roles/index.vue'),
        meta: {
          title: '角色管理',
          icon: 'fork'
        }
      },
      {
        path: 'tasks',
        name: 'GmTasks',
        component: () => import('@/views/gm/tasks/index.vue'),
        meta: {
          title: '任务清理',
          icon: 'timer'
        }
      },
      {
        path: 'email',
        name: 'GmEmail',
        component: () => import('@/views/gm/email/index.vue'),
        meta: {
          title: '邮件工具',
          icon: 'transfer'
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
