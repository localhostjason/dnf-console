<template>
  <div>
    <panel-title title="角色管理"></panel-title>

    <div class="tc-step-border">
      <div class="l">
        <panel-step num="1" title="账号选择">
          <el-form :model="form" :rules="rules" ref="formRef" label-width="100px">
            <el-form-item prop="uid">
              <el-select v-model="form.uid" filterable clearable placeholder="选择账号id" v-loading="loading">
                <el-option v-for="item in options.data" :key="item.uid" :label="item.uid" :value="item.uid">
                  <span style="float: left">{{ item.uid }}</span>
                  <span style="float: right; color: var(--el-text-color-secondary); font-size: 13px">
                    角色数:{{ item.roles }}
                  </span>
                </el-option>
              </el-select>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="selectRoles" size="small">查询</el-button>
            </el-form-item>
          </el-form>
        </panel-step>

        <panel-step num="2" title="角色信息" is-last>
          <div style="padding: 30px">
            <roles-table :data="state.data" :loading="state.loading"></roles-table>
          </div>
        </panel-step>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import PanelTitle from '@/components/PanelTitle'
import PanelStep from '@/components/PanelStep'
import RolesTable from './RolesTable'

import { reactive, ref } from 'vue'
import { FormInstance, FormRules } from 'element-plus'
import { getAccounts } from '@/api/gm/accounts'
import { validate } from '@/utils/element/form'
import { AccountOptions, RoleState, SelectForm } from './model'
import { getRoles } from '@/api/gm/role'

const formRef = ref<FormInstance>()
const form = reactive<SelectForm>({
  uid: null
})
const rules = reactive<FormRules>({
  uid: [{ required: true, message: '请选择账号', trigger: 'blur' }]
})

const state = reactive<RoleState>({
  data: [],
  loading: false
})

const options = reactive<AccountOptions>({
  data: []
})

const loading = ref<boolean>(false)

const getAccountsOptions = async () => {
  loading.value = true
  options.data = await getAccounts({})
  loading.value = false
}

const selectRoles = async () => {
  const valid = await validate(formRef)
  if (!valid) return

  state.loading = true
  state.data = await getRoles(form.uid)
  state.loading = false
}

// hook
getAccountsOptions()
</script>

<style scoped></style>
