import { RouteRecordRaw } from 'vue-router'
import { defineComponent } from 'vue'

export type Component<T extends any = any> =
  ReturnType<typeof defineComponent>
  | (() => Promise<typeof import('*.vue')>)
  | (() => Promise<T>)

export interface RouteMeta {
  // title
  title: string
  // icon on tab
  icon?: string
  affix?: boolean

  // no show tag
  hideTag?:boolean
}

// @ts-ignore
export interface AppRouteRecordRaw extends Omit<RouteRecordRaw, 'meta'> {
  name?: string
  meta?: RouteMeta
  component?: Component | string
  components?: Component
  children?: AppRouteRecordRaw[]
  props?: Recordable
  hidden?: boolean
}

// export type AppRouteModule = RouteModule | AppRouteRecordRaw;
export type AppRouteModule = AppRouteRecordRaw
