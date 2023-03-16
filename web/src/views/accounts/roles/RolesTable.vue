<template>
  <div>
    <el-table v-loading="loading" :data="data" ref="tableRef" border>
      <el-table-column prop="charac_no" label="角色编号" width="120" />
      <el-table-column prop="charac_name" label="角色名称" width="180"></el-table-column>
      <el-table-column prop="money" label="金币" width="180">
        <template #default="scope">
          <span>{{ formatPrice(scope.row.money) }}</span>
        </template>
      </el-table-column>
      <el-table-column prop="QP" label="QP" width="120"></el-table-column>
      <el-table-column prop="create_time" label="创建时间" width="180">
        <template #default="scope">
          <span>{{ dateFormat(scope.row.create_time) }}</span>
        </template>
      </el-table-column>
      <el-table-column prop="lev" label="等级"></el-table-column>
      <el-table-column prop="m_id" label="uid" width="120" />
      <el-table-column label="操作" width="120" align="right">
        <template #default="scope">
          <el-button type="primary" link size="small" @click="changeQp(scope.row)">修改QP</el-button>
        </template>
      </el-table-column>
    </el-table>

    <role-qp-dialog ref="qpDialogRef" @reloadRoles="toReloadRoles"></role-qp-dialog>
  </div>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue'
import { RoleState, Role } from '@/views/accounts/roles/model'
import { dateFormat, formatPrice } from '@/utils'
import RoleQpDialog from './components/RoleQPDialog'

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

const qpDialogRef = ref()
const changeQp = (row: Role) => {
  qpDialogRef.value.showQpDialog(row)
}

const toReloadRoles = () => {
  emit('reloadRoles')
}

const emit = defineEmits(['reloadRoles'])
</script>

<style scoped></style>
