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
            <el-form-item label="用户名" :prop="user_id ? '-' : 'username'">
              <span class="text-gray-lg sm" v-if="user_id">{{ form.username }}</span>
              <el-input v-model="form.username" v-else></el-input>
            </el-form-item>

            <el-form-item label="电子邮件" prop="email">
              <el-input v-model="form.email"></el-input>
            </el-form-item>

            <el-form-item label="密码:" prop="password" v-if="!user_id">
              <el-input type="password" v-model="form.password"></el-input>
            </el-form-item>
            <el-form-item label="确认密码:" prop="checkPassword" v-if="!user_id">
              <el-input type="password" v-model="form.checkPassword"></el-input>
            </el-form-item>

            <el-form-item label="描述" prop="desc">
              <el-input type="textarea" v-model="form.desc" :autosize="{ minRows: 4 }"></el-input>
            </el-form-item>
          </el-form>
        </el-col>
      </el-row>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialog.visible = false">取 消</el-button>
          <el-button type="primary" @click="updateUser">确 定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { nextTick, reactive, ref } from 'vue'
import { FormInstance, FormRules } from 'element-plus'

import { validate } from '@/utils/element/form'
import { successMessage } from '@/utils/element/message'
import { ModifyUserForm, User } from '@/models/user/user'

import { validateReg, setFormData, getFormDataByFields } from '@/utils'

const dialog = reactive({
  visible: false,
  title: ''
})
const formRef = ref<FormInstance>()
const user_id = ref<number | undefined>(void 0)

const form = reactive<ModifyUserForm>({
  username: '',
  email: '',
  desc: '',
  password: '',
  checkPassword: ''
})

const validatePass = (rule: any, value: any, callback: any) => {
  if (!value) callback(new Error('请输入密码'))
  if (form.checkPassword !== '') {
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
  email: [{ pattern: validateReg.email, message: '请输入正确邮箱', trigger: 'blur' }],
  username: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
  password: [{ validator: validatePass, required: true, trigger: 'blur' }],
  checkPassword: [{ validator: validatePass2, required: true, trigger: 'blur' }],
  desc: [{ max: 128, message: '最大字数不超过128个', trigger: ['blur', 'change'] }]
})

const showDialog = (row: User | null = null) => {
  dialog.visible = true

  dialog.title = row ? '编辑用户' : '创建用户'
  user_id.value = row ? row.id : void 0

  nextTick(() => {
    if (row) {
      // 注：可不这么写，虽然少写代码，用的不当，会有bug。前提 row 跟form 字段要一致或者包含
      setFormData(form, row)

      formRef.value.clearValidate()
      return
    }
    formRef.value.resetFields()
  })
}

const emit = defineEmits(['reloadUsers'])
const updateUser = async () => {
  const valid = await validate(formRef)
  if (!valid) return

  try {
    if (user_id.value) {
      const data = getFormDataByFields(form, ['email', 'desc'])
      // send update user api
      console.log('put user', data)
      successMessage('修改用户成功')
    } else {
      // send create user api
      console.log('post user', form)
      successMessage('创建用户成功')
    }
    dialog.visible = false
    emit('reloadUsers')
  } catch (e) {}
}

defineExpose({
  showDialog
})
</script>

<style></style>
