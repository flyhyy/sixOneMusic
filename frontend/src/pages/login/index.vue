<template>
  <div class="bg-linear-to-br from-[#1a1a2e] via-[#16213e] to-[#0f3460]  h-screen flex items-center ">
    <section
      class="w-[80%] box-border  border border-[rgba(255,255,255,0.1)] sm:w-125 m-auto p-6  bg-[rgba(255,255,255,0.07)] rounded-xl shadow-[0_20px_60px_rgba(0,0,0,0.4)]   ">
      <div class="flex justify-center items-center text-2xl font-extrabold tracking-tight">
        <span class="text-brand flex items-center">
          <Icon :name="iconMusicNote" class="mr-0.5"></Icon>
          SixOne
        </span>
        <span class="text-white font-normal">
          Music
        </span>
      </div>

      <!-- 登录 / 注册切换 -->
      <div class="flex justify-center mt-6 mb-1">
        <button
          class="px-8 py-1.5 text-sm font-bold rounded-l-full transition-colors"
          :class="mode === 'login' ? 'bg-[#e94560] text-white' : 'bg-white/5 text-white/60 hover:text-white'"
          @click="switchMode('login')">
          登录
        </button>
        <button
          class="px-8 py-1.5 text-sm font-bold rounded-r-full transition-colors"
          :class="mode === 'register' ? 'bg-[#e94560] text-white' : 'bg-white/5 text-white/60 hover:text-white'"
          @click="switchMode('register')">
          注册
        </button>
      </div>

      <div class="mt-3">
        <!-- 登录 -->
        <template v-if="mode === 'login'">
          <LoginInput v-model.trim="loginFormData.username" label="账号" :alert-msg="validateFormData.username"
            placeholder="账号" />
          <LoginInput v-model.trim="loginFormData.password" type="password" label="密码"
            :alert-msg="validateFormData.password" placeholder="密码" />

          <div class="flex items-center">
            <input type="checkbox" class="accent-[#e94560]" v-model="remberMe">
            <span class="text-[14px] text-[rgba(255,255,255,0.58)] ml-2">
              记住我
            </span>
          </div>
          <button
            class="active:shadow-[0_12px_32px_rgba(233,69,96,0.48)]  hover:-translate-y-px hover:shadow-[0_12px_32px_rgba(233,69,96,0.48)]  w-full mt-2  p-2 bg-linear-to-br from-[#e94560] to-[#c23152] text-white rounded-lg font-extrabold  shadow-[0_10px_28px_rgba(233,69,96,0.35)]"
            @click="onLogin">
            登录
          </button>
        </template>

        <!-- 注册 -->
        <template v-else>
          <LoginInput v-model.trim="registerFormData.username" label="账号" :alert-msg="validateFormData.username"
            placeholder="请输入账号" />
          <LoginInput v-model.trim="registerFormData.password" type="password" label="密码"
            :alert-msg="validateFormData.password" placeholder="请输入密码" />
          <LoginInput v-model.trim="registerFormData.confirmPassword" type="password" label="确认密码"
            :alert-msg="validateFormData.confirmPassword" placeholder="请再次输入密码" />

          <button
            class="active:shadow-[0_12px_32px_rgba(233,69,96,0.48)]  hover:-translate-y-px hover:shadow-[0_12px_32px_rgba(233,69,96,0.48)]  w-full mt-2  p-2 bg-linear-to-br from-[#e94560] to-[#c23152] text-white rounded-lg font-extrabold  shadow-[0_10px_28px_rgba(233,69,96,0.35)]"
            @click="onRegister">
            注册
          </button>
        </template>
      </div>
    </section>

  </div>

</template>

<script setup lang="ts">
import { ref } from 'vue';
import LoginInput from './components/input.vue'

import type { LoginFormData, RegisterFormData } from './types.ts'
import LoginApi from '@/api/login/index.ts';
import Schema, { type Rules, type ValidateError } from 'async-validator';
import { useToast } from 'vue-toastification';
import { useRouter } from 'vue-router';
import { setToken } from '@/storage/token.ts';
import { Icons } from '@/config/icons.ts';

const toast = useToast()

const router = useRouter()

const iconMusicNote = Icons.headerMusicNote

type Mode = 'login' | 'register'

const mode = ref<Mode>('login')

// 登录校验规则
const loginRules: Rules = {
  username: [{
    required: true,
    message: "请输入账号",
  }],
  password: [{
    required: true,
    message: "请输入密码"
  }]
}

// 注册校验规则
const registerRules: Rules = {
  username: [
    { required: true, message: "请输入账号" },
    { min: 3, max: 20, message: "账号长度需在 3-20 个字符之间" }
  ],
  password: [
    { required: true, message: "请输入密码" },
    { min: 6, max: 20, message: "密码长度需在 6-20 个字符之间" }
  ],
  confirmPassword: [
    { required: true, message: "请再次输入密码" },
    {
      validator: (_rule: unknown, value: string, callback: (error?: Error) => void) => {
        if (value !== registerFormData.value.password) {
          callback(new Error("两次输入的密码不一致"))
        } else {
          callback()
        }
      }
    }
  ]
}


const loginFormData = ref<LoginFormData>({
  username: 'admin',
  password: 'admin123'
})

const registerFormData = ref<RegisterFormData>({
  username: '',
  password: '',
  confirmPassword: ''
})

const validateFormData = ref({
  username: "",
  password: "",
  confirmPassword: ""
})

const remberMe = ref<boolean>(false)

// 切换登录 / 注册，清空校验提示
const switchMode = (target: Mode) => {
  mode.value = target
  validateFormData.value = {
    username: "",
    password: "",
    confirmPassword: ""
  }
}

const onLogin = async () => {
  const validator = new Schema(loginRules)

  try {

    await validator.validate(loginFormData.value)
    const res = await LoginApi.login(loginFormData.value)


    if (res.token != "") {

      setToken(res.token)

      toast.success("登录成功", {
        onClose: () => {

          router.push({
            path: "/music"
          })
        }
      })




    }

  } catch (err: any) {

    if (typeof err === "string") {
      toast.error(err)
    } else {
      const errors: ValidateError[] = err.errors
      if (errors.length) {
        let { field, message } = errors[0]
        if (field) {
          const key = field as keyof typeof validateFormData.value
          validateFormData.value[key] = message || "校验失败"
        }
      }
    }


  }
}

const onRegister = async () => {
  const validator = new Schema(registerRules)

  try {
    await validator.validate(registerFormData.value)

    await LoginApi.register({
      username: registerFormData.value.username,
      password: registerFormData.value.password
    })

    toast.success("注册成功，请登录", {
      onClose: () => {
        // 切换到登录并回填账号
        loginFormData.value.username = registerFormData.value.username
        loginFormData.value.password = ''
        registerFormData.value = {
          username: '',
          password: '',
          confirmPassword: ''
        }
        switchMode('login')
      }
    })

  } catch (err: any) {

    if (typeof err === "string") {
      toast.error(err)
    } else {
      const errors: ValidateError[] = err.errors
      if (errors.length) {
        let { field, message } = errors[0]
        if (field) {
          const key = field as keyof typeof validateFormData.value
          validateFormData.value[key] = message || "校验失败"
        }
      }
    }

  }
}




</script>
