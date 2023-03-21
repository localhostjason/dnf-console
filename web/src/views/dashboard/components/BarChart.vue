<template>
  <div>
    <div ref="chartRef" :style="{ height, width }"></div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, Ref, ref, toRefs, watch } from 'vue'
import { useECharts } from '@/utils/echarts/useEcharts'

const props = defineProps({
  width: {
    type: String as PropType<string>,
    default: '100%'
  },
  height: {
    type: String as PropType<string>,
    default: '285px'
  },
  data: {
    type: Object as PropType<Chart>,
    required: true
  }
})

const { data } = toRefs(props)

const newData = ref<Chart>()
watch(data, (val: Chart) => {
  newData.value = val
  setOps()
})

const chartRef = ref<HTMLDivElement | null>(null)
const { setOptions, echarts, getInstance } = useECharts(chartRef as Ref<HTMLDivElement>)

const loading = () => {
  const lienEcharts = getInstance()
  lienEcharts.showLoading({
    text: '正在加载数据...',
    color: '#1890FF',
    textColor: '#000',
    maskColor: 'rgba(255, 255, 255, 0.2)',
    zlevel: 0
  })
}

const hideLoading = () => {
  const lienEcharts = getInstance()
  lienEcharts.hideLoading()
}

const setOps = () => {
  if (!newData.value) {
    loading()
    return
  }
  hideLoading()
  const { cera, cera_point } = newData.value

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
      data: cera.date
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
        data: cera.total,
        areaStyle: {}
      },
      {
        name: 'D点',
        type: 'line',
        symbol: 'none',
        smooth: true,
        data: cera_point.total,
        areaStyle: {}
      }
    ]
  })
}

onMounted(() => {
  setOps()
})
</script>

<style lang="scss" scoped></style>
