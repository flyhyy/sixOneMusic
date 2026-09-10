/* prettier-ignore */ /** * @Author: hyy * @Date: 2026-06-16 15:18:37 *
@Description: 适用于发现音乐的播放按钮 */

<template>
    <button
        class="  opacity-0 absolute  bottom-1.5 right-1.5  cursor-pointer flex h-8 w-8 items-center justify-center rounded-[50%] bg-brand text-xl text-white shadow-[0_2px_8px_rgba(232,99,123,.5)] transition duration-200 group-hover:translate-y-0 group-hover:opacity-100">
        <Icon :name="Icons.play"></Icon>
    </button>
</template>
<script lang="ts" setup>
import { Icons } from "@/config/icons";
</script>
