<template>
    <button :class="[
        // 🌟 修复点：新增了 active:duration-75，让按下的瞬间立刻响应，不拖沓
        'inline-flex items-center justify-center rounded-md font-medium transition-all duration-200 active:duration-75 select-none border border-transparent outline-none cursor-pointer',

        // 点击态：按下时会有 3% 的微小缩放，手感极佳
        'active:scale-[0.97]',

        // 动态大小与颜色
        sizeClasses[size],
        typeClasses[type],

        // 其他状态
        {
            'w-full': block, // 是否撑满父容器
            'opacity-50 cursor-not-allowed! active:scale-100!': disabled // 禁用态（禁用时取消缩放）
        }
    ]" :disabled="disabled" @click="onClick">
        <!-- 左侧图标插槽 (可选) -->
        <slot name="icon"></slot>
        <!-- 按钮文字内容 -->
        <slot></slot>
    </button>
</template>

<script setup lang="ts">
// 你的 TS 类型定义保持不变
type SizeType = 'small' | 'default' | 'large'
type ButtonColorType = 'primary' | 'default' | 'danger' | 'ghost'

interface BasicButtonProps {
    size?: SizeType
    type?: ButtonColorType
    block?: boolean
    disabled?: boolean
}

const { size = "default", type = "primary", block = false, disabled = false } = defineProps<BasicButtonProps>()

const emit = defineEmits(['click'])

const onClick = () => {
    if (disabled) return
    emit('click')
}

// 📏 大小字典
const sizeClasses = {
    small: 'h-7 px-3 text-xs',
    default: 'h-9 px-4 text-sm',
    large: 'h-11 px-6 text-base'
}

// 🎨 颜色/样式字典
const typeClasses = {
    primary: 'bg-brand text-white hover:bg-[#d6576d] active:bg-[#c24c61]',
    default: 'bg-surface text-ink border-line hover:bg-gray-50 dark:hover:bg-hover active:bg-gray-100 dark:active:bg-hover',
    danger: 'bg-red-500 text-white hover:bg-red-600 active:bg-red-700',
    ghost: 'bg-transparent text-sub hover:bg-black/5 dark:hover:bg-white/5 active:bg-black/10 dark:active:bg-white/10'
}
</script>