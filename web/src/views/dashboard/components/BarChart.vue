<template>
  <div>
    <div ref="chartRef" :style="{ height, width }"></div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, Ref, ref } from 'vue'
import { useECharts } from '@/utils/echarts/useEcharts'

defineProps({
  width: {
    type: String as PropType<string>,
    default: '100%'
  },
  height: {
    type: String as PropType<string>,
    default: '285px'
  }
})

const chartRef = ref<HTMLDivElement | null>(null)
const { setOptions, echarts } = useECharts(chartRef as Ref<HTMLDivElement>)

const getX = () => {
  let data = []
  for (let i = 1; i <= 12; i++) {
    data.push(`${i}月`)
  }
  return data
}

onMounted(() => {
  setOptions({
    tooltip: {
      trigger: 'axis',
      axisPointer: {
        type: 'cross',
        label: {
          backgroundColor: '#6a7985'
        }
      }
    },
    grid: {
      left: 60,
      right: 60,
      bottom: 50
    },
    xAxis: {
      type: 'category',
      boundaryGap: false,
      data: getX()
    },
    yAxis: {
      type: 'value'
    },
    color: ['#1890FF', '#36CBCB'],
    series: [
      {
        name: 'D币',
        type: 'line',
        smooth: true,
        symbol: 'none',
        data: [5, 20, 36, 10, 10, 20, 100, 80, 100, 0, 0, 100],
        areaStyle: {}
      },
      {
        name: 'D点',
        type: 'line',
        symbol: 'none',
        smooth: true,
        data: [22, 10, 26, 70, 10, 30, 80, 10, 20, 100, 100, 0],
        areaStyle: {}
      }
    ]
  })
})
</script>

<style lang="scss" scoped></style>
