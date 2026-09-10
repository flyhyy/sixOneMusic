/* prettier-ignore */
/**
* @Author: hyy
* @Date: 2026-06-11 16:38:16
* @Description: 播放按钮
*/


<template>
    <button class=" inline-flex gap-1.5 cursor-pointer  rounded-[20px]   items-center transition duration-200   h-7 "
        :class="[colorMode, buttonSize]">

        <Icon :name="iconName" class="shrink-0"></Icon>

        <span class=" mt-px">
            <slot>
                {{ text }}
            </slot>
        </span>
    </button>
</template>

<script lang="ts" setup>
import { Icons, type IconName } from "@/config/icons";
import { computed } from "vue";

interface PlayButtonProp {
    iconName?: IconName;
    text?: string;
    theme?: ThemeType,
    size?: SizeType
}



const { iconName = Icons.play, text = "播放全部", theme = 'dark', size = 'default' } = defineProps<PlayButtonProp>();

const colorMode = computed(() => {

    if (theme === 'dark') {
        return 'bg-brand text-white hover:bg-[#f07a8f]';
    }
    return 'bg-transparent border border-line text-sub hover:border hover:border-brand hover:text-brand'

})
const _size = {
    'large': '',
    'default': 'h-9 text-sm  font-semibold px-5 ',
    'small': 'py-0 px-2.5 text-[11px]  leading-none',
}

const buttonSize = computed(() => {
    return _size[size]
})
</script>
