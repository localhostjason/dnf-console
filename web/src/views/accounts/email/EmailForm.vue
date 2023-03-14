<template>
  <div>
    <el-row :gutter="20">
      <el-col :span="8">
        <el-form :model="form" :rules="rules" ref="formRef" label-width="100px">
          <el-form-item label="物品代码" prop="code">
            <el-input v-model.number="form.code"></el-input>
          </el-form-item>

          <el-form-item label="数量" prop="number">
            <el-input v-model.number="form.number"></el-input>
          </el-form-item>

          <el-form-item label="武器锻造等级" prop="seperate_upgrade">
            <el-input v-model.number="form.seperate_upgrade"></el-input>
          </el-form-item>

          <el-form-item label="强化/增幅等级" prop="upgrade">
            <el-input v-model.number="form.upgrade"></el-input>
          </el-form-item>

          <el-form-item label="具备异界属性">
            <el-checkbox v-model="form.is_amplify" />
          </el-form-item>

          <el-form-item label="红字" v-show="form.is_amplify">
            <el-select v-model="form.amplify_option">
              <el-option v-for="item in options" :key="item.value" :label="item.label" :value="item.value" />
            </el-select>
          </el-form-item>

          <el-form-item label="追加属性" prop="amplify_value" v-show="form.is_amplify">
            <el-input v-model.number="form.amplify_value"></el-input>
          </el-form-item>

          <el-form-item label="金币" prop="gold">
            <el-input v-model.number="form.gold"></el-input>
          </el-form-item>

          <el-form-item label="是否封装">
            <el-checkbox v-model="form.seal_flag" />
            <span class="text-warning sm">ss装备不要勾选！！！</span>
          </el-form-item>

          <el-form-item>
            <el-button type="primary" size="small" @click="sendEmail">发送</el-button>
          </el-form-item>
        </el-form>
      </el-col>
      <el-col :span="13" :offset="3">
        <el-card class="box-card">
          <template #header>
            <div class="card-header">
              <span>物品代码</span>
            </div>
          </template>
          <div>
            <el-row>
              <el-form :inline="true" :model="argQuery" ref="filterFormRef" label-width="100px">
                <div>
                  <el-form-item label="物品名称:" prop="name">
                    <el-input v-model="argQuery.name" clearable></el-input>
                  </el-form-item>
                  <el-form-item>
                    <el-button type="primary" @click="filterGold">查询</el-button>
                    <el-button @click.prevent="reset">重置</el-button>
                  </el-form-item>
                </div>
              </el-form>
            </el-row>
            <el-row>
              <el-table :data="goldState.data" v-loading="goldState.loading" border fit>
                <el-table-column prop="code" label="物品编号" width="150" />
                <el-table-column prop="name" label="物品名称" />
              </el-table>
            </el-row>
            <el-row>
              <el-pagination
                :current-page="pageQuery.page"
                :page-sizes="[10, 20, 30, 40, 50]"
                :page-size="pageQuery.per_page"
                :total="goldState.total"
                background
                layout="total, sizes, prev, pager, next, jumper"
                @size-change="handleSizeChange"
                @current-change="handleCurrentChange"
              ></el-pagination>
            </el-row>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue'
import { FormInstance, FormRules } from 'element-plus'
import { Email, FilterGold } from '@/views/accounts/email/model'
import { validate } from '@/utils/element/form'
import { sendEmailByRole, getGolds } from '@/api/gm/email'
import { errorMessage, successMessage } from '@/utils/element/message'
import { PageQuery } from '@/models/page'

const formRef = ref<FormInstance>()

const form = reactive<Email>({
  code: null,
  number: 0,
  seperate_upgrade: 0,
  upgrade: 0,
  is_amplify: false,
  amplify_option: 3,
  amplify_value: 0,
  gold: 0,
  seal_flag: false
})

const rules = reactive<FormRules>({
  code: [{ required: true, message: '请输入物品代码', trigger: 'blur' }],
  number: [{ required: true, message: '请输入数量', trigger: 'blur' }],
  upgrade: [{ type: 'integer', min: 0, max: 31, message: '强化/增幅等级在0至31', trigger: 'blur' }]
})

const characNo = ref<number>(null)

const options = [
  { label: '体力', value: 1 },
  { label: '精神', value: 2 },
  { label: '力量', value: 3 },
  { label: '智力', value: 4 }
]

const goldState = reactive<any>({
  data: [],
  total: 0,
  loading: false
})

const pageQuery = reactive<PageQuery>({
  page: 1,
  per_page: 10
})

let argQuery = reactive<FilterGold>({
  name: ''
})

const filterFormRef = ref<FormInstance>()

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
  goldState.loading = true
  const params = {
    ...argQuery,
    ...pageQuery
  }
  const { items, total } = await getGolds(params)
  goldState.data = items
  goldState.total = total
  goldState.loading = false
}

const handleSizeChange = (params_limit: number) => {
  pageQuery.per_page = params_limit
  getGoldList()
}

const handleCurrentChange = (params_page: number) => {
  pageQuery.page = params_page
  getGoldList()
}

const filterGold = () => {
  pageQuery.page = 1
  getGoldList()
}

const reset = () => {
  filterFormRef.value.resetFields()
  pageQuery.page = 1
  getGoldList()
}

// hook
getGoldList()
defineExpose({
  setCharacNo
})
</script>

<style scoped></style>
