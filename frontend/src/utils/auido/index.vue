<template>
    <Transition>
        <div class="absolute bottom-0 w-full">
            <!-- 1. 使用 max-height 和 opacity 控制动画，确保方向严格垂直向下 -->
            <Transition enter-active-class="transition-all duration-300 ease-out origin-bottom"
                enter-from-class="max-h-0 opacity-0" enter-to-class="max-h-[150px] opacity-100"
                leave-active-class="transition-all duration-300 ease-in origin-bottom"
                leave-from-class="max-h-[150px] opacity-100" leave-to-class="max-h-0 opacity-0">
                <!-- 2. 外层只负责 overflow-hidden 截断，坚决不能有 padding -->
                <div v-show="showLrc" class="overflow-hidden w-full">

                    <!-- 3. 原封不动保留你的内部 UI，保持其固定高度，避免 mask-image 变形 -->
                    <div class="relative right-0 bottom-0 left-0 bg-surface/90 border-t border-t-line">
                        <div class="py-5">
                            <div
                                class="scrollbar-hide h-20 min-h-9.5 overflow-y-auto scroll-smooth px-10 py-2 text-center transition-all duration-300 ease-in-out mask-image-gradient">
                                <template v-if="lrcTextArr?.length">
                                    <div v-for="(item, index) in lrcTextArr" ref="lrcItemRefs"
                                        @click="seekAudio(item.time)" class="cursor-pointer" :class="{
                                            'text-lg font-bold text-brand': index === currentLrcIndex,
                                            'text-base text-muted': index !== currentLrcIndex,
                                        }">
                                        {{ item.text }}
                                    </div>
                                </template>
                                <template v-else>
                                    <div class="h-full flex items-center justify-center text-brand">
                                        暂无歌词!
                                    </div>
                                </template>
                            </div>
                        </div>
                        <span class="absolute right-0 top-0 cursor-pointer p-2">
                            <Icon :name="Icons.close" @click="showLrc = false"></Icon>
                        </span>
                    </div>

                </div>
            </Transition>

            <div class="relative flex h-20 items-center gap-4 border-t border-line bg-surface px-5">
                <div class="flex w-50 gap-2.5">
                    <div class="h-11.5 w-11.5 rounded-lg">
                        <Cover :name="musicItem?.singer || ''" :src="musicItem?.cover_path || ''"></Cover>
                    </div>
                    <!-- 歌手，歌曲名称 -->
                    <div>
                        <SongInfo :class="musicItem?.id || false ? 'text-brand!' : ''"
                            :primary-text="musicItem?.title || ''" :secondary-text="musicItem?.singer || ''">
                        </SongInfo>
                    </div>
                </div>

                <!-- 播放控制 -->
                <div class="mx-auto my-0 flex max-w-150 flex-1 flex-col items-center gap-1">
                    <!-- 控制 -->
                    <div class="flex items-center gap-2.5">
                        <Icon :name="iconPlayMode" @click="onPlayMode"></Icon>
                        <Icon :name="Icons.preSong" class="text-sub" @click="onSwitch('pre')"></Icon>
                        <ButtonIcon :icon-name="iconPlayPause" isFull @click="onPlayPauseCtr"></ButtonIcon>
                        <Icon :name="Icons.nextSong" class="text-sub" @click="onSwitch('next')"></Icon>
                    </div>

                    <!-- 进度条 -->
                    <div class="flex w-full items-center gap-2.5">
                        <!-- 当前时间 -->
                        <span class="shrink-0 text-center text-sm text-muted">{{
                            currentDuration
                            }}</span>

                        <!-- 轨道背景 -->
                        <SlideRail v-model="currentProgress" @drag-start="onProgressDragStart"
                            @drag-end="onProgressDragEnd">
                        </SlideRail>
                        <!-- 总时间 -->
                        <span class="shrink-0 text-center text-sm text-muted">{{
                            totalDuration
                            }}</span>
                    </div>
                </div>
                <!--  歌词,收藏,音量控制 -->
                <div class="flex w-50 shrink-0 items-center justify-end gap-2.5">
                    <Icon :name="Icons.lrc" class="text-sub" size="large" @click="showLrc = !showLrc"></Icon>
                    <Icon :name="Icons.heart" :class="musicItem?.is_collect || false ? 'text-brand' : ''"
                        class="text-sub" size="large" @click="onCollect"></Icon>
                    <div class="flex items-center gap-5">
                        <Icon :name="volumeIcon" class="text-sub" size="large" @click="onVolumeIcon"></Icon>
                        <div class="w-20">
                            <SlideRail :progress="20" :type="'volume'" v-model="volume"></SlideRail>
                        </div>
                    </div>
                </div>

                <audio :src="audioSrc" ref="audioRef" @loadedmetadata="onMetaData" @timeupdate="onTimeUpdate"
                    @pause="onPause" @play="onPlay" @ended="onPlayEnd"></audio>
            </div>
        </div>
    </Transition>
