<template>
  <div>
    <panel-title title="个人信息"></panel-title>

    <el-row>
      <el-col :span="9">
        <fieldset>
          <h5 style="margin-bottom: 10px; font-size: 13px">基本信息</h5>

          <el-form :model="form" :rules="rules" ref="formRef" label-width="120px">
            <el-form-item label="账号名">
              <span>{{ username }}</span>
            </el-form-item>
            <el-form-item label="描述" prop="desc">
              <el-input type="textarea" v-model="form.desc" :autosize="{ minRows: 4 }"></el-input>
            </el-form-item>
          </el-form>
        </fieldset>
      </el-col>
    </el-row>

    <el-button type="primary" style="margin-left: 120px" @click="saveUserInfo">保 存</el-button>
  </div>
</template>

<script setup lang="ts">
import PanelTitle from '@/components/PanelTitle'
import { reactive, ref, onBeforeMount } from 'vue'
import { getUserInfo, updateUserInfo } from '@/api/user/auth'
import { validate } from '@/utils/element/form'
import { successMessage } from '@/utils/element/message'

import type { FormInstance, FormRules } from 'element-plus'
import { EditUserInfo } from '@/models/user/auth'

const form = reactive<EditUserInfo>({
  desc: ''
})

const formRef = ref<FormInstance>()
const rules = reactive<FormRules>({
  desc: [{ max: 128, message: '最大字数不超过128个', trigger: ['blur', 'change'] }]
})
const username = ref<string>('')

onBeforeMount(async () => {
  const data = await getUserInfo()
  form.desc = data.desc
  username.value = data.username
})

const saveUserInfo = async () => {
  const valid = await validate(formRef)
  if (!valid) return

  try {
    await updateUserInfo(form)
    successMessage('个人信息更新成功')
  } catch (e) {}
}
</script>

<style scoped></style>
