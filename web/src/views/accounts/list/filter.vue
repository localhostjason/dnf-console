<template>
  <div>
    <el-form :inline="true" :model="form" ref="formRef" label-width="50px" @submit.native.prevent>
      <div>
        <el-form-item label="uid:" prop="uid">
          <el-input v-model="form.uid" clearable></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="filterAccount">查询</el-button>
          <el-button @click.prevent="reset">重置</el-button>
        </el-form-item>
      </div>
    </el-form>
  </div>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue'
import { FormInstance } from 'element-plus'
import { FilterAccountForm } from '@/views/accounts/list/model/model'

const formRef = ref<FormInstance>()
const form = reactive<FilterAccountForm>({
  uid: ''
})

// method
const filterAccount = () => {
  emit('setParams', form)
}

const reset = () => {
  formRef.value.resetFields()
  emit('setParams', {})
}

// hook
const emit = defineEmits(['setParams'])
</script>

<style scoped></style>
