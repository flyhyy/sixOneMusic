/* prettier-ignore */ /** * @Author: hyy * @Date: 2026-06-17 15:48:43 *
@Description: 带有button的icon */

<template>
    <button :title="buttonTitle" :class="fullClass"
        class="flex h-8 w-8 cursor-pointer items-center justify-center rounded-[50%]  text-base transition duration-200 shrink-0 ">
        <Icon :name="iconName"></Icon>
    </button>
</template>

<script lang="ts" setup>
import type { IconName } from "@/config/icons";
import { computed } from "vue";

interface ButtonIcon {
    iconName: IconName;
    buttonTitle?: string;
    isFull?: boolean
}

const { iconName, isFull = false, buttonTitle = "" } = defineProps<ButtonIcon>();


const fullClass = computed(() => {

    if (isFull) {

        return "bg-brand text-white"
    }
    return "border-none bg-transparent hover:bg-brand/6 hover:text-brand text-sub"

})

</script>