</template>

<script setup lang="ts">
import { Icons, type IconName } from "@/config/icons";
import { computed, nextTick, ref, useTemplateRef, watch } from "vue";
import { formatTime } from "../time";
import type {
    EmitDragEndInfo,
    EmitDragStartInfo,
} from "@/components/slideRail/index.vue";
import AudioApi from "@/api/audio";
import { shuffleArray } from "../shuffle";
import { useToast } from "vue-toastification";
import { ParseLrc, type LyricLine } from "../lrc";
import { useAudioStore } from "@/store/audio";
import type { MusicBase } from "@/pages/publicSongList/index.vue";

// 声明两个外部注入的方法
let externalOnPlay: (() => void) | null = null;
let externalOnPause: (() => void) | null = null;

const audioStore = useAudioStore()
const audioSrc = ref<string>("");

const audioRef = useTemplateRef("audioRef");

const duration = ref<number>(0);

const musicItem = ref<MusicBase>();


const toast = useToast();

const volume = ref<number>(50);

const volumeIcon = ref<IconName>(Icons.volume);
const showLrc = ref<boolean>(false)
const musicList = ref<MusicBase[]>([]);
// 打乱后的随机列表
const musicShuffleList = ref<MusicBase[]>([]);

//  列表模式是否在最后一个播放完毕
const isListPlayEnd = ref<boolean>(false);

const onCollect = () => {

}

const lrcTextArr = ref<LyricLine[]>();
const currentLrcIndex = ref<number>(-1);

const getCurrentMusicIdx = () => {
    if (!musicItem.value?.id || !musicList.value?.length) return -1;
    const currentIdx = musicList.value.findIndex(
        (item) => item.id === musicItem.value?.id,
    );
    return currentIdx;
};

watch(volume, (n) => {
    if (audioRef.value) {
        let v = n / 100;
        audioRef.value.volume = v;
        if (v === 0) {
            volumeIcon.value = Icons.mute;
        } else {
            volumeIcon.value = Icons.volume;
        }
    }
});
const isMuted = ref<boolean>(false);
const onVolumeIcon = () => {
    if (audioRef.value) {
        isMuted.value = !isMuted.value;
        audioRef.value.muted = isMuted.value;
        volumeIcon.value = isMuted.value ? Icons.mute : Icons.volume;
    }
};

// 重头播放
const firstPlay = () => {
    let item: MusicBase | undefined;
    if (musicItem.value === undefined) {
        item = musicList.value[0];
    } else if (playMode.value === "list" && isListPlayEnd.value) {
        // 获取第一个歌曲信息
        item = musicList.value[0] as MusicBase;

        isListPlayEnd.value = false;
    }

    if (item && item.id) {
        play(item);
    }
};

// 顺序播放
const listPlayMode = (isCtr: boolean = false) => {
    if (getCurrentMusicIdx() === musicList.value.length - 1) {
        console.log("顺序播放——————播放完毕");
        if (isCtr) {
            toast.warning("当前歌曲为最后一曲，无法切换到下一曲");
            return;
        }

        iconPlayPause.value = Icons.play;
        isListPlayEnd.value = true;

        return;
    }
    const nextIdx = getCurrentMusicIdx() + 1;
    const item = musicList.value[nextIdx];
    play(item);
};

// 循环播放
const loopPlayMode = (isCtr: boolean = false) => {
    const listLen = musicList.value?.length || 0;

    // 1. 边界防呆：如果列表为空，直接终止，防止取模时出现 NaN 或抛错
    if (listLen === 0) return;

    // 2. 获取当前索引。如果返回 -1 (未找到)，默认将其视作 0
    let currentIdx = getCurrentMusicIdx();
    if (currentIdx === -1) currentIdx = 0;

    // 3. 计算下一曲索引：利用步长和长度补偿，完美解决负数取模问题
    const step = isCtr ? -1 : 1;
    const nextIdx = (currentIdx + step + listLen) % listLen;

    // 4. 获取目标歌曲并执行安全播放
    const item = musicList.value[nextIdx];
    if (item && item.id) {
        play(item);
    }
};

