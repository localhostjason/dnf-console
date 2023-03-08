<template>
  <div class="container">
    <el-row class="login-content">
      <el-col :span="11">
        <div class="login-logo">
          <img src="@/assets/logo.png" style="width: 50px; height: 50px" alt="logo" class="tcxa-logo" />
        </div>
        <div>
          <img src="@/assets/logo2.png" alt="login" style='height: 272px; padding: 0 10px'/>
        </div>
      </el-col>
      <el-col :span="13">
        <div class="login-container">
          <el-form
            ref="loginForm"
            :model="form"
            :rules="rules"
            class="login-form"
            auto-complete="on"
            label-position="left"
            size="large"
          >
            <div class="title-container">
              <h3 class="title">
                <span>VUE3 后台管理系统</span>
              </h3>
            </div>

            <el-form-item prop="username">
              <el-input
                v-model="form.username"
                :autofocus="isAutoFocus"
                placeholder="请输入账号"
                name="username"
                type="text"
                auto-complete="on"
                :prefix-icon="User"
                size="large"
              ></el-input>
            </el-form-item>

            <el-form-item prop="password">
              <el-input
                :type="passwordType"
                v-model="form.password"
                placeholder="请输入密码"
                name="password"
                auto-complete="on"
                @keyup.enter.native="handleLogin(loginForm)"
                size="large"
                :prefix-icon="Unlock"
              ></el-input>
              <span class="show-pwd" @click="showPwd">
              <svg-icon :icon-class="passwordType === 'password' ? 'eye' : 'open-eye'" />
            </span>
            </el-form-item>

            <el-button
              :loading="loading"
              type="primary"
              style="width: 100%; margin-bottom: 30px;margin-top: 20px"
              @click.native.prevent="handleLogin(loginForm)"
              size="large"
              >登录
            </el-button>
          </el-form>
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref } from 'vue'
import { FormInstance, FormRules } from 'element-plus'
import { useRouter } from 'vue-router'
import { getUserInfo, login } from '@/api/user/auth'
import { ElNotification } from 'element-plus'
import { useUserStore } from '@/store/modules/user'
import { Avatar ,Unlock, User } from '@element-plus/icons-vue'
import { Login } from '@/models/user/auth'

export default defineComponent({
  name: 'Login',
  setup() {
    const userStore = useUserStore()
    const router = useRouter()

    const form = reactive<Login>({
      username: '',
      password: ''
    })

    const rules = reactive<FormRules>({
      username: [{ required: true, trigger: 'blur', message: '请输入账号' }],
      password: [{ required: true, trigger: 'blur', message: '请输入密码' }]
    })
    const loginForm = ref<FormInstance>()

    const passwordType = ref<string>('password')
    const loading = ref<boolean>(false)
    const isAutoFocus = ref<boolean>(true)

    const handleLogin = async (formEl: FormInstance | undefined) => {
      if (!formEl) return
      await formEl.validate(async (valid, fields) => {
        if (!valid) return

        loading.value = true

        try {
          const { token } = await login(form)
          userStore.setToken(token)
        } catch (e) {}
        loading.value = false

        await router.push({ path: '/' })
        ElNotification.success({
          title: '登录成功',
          message: '欢迎您进入管理中心'
        })
      })
    }

    const showPwd = () => {
      if (passwordType.value === "password") {
        passwordType.value = "";
      } else {
        passwordType.value = "password";
      }
    }

    return {
      form,
      rules,
      loginForm,
      passwordType,
      loading,
      isAutoFocus,
      handleLogin,
      Avatar,
      Unlock,
      User,
      showPwd,
    }
  }
})
</script>

<style rel="stylesheet/scss" lang="scss" scoped>
@import './login.scss';
$dark_gray: #889aa4;

.container {
  background-color: #f2f3f5;
  height: 100%;
  width: 100%;
}


.svg-container {
  padding: 6px 5px 6px 15px;
  color: $dark_gray;
  vertical-align: middle;
  width: 30px;
  display: inline-block;

  .svg-icon {
    font-size: 20px;
  }
}
</style>
