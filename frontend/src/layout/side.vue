/** * @Author: hyy * @Date: 2026-06-04 15:55:17 * @Description: 侧边栏 */

<template>
  <aside class="flex h-full w-64 flex-col gap-1 border-r border-line bg-surface px-3 py-4">
    <SideOptions></SideOptions>

    <div class="mx-3 my-2 h-px bg-line"></div>

    <SideMyPlayList></SideMyPlayList>
  </aside>
</template>
<script setup lang="ts">
import SideOptions from "./components/sideOptions/index.vue";
import SideMyPlayList from "./components/sideMyPlayList/index.vue";
</script>
