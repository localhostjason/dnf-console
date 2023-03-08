import NProgress from '@/utils/progress'
import router, { resetRouter } from './router'
import { getUserInfo } from '@/api/user/auth'
import { useUserStoreWithOut } from '@/store/modules/user'
import { usePermissionStoreWithOut } from '@/store/modules/permission'
import type { Router } from 'vue-router'
import { getPageTitle } from '@/utils/get-page-title'

const whiteList = ['/login']
const permissionStore = usePermissionStoreWithOut()
const userStore = useUserStoreWithOut()

export const setupPermissionRouter = (router: Router) => {
  router.beforeEach(async (to, _from, next) => {
    NProgress.start()

    // set page title
    document.title = getPageTitle(to.meta.title)

    const token = userStore.getToken
    const username = userStore.getUsername
    if (!token) {
      whiteList.includes(to.path) ? next() : next(`/login`)
      NProgress.done()
      return
    }

    if (to.path === '/login') {
      next({ path: '/' })
      NProgress.done()
      return
    }

    if (username) {
      next()
      return
    }

    try {
      const { username, role } = await getUserInfo()
      userStore.setUserInfo(username, role)
      // generate accessible routes map based on roles
      const menus = [] // menus = ["Apis"] // 此menus 可通過接口獲得
      // generate accessible routes map based on roles
      const accessRoutes = permissionStore.generateRoutes(menus)
      // dynamically add accessible routes
      resetRouter()
      accessRoutes.forEach(val => {
        router.addRoute(val)
      })

      next({ ...to, replace: true })
    } catch (error) {
      console.log('err:', error)
      // remove token and go to login page to re-login
      userStore.removeUserStore()
      next(`/login`)
      NProgress.done()
    }
  })

  router.afterEach(() => {
    NProgress.done()
  })
}
