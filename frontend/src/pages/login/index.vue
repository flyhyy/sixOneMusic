<template>
  <div class="bg-linear-to-br from-[#1a1a2e] via-[#16213e] to-[#0f3460]  h-screen flex items-center ">
    <section
      class="w-[80%] box-border  border border-[rgba(255,255,255,0.1)] sm:w-125 m-auto p-6  bg-[rgba(255,255,255,0.07)] rounded-xl shadow-[0_20px_60px_rgba(0,0,0,0.4)]   ">
      <div
        class="w-20 h-20 rounded-xl   bg-linear-to-br from-[#e94560] to-[#f6b17a] flex  justify-center items-center m-auto">
        <svg width="34" height="34" viewBox="0 0 24 24" fill="#fff">
          <path d="M12 3v10.55A4 4 0 1 0 14 17V7h4V3h-6z" />
        </svg>
      </div>
      <h5 class="tracking-[2px]   text-white text-xl font-bold text-center mt-4">
        六一音乐
      </h5>


      <div class="mt-4">
        <LoginInput v-model.trim="formData.username" label="账号" :alert-msg="validateFormData.username"
          placeholder="账号" />
        <LoginInput v-model.trim="formData.password" label="密码" :alert-msg="validateFormData.password"
          placeholder="密码" />

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
      </div>
    </section>

  </div>

</template>

<script setup lang="ts">
import { ref } from 'vue';
import LoginInput from './components/input.vue'

import type { LoginFormData } from './types.ts'
import LoginApi from '@/api/login/index.ts';
import Schema, { type Rules, type ValidateError } from 'async-validator';
import { useToast } from 'vue-toastification';
import { useRouter } from 'vue-router';
import { setToken } from '@/storage/token.ts';

const toast = useToast()

const router = useRouter()

const rules: Rules = {
  username: [{
    required: true,
    message: "请输入账号",
  }],
  password: [{
    required: true,
    message: "请输入密码"
  }]
}


const formData = ref<LoginFormData>({
  username: 'admin',
  password: 'admin123'
})

const validateFormData = ref({
  username: "",
  password: ""
})

const remberMe = ref<boolean>(false)

const onLogin = async () => {
  const validator = new Schema(rules)

  try {

    await validator.validate(formData.value)
    const res = await LoginApi.login(formData.value)


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





</script>
