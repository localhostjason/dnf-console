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

            <el-form-item label="QQ:" prop="qq">
              <el-input v-model="form.qq"></el-input>
            </el-form-item>
          </el-form>
        </el-col>
      </el-row>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialog.visible = false">取 消</el-button>
          <el-button type="primary" @click="updateAccount">确 定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { nextTick, reactive, ref } from 'vue'
import { FormInstance, FormRules } from 'element-plus'
import { EditAccountForm } from '@/views/gm/account/model/info'
import { AccountDetail } from '@/views/gm/account/model/model'
import { validate } from '@/utils/element/form'
import { updateAccountInfo } from '@/api/gm/accounts'
import { successMessage } from '@/utils/element/message'
import { setFormData } from '@/utils'

const dialog = reactive({
  visible: false,
  title: ''
})
const formRef = ref<FormInstance>()
const form = reactive<EditAccountForm>({
  qq: ''
})
const rules = reactive<FormRules>({})

let accountInfo = {
  uid: 0,
  account_name: ''
}
const showAccountDialog = (row: AccountDetail) => {
  dialog.visible = true

  accountInfo.uid = row.uid
  accountInfo.account_name = row.account_name

  dialog.title = '修改账号基本信息'

  nextTick(() => {
    setFormData(form, row)

    formRef.value.clearValidate()
  })
}

const updateAccount = async () => {
  const valid = await validate(formRef)
  if (!valid) return
  try {
    await updateAccountInfo(accountInfo.uid, form)
    successMessage('更新成功')
    dialog.visible = false
    emit('reloadAccount')
  } catch (e) {}
}

// hook
defineExpose({
  showAccountDialog
})

const emit = defineEmits(['reloadAccount'])
</script>

<style scoped></style>
