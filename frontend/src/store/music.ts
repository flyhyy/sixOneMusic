import MusicApi from "@/api/music";
import { defineStore } from "pinia";
import { computed, ref } from "vue";
import { useAudioStore } from "./audio";
import { useToast } from "vue-toastification";
import { getKeyIndex } from "@/utils";
import { formatTime } from "@/utils/time";
import { PlayListApi } from "@/api/playList";
import type { MusicBase } from "@/pages/publicSongList/index.vue";


export const useMusicStore = defineStore("music", () => {

    const toast = useToast()



    const audioStore = useAudioStore()
    // 播放队列
    const playQueueList = ref<MusicBase[]>([])
    // 全部音乐
    const songsList = ref<MusicBase[]>([])
    //  收藏音乐


    const totalTime = computed(() => {
        const t = songsList.value.reduce((total, b) => {
            return total + b.duration
        }, 0)

        return formatTime(t)
    })



    // 设置音乐列表数据
    function setSongList(list: MusicBase[]) {
        songsList.value = list
        audioStore.setAudioList(list)
    }

    // 获取所有列表数据
    async function getAllMusicList() {
        const data = await MusicApi.getMusicListAll(100, 1)
        const val: MusicBase[] = data.data.map(item => {
            return {
                isPlay: false,
                ...item
            }
        })
        songsList.value = val

        audioStore.setAudioList(val)



    }
    // 获取歌单列表数据
    const getSongPlayList = async (playListId: number) => {
        const res = await PlayListApi.getSongs(playListId)
        songsList.value = res
        audioStore.setAudioList(res)


    }

    // 设置播放队列
    function setPlayQueueList(item: MusicBase) {

        if (!playQueueList.value.length) {
            playQueueList.value.push(item)
            console.log(playQueueList.value);
        } else {
            let fIdx = playQueueList.value.findIndex((p) => p.id === item.id)
            if (fIdx === -1) {
                playQueueList.value.push(item)
            }
        }
    }
    // 收藏
    async function collecHandler(item: MusicBase) {
        let is_collect = !item.is_collect

        const res = await MusicApi.collect(item.id, is_collect)
        if (typeof res === "number") {
            const index = getKeyIndex(songsList.value, 'id', res)
            if (index > -1) {
                songsList.value[index].is_collect = is_collect
                if (is_collect) {

                    toast.success("收藏成功！")
                } else {
                    toast.success("取消收藏！")

                }
            }
        } else {
            toast.error("收藏失败")
        }

    }
    // 获取收藏歌曲
    async function getCollectSongs() {
        const res = await MusicApi.getCollectSongs()
        songsList.value = res
        audioStore.setAudioList(res)

    }



    return {
        songsList,
        totalTime,
        playQueueList,
        setSongList,
        getAllMusicList,
        getSongPlayList,
        setPlayQueueList,
        getCollectSongs,
        collecHandler
    }


})