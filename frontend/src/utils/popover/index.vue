<template>
    <Teleport to="body">
        <Transition enter-active-class="transition-all duration-200 ease-out"
            enter-from-class="opacity-0 scale-95 translate-y-2" enter-to-class="opacity-100 scale-100 translate-y-0"
            leave-active-class="transition-all duration-150 ease-in"
            leave-from-class="opacity-100 scale-100 translate-y-0" leave-to-class="opacity-0 scale-95 translate-y-2">

            <div ref="PopoverContainerRef"
                class=" absolute z-50 bg-surface rounded-xl shadow-[0_12px_48px_rgba(0,0,0,0.5)] p-2 min-w-50 max-w-65 max-h-75 overflow-y-auto "
                v-show="visible" :style="{ top: y + 'px', left: x + 'px' }" @click.stop>
                <component :is="dynamicComponent" v-bind="dynamicProps"></component>

            </div>
        </Transition>
    </Teleport>

</template>

<script lang="ts" setup>
import { nextTick, ref, useTemplateRef, onUnmounted, shallowRef, type Component } from "vue";

const PopoverContainerRef = useTemplateRef<HTMLElement>("PopoverContainerRef");
const visible = ref<boolean>(false);
const x = ref<number>(0);
const y = ref<number>(0);

const dynamicComponent = shallowRef<Component | null>(null)
const dynamicProps = ref<Record<string, any>>({})



// ================= 新增：记录当前绑定的触发节点 =================
const currentTarget = ref<HTMLElement | null>(null);

const hide = () => {
    visible.value = false;
    currentTarget.value = null; // 清空当前节点

    setTimeout(() => {
        dynamicComponent.value = null
        dynamicProps.value = {}
    }, 200);

    document.removeEventListener("click", onClickOutside);
};

const onClickOutside = (event: MouseEvent) => {
    const clickTarget = event.target as Node;
    // 核心判断：点击不能在 Popover 内部，且不能在当前的触发按钮内部
    if (
        PopoverContainerRef.value &&
        !PopoverContainerRef.value.contains(clickTarget) &&
        !(currentTarget.value && currentTarget.value.contains(clickTarget))
    ) {
        hide();
    }
};

const show = async (targetEl: HTMLElement, com: Component, props?: Record<string, any>) => {
    if (!targetEl) return;

    // 1. Toggle 逻辑：如果点击的是同一个按钮，且弹窗开着，则执行关闭
    if (visible.value && currentTarget.value === targetEl) {
        hide();
        return;
    }

    dynamicComponent.value = com
    if (props) {

        dynamicProps.value = props
    }


    // 2. 立即移除旧的全局监听，防止上一个弹窗的冒泡事件误杀本次打开
    document.removeEventListener("click", onClickOutside);

    // 3. 记录新的触发节点
    currentTarget.value = targetEl;
    visible.value = true;

    await nextTick();

    if (PopoverContainerRef.value) {
        const targetRect = targetEl.getBoundingClientRect();
        const popoverRect = PopoverContainerRef.value.getBoundingClientRect();

        const viewportWidth = window.innerWidth;
        const viewportHeight = window.innerHeight;
        const offset = 8;

        // X 轴智能计算
        let calcX = targetRect.left - popoverRect.width - offset;
        if (calcX < 0) {
            calcX = targetRect.right + offset;
            if (calcX + popoverRect.width > viewportWidth) {
                calcX = viewportWidth - popoverRect.width - offset;
            }
        }

        // Y 轴智能计算
        let calcY = targetRect.top;
        if (calcY + popoverRect.height > viewportHeight) {
            calcY = targetRect.bottom - popoverRect.height;
            if (calcY < 0) {
                calcY = offset;
            }
        }

        x.value = calcX;
        y.value = calcY;

        // 4. 延迟挂载新的全局监听，完美避开本次点击的冒泡
        setTimeout(() => {
            document.addEventListener("click", onClickOutside);
        }, 0);
    }
};

onUnmounted(() => {
    document.removeEventListener("click", onClickOutside);
});

defineExpose({
    show,
    hide
});
</script>

<style scoped></style>
