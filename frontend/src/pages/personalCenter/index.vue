/* prettier-ignore */
/**
* @Author: hyy
* @Date: 2026-06-18 13:14:26
* @Description: 个人中心
*/
<template>
    <TitleInfo :icon-name="Icons.sidePersonalCenter" title="个人中心" size="large">

    </TitleInfo>

    <div class="bg-surface  rounded-2xl p-8 flex justify-center items-center gap-6 mt-6">
        <div class="flex-1 ">
            <h2 class="text-[22px] mt-1 flex items-center gap-2 text-ink">admin</h2>
            <p class="text-muted text-sm mt-1 ">
                192.169.1.1
            </p>

        </div>
        <div class="flex gap-2.5 ">

            <PlayButton theme="light" :icon-name="Icons.logout" size="small">
                退出
            </PlayButton>
        </div>


    </div>

    <div class="grid grid-cols-[repeat(auto-fill,minmax(150px,1fr))] gap-4  mt-6">
        <Card :icon-name="Icons.heartLine" value="1" label="收藏歌曲"></Card>
        <Card :icon-name="Icons.stackLine" value="1" label="曲库总数"></Card>
        <Card :icon-name="Icons.history" value="1" label="播放历史"></Card>
        <Card :icon-name="Icons.time" value="1" label="总播放时长"></Card>

    </div>

    <div class="mt-6">
        <TitleInfo :icon-name="Icons.history" title="最近播放" size="default" show-line>

        </TitleInfo>
    </div>
    <div>
        <!-- 
        <SongCard v-for="(item, index) in mockData" :album="item.album" :cover-name="item.coverName"
            :cover-src="item.coverSrc" :is-collet="item.isCollet" :singer="item.singer" :song-name="item.songName"
            :time="item.time" :index :acitve="item.acitve"></SongCard> -->
    </div>


</template>

<script lang="ts" setup>
import { Icons } from '@/config/icons';
import Card from './components/card.vue' 
</script>