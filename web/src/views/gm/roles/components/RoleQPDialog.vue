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
            <el-form-item label="现有QP">
              <span class="text-gray-lg sm">{{ qp }}</span>
            </el-form-item>

            <el-form-item label="修改QP" prop="qp">
              <el-input-number :min="0" v-model.number="form.qp"></el-input-number>
            </el-form-item>
          </el-form>
        </el-col>
      </el-row>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialog.visible = false">取 消</el-button>
          <el-button type="primary" @click="updateRoleQp">确 定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { nextTick, reactive, ref } from 'vue'
import { FormInstance, FormRules } from 'element-plus'
import { QpForm } from '@/views/gm/roles/model/qp'
import { Role } from '@/views/gm/roles/model'
import { validate } from '@/utils/element/form'
import { updateQPbyRole } from '@/api/gm/role'
import { successMessage } from '@/utils/element/message'

const dialog = reactive({
  visible: false,
  title: ''
})
const formRef = ref<FormInstance>()
const form = reactive<QpForm>({
  qp: 500
})

const rules = reactive<FormRules>({
  qp: [{ type: 'integer', min: 0, message: 'QP值至少未0', trigger: 'blur' }]
})

const qp = ref<number>(null)
let characNo: number = null

// method
const showQpDialog = (row: Role) => {
  dialog.visible = true
  dialog.title = '修改QP'

  qp.value = row.QP
  characNo = row.charac_no

  nextTick(() => {
    formRef.value.resetFields()
  })
}

const updateRoleQp = async () => {
  const valid = await validate(formRef)
  if (!valid) return

  try {
    await updateQPbyRole(characNo, form)
    successMessage('更新成功')
    dialog.visible = false
    emit('reloadRoles')
  } catch (e) {}
}

// hook

defineExpose({
  showQpDialog
})
const emit = defineEmits(['reloadRoles'])
</script>

<style scoped></style>
