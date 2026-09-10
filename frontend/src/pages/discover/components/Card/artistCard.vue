/* prettier-ignore */ /** * @Author: hyy * @Date: 2026-06-16 13:37:12 *
@Description: 艺术家卡片 */
<template>
    <div
        class="cursor-pointer rounded-xl bg-surface px-2 py-3.5 text-center transition duration-200 hover:-translate-y-0.5 hover:bg-hover hover:shadow-[0_4px_24px_rgba(0,0,0,0.08)]">
        <div class="mx-auto mb-2.5 flex h-17 w-17 items-center justify-center rounded-[50%] text-[26px] font-bold text-white transition duration-200"
            :style="firstInfo.stryle">
            <span>
                {{ firstInfo.firstChar }}
            </span>
        </div>
        <span class="overflow-hidden text-sm font-semibold text-ellipsis whitespace-normal text-ink">
            {{ singer }}
        </span>

        <p class="mt-0.5 text-[11px] text-muted">{{ musicNum }}首</p>
    </div>
</template>

<script lang="ts" setup>
import { StringToGradient } from "@/utils/color";
import { GetFirstChar } from "@/utils/str";
import { computed } from "vue";

export interface ArtistCardProps {
    singer: string;
    musicNum: string | number;
}
const { singer, musicNum } = defineProps<ArtistCardProps>();

const firstInfo = computed(() => {
    let firstChar = GetFirstChar(singer)

    return {
        stryle: `background:${StringToGradient(firstChar)}`,
        firstChar,
    };
});
</script>
