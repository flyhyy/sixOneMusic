/* prettier-ignore */
/**
* @Author: hyy
* @Date: 2026-06-17 14:50:57
* @Description: 歌曲卡片
*/


<template>
    <div class="flex items-center gap-3.5 py-2.75 px-3.5 rounded-[10px] cursor-pointer  relative hover:bg-brand/12 hover:text-brand transition-colors duration-200 ease-in-out "
        :class="[activeClass, index % 2 === 0 ? '' : 'bg-surface']">
        <span class="inline-block w-8 text-center text-muted shrink-0 font-medium tabular-nums">
            {{ index }}
        </span>
        <!-- 封面 -->
        <div class="w-11.5 h-11.5 rounded-lg ">
            <Cover :name="coverName" :src="coverSrc"></Cover>
        </div>
        <!-- 歌手，歌曲名称 -->
        <div class="flex-1">

            <SongInfo :class="acitve ? 'text-brand!' : ''" :primary-text="songName" :secondary-text="singer">
            </SongInfo>
        </div>
        <!-- 专辑 -->
        <div class="w-32.5 text-sub whitespace-normal overflow-hidden text-ellipsis shrink-0 text-center">
            {{ album }}
        </div>
        <div class="w-12 text-right text-muted shrink-0 tabular-nums">
            {{ formatTime(time) }}

        </div>

        <div class="flex gap-1.5 ">

            <ButtonIcon :icon-name="isPlay ? Icons.pause : Icons.play" button-title="播放" @click="$emit('play', isPlay)">
            </ButtonIcon>
            <ButtonIcon :icon-name="Icons.heart" :class="isCollect ? 'text-brand' : ''"
                :button-title="isCollect ? '取消收藏' : '添加收藏'" @click="$emit('collect')"></ButtonIcon>
            <ButtonIcon :icon-name="Icons.add" button-title="加入播放队列" @click="$emit('addPlayQueue')"></ButtonIcon>
            <ButtonIcon :icon-name="Icons.folder" button-title="添加到歌单"
                @click="$emit('addPlayList', $event.currentTarget)"></ButtonIcon>



        </div>
    </div>

</template>

<script lang="ts" setup>
import type { SongInfoProps } from '@/components/songInfo/index.vue'
import { Icons } from '@/config/icons'
import { formatTime } from '@/utils/time'
import { computed } from 'vue'

export type SongCardProps = Omit<SongInfoProps, 'primaryText' | 'secondaryText'> & {
    coverSrc: string
    coverName?: string
    songName: string
    singer: string
    album: string
    time: number
    isCollect: boolean
    acitve?: boolean
    index: number
    isPlay: boolean
}
const { coverSrc, coverName = "", singer, songName, album, acitve = false, isPlay } = defineProps<SongCardProps>()

const activeClass = computed(() => {

    if (acitve) {
        return 'border-l-brand  border-l-[3px] !bg-brand/12  '
    } else return ''

})


</script>