// 单曲循环
const oneLoopPlayMode = () => {
    if (audioRef.value) {
        audioRef.value.currentTime = 0;
        audioRef.value.play();
        resetLrc()

    }
};
// 切换到“随机播放”模式时的初始化动作
const initRandomMode = () => {
    if (!musicList.value || musicList.value.length === 0) return;

    // 1. 先生成一个全打乱的随机队列
    let tempShuffled = shuffleArray(musicList.value);

    // 2. 【核心细节】：为了绝对不重复，如果当前正在放某首歌，
    //    我们需要把“当前这首歌”强制推到打乱数组的第 0 位（首位）！
    if (musicItem.value?.id) {
        const currentIdx = tempShuffled.findIndex(
            (item) => item.id === musicItem.value?.id,
        );
        if (currentIdx !== -1) {
            // 从数组中删除当前歌曲，并在头部重新插入
            const [current] = tempShuffled.splice(currentIdx, 1);
            tempShuffled.unshift(current);
        }
    }

    musicShuffleList.value = tempShuffled;
};
// 随机播放
const randomPlayMode = () => {
    // 如果当前没有任何数据或只有 1 首歌，直接强制循环当前歌曲
    if (!musicShuffleList.value.length) return;
    if (musicShuffleList.value.length === 1) {
        play(musicShuffleList.value[0]);
        return;
    }

    // 1. 在“随机队列”里找当前音频的坐标
    let currentIdx = musicShuffleList.value.findIndex(
        (item) => item.id === musicItem.value?.id,
    );

    // 2. 下一首序号自增 1
    let nextIdx = currentIdx + 1;

    // 3. 如果整整一轮“随机打乱的列表”都已经放完了，自动重新洗一遍牌！
    if (nextIdx >= musicShuffleList.value.length) {
        console.log("当前洗牌队列已全部播完，进行第二轮新洗牌...");
        initRandomMode();
        nextIdx = 1; // 因为新洗牌的第0位会被保持为“当前刚放完的曲目”，所以直接跳去听第1位
    }

    // 4. 播放下一曲，天然保证不同于上一曲！
    const nextSong = musicShuffleList.value[nextIdx];
    play(nextSong);
};

const playModeConfig = {
    list: {
        label: "顺序播放",
        icon: Icons.sideMyPlayList,
        fun: listPlayMode,
    },
    loop: {
        label: "循环播放",
        icon: Icons.listRepeat,
        fun: loopPlayMode,
    },
    oneLoop: {
        label: "单曲循环",
        icon: Icons.oneRepeat,
        fun: oneLoopPlayMode,
    },
    random: {
        label: "随机播放",
        icon: Icons.random,
        fun: randomPlayMode,
    },
};

export type PlayMode = keyof typeof playModeConfig;

const currentTime = ref<number>(0);

const iconPlayPause = ref<IconName>(Icons.play);

// 默认列表模式
const iconPlayMode = ref<IconName>(Icons.sideMyPlayList);
//  'list' | 'loop' | 'oneLoop' | 'random'

const playMode = ref<PlayMode>("list");

// 播放结束监听
const onPlayEnd = () => {
    playModeConfig[playMode.value].fun();
};

const playModelConfigkeys = Object.keys(playModeConfig);

const playModelConfigkeysLen = playModelConfigkeys.length;

const setMusicList = (list: MusicBase[]) => {
    musicList.value = list;
    initRandomMode();
};

const setPlayMode = (mode: PlayMode) => {
    playMode.value = mode;
    iconPlayMode.value = playModeConfig[mode].icon;
};

// 播放模式
const onPlayMode = () => {
    const cIdx = playModelConfigkeys.indexOf(playMode.value);

    const nextIdx = (cIdx + 1) % playModelConfigkeysLen;
    const nextKey = playModelConfigkeys[nextIdx] as PlayMode;

    playMode.value = nextKey;

    iconPlayMode.value = playModeConfig[nextKey].icon;
};

const isDragging = ref<boolean>(false);

// 监听进度条拖动
const onProgressDragStart = (obj: EmitDragStartInfo) => {
    isDragging.value = true;
    watchCurrentProgress.pause();
    currentTime.value = (obj.value / 100) * duration.value;
};
// 监听进度条停止拖动
const onProgressDragEnd = (obj: EmitDragEndInfo) => {
    isDragging.value = false;
    if (audioRef.value) {
        audioRef.value.currentTime = (obj.value / 100) * duration.value;
        watchCurrentProgress.resume();
    }
};

// 总时长
const totalDuration = computed(() => {
    return formatTime(duration.value);
});

