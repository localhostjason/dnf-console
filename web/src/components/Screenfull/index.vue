<template>
  <div>
    <svg-icon :icon-class="isFullscreen ? 'exit-fullscreen' : 'fullscreen'" @click="click"></svg-icon>
  </div>
</template>

<script lang="ts">
import screenfull from 'screenfull'
import { defineComponent, ref, onMounted, onBeforeUnmount } from 'vue'
import { warnMessage } from '@/utils/element/message'

export default defineComponent({
  name: 'ScreenFull',
  setup() {
    const isFullscreen = ref(false)

    const click = async () => {
      if (!screenfull.isEnabled) {
        warnMessage('当前浏览器不支持')
        return false
      }
      await screenfull.toggle()
    }

    const change = (): void => {
      isFullscreen.value = screenfull.isFullscreen
    }

    const init = () => {
      if (screenfull.isEnabled) {
        screenfull.on('change', change)
      }
    }

    const destroy = () => {
      if (screenfull.isEnabled) {
        screenfull.off('change', change)
      }
    }

    onMounted(() => {
      init()
    })

    onBeforeUnmount(() => {
      destroy()
    })

    return {
      isFullscreen,
      click
    }
  }
})
</script>

<style scoped>
.screenfull-svg {
  display: inline-block;
  cursor: pointer;
  fill: #5a5e66;
  width: 20px;
  height: 20px;
  vertical-align: 10px;
}
</style>
