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
            <el-form-item label="账号用户名" prop="account_name">
              <el-input v-model="form.account_name"></el-input>
            </el-form-item>

            <el-form-item label="QQ" prop="qq">
              <el-input v-model="form.qq"></el-input>
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
          <el-button type="primary" @click="createAccount">确 定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { nextTick, reactive, ref } from 'vue'
import { FormInstance, FormRules } from 'element-plus'
import { CreateAccountForm } from '@/views/gm/account/model/info'
import { AccountDetail } from '@/views/gm/account/model/model'
import { setFormData } from '@/utils'
import { validate } from '@/utils/element/form'
import { createAccountInfo } from '@/api/gm/accounts'
import { successMessage } from '@/utils/element/message'

const dialog = reactive({
  visible: false,
  title: ''
})
const formRef = ref<FormInstance>()
const form = reactive<CreateAccountForm>({
  account_name: '',
  qq: '',
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
  account_name: [{ required: true, message: '请输入账号名称', trigger: 'blur' }],
  password: [{ validator: validatePass, required: true, trigger: 'blur' }],
  check_password: [{ validator: validatePass2, required: true, trigger: 'blur' }]
})

const createAccount = async () => {
  const valid = await validate(formRef)
  if (!valid) return

  try {
    await createAccountInfo(form)
    successMessage('创建成功')
    dialog.visible = false
    emit('reloadAccount')
  } catch (e) {}
}

const showCreateAccountDialog = () => {
  dialog.visible = true
  dialog.title = '创建账号'

  nextTick(() => {
    formRef.value.resetFields()
  })
}

// hook
defineExpose({
  showCreateAccountDialog
})

const emit = defineEmits(['reloadAccount'])
</script>

<style scoped></style>
