import { createApp } from 'vue'
import App from './App.vue'

import router from './router'
import store from './store'

import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

import * as ElIcon from '@element-plus/icons-vue'

import SvgIcon from './components/SvgIcon/index.vue'
import './icons'
import { setupPermissionRouter } from './permission'

// 导入公共样式
import 'normalize.css/normalize.css'
import './styles/index.scss'

// init

;(async () => {
  const app = createApp(App)

  // load vue router
  app.use(router)

  //  load pinia store
  app.use(store)

  // load element plus ui
  app.use(ElementPlus, {
    size: 'default'
  })
  for (const iconName in ElIcon) {
    app.component(iconName, ElIcon[iconName])
  }

  // load 路由守卫
  setupPermissionRouter(router)

  // load svg组件
  app.component('svg-icon', SvgIcon)


  app.mount('#app', true)
})()
