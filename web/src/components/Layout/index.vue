<template>
  <div :class="classes" class="app-wrapper">
    <div v-if="device === 'mobile' && sidebar.opened" class="drawer-bg" @click="handleClickOutside"></div>

    <!-- 侧边栏 -->
    <sidebar class="sidebar-container"></sidebar>
    <div class="main-container">
      <navbar />

      <tags-view></tags-view>
      <!-- 主体内容 -->
      <app-main />
    </div>
  </div>
</template>

<script lang="ts">
import { Navbar, AppMain, Sidebar, TagsView } from './components'
import { ref, reactive, computed, toRefs, watchEffect, onMounted, onBeforeMount } from 'vue'
import { useEventListener } from '@vueuse/core'
import settings from '@/settings'
import { useAppStore } from '@/store/modules/app'

interface setInter {
  sidebar: any
  device: String
  classes: any
}

export default {
  name: 'layout',
  components: {
    TagsView,
    Navbar,
    AppMain,
    Sidebar
  },
  setup() {
    const appStore = useAppStore()

    const WIDTH = ref(992)

    const set: setInter = reactive({
      sidebar: computed(() => {
        return appStore.sidebar
      }),

      device: computed(() => {
        return appStore.device
      }),

      classes: computed(() => {
        return {
          hideSidebar: !set.sidebar.opened,
          openSidebar: set.sidebar.opened,
          withoutAnimation: set.sidebar.withoutAnimation,
          mobile: set.device === 'mobile'
        }
      })
    })

    watchEffect(() => {
      if (set.device === 'mobile' && !set.sidebar.opened) {
        appStore.closeSideBar(false)
      }
    })

    const handleClickOutside = () => {
      appStore.closeSideBar(false)
    }

    const $_isMobile = () => {
      const rect = document.body.getBoundingClientRect()
      return rect.width - 1 < WIDTH.value
    }

    const $_resizeHandler = () => {
      if (!document.hidden) {
        const isMobile = $_isMobile()
        appStore.toggleDevice(isMobile ? 'mobile' : 'desktop')

        if (isMobile) {
          appStore.closeSideBar(true)
        }
      }
    }

    onMounted(() => {
      const isMobile = $_isMobile()
      if (isMobile) {
        appStore.toggleDevice('mobile')
        appStore.closeSideBar(true)
      }
    })

    onBeforeMount(() => {
      useEventListener('resize', $_resizeHandler)
    })

    return {
      ...toRefs(set),
      handleClickOutside,
      settings
    }
  }
}
</script>

<style rel="stylesheet/scss" lang="scss" scoped>
@import '../../styles/mixin.scss';

.app-wrapper {
  @include clearfix;
  position: relative;
  height: 100%;
  width: 100%;

  &.mobile.openSidebar {
    position: fixed;
    top: 0;
  }
}

.drawer-bg {
  background: #000;
  opacity: 0.3;
  width: 100%;
  top: 0;
  height: 100%;
  position: absolute;
  z-index: 999;
}
</style>
