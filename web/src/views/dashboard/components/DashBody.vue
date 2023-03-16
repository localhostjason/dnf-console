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
            <bar-chart></bar-chart>
          </div>
        </el-card>
      </el-col>

      <el-col :span="6">
        <el-card class="box-card" shadow="hover">
          <template #header>
            <div class="card-header">
              <svg-icon icon-class="list"></svg-icon>
              <span style="position: relative; left: 10px">充值 Top5</span>
            </div>
          </template>
          <div class="dash-table">
            <el-table :data="tableData" style="width: 100%" :show-header="false" max-height="245px">
              <el-table-column prop="name" label="名称" width="130" />
              <el-table-column prop="number" label="金额">
                <template #default="scope">
                  <el-progress :percentage="scope.row.percent">
                    <el-button text>{{ scope.row.number }}</el-button>
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
              <span style="position: relative; left: 10px">在线Top5</span>
            </div>
          </template>
          <div class="dash-table">
            <el-table :data="tableData" style="width: 100%" :show-header="false" max-height="245px">
              <el-table-column prop="name" label="名称" width="130" />
              <el-table-column prop="number" label="金额">
                <template #default="scope">
                  <el-progress :percentage="scope.row.percent">
                    <el-button text>{{ scope.row.number }}</el-button>
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

const tableData = [
  {
    name: 'TomAdmin',
    percent: 80,
    number: '12,000'
  },
  {
    name: 'Test',
    percent: 70,
    number: '9,999'
  },
  {
    name: 'Tom',
    percent: 60,
    number: '6,000'
  },
  {
    name: 'admin',
    percent: 50,
    number: '3,000'
  },
  {
    name: 'admin123',
    percent: 10,
    number: '1,000'
  }
]
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
