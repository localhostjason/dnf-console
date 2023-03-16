<template>
  <div>
    <panel-title title="角色管理"></panel-title>

    <div class="tc-step-border">
      <div class="l">
        <panel-step num="1" title="账号选择">
          <select-account @setUid="setUid"></select-account>
        </panel-step>

        <panel-step num="2" title="角色信息" is-last>
          <roles-table :data="state.data" :loading="state.loading" @reloadRoles="getRoleList"></roles-table>
        </panel-step>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import PanelTitle from '@/components/PanelTitle'
import PanelStep from '@/components/PanelStep'
import RolesTable from './RolesTable'
import SelectAccount from '../tasks/SelectAccount'

import { reactive, ref } from 'vue'
import { AccountOptions, RoleState, SelectForm } from './model'
import { getRoles } from '@/api/gm/role'

const state = reactive<RoleState>({
  data: [],
  loading: false
})

let uid: number = null
const setUid = (id: number) => {
  uid = id
  getRoleList()
}

const getRoleList = async () => {
  try {
    state.loading = true
    state.data = await getRoles(uid)
    state.loading = false
  } catch (e) {
    state.loading = false
  }
}
</script>

<style scoped></style>
