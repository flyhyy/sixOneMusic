import { PlayListApi } from "@/api/playList";
import type { PlayListItem, PlayListItemBase } from "@/layout/components/sideMyPlayList/index.vue";
import type { MusicBase } from "@/pages/publicSongList/index.vue";
import { defineStore } from "pinia";
import { ref } from "vue";


export const usePlayListStore = defineStore("playList", () => {

    const playList = ref<PlayListItem[]>([]);

    const playListSongs = ref<MusicBase[]>([])

    const getPlayList = async () => {
        const data = await PlayListApi.list()

        if (data && data.length) {
            const arr = data.map(item => {
                return {
                    isEdit: false,
                    ...item
                }
            }) as PlayListItem[]
            playList.value = arr
        } else {
            playList.value = []
        }

    }


    const updatePlayList = async (item: PlayListItemBase) => {
        await PlayListApi.Update(item)
        getPlayList()
    }
    const addPlayList = async (name: string) => {

        await PlayListApi.Add(name)
        getPlayList()

    }
    const delPlayList = async (id: number) => {
        await PlayListApi.Del(id)
        getPlayList()

    }

    const addSongPlayList = async (playListId: number, songId: number) => {
        await PlayListApi.addSong(playListId, songId)
    }

    const getSongPlayList = async (playListId: number) => {
        const res = await PlayListApi.getSongs(playListId)
        playListSongs.value = res

    }

    getPlayList()


    return {
        playList,
        playListSongs,
        getPlayList,
        updatePlayList,
        addPlayList,
        delPlayList,
        addSongPlayList,
        getSongPlayList
    }


})