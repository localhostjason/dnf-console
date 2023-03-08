import { storageLocal } from '@/utils/storage'
import { defineStore } from 'pinia'
import store from '@/store'

interface UserState {
  token?: string
  username?: string
  role?: string
}

export const useUserStore = defineStore({
  id: 'app-user',
  state: (): UserState => ({
    token: undefined,
    username: null,
    role: null
  }),
  getters: {
    getToken(): string {
      return this.token || storageLocal.getItem('token')
    },
    getUsername(): string {
      return this.username
    },
    getUserRole(): string {
      return this.role
    }
  },
  actions: {
    setToken(token: string) {
      this.token = token
      storageLocal.setItem('token', token)
    },
    setUserInfo(username: string, role: string) {
      this.username = username
      this.role = role
    },
    removeUserStore() {
      this.token = ''
      this.username = ''
      this.role = ''
      storageLocal.clear()
    }
  }
})

// Need to be used outside the setup
export function useUserStoreWithOut() {
  return useUserStore(store)
}
