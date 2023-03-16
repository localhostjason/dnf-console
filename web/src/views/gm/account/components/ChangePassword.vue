<template>
  <div>
    <el-dialog
      v-model="dialog.visible"
      :title="dialog.title"
      width="600px"
      :close-on-press-escape="false"
      :close-on-click-modal="false"
    >
      <el-row>
        <el-col :span="22">
          <el-form :model="form" :rules="rules" ref="formRef" label-width="100px">
            <el-form-item label="账号ID">
              <span class="text-gray-lg sm">{{ accountInfo.uid }}</span>
            </el-form-item>
            <el-form-item label="账号用户名">
              <span class="text-gray-lg sm">{{ accountInfo.account_name }}</span>
            </el-form-item>

            <el-form-item label="密码" prop="password">
              <el-input type="password" v-model="form.password"></el-input>
            </el-form-item>
            <el-form-item label="确认密码" prop="check_password">
              <el-input type="password" v-model="form.check_password"></el-input>
            </el-form-item>
          </el-form>
        </el-col>
      </el-row>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialog.visible = false">取 消</el-button>
          <el-button type="primary" @click="updateAccountPassword">确 定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { nextTick, reactive, ref } from 'vue'
import { FormInstance, FormRules } from 'element-plus'
import { ChangePasswordForm } from '@/views/gm/account/model/info'
import { AccountDetail } from '@/views/gm/account/model/model'
import { validate } from '@/utils/element/form'
import { updateAccountPasswordByUid } from '@/api/gm/accounts'
import { successMessage } from '@/utils/element/message'

const dialog = reactive({
  visible: false,
  title: ''
})
const formRef = ref<FormInstance>()
const form = reactive<ChangePasswordForm>({
  password: '',
  check_password: ''
})

const validatePass = (rule: any, value: any, callback: any) => {
  if (!value) callback(new Error('请输入密码'))
  if (form.check_password !== '') {
    formRef.value.validateField('checkPassword')
  }
  callback()
}
const validatePass2 = (rule: any, value: any, callback: any) => {
  if (!value) callback(new Error('请再次输入密码'))
  if (value !== form.password) callback(new Error('两次输入密码不一致!'))
  callback()
}

const rules = reactive<FormRules>({
  password: [{ validator: validatePass, required: true, trigger: 'blur' }],
  check_password: [{ validator: validatePass2, required: true, trigger: 'blur' }]
})

// method
let accountInfo = {
  uid: 0,
  account_name: ''
}
const showAccountPasswordDialog = (row: AccountDetail) => {
  dialog.visible = true

  accountInfo.uid = row.uid
  accountInfo.account_name = row.account_name

  dialog.title = '修改账号密码'

  nextTick(() => {
    formRef.value.resetFields()
  })
}

const updateAccountPassword = async () => {
  const valid = await validate(formRef)
  if (!valid) return

  try {
    await updateAccountPasswordByUid(accountInfo.uid, form)
    successMessage('更新密码成功')
    dialog.visible = false
    emit('reloadAccount')
  } catch (e) {}
}

// hook

defineExpose({
  showAccountPasswordDialog
})

const emit = defineEmits(['reloadAccount'])
</script>

<style scoped></style>
