<template>
  <div class="dash-foot">
    <el-row>
      <el-col :span="24">
        <el-card class="box-card" shadow="hover">
          <div>
            <el-tabs v-model="activeName" type="card">
              <el-tab-pane label="用户管理" name="first">
                <el-table :data="tableData" border>
                  <el-table-column prop="date" label="日期"></el-table-column>
                  <el-table-column label="配送信息">
                    <el-table-column prop="name" label="姓名"></el-table-column>
                    <el-table-column label="地址">
                      <el-table-column prop="province" label="省份"></el-table-column>
                      <el-table-column prop="city" label="市区"></el-table-column>
                      <el-table-column prop="address" label="地址"></el-table-column>
                      <el-table-column prop="zip" label="邮编"></el-table-column>
                    </el-table-column>
                  </el-table-column>
                </el-table>
              </el-tab-pane>
              <el-tab-pane label="配置管理" name="second">
                <el-table :data="tableData" border>
                  <el-table-column prop="date" label="日期"></el-table-column>
                  <el-table-column label="配送信息">
                    <el-table-column prop="name" label="姓名"></el-table-column>
                    <el-table-column label="地址">
                      <el-table-column prop="province" label="省份"></el-table-column>
                      <el-table-column prop="city" label="市区"></el-table-column>
                      <el-table-column prop="address" label="地址"></el-table-column>
                      <el-table-column prop="zip" label="邮编"></el-table-column>
                    </el-table-column>
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
import { ref } from 'vue'

const activeName = ref<string>('first')

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
