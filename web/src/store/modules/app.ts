import { storageLocal } from '@/utils/storage'
import { defineStore } from 'pinia'

interface AppState {
  sidebar: {
    opened: boolean
    withoutAnimation: boolean
  }
  device: string
}

function getSideBarOpened(): boolean {
  const data = storageLocal.getItem('sidebarStatus')
  if (data === null) {
    return true
  }
  return Boolean(Number(data))
}

export const useAppStore = defineStore({
  id: 'app',
  state: (): AppState => ({
    sidebar: {
      opened: getSideBarOpened(),
      withoutAnimation: false
    },
    device: 'desktop'
  }),
  getters: {
    getSidebarOpened() {
      return this.sidebar.opened
    }
  },
  actions: {
    toggleSideBar(): void {
      this.sidebar.opened = !this.sidebar.opened
      this.sidebar.withoutAnimation = false
      if (this.sidebar.opened) {
        storageLocal.setItem('sidebarStatus', 1)
      } else {
        storageLocal.setItem('sidebarStatus', 0)
      }
    },
    closeSideBar(withoutAnimation: boolean): void {
      storageLocal.setItem('sidebarStatus', 0)
      this.sidebar.opened = false
      this.sidebar.withoutAnimation = withoutAnimation
    },
    toggleDevice(device: string): void {
      this.device = device
    }
  }
})
