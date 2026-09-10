/* prettier-ignore */
/**
* @Author: hyy
* @Date: 2026-06-18 13:26:04
* @Description: 卡片
*/
<template>
    <div
        class="bg-surface rounded-xl p-5 text-center transition duration-200 cursor-default hover:-translate-y-0.5 hover:shadow-[0_4px_24px_rgba(0,0,0,0.08)]">
        <div class="text-[28px] font-bold text-brand">
            {{ value }}
        </div>
        <div class="text-[13px] text-muted mt-1 flex justify-center items-center gap-1">
            <Icon :name="iconName"></Icon>
            {{ label }}
        </div>
    </div>


</template>
<script setup lang="ts">
import type { IconName } from '@/config/icons';

interface PersonalCenterCardProps {
    iconName: IconName
    value: string | number
    label: string
}
const { iconName, value, label } = defineProps<PersonalCenterCardProps>()

</script>
