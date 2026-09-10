/* prettier-ignore */
/**
* @Author: hyy
* @Date: 2026-06-15 16:31:48
* @Description: 标题信息
*/
<template>
    <div class=" flex justify-between items-center"
        :class="showLine ? 'after:content-[\'\'] after:flex-1 after:h-px after:bg-line gap-2' : ''">
        <!-- 标题展示 -->
        <div>
            <TitleInfoItem :icon-name="iconName" :size="size">
                {{ title }}
            </TitleInfoItem>
        </div>
        <!-- 对应操作 -->
        <div v-if="$slots.default">
            <slot> </slot>
        </div>

    </div>

</template>

<script lang="ts" setup>
import type { IconName } from '@/config/icons.ts';
import TitleInfoItem from './itemIndex.vue'
import type { SizeType } from './itemIndex.vue'
interface TitleInfoProps {
    iconName: IconName
    title: string,
    size?: SizeType
    showLine?: boolean

}

const { title, iconName, size, showLine = false } = defineProps<TitleInfoProps>()

</script>
