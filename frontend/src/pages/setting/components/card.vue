/* prettier-ignore */
/**
* @Author: hyy
* @Date: 2026-06-18 15:02:02
* @Description: 设置页面卡片
*/


<template>

    <div class="bg-surface rounded-xl p-6  mb-6">

        <h3 class="text-base mb-4 flex items-center gap-2 text-ink">
            <Icon :name="iconName"></Icon>
            <span>{{ title }}</span>
        </h3>
        <slot></slot>

    </div>

</template>

<script lang="ts" setup>
import type { IconName } from '@/config/icons';


interface SettingCardProps {
    iconName: IconName
    title: string
}

const { iconName, title } = defineProps<SettingCardProps>()
</script>