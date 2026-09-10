/* prettier-ignore */
/**
* @Author: hyy
* @Date: 2026-06-15 18:14:52 *
@Description: 歌曲卡片
*/

<template>
  <div
    class="group relative max-w-37.5 min-w-37.5 shrink-0  snap-start rounded-xl bg-surface p-3 transition duration-200 hover:bg-hover hover:-translate-y-0.5 hover:shadow-[0_4px_24px_rgba(0,0,0,0.08)]">
    <div class="relative flex aspect-square items-center justify-center overflow-hidden rounded-lg text-[28px]" :style="{
      backgroundImage: coverSrc ? '' : StringToGradient(songName),
    }">
      <DiscoverPlayButton></DiscoverPlayButton>

      {{ coverSrc }}
      <!-- 有封面图 -->
      <img v-if="coverSrc" :src="coverSrc" alt="" />

      <!-- 没有封面图 -->
      <Icon v-else :name="Icons.music" class="text-[#ffffff59]"></Icon>
    </div>

    <!-- 歌曲名称 -->
    <div class="mt-2.5">
      <span class="overflow-hidden font-semibold text-ellipsis whitespace-nowrap text-ink">
        {{ songName }}
      </span>
    </div>
    <!-- 歌手名称 -->
    <div class="mt-0.5">
      <span class="overflow-hidden text-[11px] text-ellipsis whitespace-nowrap text-muted">
        {{ singer }}
      </span>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { Icons } from "@/config/icons";
import { StringToGradient } from "@/utils/color";
import { computed } from "vue";
import DiscoverPlayButton from '@/pages/discover/components/DiscoverPlayButton/index.vue'

export interface CardProps {
  // 封面
  cover: string;
  // 歌手
  singer: string;
  // 歌曲名称
  songName: string;
}
const { cover, singer, songName } = defineProps<CardProps>();

const coverSrc = computed(() => {
  if (cover != '') return cover
  return undefined
})


</script>
