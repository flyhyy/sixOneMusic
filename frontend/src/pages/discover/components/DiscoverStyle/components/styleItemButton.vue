/* prettier-ignore */
/**
* @Author: hyy
* @Date: 2026-06-12 17:07:33
* @Description: 音乐风格组件
*/


<template>
    <button
        class="flex cursor-pointer items-center justify-center gap-1 rounded-3xl border border-line  px-3.5 py-1.5 text-sm whitespace-nowrap text-sub transition duration-200"
        :class="active ? activeClass : baseClass">
        <Icon :name="iconName"></Icon>
        <slot></slot>
    </button>
</template>

<script setup lang="ts">
import type { IconName } from '@/config/icons';

interface StyleItemButtonProp {
    iconName: IconName;
    active?: boolean
}

const baseClass = 'hover:border-brand hover:text-brand bg-transparent'
const activeClass = 'bg-brand border-brand text-white'



const { iconName, active = false } = defineProps<StyleItemButtonProp>();
</script>
