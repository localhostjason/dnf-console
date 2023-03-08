import { PropType as VuePropType } from 'vue'

declare global {
  declare type Nullable<T> = T | null

  declare interface Fn<T = any, R = T> {
    (...arg: T[]): R
  }

  declare type TimeoutHandle = ReturnType<typeof setTimeout>

  // vue
  declare type PropType<T> = VuePropType<T>

  declare type Recordable<T = any> = Record<string, T>
}
