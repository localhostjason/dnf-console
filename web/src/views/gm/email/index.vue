<template>
  <div>
    <el-row>
      <panel-title title="邮件工具"></panel-title>
    </el-row>

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
            @setCharacNo="setCharacNo"
          ></role-table>
        </panel-step>
        <panel-step num="3" title="发送邮件" is-last>
          <email-form ref="emailFormRef"></email-form>
        </panel-step>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import PanelTitle from '@/components/PanelTitle'
import PanelStep from '@/components/PanelStep'
import SelectAccount from '../tasks/SelectAccount'
import EmailForm from './EmailForm'

import { getRoles } from '@/api/gm/role'
import { reactive, ref } from 'vue'
import { RoleState } from '@/views/gm/roles/model'
import RoleTable from '../tasks/RoleTable'
import { any } from 'vue-types'

const state = reactive<RoleState>({
  data: [],
  loading: false
})

const setUid = async (uid: number) => {
  try {
    state.loading = true
    state.data = await getRoles(uid)
    state.loading = false
  } catch (e) {
    state.loading = false
  }
}

const emailFormRef = ref<any>(null)
const setCharacNo = (characNo: number) => {
  emailFormRef.value.setCharacNo(characNo)
}
</script>

<style scoped></style>
