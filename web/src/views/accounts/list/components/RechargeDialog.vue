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
        <div class="div_title">
          <span>基本信息</span>
        </div>
      </el-row>

      <div style="margin-bottom: 30px">
        <el-descriptions class="margin-top" :column="3" border>
          <el-descriptions-item label="账号ID">
            <span>{{ data.uid }}</span>
          </el-descriptions-item>
          <el-descriptions-item label="D币">
            <span>{{ data.cera }}</span>
          </el-descriptions-item>
          <el-descriptions-item label="D点">
            <span>{{ data.cera_point }}</span>
          </el-descriptions-item>
        </el-descriptions>
      </div>

      <el-row>
        <div class="div_title">
          <span>充值项目</span>
        </div>
      </el-row>

      <el-row>
        <el-col :span="22">
          <el-form :model="form" :rules="rules" ref="formRef" label-width="100px">
            <el-form-item label="充值类型" prop="cera_option">
              <el-select v-model="form.cera_option">
                <el-option v-for="item in options" :key="item.value" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>

            <el-form-item label="D币" prop="cera" v-show="form.cera_option === 'cera'">
              <el-input v-model.number="form.cera"></el-input>
            </el-form-item>

            <el-form-item label="D点" prop="cera_point" v-show="form.cera_option === 'cera_point'">
              <el-input v-model.number="form.cera_point"></el-input>
            </el-form-item>
          </el-form>
        </el-col>
      </el-row>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialog.visible = false">取 消</el-button>
          <el-button type="primary" @click="rechargeAccount">确 定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue'
import { FormInstance, FormRules } from 'element-plus'
import { AccountDetail, RechargeForm } from '@/views/accounts/list/model/model'
import { assignWith } from 'lodash'
import { rechargeByUid } from '@/api/gm/accounts'
import { successMessage } from '@/utils/element/message'

const dialog = reactive({
  visible: false,
  title: '充值'
})
const formRef = ref<FormInstance>()

const options = [
  { label: 'D币', value: 'cera' },
  { label: 'D点', value: 'cera_point' }
]

const form = reactive<RechargeForm>({
  cera_point: 500,
  cera: 100,
  cera_option: 'cera'
})

type Info = {
  uid: number
  cera_point: number
  cera: number
}

const data = reactive<Info>({
  uid: 0,
  cera_point: 0,
  cera: 0
})

const rules = reactive<FormRules>({
  cera: [{ type: 'integer', min: 0, message: 'D币不能少于0', trigger: 'blur' }],
  cera_point: [{ type: 'integer', min: 0, message: 'D点不能少于0', trigger: 'blur' }]
})

const rechargeAccount = async () => {
  const postData: RechargeForm = {
    cera_point: data.cera_point + form.cera_point,
    cera: data.cera + form.cera,
    cera_option: form.cera_option
  }
  try {
    await rechargeByUid(data.uid, postData)
    emit('reloadAccount')
    dialog.visible = false
    successMessage('充值成功')
  } catch (e) {}
}

const showRechargeDialog = (row: AccountDetail) => {
  data.uid = row.uid
  data.cera_point = row.cera_point
  data.cera = row.cera
  dialog.visible = true
}

// hook
defineExpose({
  showRechargeDialog
})

const emit = defineEmits(['reloadAccount'])
</script>

<style scoped></style>
