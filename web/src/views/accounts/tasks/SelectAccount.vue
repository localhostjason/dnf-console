<template>
  <div>
    <el-form :model="form" :rules="rules" ref="formRef" label-width="100px">
      <el-form-item label="账号ID" prop="uid">
        <el-select
          v-model="form.uid"
          filterable
          clearable
          placeholder="选择账号id"
          :loading="loading"
          loading-text="正在加载账号"
        >
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
  </div>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue'
import { FormInstance, FormRules } from 'element-plus'
import { AccountOptions, SelectForm } from '@/views/accounts/roles/model'
import { getAccounts } from '@/api/gm/accounts'
import { validate } from '@/utils/element/form'

const formRef = ref<FormInstance>()
const form = reactive<SelectForm>({
  uid: null
})
const rules = reactive<FormRules>({
  uid: [{ required: true, message: '请选择账号', trigger: 'blur' }]
})

const options = reactive<AccountOptions>({
  data: []
})

const loading = ref<boolean>(false)

const getAccountsOptions = async () => {
  loading.value = true
  options.data = await getAccounts({ has_roles: true })
  loading.value = false
}

const selectRoles = async () => {
  const valid = await validate(formRef)
  if (!valid) return
  emit('setUid', form.uid)
}

// hook
getAccountsOptions()

const emit = defineEmits(['setUid'])
</script>

<style scoped></style>
