<template>
  <div>
    <el-row>
      <el-button type="primary" size="small" @click="doneTask" :disabled="!multipleSelection.length"
        >批量完成
      </el-button>
    </el-row>
    <el-table v-loading="loading" :data="data" ref="tableRef" border @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="55" :selectable="selectable" />
      <el-table-column prop="play_1" label="任务ID" width="180"></el-table-column>
      <el-table-column prop="play_1_trigger" label="是否已完成" width="180">
        <template #default="scope">
          <span>{{ scope.row.play_1_trigger ? '未完成' : '已完成' }}</span>
        </template>
      </el-table-column>
      <el-table-column prop="charac_no" label="角色编号"></el-table-column>
    </el-table>
  </div>
</template>

<script setup lang="ts">
import { Task } from './model'
import { ref } from 'vue'
import { confirmWarning } from '@/utils/element/messageBox'
import { updateTasks } from '@/api/gm/task'
import { errorMessage, successMessage } from '@/utils/element/message'
import { getIdFormArray } from '@/utils'

const multipleSelection = ref<Task[]>([])
const characNo = ref<number>(null)

// method
const handleSelectionChange = (val: Task[]) => {
  multipleSelection.value = val
}

const doneTask = async () => {
  try {
    await confirmWarning('是否批量完成任务？')
    if (!characNo.value) {
      errorMessage('未获取到角色ID')
      return
    }

    const ids = getIdFormArray(multipleSelection.value, 'play_1')
    await updateTasks(characNo.value, { ids: ids })
    successMessage('批量更新成功！')
    emit('reloadTask')
  } catch (e) {}
}

const setCharacNo = (id: number) => {
  characNo.value = id
}

const selectable = row => {
  return row.play_1_trigger !== 0
}

// hook
defineProps({
  data: {
    type: Array as PropType<Task[]>,
    required: true,
    default: () => []
  },
  loading: {
    type: Boolean,
    required: true,
    default: false
  }
})

defineExpose({
  setCharacNo
})

const emit = defineEmits(['reloadTask'])
</script>

<style scoped></style>
