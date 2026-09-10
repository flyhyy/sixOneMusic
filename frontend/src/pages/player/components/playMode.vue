/**
* @Author: 你的名字
* @Date: 2026-06-02 15:43:54
* @Description: 播放模式控制
*/

<template>
    <button 
    @click="onClick"
    class=" border-none text-sm rounded-lg  py-1 px-3"
        :class="active ? 'text-[#e94560] bg-[#e945601f]' : 'bg-transparent text-white/35 '">
        {{ label }}
    </button>

</template>

<script setup lang="ts">
interface prop {
    label: string
    active: boolean
}

const { label = 'label', active = true } = defineProps<prop>()

const emits = defineEmits(['click'])
const onClick = ()=>{
        emits('click')
}
</script>
