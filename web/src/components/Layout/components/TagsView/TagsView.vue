<template>
  <div ref="tagsViewRef" id="tags-view-container" class="tags-view-container" v-if="showTag">
    <scroll-pane ref="scrollPane" class="tags-view-wrapper">
      <!-- 暂时无右击行为  -->
      <router-link
        v-for="tag in visitedViews"
        ref="tagRef"
        :key="tag.path"
        :class="isActive(tag) ? 'active' : ''"
        :to="{ path: tag.path, query: tag.query, fullPath: tag.fullPath }"
        class="tags-view-item"
        @click.middle.native="closeSelectedTag(tag)"
        @contextmenu.prevent.native="openMenu(tag, $event)"
      >
        {{ tag.title }}

        <Close
          v-if="!tag.meta.affix"
          class="el-icon-close"
          style="width: 10px; height: 10px; margin-right: 2px; position: relative; top: 3px; left: 5px"
          @click.prevent.stop="closeSelectedTag(tag)"
        />
      </router-link>
    </scroll-pane>
    <ul ref="contextmenuRef" v-show="visible" :style="{ left: left + 'px', top: top + 'px' }" class="contextmenu">
      <li @click="refreshSelectedTag(selectedTag)">刷新</li>
      <li v-if="!(selectedTag.meta && selectedTag.meta.affix)" @click="closeSelectedTag(selectedTag)">关闭</li>
      <li @click="closeOthersTags">关闭其它</li>
      <li @click="closeAllTags(selectedTag)">关闭所有</li>
    </ul>
  </div>
</template>

<script setup lang="ts">
import ScrollPane from './ScrollPane.vue'
import { computed, ref, onMounted, watch, nextTick } from 'vue'
import { useTagsStore } from '@/store/modules/tagsView'
import { usePermissionStore } from '@/store/modules/permission'
import { storeToRefs } from 'pinia'
import { useRoute, useRouter } from 'vue-router'
import path from 'path'
import { toRaw } from '@vue/reactivity'

const route = useRoute()
const router = useRouter()

const showTag = ref<boolean>(true)

const tagsStore = useTagsStore()
const permissionStore = usePermissionStore()

const { visitedViews } = storeToRefs(tagsStore)

const { routers } = storeToRefs(permissionStore)
const routes = routers.value

const visible = ref<boolean>(false)
const top = ref<number>(0)
const left = ref<number>(0)
const selectedTag = ref<object>({})
let affixTags = ref<any[]>([])

const contextmenuRef = ref<HTMLElement | null>(null)
const tagRef = ref(null)
const tagsViewRef = ref(null)
const scrollPane = ref(null)

onMounted(() => {
  initTags()
  addTags()

  const { hideTag } = route.meta
  showTag.value = !hideTag;
})

watch(
  () => route.path,
  () => {
    addTags()
    moveToCurrentTag()
  }
)

watch(
  () => route.meta,
  value => {
    showTag.value = !value.hideTag;
  }
)

watch(
  () => visible.value,
  value => {
    if (value) {
      document.body.addEventListener('click', closeMenu)
    } else {
      document.body.removeEventListener('click', closeMenu)
    }
  }
)

const isActive = r => {
  return r.path === route.path
}
const filterAffixTags = (routes, basePath = '/'): any[] => {
  let tags = []
  routes.forEach(route => {
    if (route.meta && route.meta.affix) {
      const tagPath = path.resolve(basePath, route.path)
      tags.push({
        fullPath: tagPath,
        path: tagPath,
        name: route.name,
        meta: { ...route.meta }
      })
    }
    if (route.children) {
      const tempTags = filterAffixTags(route.children, route.path)
      if (tempTags.length >= 1) {
        tags = [...tags, ...tempTags]
      }
    }
  })
  return tags
}
const initTags = () => {
  affixTags.value = filterAffixTags(routes)

  for (const tag of affixTags.value) {
    // Must have tag name
    if (tag.name) {
      tagsStore.addVisitedView(tag)
    }
  }
}
const addTags = () => {
  const { name, meta } = route
  if (name && !meta.hideTag) {
    tagsStore.addView(route)
  }
  return false
}
const moveToCurrentTag = () => {
  nextTick(() => {
    for (const tag of tagRef.value) {
      if (tag.to.path === route.path) {
        scrollPane.value.moveToTarget(tag, tagRef.value)
        // when query is different , then update
        if (tag.to.fullPath !== route.fullPath) {
          tagsStore.updateVisitedView(route)
        }
        break
      }
    }
  })
}
const refreshSelectedTag = async view => {
  const { fullPath } = view
  await nextTick(() => {
    router.replace({
      path: '/redirect' + fullPath
    })
  })
}

