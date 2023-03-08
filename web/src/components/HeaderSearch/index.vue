<template>
  <div :class="{ show: show }" class="header-search">
    <svg-icon class-name="search-icon" icon-class="search" @click.stop="click" />
    <el-select
      ref="headerSearchSelectRef"
      v-model="search"
      :remote-method="querySearch"
      filterable
      default-first-option
      remote
      placeholder="请输入菜单名"
      class="header-search-select"
      @change="change"
    >
      <el-option
        v-for="info in options"
        :key="info.item.path"
        :value="info.item.path"
        :label="setItemLabel(info)"
      ></el-option>
    </el-select>
  </div>
</template>

<script setup lang="ts">
// fuse is a lightweight fuzzy-search module
// make search results more in line with expectations
import Fuse from 'fuse.js'
import * as path from 'path'

import { ref, computed, watch, nextTick, onMounted, toRaw } from 'vue'
import { storeToRefs } from 'pinia'
import { useRouter } from 'vue-router'
import { usePermissionStore } from '@/store/modules/permission'

const search = ref('')
const options = ref([])
const searchPool = ref([])
const show = ref(false)
let fuse = undefined

const { routers } = storeToRefs(usePermissionStore())
const routes = routers.value

const router = useRouter()

const headerSearchSelectRef = ref(null)

// Filter out the routes that can be displayed in the sidebar
// And generate the internationalized title
const generateRoutes = (routes, basePath = '/', prefixTitle = []) => {
  let res = []
  for (const router of routes) {
    // skip hidden router
    if (router.hidden) {
      continue
    }

    const data = {
      path: path.resolve(basePath, router.path),
      title: [...prefixTitle]
    }

    if (router.meta && router.meta.title) {
      // generate internationalized title
      const i18ntitle = router.meta.title

      data.title = [...data.title, i18ntitle]

      if (router.redirect !== 'noredirect') {
        // only push the routes with title
        // special case: need to exclude parent router without redirect
        res.push(data)
      }
    }

    // recursive child routes
    if (router.children) {
      const tempRoutes = generateRoutes(router.children, data.path, data.title)
      if (tempRoutes.length >= 1) {
        res = [...res, ...tempRoutes]
      }
    }
  }
  return res
}

const click = () => {
  show.value = !show.value
  if (show.value) {
    headerSearchSelectRef.value && headerSearchSelectRef.value.focus()
  }
}

const close = () => {
  headerSearchSelectRef.value && headerSearchSelectRef.value.blur()
  options.value = []
  show.value = false
}

const change = path => {
  router.push(path)
  search.value = ''
  options.value = []
  nextTick(() => {
    show.value = false
  })
}

const querySearch = (query: any) => {
  if (query !== '') {
    options.value = fuse.search(query)
  } else {
    options.value = []
  }
}

const setItemLabel = (data: any) => {
  return data.item.title.join(' > ')
}

const initFuse = (list: any) => {
  fuse = new Fuse(list, {
    shouldSort: true,
    threshold: 0.4,
    location: 0,
    distance: 100,

    // maxPatternLength: 32, todo 废弃了？
    minMatchCharLength: 1,
    keys: [
      {
        name: 'title',
        weight: 0.7
      },
      {
        name: 'path',
        weight: 0.3
      }
    ]
  })
}

onMounted(() => {
  searchPool.value = generateRoutes(routes)
})

watch(
  () => routes,
  () => {
    searchPool.value = generateRoutes(routes)
  }
)

watch(
  () => searchPool.value,
  list => {
    initFuse(list)
  }
)

watch(show, value => {
  if (value) {
    document.body.addEventListener('click', close)
  } else {
    document.body.removeEventListener('click', close)
  }
})
</script>

<style lang="scss" scoped>
.header-search {
  font-size: 0 !important;

  .search-icon {
    cursor: pointer;
    font-size: 18px;
    vertical-align: middle;
  }

  .header-search-select {
    font-size: 18px;
    transition: width 0.2s;
    width: 0;
    overflow: hidden;
    background: transparent;
    border-radius: 0;
    display: inline-block;
    vertical-align: middle;

    ::v-deep(.el-input__inner) {
      border-radius: 0;
      border: 0;
      padding-left: 0;
      padding-right: 0;
      box-shadow: none !important;
      border-bottom: 1px solid #d9d9d9;
      vertical-align: middle;
    }
  }

  &.show {
    .header-search-select {
      width: 210px;
      margin-left: 10px;
    }
  }
}
</style>
