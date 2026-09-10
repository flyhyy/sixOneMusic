/* prettier-ignore */
/**
* @Author: hyy
* @Date: 2026-06-17 15:05:59
* @Description: 歌曲信息
*/
<template>
    <div class="flex flex-col">

        <div class="overflow-hidden text-[14px] font-semibold text-ellipsis whitespace-normal text-ink">
            {{ primaryText }}
        </div>
        <div class=" overflow-hidden text-[12px] text-ellipsis whitespace-normal text-muted mt-0.5">
            {{ secondaryText }}
        </div>
    </div>

</template>

<script lang="ts" setup>

export interface SongInfoProps {
    // 主要
    primaryText: string
    // 次要
    secondaryText: string
}
const { primaryText, secondaryText } = defineProps<SongInfoProps>()

</script>
