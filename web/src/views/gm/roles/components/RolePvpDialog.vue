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
            <el-form-item label="PK段位" prop="pvp_grade">
              <el-input-number :min="0" :max="34" v-model.number="form.pvp_grade"></el-input-number>
            </el-form-item>

            <el-form-item label="胜利场次" prop="win">
              <el-input-number :min="0" v-model.number="form.win"></el-input-number>
            </el-form-item>

            <el-form-item label="胜点" prop="pvp_point">
              <el-input-number :min="0" :max="99" v-model.number="form.pvp_point"></el-input-number>
            </el-form-item>
          </el-form>
        </el-col>
      </el-row>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialog.visible = false">取 消</el-button>
          <el-button type="primary" @click="updateRolePvp">确 定</el-button>
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
import { updatePvpByRole, updateQPbyRole } from '@/api/gm/role'
import { successMessage } from '@/utils/element/message'
import { PvpForm } from '@/views/gm/roles/model/pvp'
import { setFormData } from '@/utils'

const dialog = reactive({
  visible: false,
  title: ''
})
const formRef = ref<FormInstance>()
let form = reactive<PvpForm>({
  win: 0,
  pvp_grade: 0,
  pvp_point: 0
})

const rules = reactive<FormRules>({
  pvp_grade: [{ type: 'integer', min: 0, max: 34, message: 'pk段位0至34', trigger: 'blur' }]
})

let characNo: number = null

// method
const showPvpDialog = (row: Role) => {
  dialog.visible = true
  dialog.title = '修改PK段位'

  characNo = row.charac_no

  nextTick(() => {
    setFormData(form, row.pvp)
    formRef.value.clearValidate()
  })
}

const updateRolePvp = async () => {
  const valid = await validate(formRef)
  if (!valid) return

  try {
    await updatePvpByRole(characNo, form)
    successMessage('更新成功')
    dialog.visible = false
    emit('reloadRoles')
  } catch (e) {}
}

// hook

defineExpose({
  showPvpDialog
})
const emit = defineEmits(['reloadRoles'])
</script>

<style scoped></style>
