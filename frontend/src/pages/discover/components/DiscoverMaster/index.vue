/* prettier-ignore */
/**
* @Author: hyy
* @Date: 2026-06-11 17:44:18
* @Description: 发现音乐-信息
*/


<template>
    <div
        class="relative flex flex-wrap items-center gap-6 overflow-hidden rounded-2xl bg-surface p-8 after:absolute after:h-55 after:w-55  after:bg-brand after:opacity-[0.05] after:content-[''] after:pointer-events-none after:rounded-full after:-right-16 after:-top-16">
        <div class="z-1 min-w-50 flex-1">
            <h2 class="flex items-center gap-2 text-2xl font-extrabold text-ink">
                <Icon :name="currentHourData.icon"></Icon>
                <span>
                    {{ currentHourData.hint }}
                </span>
            </h2>
        </div>
        <div>
            <PlayButton class="mr-2.5"></PlayButton>
            <PlayButton :icon-name="Icons.random" theme="light">
                随机播放
            </PlayButton>
        </div>
    </div>

</template>

<script lang="ts" setup>
import { Icons } from "@/config/icons";
import { computed, onMounted, onUnmounted, ref } from "vue";

interface TimeRangeIcon {
    icon: string;
    startTime: number;
    endTime: number;
    hint: string;
}

// 时间范围图标设定
const timeRangeIcon: TimeRangeIcon[] = [
    {
        // 凌晨
        icon: Icons.discoverCurrentTimeEarlyMorning,
        startTime: 0,
        endTime: 5,
        hint: "夜深了，请注意多休息。",
    },
    {
        // 早上
        icon: Icons.discoverCurrentTimeMorning,
        startTime: 6,
        endTime: 11,
        hint: "早安！新的一天，别忘了吃顿美味的早餐给自己充充电。",
    },
    {
        // 中午
        icon: Icons.discoverCurrentTimeNoon,
        startTime: 12,
        endTime: 13,
        hint: "中午好！记得按时吃午饭哦。",
    },
    {
        // 下午
        icon: Icons.discoverCurrentTimeAfternoon,
        startTime: 14,
        endTime: 18,
        hint: "下午好！站起来伸个懒腰，喝杯水稍微放松一下吧。",
    },
    {
        // 晚上
        icon: Icons.discoverCurrentTimeNight,
        startTime: 19,
        endTime: 23,
        hint: "晚上好！享受属于你的休闲时光。",
    },
];

let timer: ReturnType<typeof setTimeout> | undefined = undefined;

const currentHour = ref<number | null>(null);

const currentHourData = computed(() => {
    const _currentHour = currentHour.value ?? new Date().getHours();
    return (
        getHourData(_currentHour) ?? {
            icon: "",
            startTime: 0,
            endTime: 23,
            hint: "享受音乐时光",
        }
    );
});

// 获取时间图标
const getHourData = (hour: number): TimeRangeIcon | undefined => {
    const matchedItem = timeRangeIcon.find((item) => {
        if (hour >= item.startTime && hour <= item.endTime) {
            return item;
        }
    });
    return matchedItem;
};
// 当前小时
const currentTimeHour = () => {
    const now = new Date();
    const nextHour = new Date(now.getTime());
    nextHour.setHours(now.getHours() + 1);
    nextHour.setMinutes(0, 0, 0);

    // 计算下一整点的毫秒数
    const delay = nextHour.getTime() - now.getTime();
    console.log(delay);

    timer = setTimeout(() => {
        currentHour.value = new Date().getHours();

        currentTimeHour();
    }, delay);
};

onMounted(() => {
    currentTimeHour();
});
onUnmounted(() => {
    clearTimeout(timer);
});

</script>