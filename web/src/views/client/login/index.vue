<template>
  <div>
    <panel-title title="客户端登录"></panel-title>

    <el-row style="margin-top: 50px" :gutter="120">
      <el-col :span="8">
        <el-form :model="form" :rules="rules" ref="formRef" label-width="120px">
          <el-form-item prop="loginType">
            <el-switch
              v-model="form.login_type"
              active-value="password"
              inactive-color="#36CBCB"
              inactive-value="uid"
              class="mb-2"
              active-text="账号登录"
              inactive-text="账号ID登录"
            />
          </el-form-item>

          <div v-show="form.login_type === 'password'">
            <el-form-item label="账号" prop="username">
              <el-input v-model="form.username"></el-input>
            </el-form-item>

            <el-form-item label="密码:" prop="password">
              <el-input type="password" v-model="form.password"></el-input>
            </el-form-item>
          </div>
          <div v-show="form.login_type === 'uid'">
            <select-account
              :has-role="false"
              label-width="120px"
              :has-btn="false"
              @setUid="setUid"
              enable-event-change
            ></select-account>
          </div>

          <!--          <el-form-item label="RSA加密公钥:" prop="public_pem">-->
          <!--            <el-input v-model="form.public_pem"></el-input>-->
          <!--          </el-form-item>-->

          <!--          <el-form-item label="游戏目录:" prop="game_dir">-->
          <!--            <el-input v-model="form.game_dir"></el-input>-->
          <!--          </el-form-item>-->

          <el-form-item>
            <el-button type="primary" @click="loginClientHandler">登 录</el-button>
          </el-form-item>
        </el-form>
      </el-col>
      <el-col :span="12">
        <fieldset>
          <h5 style="margin-bottom: 15px; font-size: 13px; font-weight: bold">RSA加密hash</h5>

          <div>
            <el-input type="textarea" v-model="tokenRef" :autosize="{ minRows: 8 }"></el-input>
          </div>
        </fieldset>
      </el-col>
    </el-row>

    <el-row style="margin-top: 30px">
      <el-col :span="24">
        <el-card class="box-card" shadow="hover">
          <template #header>
            <div class="card-header">
              <span>注意点</span>
            </div>
          </template>
          <div class="card-body">
            <p>
              1. 执行 "main.exe -pem" 会生成“publickey.pem”和 “private.pem” 文件，将 publickey.pem 复制到服务器 且 重启
            </p>
            <p>2. 点击登录，拿到hash。 在客户端目录下 运行 “dnf.exe xx” ，xx就是加密的hash</p>
            <p>3. 最终：还是推荐使用登录器 方便。</p>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup lang="ts">
import PanelTitle from '@/components/PanelTitle'
import { reactive, ref } from 'vue'
import { FormInstance, FormRules } from 'element-plus'
import { ClientLoginForm } from '@/views/client/login/models/login'
import SelectAccount from '@/views/gm/tasks/SelectAccount'
import { validate, validateField } from '@/utils/element/form'
import { loginClient } from '@/api/client/login'

const formRef = ref<FormInstance>()

const form = reactive<ClientLoginForm>({
  login_type: 'password',
  uid: 18000002,
  username: 'bailbo9',
  password: '123456',
  public_pem: '',
  game_dir: ''
})

const tokenRef = ref<any>('')

const rules = reactive<FormRules>({
  username: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
  password: [{ required: true, message: '请输入密码', trigger: 'blur' }],
  public_pem: [{ required: true, message: '请输入RSA加密公钥', trigger: 'blur' }],
  game_dir: [{ required: true, message: '请输入游戏目录', trigger: 'blur' }]
})

const setUid = id => {
  form.uid = id
}

const loginClientHandler = async () => {
  if (form.login_type === 'uid') {
    const valid = await validateField(formRef, ['public_pem', 'game_dir'])
    if (!valid) return
  } else {
    const valid = await validate(formRef)
    if (!valid) return
  }

  try {
    const { token } = await loginClient(form)
    tokenRef.value = token
  } catch (e) {}

  console.log(111, form)
}
</script>

<style scoped></style>
