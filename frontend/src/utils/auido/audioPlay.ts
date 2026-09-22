import { createVNode, render, type VNode } from 'vue'
import AudioPlay, { type PlayMode } from './index.vue'
import type { MusicBase } from '@/pages/publicSongList/index.vue'
type AudioInstanceType = InstanceType<typeof AudioPlay>

let audioInstance: AudioInstanceType | null = null
let container: HTMLDivElement | null = null
let vNode: VNode | null = null

// 保存外部 hooks 回调，播放控制器重建后自动重新绑定
let externalOnPlay: (() => void) | null = null
let externalOnPause: (() => void) | null = null


function getInstance(): AudioInstanceType {
    if (!audioInstance) {
        container = document.createElement("div")
        document.body.append(container)
        vNode = createVNode(AudioPlay)
        render(vNode, container)
        audioInstance = (vNode.component?.exposed as AudioInstanceType) || null

        // 重建后重新绑定之前保存的 hooks
        audioInstance.setExternalHooks(
            externalOnPlay || (() => {}),
            externalOnPause || (() => {})
        )
    }
    return audioInstance
}

/**
 * 销毁播放控制器：卸载组件、移除 DOM、清空实例引用。
 * 退出登录时调用，下次播放时会自动重新创建。
 */
export function destroyAudioPlay() {
    if (vNode && container) {
        render(null, container)
        container.remove()
    }
    vNode = null
    container = null
    audioInstance = null
}


export function useAudioPlayUtil() {
    const initHooks = (onPlay: () => void, onPause: () => void) => {
        externalOnPlay = onPlay
        externalOnPause = onPause
        getInstance().setExternalHooks(onPlay, onPause)
    }
    // 播放
    const play = (musicItem: MusicBase) => {
        getInstance().play(musicItem)


    }
    // 随机播放
    const randomPlay = () => {
        getInstance().randomPlayMode()
    }

    // 暂停
    const pause = () => {
        getInstance().pauseAudio()

    }
    // 设置音乐列表
    const setMusicList = (musicList: MusicBase[]) => {
        getInstance().setMusicList(musicList)

    }
    // 设置播放模式
    const setPlayMode = (playMode: PlayMode) => {
        getInstance().setPlayMode(playMode)

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
