/* prettier-ignore */
/**
* @Author: hyy
* @Date: 2026-06-15 16:29:51
* @Description: 主要数据展示
*/

<template>

    <!-- 最近播放 -->
    <!-- <TitleInfo title="最近播放" :icon-name="Icons.history">
        <div class="flex items-center gap-3.5">
            <PlayButton theme="light" size="small"> 播放</PlayButton>
            <span @click="onCheckAll"
                class="text-sm  cursor-pointer flex items-center text-muted gap-1 hover:text-brand  transition duration-200">
                查看全部
                <Icon :name="Icons.arrowRight"></Icon>
            </span>
        </div>
    </TitleInfo>

    <div class="mt-4 flex  overflow-x-auto gap-3.5 snap-x pb-1.5">
        <MusicCard v-for="item in mockData" :cover="item.cover" :singer="item.singer" :song-name="item.songName">
        </MusicCard>

    </div> -->

    <!-- 艺术家 -->

    <TitleInfo title="艺术家" :icon-name="Icons.artist">
        <span @click="onShowAllArtist"
            class="text-sm  cursor-pointer flex items-center text-muted gap-1 hover:text-brand  transition duration-200">
            {{ isShowArtist ? '收起' : '展开全部' }}
            <Icon :name="isShowArtist ? Icons.arrowUp : Icons.arrowDown"></Icon>
        </span>
    </TitleInfo>

    <div class="mt-4 grid gap-3.5  grid-cols-[repeat(auto-fill,minmax(110px,1fr))]">
        <ArtistCard v-for="item in singerList" :music-num="item.total_num" :singer="item.name"
            @click="onArtListItem(item)"></ArtistCard>
    </div>

    <!-- 专辑收藏 -->

    <TitleInfo title="专辑收藏" :icon-name="Icons.album" class="mt-4">

        <span @click="onShowAllAlbum"
            class="text-sm  cursor-pointer flex items-center text-muted gap-1 hover:text-brand  transition duration-200">
            {{ isShowAlbum ? '收起' : '展开全部' }}
            <Icon :name="isShowAlbum ? Icons.arrowUp : Icons.arrowDown"></Icon>
        </span>

    </TitleInfo>
    <div class="mt-4 grid gap-4.5 grid-cols-[repeat(auto-fill,minmax(200px,1fr))]">

        <AlbumCard v-for="item in albumList" :album-name="item.name" :singer="item.singer"
            @click="onAblumListItem(item)" :total-time="formatTime(item.totoal_duration)" :total-num="item.total_num">
        </AlbumCard>
    </div>


</template>

<script lang="ts" setup>
import { Icons } from '@/config/icons';
import ArtistCard from '@/pages/discover/components/Card/artistCard.vue'
import AlbumCard from '@/pages/discover/components/Card/albumCard.vue';




import { ref } from 'vue';
import { DiscoverApi } from '@/api/discover';
import type { AblumListRespose, SingerListResponse } from '../../types';
import { formatTime } from '@/utils/time';
import { useAudioStore } from '@/store/audio';

const singerList = ref<SingerListResponse[]>([])
const audioStore = useAudioStore()

const getSingerList = async () => {
    const data = await DiscoverApi.GetSingerList()
    singerList.value = data
}
getSingerList()


const albumList = ref<AblumListRespose[]>([])
const getAlbumList = async () => {

    const data = await DiscoverApi.GetAlbumList()
    albumList.value = data
}
getAlbumList()



const isShowArtist = ref(false)
const isShowAlbum = ref(false)




// 艺术家 展开全部
const onShowAllArtist = () => {
    isShowArtist.value = !isShowArtist.value
}
// 专辑 展开全部
const onShowAllAlbum = () => {
    isShowAlbum.value = !isShowAlbum.value
}


const onArtListItem = async (item: SingerListResponse) => {
    const data = await DiscoverApi.GetSingerSongList(item.name)

    audioStore.play(data[0])
    audioStore.setAudioList(data)
}


const onAblumListItem = async (item: AblumListRespose) => {
    const data = await DiscoverApi.GetAlbumSongList(item.name)
    audioStore.play(data[0])
    audioStore.setAudioList(data)
}

</script>
