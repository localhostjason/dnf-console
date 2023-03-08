<template>
  <div v-if="!item.hidden && item.children && item.children.filter(v => !v.hidden).length" class="menu-wrapper">
    <template v-if="hasOneShowingChild(item.children) && !onlyOneChild.children">
      <app-link v-if="onlyOneChild.meta" :to="resolvePath(onlyOneChild.path)">
        <el-menu-item :index="resolvePath(onlyOneChild.path)" :class="{ 'submenu-title-noDropdown': !isNest }">
          <svg-icon v-if="onlyOneChild.meta && onlyOneChild.meta.icon" :icon-class="onlyOneChild.meta.icon"></svg-icon>
          <template #title>
            <span>{{ onlyOneChild.meta.title }}</span>
          </template>
        </el-menu-item>
      </app-link>
    </template>

    <el-sub-menu v-else ref="subMenu" :index="resolvePath(item.path)" teleported>
      <template #title>
        <svg-icon v-if="item.meta && item.meta.icon" :icon-class="item.meta.icon" />
        <span v-if="item.meta && item.meta.title && sidebar.opened" slot="title">{{ item.meta.title }}</span>
      </template>

      <template v-for="child in item.children.filter(v => !v.hidden)">
        <sidebar-item
          v-if="child.children && child.children.length > 0"
          :is-nest="true"
          :item="child"
          :key="child.path"
          :base-path="resolvePath(child.path)"
          class="nest-menu"
        ></sidebar-item>

        <router-link v-else :to="resolvePath(child.path)" :key="child.name">
          <el-menu-item :index="resolvePath(child.path)">
            <svg-icon v-if="child.meta && child.meta.icon" :icon-class="child.meta.icon"></svg-icon>
            <span v-if="child.meta && child.meta.title" slot="title">{{ child.meta.title }}</span>
          </el-menu-item>
        </router-link>
      </template>
    </el-sub-menu>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import AppLink from './Link.vue'
import { AppRouteRecordRaw } from '@/router/types'
import { useAppStore } from '@/store/modules/app'
import { storeToRefs } from 'pinia'

export default defineComponent({
  name: 'SidebarItem',
  components: { AppLink },
  props: {
    item: {
      type: Object,
      required: true
    },
    isNest: {
      type: Boolean,
      default: false
    },
    basePath: {
      type: String,
      default: ''
    }
  },
  setup(props) {
    const appStore = useAppStore()
    const { sidebar } = storeToRefs(appStore)

    const onlyOneChild = ref<AppRouteRecordRaw>({} as any)

    const resolvePath = (routePath: string) => {
      const path = require('path')
      return path.resolve(props.basePath, routePath)
    }

    function hasOneShowingChild(children: AppRouteRecordRaw[] = []) {
      const showingChildren = children.filter(item => {
        if (item.hidden) {
          return false
        } else {
          onlyOneChild.value = item
          return true
        }
      })

      return showingChildren.length === 1
    }

    return {
      sidebar,
      resolvePath,
      onlyOneChild,
      hasOneShowingChild
    }
  }
})
</script>

<style rel="" lang="scss" scoped>
.menu-wrapper {
  .font_icon {
    width: 1em;
    height: 1em;
    vertical-align: -0.15em;
    fill: currentColor;
    /*overflow: hidden;*/
    margin-right: 16px;
    position: relative;
    bottom: 1px;
  }
}
</style>
