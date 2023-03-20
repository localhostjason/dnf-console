<template>
  <div class="dash-body">
    <el-row :gutter="20">
      <el-col :span="12">
        <el-card class="box-card" shadow="hover">
          <template #header>
            <div class="card-header">
              <svg-icon icon-class="pulse"></svg-icon>
              <span style="position: relative; left: 10px">趋势</span>
            </div>
          </template>
          <div>
            <bar-chart :data="chartData"></bar-chart>
          </div>
        </el-card>
      </el-col>

      <el-col :span="6">
        <el-card class="box-card" shadow="hover">
          <template #header>
            <div class="card-header">
              <svg-icon icon-class="list"></svg-icon>
              <span style="position: relative; left: 10px">充值D币 Top5</span>
            </div>
          </template>
          <div class="dash-table">
            <el-table :data="state.cera" style="width: 100%" :show-header="false" max-height="245px" height="245px">
              <el-table-column prop="uid" label="UID" width="130" />
              <el-table-column prop="number" label="金额">
                <template #default="scope">
                  <el-progress :percentage="scope.row.percent">
                    <el-button text>{{ formatPrice(scope.row.total) }}</el-button>
                  </el-progress>
                </template>
              </el-table-column>
            </el-table>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card class="box-card" shadow="hover">
          <template #header>
            <div class="card-header">
              <svg-icon icon-class="timer"></svg-icon>
              <span style="position: relative; left: 10px">充值D点 Top5</span>
            </div>
          </template>
          <div class="dash-table">
            <el-table
              :data="state.cera_point"
              style="width: 100%"
              :show-header="false"
              max-height="245px"
              height="245px"
            >
              <el-table-column prop="uid" label="UID" width="130" />
              <el-table-column prop="number" label="金额">
                <template #default="scope">
                  <el-progress :percentage="scope.row.percent">
                    <el-button text>{{ formatPrice(scope.row.total) }}</el-button>
                  </el-progress>
                </template>
              </el-table-column>
            </el-table>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup lang="ts">
import BarChart from './BarChart'
import { getDashChart, getDashTop5 } from '@/api/dash'
import { onBeforeMount, reactive, ref } from 'vue'
import { formatPrice } from '@/utils'

const state = reactive({
  cera: [],
  cera_point: [],
  cera_total: 0,
  ccera_point_total: 0
})

const getTop5 = async () => {
  const data = await getDashTop5()
  state.cera = data.cera.map(val => {
    val['percent'] = (val.total / data.cera_total) * 100
    return val
  })
  state.cera_point = data.cera_point.map(val => {
    val['percent'] = (val.total / data.cera_point_total) * 100
    return val
  })
  state.cera_total = data.cera_total
  state.ccera_point_total = data.ccera_point_total
}

const chartData = ref<Chart>({
  cera: {
    date: [],
    total: []
  },
  cera_point: {
    date: [],
    total: []
  }
})

const getChartData = async () => {
  chartData.value = await getDashChart()
}

getChartData()

getTop5()
</script>

<style lang="scss" scoped>
.dash-body {
  margin-top: 18px;

  .card-header {
    font-size: 14px;
  }

  ::v-deep(.el-card__body) {
    padding: 0 !important;
  }

  .dash-table {
    padding: 20px;
  }
}
</style>
