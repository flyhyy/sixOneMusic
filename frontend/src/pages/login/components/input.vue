<template>
    <div>
        <span class=" text-[#ffffffb8]  text-lg font-bold text-[13px]">{{ label }} </span>
        <input :type="type" v-model="value"
            class=" focus:outline-none focus:ring-1 focus:border-transparent focus:ring-[#e94560b3]  mt-2 w-full border p-2 text-white   border-[#ffffff1a] rounded-sm  caret-white placeholder:text-[#ccc] placeholder:text-[15px]"
            :placeholder="placeholder" />
        <small class="text-[#ff9aa8] block h-6 leading-6">{{ alertMsg }}</small>
    </div>

</template>

<script setup lang="ts">
interface prop {
    label: string
    alertMsg: string
    placeholder: string
    type?: string
}

const { label = 'label', alertMsg = "提示信息", placeholder = "请输入", type = 'text' } = defineProps<prop>()

let value = defineModel<string>({
    default: ""
})



</script>