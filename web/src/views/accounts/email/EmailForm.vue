<template>
  <div>
    <el-row :gutter="20">
      <el-col :span="5">
        <el-form :model="form" :rules="rules" ref="formRef" label-width="100px">
          <el-form-item label="物品代码" prop="code">
            <el-input v-model.number="form.code"></el-input>
          </el-form-item>

          <el-form-item label="数量" prop="number">
            <el-input v-model.number="form.number"></el-input>
          </el-form-item>

          <el-form-item>
            <el-button type="primary" size="small" @click="sendEmail">发送</el-button>
          </el-form-item>
        </el-form>
      </el-col>
      <el-col :span="6" :offset="13">
        <el-card class="box-card">
          <template #header>
            <div class="card-header">
              <span>物品代码</span>
            </div>
          </template>
          <div v-for="o in 4" :key="o" class="text item">{{ 'List item ' + o }}</div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue'
import { FormInstance, FormRules } from 'element-plus'
import { Email } from '@/views/accounts/email/model'
import { validate } from '@/utils/element/form'
import { sendEmailByRole, getGolds } from '@/api/gm/email'
import { errorMessage, successMessage } from '@/utils/element/message'

const formRef = ref<FormInstance>()

const form = reactive<Email>({
  code: null,
  number: 0
})

const rules = reactive<FormRules>({
  code: [{ required: true, message: '请输入物品代码', trigger: 'blur' }],
  number: [{ required: true, message: '请输入数量', trigger: 'blur' }]
})

const characNo = ref<number>(null)

const sendEmail = async () => {
  const valid = await validate(formRef)
  if (!valid) return

  if (!characNo.value) {
    errorMessage('未获取到角色ID')
    return
  }

  try {
    await sendEmailByRole(characNo.value, form)
    successMessage('发送成功，请小退一下！')
  } catch (e) {}
}

const setCharacNo = (id: number) => {
  characNo.value = id
}

const getGoldList = async () => {
  try {
    await getGolds()
  } catch (e) {}
}

// hook
getGoldList()
defineExpose({
  setCharacNo
})
</script>

<style scoped></style>
