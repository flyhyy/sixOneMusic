/* prettier-ignore */
/**
* @Author: hyy
* @Date: 2026-06-17 13:22:57
* @Description: 公共列表页面
*/

<template>

    <TitleInfo :icon-name="Icons.sidePlayQueue" :title="currentConfig?.title || ''" size="large">

    </TitleInfo>

    <div class="flex items-center mt-6 gap-2">
        <PlayButton v-for="item in currentConfig?.buttons || []" @click="item.fn()" :icon-name="item.iconName">
            {{ item.name }}
        </PlayButton>

        <span class="ml-auto text-muted text-[13px]">
            共{{ songsList.length }}首 {{ totalTime }}
        </span>

    </div>

    <!-- 空数据提示 -->
    <!-- <EmptyData :icon-name="Icons.sideMyPlayList">
        收藏为空
        <template #other>
            <PlayButton :icon-name="Icons.compass" @click="onJumpDiscoverPage">去发现音乐</PlayButton>
        </template>
</EmptyData> -->
    <!-- 有数据处理 -->
    <div class="mt-4">
        <SongCard v-for="(item, index) in songsList" :album="item.album_name" :cover-src="item.cover_path"
            :isCollect="item.is_collect" :singer="item.singer"
            :is-play="audioStore.currentMusicId === item.id && audioStore.audioPlayState" :song-name="item.title"
            :time="item.duration" :index="index + 1" @play="onPlay(item, $event)" @addPlayQueue="onAddPlayQueue(item)"
            @collect="musicStore.collecHandler(item)" @addPlayList="onAddPlayList(item, $event)">
        </SongCard>
    </div>



</template>


<script lang="ts" setup>
import { Icons, type IconName } from '@/config/icons';
import { useAudioStore } from '@/store/audio';
import { useMusicStore } from '@/store/music';
import { storeToRefs } from 'pinia';
import { usePopoverUtil } from '@/utils/popover/popover';
import PlayList from '@/layout/components/sideMyPlayList/playList.vue'
import { useRoute } from 'vue-router';
import { ref, watch } from 'vue';

export interface MusicBase {
    title: string
    singer: string
    duration: number
    size: number
    id: number
    cover_path: string
    is_collect: boolean
    album_name: string

}


export interface MusicListAllResponse {

    page: number,
    pageSize: number,
    data: MusicBase[]

}
interface ButtonConfig {
    name: string,
    iconName: IconName
    fn: ((...args: any[]) => void)
}
interface PageTypeConfigFace {
    buttons: ButtonConfig[],
    title: string,

    initFn: ((...args: any[]) => void)
}

export type PageType = "AllMusic" | "PlayQueue" | "PlayList" | "MyCollect"




const musicStore = useMusicStore()
const audioStore = useAudioStore()
const popoverUtil = usePopoverUtil(PlayList)
const route = useRoute()



const { songsList, totalTime, playQueueList } = storeToRefs(musicStore)

const playAll = () => {
    audioStore.play(songsList.value[0])

}

const onPlay = (item: MusicBase, isPlay: boolean) => {

    if (isPlay) {
        audioStore.pause()
    } else {

        audioStore.play(item)
    }
}
// 加入播放队列
const onAddPlayQueue = (item: MusicBase) => {

    musicStore.setPlayQueueList(item)
}

const onAddPlayList = (item: MusicBase, node: HTMLElement) => {
    popoverUtil.show(node, item)
}

const PageTypeConfig: Record<PageType, PageTypeConfigFace> = {

    "AllMusic": {
        title: "全部音乐",
        buttons: [
            {
                name: '播放全部',
                fn: playAll,
                iconName: Icons.play,
            },
            {
                name: '随机播放',
                fn: audioStore.randomPlay,
                iconName: Icons.random
            },

        ],

        initFn: musicStore.getAllMusicList
    },
    "PlayQueue": {
        title: "播放队列",
        buttons: [
            {
                name: '播放全部',
                fn: playAll,
                iconName: Icons.play,

            },
            // {
            //     name: '清空',
            //     fn: () => { },
            //     iconName: Icons.delete
            // },
        ],
        initFn: () => {
            musicStore.setSongList(playQueueList.value)
        }

    },
    "PlayList": {
        title: "歌单列表",
        buttons: [
            {
                name: '播放全部',
                fn: playAll,
                iconName: Icons.play,

            },
        ],
        initFn: musicStore.getSongPlayList

    },
    "MyCollect": {
        title: "我的收藏",
        buttons: [
            {
                name: '播放全部',
                fn: playAll,
                iconName: Icons.play,

            },
            {
                name: '随机播放',
                fn: audioStore.randomPlay,
                iconName: Icons.random
            },
        ],
        initFn: musicStore.getCollectSongs

    }



}


const currentConfig = ref<PageTypeConfigFace>()
const init = () => {

    const t = route.query.type as PageType

    if (t === "PlayList") {
        const playListId = route.query.playListId as string
        const playListName = route.query.playListName
        PageTypeConfig[t].title = "歌单列表"
        PageTypeConfig[t].title = playListName + PageTypeConfig[t].title
        PageTypeConfig[t].initFn(playListId)
    } else {
        PageTypeConfig[t].initFn()

    }

    currentConfig.value = PageTypeConfig[t]

}

init()

watch(() => [route.path, route.query], (n) => {
    console.log(n);
    init()

})

</script>