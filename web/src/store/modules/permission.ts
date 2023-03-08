import { asyncRoutes, basicRoutes } from '@/router'
import { defineStore } from 'pinia'
import store from '@/store'

/**
 * 匹配 router：name，生成新的router
 * @param routes_name
 * @param route
 */
function hasPermission(routes_name, route) {
  if (route.name) {
    // console.log(routes_map, route.path, routes_map.some(val => val.includes(route.path)));
    // return routes_map.includes(route.path)
    return Boolean(!routes_name.includes(route.name) || route.hidden)
  } else {
    return true
  }
}

/**
 * 从all 的路由数据 ，刷选出满足用户 hasPermission 的路由
 * @param routes
 * @param routes_map
 */
export function filterAsyncRoutes(routes, routes_map) {
  const res = []

  routes.forEach(route => {
    const tmp = { ...route }
    if (hasPermission(routes_map, tmp)) {
      if (tmp.children) {
        tmp.children = filterAsyncRoutes(tmp.children, routes_map)
      }
      res.push(tmp)
    }
  })

  return res
}

/***
 * constantRouterMap : 公共路由，比如login ，404，500
 * asyncRouterMap ：   动态需添加路由，自定义加权限
 */
const state = {
  routers: [],
  addRouters: []
}

interface PermissionState {
  routers: any[]
  addRouters: any[]
}

export const usePermissionStore = defineStore({
  id: 'app-permission',
  state: (): PermissionState => ({
    routers: [],
    addRouters: []
  }),
  getters: {
    getPermissionAddRouters(): any[] {
      return this.addRouters
    }
  },
  actions: {
    generateRoutes(data: string[]) {
      let accessedRoutes = filterAsyncRoutes(asyncRoutes, data)
      this.addRouters = accessedRoutes
      this.routers = basicRoutes.concat(accessedRoutes)
      return accessedRoutes
    }
  }
})

// Need to be used outside the setup
export function usePermissionStoreWithOut() {
  return usePermissionStore(store)
}
