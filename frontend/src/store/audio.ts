import type { MusicBase } from "@/pages/publicSongList/index.vue";
import { useAudioPlayUtil } from "@/utils/auido/audioPlay";
import { defineStore } from "pinia";
import { ref } from "vue";



export const useAudioStore = defineStore("audio", () => {
    const audioPlayUtil = useAudioPlayUtil()
    //  播放状态
    const audioPlayState = ref<boolean>(false)

    const currentMusicId = ref<number | null>(null)

    function setCurrentMusicId(id: number) {
        currentMusicId.value = id
    }

    function play(item: MusicBase) {

        audioPlayUtil.play(item)
        setCurrentMusicId(item.id)
    }
    function randomPlay() {
        audioPlayUtil.randomPlay()

    }
    function pause() {
        audioPlayState.value = false
        audioPlayUtil.pause()
    }

    function setAudioList(items: MusicBase[]) {
        audioPlayUtil.setMusicList(items)
    }

    // 设置播放状态
    function setAudioPlayState(state: boolean) {
        audioPlayState.value = state
    }

    audioPlayUtil.initHooks(() => setAudioPlayState(true), () => setAudioPlayState(false))


    return {
        audioPlayState,
        currentMusicId,
        setCurrentMusicId,
        play,
        randomPlay,
        pause,
        setAudioPlayState,
        setAudioList
    }

})