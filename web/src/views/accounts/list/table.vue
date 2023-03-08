<template>
  <div>
    <el-row>
      <el-button type="primary">注册账号</el-button>
    </el-row>
    <el-row>
      <el-table v-loading="state.loading" :data="state.data" ref="tableRef" border>
        <el-table-column prop="uid" label="UID" width="180" />
        <el-table-column prop="account_name" label="账号名" width="200" />
        <el-table-column prop="roles" label="创建角色总数" width="120" align="center" />
        <el-table-column prop="money" label="携带金钱" width="180">
          <template #default="scope">
            <span>{{ formatPrice(scope.row.money) }}</span>
          </template>
        </el-table-column>
        <el-table-column label="操作" align="right">
          <template #default="scope">
            <el-button type="text">修改密码</el-button>
            <el-button type="text">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-row>
    <el-row>
      <el-pagination
        :current-page="pageQuery.page"
        :page-sizes="[10, 20, 30, 40, 50]"
        :page-size="pageQuery.per_page"
        :total="state.total"
        background
        layout="total, sizes, prev, pager, next, jumper"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      ></el-pagination>
    </el-row>
  </div>
</template>

<script setup lang="ts">
import { getAccounts } from '@/api/gm/accounts'
import { reactive } from 'vue'
import { AccountState, FilterAccountForm } from '@/views/accounts/list/model'
import { PageQuery } from '@/models/page'

// define
const state = reactive<AccountState>({
  data: [],
  total: 0,
  loading: false
})

const pageQuery = reactive<PageQuery>({
  page: 1,
  per_page: 10
})

let argQuery = reactive<FilterAccountForm>({
  uid: ''
})

const getAccountList = async () => {
  state.loading = true
  const params = {
    ...argQuery,
    ...pageQuery
  }
  const { items, total } = await getAccounts(params)
  state.data = items
  state.total = total
  state.loading = false
}

// methods
const handleSizeChange = (params_limit: number) => {
  pageQuery.per_page = params_limit
  getAccountList()
}

const handleCurrentChange = (params_page: number) => {
  pageQuery.page = params_page
  getAccountList()
}

const setParams = (data: FilterAccountForm) => {
  argQuery = data
  getAccountList()
}

function formatPrice(price) {
  try {
    return String(price).replace(/\B(?=(\d{3})+(?!\d))/g, ',')
  } catch (e) {
    return price
  }
}

// hook
getAccountList() // created

defineExpose({
  setParams
})
</script>

<style scoped></style>
