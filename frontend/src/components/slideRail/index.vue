/* prettier-ignore */
/**
* @Author: hyy
* @Date: 2026-07-28 15:21:24
* @Description: 滑轨拖动组件
*/


<template>
    <div class="flex-1 h-1 bg-line rounded-full cursor-pointer relative hover:h-1.5 group transition-all duration-50"
        @mousedown.prevent="handleMouseDown" ref="railRef">

        <!-- 实际进度 (这里用 style 模拟动态宽度，Vue 中使用 :style="{ width: progress + '%' }") -->
        <div class="h-full   rounded-full relative " :style="progressStyle" :class="modelClass.wrap">

            <!-- 拖拽滑块 (小圆球) -->
            <div class="absolute -right-1.5 top-1/2 -translate-y-1/2 w-3 h-3 rounded-full shadow-[0_1px_4px_rgba(0,0,0,0.3)]  transition-opacity duration-100"
                :class="modelClass.ball">
            </div>

        </div>
    </div>

</template>
<script setup lang="ts">
import { computed, onUnmounted, useTemplateRef } from 'vue';


interface SlideRailProps {
    // 模式  播放|音量
    type?: "play" | "volume"
}

const model = defineModel<number>({ default: 0 })

const { type = "play" } = defineProps<SlideRailProps>()

const progressStyle = computed(() => {

    return `width:${model.value > 100 ? 100 : model.value < 0 ? 0 : model.value}%`


})

const railRef = useTemplateRef("railRef")
export interface EmitDragEndInfo {
    value: number
}
export interface EmitDragStartInfo {
    value: number
}
const emits = defineEmits<{
    "drag-start": [info: EmitDragStartInfo]
    "drag-end": [info: EmitDragEndInfo]
}>()

const modelClass = computed(() => {

    if (type === "play") {

        return {
            wrap: "bg-brand  ",
            ball: "bg-brand  opacity-0 group-hover:opacity-100 "
        }

    }
    return {
        wrap: "bg-sub  group-hover:bg-brand",
        ball: "bg-white  opacity-100 group-hozzver:bg-brand"
    }

})

const updateVal = (e: MouseEvent) => {

    const slideRailVal = railRef.value
    if (!slideRailVal) {
        return
    }
    const r = slideRailVal.getBoundingClientRect()
    const ratio = Math.max(0, Math.min(1, (e.clientX - r.left) / r.width))

    model.value = ratio * 100
    emits("drag-start", { value: model.value })
}




const stop = () => {
    emits("drag-end", { value: model.value })

    document.removeEventListener("mousemove", updateVal)
    document.removeEventListener("mouseup", stop)
}

onUnmounted(() => {
    stop()
})

const handleMouseDown = (e: MouseEvent) => {

    updateVal(e)

    document.addEventListener("mousemove", updateVal)
    document.addEventListener("mouseup", stop)


}


</script>