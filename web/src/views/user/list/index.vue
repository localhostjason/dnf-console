<template>
  <div>
    <panel-title title="用户管理"></panel-title>

    <el-row>
      <el-button type="primary" @click="download">下载</el-button>
      <el-button type="primary" @click="createUser">创建</el-button>
    </el-row>
    <el-row>
      <el-table v-loading="state.loading" :data="state.data" ref="tableRef" border>
        <el-table-column prop="username" label="用户名" width="180" />
        <el-table-column prop="role" label="角色" width="180" />
        <el-table-column prop="time" label="创建时间" width="180">
          <template #default="scope">
            <span>{{ dateFormat(scope.row.time) }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="last_login_time" label="上次登录时间" width="180">
          <template #default="scope">
            <span v-if="scope.row.last_login_time">{{ dateFormat(scope.row.last_login_time) }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="desc" label="描述" />
        <el-table-column label="操作" width="160">
          <template #default="scope">
            <el-button type="primary" plain size="small" @click="editUser(scope.row)">编辑</el-button>
            <el-button type="danger" size="small" @click="deleteUser(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-row>

    <user-dialog ref="userDialogRef" @reloadUsers="getUsersData"></user-dialog>
  </div>
</template>

<script setup lang="ts">
import PanelTitle from '@/components/PanelTitle'
import UserDialog from './EditCreateUserdialog.vue'
import { onBeforeMount, reactive, ref, toRaw } from 'vue'
import { getUsers, getUsersList, downloadFile } from '@/api/user/users'
import { dateFormat } from '@/utils/filters'
import { downloadFileByBlob } from '@/utils/download'
import { successMessage } from '@/utils/element/message'
import { confirmWarning } from '@/utils/element/messageBox'
import { User, UserState } from '@/models/user/user'

// 定义
const userDialogRef = ref<any>(null)

const state = reactive<UserState>({
  data: [],
  loading: false
})

// method
const getUsersData = async (): Promise<void> => {
  state.loading = true
  state.data = await getUsers()
  state.loading = false
}

const download = async (): Promise<void> => {
  try {
    const { data, filename } = await downloadFile()
    await downloadFileByBlob(data, filename)
    successMessage('下载成功')
  } catch (e) {
    console.log('e', e)
  }
}

const deleteUser = async (row: User): Promise<void> => {
  try {
    const data = await confirmWarning('此操作将永久删除用户，是否继续？')
    console.log('data', data)
  } catch (e) {
    console.log('e', e)
  }
}

const editUser = (row: User): void => {
  userDialogRef.value.showDialog(row)
}
const createUser = (): void => {
  userDialogRef.value.showDialog()
}

// hook
onBeforeMount(() => {
  getUsersData()
})
</script>

<style scoped></style>
