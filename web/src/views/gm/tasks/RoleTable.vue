<template>
  <div>
    <el-table v-loading="loading" :data="data" ref="tableRef" border>
      <el-table-column width="50" label="#" align="center">
        <template #default="scope">
          <el-radio v-model="characNo" :label="scope.row.charac_no" class="radio" @change="selectCharacNo"></el-radio>
        </template>
      </el-table-column>
      <el-table-column prop="charac_no" label="角色编号" width="180" />
      <el-table-column prop="charac_name" label="角色名称" width="180"></el-table-column>
      <el-table-column prop="create_time" label="创建时间" width="180">
        <template #default="scope">
          <span>{{ dateFormat(scope.row.create_time) }}</span>
        </template>
      </el-table-column>
      <el-table-column prop="lev" label="等级"></el-table-column>
      <el-table-column prop="m_id" label="uid" width="150" />
    </el-table>
  </div>
</template>

<script setup lang="ts">
import { defineExpose, reactive, ref } from 'vue'
import { RoleState, Role } from '@/views/gm/roles/model'
import { dateFormat } from '@/utils'

const characNo = ref<number>(null)

defineProps({
  data: {
    type: Array as PropType<Role[]>,
    required: true,
    default: () => []
  },
  loading: {
    type: Boolean,
    required: true,
    default: false
  }
})

const selectCharacNo = () => {
  emit('setCharacNo', characNo.value)
}

const resetCharacNo = () => {
  characNo.value = null
}

// hook
const emit = defineEmits(['setCharacNo'])

defineExpose({
  resetCharacNo
})
</script>

<style scoped></style>
