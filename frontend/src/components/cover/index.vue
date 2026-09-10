/* prettier-ignore */ /** * @Author: hyy * @Date: 2026-06-17 14:54:57 *
@Description: 封面 */
<template>
    <div class="relative flex aspect-square items-center justify-center overflow-hidden rounded-lg text-[28px] shadow-[0_2px_6px_rgba(0,0,0,.15)]"
        :style="{
            backgroundImage: coverSrc ? '' : StringToGradient(name),
        }">
        <DiscoverPlayButton></DiscoverPlayButton>

        <!-- 有封面图 -->
        <img v-if="coverSrc" :src="coverSrc" alt="" />

        <!-- 没有封面图 -->
        <Icon v-else :name="Icons.music" class="text-[#ffffff59]"></Icon>
    </div>
</template>

<script lang="ts" setup>
import { Icons } from "@/config/icons";
import { StringToGradient } from "@/utils/color";
import { computed } from "vue";
import DiscoverPlayButton from '@/pages/discover/components/DiscoverPlayButton/index.vue'

interface CoverProps {
    src: string;
    name: string;
}
const { src } = defineProps<CoverProps>();

const coverSrc = computed(() => {
    if (src != "") return src;
    return undefined;
});
</script>
