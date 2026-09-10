/* prettier-ignore */
/**
* @Author: hyy
* @Date: 2026-06-12 18:06:51
* @Description: 空数据提示
*/


<template>

    <div class="text-center text-muted" :class="sizeClass.warp">
        <Icon :name="iconName" class="m-auto  mb-4 " :class="sizeClass.icon"></Icon>

        <span class="text-base">
            <slot></slot>
        </span>
        <div class="mt-4" v-if="$slots.other">
            <slot name="other"></slot>
        </div>

    </div>

</template>

<script lang="ts" setup>
import type { IconName } from '@/config/icons';
import { computed } from 'vue';


interface EmptyProps {
    iconName: IconName;
    size?: SizeType
}

const { iconName, size = "default" } = defineProps<EmptyProps>()

const sizeClass = computed(() => {

    if (size === "default") {
        return {
            warp: "py-15 px-5",
            icon: "text-[56px] mb-4"
        }
    } else if (size === "large") {
        return {
            warp: "",
            icon: ""
        }
    } else {
        return {
            warp: "py-5",
            icon: "text-xl"

        }
    }


})


</script>