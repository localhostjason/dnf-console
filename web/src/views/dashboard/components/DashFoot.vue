<template>
  <div class="dash-foot">
    <el-row>
      <el-col :span="24">
        <el-card class="box-card" shadow="hover">
          <div>
            <el-tabs v-model="activeName" type="card">
              <el-tab-pane label="最近充值记录" name="first">
                <el-table :data="rechargeList.data" border>
                  <el-table-column prop="time" label="日期">
                    <template #default="scope">
                      <span>{{ dateFormat(scope.row.time) }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column label="充值详情">
                    <el-table-column prop="uid" label="游戏账号"></el-table-column>
                    <el-table-column prop="ip" label="IP地址"></el-table-column>
                    <el-table-column prop="action" label="充值类型">
                      <template #default="scope">
                        <el-tag v-if="scope.row.action === 1">D币</el-tag>
                        <el-tag v-else type="info">D点</el-tag>
                      </template>
                    </el-table-column>
                    <el-table-column prop="number" label="金额"></el-table-column>
                  </el-table-column>
                </el-table>
              </el-tab-pane>
              <el-tab-pane label="最近登录记录" name="second">
                <el-table :data="operateList.data" border>
                  <el-table-column prop="time" label="日期">
                    <template #default="scope">
                      <span>{{ dateFormat(scope.row.time) }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column label="操作信息">
                    <el-table-column prop="username" label="用户名"></el-table-column>
                    <el-table-column prop="ip" label="IP地址"></el-table-column>
                    <el-table-column prop="action" label="操作"></el-table-column>
                    <el-table-column prop="result" label="详情"></el-table-column>
                  </el-table-column>
                </el-table>
              </el-tab-pane>
            </el-tabs>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup lang="ts">
import { getRechargeLog, getOperateLog } from '@/api/log'
import { reactive, ref } from 'vue'
import { PageQuery } from '@/models/page'
import { dateFormat } from '@/utils'

const activeName = ref<string>('first')

const pageQuery = reactive<PageQuery>({
  page: 1,
  per_page: 10
})

const rechargeList = reactive({
  data: []
})
const operateList = reactive({
  data: []
})

const getRechargeData = async () => {
  const params = {
    ...pageQuery
  }
  const { items } = await getRechargeLog(params)
  rechargeList.data = items
}

const getOperateData = async () => {
  const params = {
    ...pageQuery
  }
  const { items } = await getOperateLog(params)
  operateList.data = items
}

getRechargeData()
getOperateData()

const tableData = [
  {
    date: '2016-05-03',
    name: '王小虎',
    province: '上海',
    city: '普陀区',
    address: '上海市普陀区金沙江路 1518 弄',
    zip: 200333
  },
  {
    date: '2016-05-02',
    name: '王小虎',
    province: '上海',
    city: '普陀区',
    address: '上海市普陀区金沙江路 1518 弄',
    zip: 200333
  },
  {
    date: '2016-05-04',
    name: '王小虎',
    province: '上海',
    city: '普陀区',
    address: '上海市普陀区金沙江路 1518 弄',
    zip: 200333
  },
  {
    date: '2016-05-01',
    name: '王小虎',
    province: '上海',
    city: '普陀区',
    address: '上海市普陀区金沙江路 1518 弄',
    zip: 200333
  },
  {
    date: '2016-05-08',
    name: '王小虎',
    province: '上海',
    city: '普陀区',
    address: '上海市普陀区金沙江路 1518 弄',
    zip: 200333
  },
  {
    date: '2016-05-06',
    name: '王小虎',
    province: '上海',
    city: '普陀区',
    address: '上海市普陀区金沙江路 1518 弄',
    zip: 200333
  },
  {
    date: '2016-05-07',
    name: '王小虎',
    province: '上海',
    city: '普陀区',
    address: '上海市普陀区金沙江路 1518 弄',
    zip: 200333
  }
]
</script>

<style scoped>
.dash-foot {
  margin-top: 18px;
}
</style>
