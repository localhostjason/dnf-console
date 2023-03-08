<template>
  <svg :class="svgClass" aria-hidden="true">
    <use :xlink:href="iconName"></use>
  </svg>
</template>

<script lang="ts">
import { propTypes } from '@/utils/propTypes'
import { defineComponent, computed } from 'vue'

export default defineComponent({
  props: {
    iconClass: propTypes.string.def(''),
    className: propTypes.string.def('')
  },
  setup(props, ctx) {
    const iconName = computed(() => {
      return `#icon-${props.iconClass}`
    })

    const svgClass = computed(() => {
      if (props.className) {
        return 'svg-icon ' + props.className
      } else {
        return 'svg-icon'
      }
    })

    return {
      iconName,
      svgClass
    }
  }
})
</script>

<style scoped>
.svg-icon {
  width: 1em;
  height: 1em;
  vertical-align: -0.15em;
  fill: currentColor;
  overflow: hidden;
}
</style>
