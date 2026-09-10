<template>
    <section class="bg-linear-to-br from-[#1a1a2e] via-[#16213e] to-[#0f3460]  h-full p-1">
        <div
            class=" w-full rounded-xl bg-[rgba(255,255,255,0.06)] shadow-[0_20px_60px_rgba(0,0,0,0.4)]  p-2 border border-[rgba(255,255,255,0.1)]">

            <div class="flex">
                <div class="flex-1 flex items-center">
                    <span class="text-[20px] text-white  font-extrabold">
                        <small class="text-[#ccc]">
                            地址 :
                        </small>
                        127.0.0.1
                    </span>
                </div>
                <div class="flex-1 text-right">
                    <span class="inline-block  bg-[#363D53] p-2 rounded-[50%] ">
                        <svg t="1780306661860" width="26" height="26" viewBox="0 0 1024 1024" version="1.1"
                            fill="#CDCFD4" xmlns="http://www.w3.org/2000/svg" p-id="4994">
                            <path
                                d="M439.264 208a16 16 0 0 0-16 16v67.968a239.744 239.744 0 0 0-46.496 26.896l-58.912-34a16 16 0 0 0-21.856 5.856l-80 138.56a16 16 0 0 0 5.856 21.856l58.896 34a242.624 242.624 0 0 0 0 53.728l-58.88 34a16 16 0 0 0-6.72 20.176l0.848 1.68 80 138.56a16 16 0 0 0 21.856 5.856l58.912-34a239.744 239.744 0 0 0 46.496 26.88V800a16 16 0 0 0 16 16h160a16 16 0 0 0 16-16v-67.968a239.744 239.744 0 0 0 46.512-26.896l58.912 34a16 16 0 0 0 21.856-5.856l80-138.56a16 16 0 0 0-4.288-20.832l-1.568-1.024-58.896-34a242.624 242.624 0 0 0 0-53.728l58.88-34a16 16 0 0 0 6.72-20.176l-0.848-1.68-80-138.56a16 16 0 0 0-21.856-5.856l-58.912 34a239.744 239.744 0 0 0-46.496-26.88V224a16 16 0 0 0-16-16h-160z m32 48h96v67.376l28.8 12.576c13.152 5.76 25.632 12.976 37.184 21.52l25.28 18.688 58.448-33.728 48 83.136-58.368 33.68 3.472 31.2a194.624 194.624 0 0 1 0 43.104l-3.472 31.2 58.368 33.68-48 83.136-58.432-33.728-25.296 18.688c-11.552 8.544-24.032 15.76-37.184 21.52l-28.8 12.576V768h-96v-67.376l-28.784-12.576c-13.152-5.76-25.632-12.976-37.184-21.52l-25.28-18.688-58.448 33.728-48-83.136 58.368-33.68-3.472-31.2a194.624 194.624 0 0 1 0-43.104l3.472-31.2-58.368-33.68 48-83.136 58.432 33.728 25.296-18.688a191.744 191.744 0 0 1 37.184-21.52l28.8-12.576V256z m47.28 144a112 112 0 1 0 0 224 112 112 0 0 0 0-224z m0 48a64 64 0 1 1 0 128 64 64 0 0 1 0-128z"
                                p-id="4995"></path>
                        </svg>
                    </span>
                </div>


            </div>
            <!--  封面 -->

            <div class="w-40 h-40 rounded-[20px] overflow-hidden m-auto shadow-[0_2px_40px_rgba(0,0,0,0.4)]">
                <img src="https://picsum.photos/seed/song4/400/400" alt="">
            </div>
            <!-- 标题 -->
            <div class="text-center mt-4">
                <span class="text-white text-lg font-bold">

                    七里香
                </span>
            </div>
            <!-- 歌手名 -->
            <div class="text-center">
                <span class="text-[rgba(255,255,255,0.55)] text-sm">
                    周杰伦
                </span>
            </div>

            <!-- 进度条 -->
            <div class="p-4 ">
                <!-- 外层 -->
                <div class="bg-[rgba(255,255,255,0.12)] h-2 rounded-sm   relative w-full">
                    <!-- 进度 -->
                    <div
                        class="h-full w-[50%]   bg-linear-to-l  from-[#e94560] to-[#f06b7c]  relative  rounded-sm   after:content-[''] after:w-4  after:h-4 after:bg-white after:absolute after:-right-2  after:rounded-[50%]  after:translate-y-[-50%] after:top-[50%] after:shadow-[0_2px_6px_rgba(0,0,0,0.3)] ">
                    </div>
                </div>
                <!--  时间 -->
                <div class="flex mt-1">
                    <span class="flex-1  text-[rgba(255,255,255,0.4)] text-xs">0:00</span>
                    <span class="flex-1  text-[rgba(255,255,255,0.4)] text-xs text-right">0:00</span>

                </div>
            </div>
            <!-- 播放控制 -->
            <div class="pl-6 pr-6 flex justify-center items-center gap-8">
                <!-- 上一曲 -->

                <button class="w-6 h-6 bg-none  text-white ">
                    <svg viewBox="0 0 24 24" fill="currentColor">
                        <path d="M6 6h2v12H6zm3.5 6l8.5 6V6z" />
                    </svg>
                </button>

                <!-- 播放-暂停 -->
                <button @click="onPlay"
                    class=" w-12 h-12 bg-linear-to-br from-[#e94560] to-[#c23152] rounded-[50%]  text-white  flex items-center justify-center shadow-[0_6px_24px_rgba(233,69,96,0.4)]">
                    <svg v-if="!isPlaying" viewBox="0 0 24 24" fill="currentColor" class="w-8 h-8">
                        <path d="M8 5v14l11-7z" />
                    </svg>
                    <svg v-else viewBox="0 0 24 24" fill="currentColor" class="w-8 h-8">
                        <path d="M6 19h4V5H6v14zm8-14v14h4V5h-4z" />
                    </svg>
                </button>
                <!-- 下一曲 -->
                <button class="w-6 h-6 bg-none  text-white">
                    <svg viewBox="0 0 24 24" fill="currentColor">
                        <path d="M6 18l8.5-6L6 6v12zM16 6v12h2V6h-2z" />
                    </svg>
                </button>

            </div>


            <!-- 音量控制 -->

            <div class="flex items-center gap-4 pl-4 pr-4 mt-6">
                <!-- 喇叭 -->
                <span class="shrink-0 text-white/50 w-5 h-5">
                    <svg viewBox="0 0 24 24" fill="currentColor" v-if="volume > 0 && muted">
                        <path
                            d="M3 9v6h4l5 5V4L7 9H3zm13.5 3c0-1.77-1.02-3.29-2.5-4.03v8.05c1.48-.73 2.5-2.25 2.5-4.02zM14 3.23v2.06c2.89.86 5 3.54 5 6.71s-2.11 5.85-5 6.71v2.06c4.01-.91 7-4.49 7-8.77s-2.99-7.86-7-8.77z" />
                    </svg>
                    <svg viewBox="0 0 24 24" fill="currentColor" v-else>
                        <path
                            d="M16.5 12c0-1.77-1.02-3.29-2.5-4.03v2.21l2.45 2.45c.03-.2.05-.41.05-.63zm2.5 0c0 .94-.2 1.82-.54 2.64l1.51 1.51C20.63 14.91 21 13.5 21 12c0-4.28-2.99-7.86-7-8.77v2.06c2.89.86 5 3.54 5 6.71zM4.27 3L3 4.27 7.73 9H3v6h4l5 5v-6.73l4.25 4.25c-.67.52-1.42.93-2.25 1.18v2.06c1.38-.31 2.63-.95 3.69-1.81L19.73 21 21 19.73l-9-9L4.27 3zM12 4L9.91 6.09 12 8.18V4z" />
                    </svg>
                </span>
                <!-- 音量 -->
                <input type="range" class="flex-1 bg-white/12 appearance-none h-1 rounded-sm accent-white" min="0"
                    max="100" @input="onVolume" v-model="volume">
                <!-- 音量数值 -->
                <span class="w-10 text-right text-xs text-[rgba(255,255,255,0.4)]">
                    100%
                </span>

            </div>

            <!-- 如何播放 -->
            <div class="gap-4 flex justify-center mt-4">
                <playMode v-for="(item) in modeList" :label="item.label" :active="item.code === defaultMode"
                    @click="onModePlay(item)"></playMode>
            </div>


        </div>


    </section>

</template>

<script setup lang="ts">
import { ref } from 'vue';
import playMode from './components/playMode.vue';

interface ModeItem {
    label: string,
    code: string
    active: boolean
}

const defaultMode = ref<string>('list')

const modeList = ref<ModeItem[]>([
    {
        label: '列表循环',
        code: 'list',
        active: true
    },
    {
        label: '单曲循环',
        code: 'one',
        active: false
    }, {
        label: '随机播放',
        code: 'random',
        active: false
    }
])

const onModePlay = (item: ModeItem) => {
    defaultMode.value = item.code
}


const isPlaying = ref<boolean>(false)

// 音量
const volume = ref<number>(50)
// 静音
const muted = ref<boolean>(true)

// 播放事件 
const onPlay = () => {
    isPlaying.value = !isPlaying.value
}
// 音量拖动事件
const onVolume = (event: Event) => {
    const target = event.target as HTMLInputElement
    console.log(target.value);
    let val = Number(target.value)
    if (val === 0) {
        muted.value = false
    } else {
        muted.value = true
    }

}

</script>