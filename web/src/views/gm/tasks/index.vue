<template>
  <div>
    <panel-title title="任务清理"></panel-title>

    <div class="tc-step-border">
      <div class="l">
        <panel-step num="1" title="账号选择">
          <select-account @setUid="setUid"></select-account>
        </panel-step>

        <panel-step num="2" title="角色选择">
          <role-table
            ref="roleTableRef"
            :loading="state.loading"
            :data="state.data"
            @setCharacNo="getTaskList"
          ></role-table>
        </panel-step>

        <panel-step num="3" title="任务详情" is-last>
          <task-detail
            ref="taskRef"
            :loading="taskState.loading"
            :data="taskState.data"
            @reloadTask="reloadTask"
          ></task-detail>
        </panel-step>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import PanelTitle from '@/components/PanelTitle'
import PanelStep from '@/components/PanelStep'
import SelectAccount from './SelectAccount'
import RoleTable from './RoleTable'
import TaskDetail from './TaskDetail'

import { getRoles } from '@/api/gm/role'
import { reactive, ref } from 'vue'
import { RoleState } from '@/views/gm/roles/model'
import { getTasks } from '@/api/gm/task'
import { warnMessage } from '@/utils/element/message'
import { TaskState } from './model'

const state = reactive<RoleState>({
  data: [],
  loading: false
})

const taskState = reactive<TaskState>({
  data: [],
  loading: false
})

const taskRef = ref<any>(null)
const roleTableRef = ref<any>(null)

const resetTaskTable = () => {
  state.data = []
  taskState.data = []
  roleTableRef.value.resetCharacNo()
}

const setUid = async (uid: number) => {
  resetTaskTable()

  try {
    state.loading = true
    state.data = await getRoles(uid)
    state.loading = false
  } catch (e) {
    state.loading = false
  }
}

let characId = null

const getTaskList = async (characNo: number) => {
  if (!characNo) {
    warnMessage('请选择角色')
    return
  }

  characId = characNo
  try {
    taskState.loading = true
    taskState.data = await getTasks(characNo)
    taskRef.value.setCharacNo(characNo)
    taskState.loading = false
  } catch (e) {
    taskState.loading = false
  }
}

const reloadTask = () => {
  getTaskList(characId)
}
</script>

<style scoped></style>
