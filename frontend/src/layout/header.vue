/**
* @Author: hyy
* @Date: 2026-06-04 15:55:50
* @Description: 主页头部组件
*/

<template>
  <header class="py-1.5 h-14 border-b border-line flex items-center static bg-surface">
    <!-- logo -->
    <div class="text-brand flex  h-full  items-center text-2xl font-extrabold tracking-tight ">
      <Icon :name="iconMusicNote" class="mr-0.5"></Icon>
      SixOne
      <span class="text-ink font-normal">
        Music
      </span>
    </div>
    <HeaderSearch class="flex-1 ml-5"></HeaderSearch>
    <HeaderOptions></HeaderOptions>
  </header>
</template>

<script lang="ts" setup>
import HeaderSearch from './components/headerSearch/index.vue'
import HeaderOptions from './components/headerActions/index.vue'
import { Icons } from '@/config/icons.ts';

const iconMusicNote = Icons.headerMusicNote


</script>
