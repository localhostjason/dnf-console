<template>
  <div>
    <panel-title title="测试"></panel-title>

    <div ref="chartRef" :style="{ height, width }"></div>
  </div>
</template>

<script setup lang="ts">
import PanelTitle from '@/components/PanelTitle'
import { useECharts } from '@/utils/echarts/useEcharts'

import { ref, Ref, onMounted } from 'vue'

defineProps({
  width: {
    type: String as PropType<string>,
    default: '100%'
  },
  height: {
    type: String as PropType<string>,
    default: '500px'
  }
})

const chartRef = ref<HTMLDivElement | null>(null)
const { setOptions, echarts } = useECharts(chartRef as Ref<HTMLDivElement>)

onMounted(() => {
  setOptions({
    title: {
      text: 'ECharts 入门示例'
    },
    tooltip: {},
    legend: {
      data: ['销量']
    },
    xAxis: {
      data: ['衬衫', '羊毛衫', '雪纺衫', '裤子', '高跟鞋', '袜子']
    },
    yAxis: {},
    series: [
      {
        name: '销量',
        type: 'bar',
        data: [5, 20, 36, 10, 10, 20]
      }
    ]
  })
})
</script>

<style scoped></style>
