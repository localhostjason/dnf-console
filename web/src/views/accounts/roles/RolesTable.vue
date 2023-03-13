<template>
  <div>
    <el-table v-loading="loading" :data="data" ref="tableRef" border>
      <el-table-column prop="m_id" label="uid" width="180" />
      <el-table-column prop="charac_no" label="角色编号" width="180" />
      <el-table-column prop="charac_name" label="角色名称" width="180"></el-table-column>
      <el-table-column prop="create_time" label="创建时间" width="180">
        <template #default="scope">
          <span>{{ dateFormat(scope.row.create_time) }}</span>
        </template>
      </el-table-column>
      <el-table-column prop="lev" label="等级"></el-table-column>
    </el-table>
  </div>
</template>

<script setup lang="ts">
import { reactive } from 'vue'
import { RoleState, Role } from '@/views/accounts/roles/model'
import { dateFormat } from '@/utils'

// const state = reactive<RoleState>({
//   data: [],
//   loading: false
// })

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
</script>

<style scoped></style>
