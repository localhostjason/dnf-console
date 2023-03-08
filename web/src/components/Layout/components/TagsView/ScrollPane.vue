<template>
  <div ref="scrollContainer">
    <el-scrollbar ref="bar" :vertical="false" class="scroll-container" @wheel.native.prevent="handleScroll">
      <slot />
    </el-scrollbar>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, useSlots } from 'vue'

const tagAndTagSpacing = 4 // tagAndTagSpacing

const left = ref<number>(0)
const scrollContainer = ref(null)
const bar = ref(null)

const scrollWrapper = computed(() => {
  return bar.value.wrapRef
})

const handleScroll = e => {
  const eventDelta = e.wheelDelta || -e.deltaY * 40
  const $scrollWrapper = scrollWrapper.value
  $scrollWrapper.scrollLeft = $scrollWrapper.scrollLeft + eventDelta / 4
}

const moveToTarget = (currentTag, tagList) => {
  const $container = scrollContainer.value
  const $containerWidth = $container.offsetWidth
  const $scrollWrapper = scrollWrapper.value

  if (!tagList.length) {
    return
  }

  let firstTag = null
  let lastTag = null

  // find first tag and last tag
  if (tagList.length > 0) {
    firstTag = tagList[0]
    lastTag = tagList[tagList.length - 1]
  }

  if (firstTag && firstTag.to.path === currentTag.to.path) {
    $scrollWrapper.scrollLeft = 0
  } else if (lastTag && lastTag.to.path === currentTag.to.path) {
    $scrollWrapper.scrollLeft = $scrollWrapper.scrollWidth - $containerWidth
  } else {
    const currentIndex = tagList.findIndex(item => item.to.path === currentTag.to.path)
    const prevTag = tagList[currentIndex - 1]
    const nextTag = tagList[currentIndex + 1]

    // the tag's offsetLeft after of nextTag
    const afterNextTagOffsetLeft = nextTag.$el.offsetLeft + nextTag.$el.offsetWidth + tagAndTagSpacing

    // the tag's offsetLeft before of prevTag
    const beforePrevTagOffsetLeft = prevTag.$el.offsetLeft - tagAndTagSpacing

    if (afterNextTagOffsetLeft > $scrollWrapper.scrollLeft + $containerWidth) {
      $scrollWrapper.scrollLeft = afterNextTagOffsetLeft - $containerWidth
    } else if (beforePrevTagOffsetLeft < $scrollWrapper.scrollLeft) {
      $scrollWrapper.scrollLeft = beforePrevTagOffsetLeft
    }
  }
}

defineExpose({
  moveToTarget
})
</script>

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