const closeSelectedTag = view => {
  const data = tagsStore.delView(view)

  if (isActive(view)) {
    toLastView(data['visitedViews'], view)
  }
}

// 右击：关闭其他
const closeOthersTags = () => {
  router.push(selectedTag.value)
  tagsStore.delOthersViews(selectedTag.value)

  moveToCurrentTag()
}

// 右击：关闭所有
const closeAllTags = async view => {
  const data = await tagsStore.delAllViews()
  const views = data['visitedViews']
  if (affixTags.value.some(tag => tag.path === view.path)) {
    return
  }
  toLastView(views, view)
}

const toLastView = (visitedViews, view) => {
  const latestView = visitedViews.slice(-1)[0]
  if (latestView) {
    router.push(latestView)
  } else {
    // now the default is to redirect to the home page if there is no tags-view,
    // you can adjust it according to your needs.
    if (view.name === 'Dashboard') {
      // to reload home page
      router.replace({ path: '/redirect' + view.fullPath })
    } else {
      router.push('/')
    }
  }
}

const openMenu = (tag, e) => {
  const menuMinWidth = 105

  const offsetLeft = tagsViewRef.value.getBoundingClientRect().left
  const offsetWidth = tagsViewRef.value.offsetWidth
  const maxLeft = offsetWidth - menuMinWidth
  const left_ = e.clientX - offsetLeft + 15

  if (left_ > maxLeft) {
    left.value = maxLeft
  } else {
    left.value = left_
  }

  top.value = e.clientY - 37
  visible.value = true
  selectedTag.value = tag
}
const closeMenu = () => {
  visible.value = false
}
</script>

<style lang="scss" scoped>
.tags-view-container {
  position: relative;
  //top: 50px;

  .tags-view-wrapper {
    background: #fff;
    height: 40px;
    line-height: 38px;

    border-bottom: 1px solid #d8dce5;

    .tags-view-item {
      display: inline-block;
      position: relative;
      cursor: pointer;
      height: 26px;
      line-height: 23px;
      border: 1px solid #d8dce5;
      color: #495060;
      background: #fff;
      padding: 2px 8px 0;
      font-size: 12px;
      margin-left: 5px;
      margin-top: 4px;

      &:first-of-type {
        margin-left: 15px;
      }

      &:last-of-type {
        margin-right: 15px;
      }

      &.active {
        background-color: #42b983;
        color: #fff;
        border-color: #42b983;

        &::before {
          content: '';
          background: #fff;
          display: inline-block;
          width: 8px;
          height: 8px;
          border-radius: 50%;
          position: relative;
          margin-right: 2px;
        }
      }
    }
  }

  .contextmenu {
    margin: 0;
    background: #fff;
    z-index: 3000;
    position: absolute;
    list-style-type: none;
    padding: 5px 0;
    border-radius: 4px;
    font-size: 12px;
    font-weight: 400;
    color: #333;
    box-shadow: 2px 2px 3px 0 rgba(0, 0, 0, 0.3);

    li {
      margin: 0;
      padding: 7px 16px;
      cursor: pointer;

      &:hover {
        background: #eee;
      }
    }
  }
}
</style>

<style lang="scss">
//reset element css of el-icon-close
.tags-view-wrapper {
  .tags-view-item {
    .el-icon-close {
      width: 16px;
      height: 16px;
      vertical-align: 2px;
      border-radius: 50%;
      text-align: center;
      transition: all 0.3s cubic-bezier(0.645, 0.045, 0.355, 1);
      transform-origin: 100% 50%;

      &:before {
        transform: scale(0.6);
        display: inline-block;
        vertical-align: -3px;
      }

      &:hover {
        background-color: #b4bccc;
        color: #fff;
      }
    }
  }
}
</style>

<style lang="scss" scoped>
.scroll-container {
  white-space: nowrap;
  position: relative;
  overflow: hidden;
  width: 100%;

  ::v-deep(.el-scrollbar__bar) {
    bottom: 0;
  }

  ::v-deep(.el-scrollbar__wrap) {
    height: 49px;
  }
}
</style>
