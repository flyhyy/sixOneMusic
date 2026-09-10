/**
* @Author: hyy
* @Date: 2026-06-09 11:36:00
* @Description: 头部搜索栏
*/

<template>

    <div class="relative max-w-2xs  overflow-hidden  ">
        <input v-model="model" placeholder="搜索歌曲、歌手、专辑..."
            class=" focus:border-brand bg-input h-9 w-full  pt-0 pr-10  pb-0 pl-4 border border-line outline-none rounded-3xl placeholder:text-sm placeholder:text-muted" />
        <Icon :name="iconSearch" class="absolute   text-2xl text-muted top-[50%] translate-y-[-50%] right-2"></Icon>
    </div>

</template>

<script lang="ts" setup>
import { Icons } from '@/config/icons';


const model = defineModel()

const iconSearch = Icons.headerSearch

</script>