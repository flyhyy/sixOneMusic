import { createVNode, render } from 'vue'
import AudioPlay, { type PlayMode } from './index.vue'
import type { MusicBase } from '@/pages/publicSongList/index.vue'
type AudioInstanceType = InstanceType<typeof AudioPlay>

let audioInstance: AudioInstanceType | null = null


function getInstance(): AudioInstanceType {
    if (!audioInstance) {

        const container = document.createElement("div")
        document.body.append(container)
        const vNode = createVNode(AudioPlay)
        render(vNode, container)
        audioInstance = (vNode.component?.exposed as AudioInstanceType) || null


    }
    return audioInstance
}


export function useAudioPlayUtil() {
    const audioInstance = getInstance()

    const initHooks = (onPlay: () => void, onPause: () => void) => {
        if (audioInstance?.setExternalHooks) {
            audioInstance.setExternalHooks(onPlay, onPause)
        }
    }
    // 播放
    const play = (musicItem: MusicBase) => {
        audioInstance.play(musicItem)


    }
    // 随机播放
    const randomPlay = () => {
        audioInstance.randomPlayMode()
    }

    // 暂停
    const pause = () => {
        audioInstance.pauseAudio()

    }
    // 设置音乐列表
    const setMusicList = (musicList: MusicBase[]) => {
        audioInstance.setMusicList(musicList)

    }
    // 设置播放模式
    const setPlayMode = (playMode: PlayMode) => {
        audioInstance.setPlayMode(playMode)

    }

    return {
        play,
        pause,
        randomPlay,
        setMusicList,
        setPlayMode,
        initHooks
    }
}