//当前播放时长
const currentDuration = computed(() => {
    return formatTime(currentTime.value);
});
// 当前播放进度
const currentProgress = ref<number>(0);

// 监听当前进度
const watchCurrentProgress = watch(currentTime, (n) => {
    currentProgress.value = (n / duration.value) * 100;
});

const lrcItemRefs = useTemplateRef<HTMLDivElement[]>("lrcItemRefs");
watch(currentLrcIndex, async (n) => {
    // 确保 DOM 已经更新
    await nextTick()
    if (lrcItemRefs.value) {
        const acitveItem = lrcItemRefs.value[n];
        if (acitveItem) {
            acitveItem.scrollIntoView({
                behavior: "smooth",
                block: "center",
            });
        }
    }
});

const seekAudio = (targetTime: number) => {
    // 假设你引用了之前的 audioRef
    if (audioRef.value) {
        audioRef.value.currentTime = targetTime
        audioRef.value.play() // 顺便确保它开始播放
    }
}

// 当前播放时长
const onTimeUpdate = (e: Event) => {
    if (isDragging.value) return;
    const ev = e.target as HTMLAudioElement;
    currentTime.value = ev.currentTime;

    if (lrcTextArr.value?.length) {
        // 从最后一句往前找
        for (let i = lrcTextArr.value.length - 1; i >= 0; i--) {
            if (currentTime.value >= lrcTextArr.value[i].time) {
                // 找到了当前行，如果索引没变就不做多余的操作
                if (currentLrcIndex.value !== i) {
                    currentLrcIndex.value = i;
                }
                break; // 找到就退出循环
            }
        }
    }
};

// 播放 暂停 按钮控制
const onPlayPauseCtr = () => {
    if (audioRef.value) {
        if (musicList.value.length === 0) {
            console.warn("播放列表为空，无法播放");
            toast.error("播放列表为空，无法播放");
            return;
        }
        firstPlay();

        if (!audioRef.value.paused) {
            audioRef.value.pause();
        } else {
            audioRef.value.play();
        }

    }
};

// 切换控制
const onSwitch = (type: "pre" | "next") => {
    console.log(type);

    if (type === "next") {
        playModeConfig[playMode.value].fun(true);
    } else {
        // 上一曲操作
        if (playMode.value === "list") {
            const preIdx = getCurrentMusicIdx() - 1;
            const item = musicList.value[preIdx];
            play(item);
        } else if (playMode.value === "oneLoop") {
            oneLoopPlayMode();
        } else if (playMode.value === "random") {
            randomPlayMode();
        } else {
            loopPlayMode(true);
        }
    }
};

//  暂停事件监听
const onPause = () => {
    iconPlayPause.value = Icons.play

    if (externalOnPause) externalOnPause()

};

//  播放事件监听
const onPlay = () => {

    iconPlayPause.value = Icons.pause

    audioStore.setCurrentMusicId(musicItem.value?.id || -1)
    if (externalOnPlay) externalOnPlay()
};

// 暴露注入方法给 Util
const setExternalHooks = (onPlay: () => void, onPause: () => void) => {
    externalOnPlay = onPlay;
    externalOnPause = onPause;
};

// 3. 增加一个纯粹的暂停方法
const pauseAudio = () => {
    if (audioRef.value && !audioRef.value.paused) {
        audioRef.value.pause();
    }
};
const play = (item: MusicBase) => {
    musicItem.value = item;
    const src = AudioApi.audioPath(item.id,);

    audioSrc.value = src;
    nextTick(async () => {
        try {
            const lrc = await AudioApi.audioLrc(item.id);


            lrcTextArr.value = ParseLrc(lrc);

            await audioRef.value?.play();
            resetLrc()



        } catch (err) {
            console.log("播放异常");
            onSwitch("next")
        }
    });
};
const resetLrc = () => {
    currentLrcIndex.value = -1
}

const onMetaData = (e: Event) => {
    const ev = e.target as HTMLAudioElement;
    duration.value = ev.duration;
};

defineExpose({
    play,
    randomPlayMode,
    onPlayPauseCtr,
    pauseAudio,
    setMusicList,
    setPlayMode,
    setExternalHooks,
    audioRef,

});
</script>

<style scoped>
/* 渐变遮罩：让顶部和底部的歌词有一种渐渐消失（淡出）的视觉效果 */
.mask-image-gradient {
    mask-image: linear-gradient(to bottom,
            transparent 0%,
            black 20%,
            black 80%,
            transparent 100%);
    -webkit-mask-image: linear-gradient(to bottom,
            transparent 0%,
            black 20%,
            black 80%,
            transparent 100%);
}
</style>