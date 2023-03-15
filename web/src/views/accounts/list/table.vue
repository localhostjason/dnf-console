<template>
  <div>
    <el-row>
      <el-button type="primary" @click="reg">注册账号</el-button>
    </el-row>
    <el-row>
      <el-table v-loading="state.loading" :data="state.data" ref="tableRef" border>
        <el-table-column prop="uid" label="UID" width="180" />
        <el-table-column prop="account_name" label="账号名" width="200" />
        <el-table-column prop="roles" label="创建角色总数" width="120" align="center" />
        <el-table-column prop="cera_point" label="D点" width="180">
          <template #default="scope">
            <span>{{ formatPrice(scope.row.cera_point) }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="cera" label="D币" width="180">
          <template #default="scope">
            <span>{{ formatPrice(scope.row.cera) }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="money" label="账号金库存款" width="180">
          <template #default="scope">
            <span>{{ formatPrice(scope.row.money) }}</span>
          </template>
        </el-table-column>
        <el-table-column label="操作" align="right">
          <template #default="scope">
            <el-button type="primary" link @click="recharge(scope.row)" size="small">充值</el-button>
            <el-button type="primary" link @click="resetCreateCharacHandler(scope.row)" size="small"
              >重置创建角色</el-button
            >
            <el-button type="primary" link @click="updatePwd" size="small">修改密码</el-button>
            <el-button type="primary" link @click="deleteAccount" size="small">删除</el-button>
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

    <recharge-dialog ref="rechargeDialogRef" @reloadAccount="getAccountList"></recharge-dialog>
  </div>
</template>

<script setup lang="ts">
import { getAccounts, resetCreateCharac } from '@/api/gm/accounts'
import { reactive, ref } from 'vue'
import { AccountState, FilterAccountForm, Account } from '@/views/accounts/list/model'
import { PageQuery } from '@/models/page'
import { successMessage, warnMessage } from '@/utils/element/message'
import RechargeDialog from './components/RechargeDialog'
import { confirmWarning } from '@/utils/element/messageBox'

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

const rechargeDialogRef = ref<any>(null)

const getAccountList = async () => {
  state.loading = true
  const params = {
    ...argQuery,
    ...pageQuery
  }
  const data: Account = await getAccounts(params)
  state.data = data.items
  state.total = data.total
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

const reg = () => {
  warnMessage('未实现')
}

const updatePwd = () => {
  warnMessage('未实现')
}

const deleteAccount = () => {
  warnMessage('未实现')
}

// 充值
const recharge = row => {
  rechargeDialogRef.value.showRechargeDialog(row)
}

const resetCreateCharacHandler = async row => {
  try {
    await confirmWarning('是否重置 创建角色 限制？')
    await resetCreateCharac(row.uid)
    successMessage('重置成功！')
  } catch (e) {}
}

// hook
getAccountList() // created

defineExpose({
  setParams
})
</script>

<style scoped></style>
