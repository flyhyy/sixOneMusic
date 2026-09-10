/* * @Author: hyy * @Date: 2026-06-04 15:58:43 * @Description: 总体布局 */

<template>
  <section class="flex h-screen flex-col ">
    <header-container> </header-container>

    <main class="flex flex-1">
      <side-container> </side-container>
      <main-container>
        <router-view></router-view>
      </main-container>
    </main>
  </section>
</template>

<script setup lang="ts">
import MainContainer from "./main.vue";
import SideContainer from "./side.vue";
import HeaderContainer from "./header.vue";
</script>
