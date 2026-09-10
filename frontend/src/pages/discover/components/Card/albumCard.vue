/* prettier-ignore */ /** * @Author: hyy * @Date: 2026-06-16 14:10:30 *
@Description: 专辑收藏 */

<template>
    <div
        class=" group cursor-pointer overflow-hidden rounded-xl bg-surface transition duration-200 hover:hover:-translate-y-0.5 hover:shadow-[0_12px_48px_rgba(0,0,0,0.12)]">
        <div class="relative flex aspect-square w-full items-center justify-center text-[52px]"
            :style="backgroundStyle">

            <DiscoverPlayButton></DiscoverPlayButton>

            <Icon :name="Icons.disc" class="text-[#ffffff4d]"></Icon>
        </div>

        <div class="px-4 pt-3.5 pb-4">
            <div class="overflow-hidden text-[15px] font-bold text-ellipsis whitespace-normal text-ink">
                {{ albumName }}
            </div>
            <div class="mt-0.75 overflow-hidden text-[13px] text-ellipsis whitespace-normal text-muted">
                {{ singer }}
            </div>
            <div class="mt-2 flex gap-3.5 text-[11px] text-muted">
                <span class="flex items-center gap-0.5">
                    <Icon :name="Icons.musicLine"></Icon> {{ totalNum }} 首
                </span>
                <span class="flex items-center gap-0.5">
                    <Icon :name="Icons.time"></Icon>
                    {{ totalTime }}
                </span>
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { Icons } from "@/config/icons";
import { StringToGradient } from "@/utils/color";
import { GetFirstChar } from "@/utils/str";
import { computed } from "vue";
import DiscoverPlayButton from '@/pages/discover/components/DiscoverPlayButton/index.vue'

export interface AlbumCardProps {
    // 专辑名称
    albumName: string;
    // 歌手
    singer: string;
    // 总数
    totalNum: string | number;
    // 总时长
    totalTime: string;
}

const { albumName, singer, totalNum, totalTime } =
    defineProps<AlbumCardProps>();

const backgroundStyle = computed(() => {
    return `background:${StringToGradient(GetFirstChar(albumName))}`;
});
</script>
