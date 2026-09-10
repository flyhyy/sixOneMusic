/**
* @Author: hyy
* @Date: 2026-06-09 11:35:33
* @Description: 头部右侧操作栏
*/


<template>
    <div class="flex gap-4 flex-1 justify-end">
        <button v-for="item in icons" :key="item.title" :title="getTitle(item)" @click="onClick(item)"
            class="text-center w-9 h-9 rounded-[50%] text-sub  hover:text-brand hover:bg-brand-soft flex items-center justify-center transition-all cursor-pointer">
            <Icon :name="getIcon(item)" class="text-2xl "></Icon>
        </button>
    </div>


</template>


<script lang="ts" setup>
import { Icons } from '@/config/icons';
import { useThemeStore } from '@/store/theme';
import { ref } from 'vue';

interface Icon {

    base: string
    active?: string,
    link?: string,
    title: string
    type?: 'theme' | 'link'

}

const themeStore = useThemeStore()

const icons = ref<Icon[]>([
    {
        base: Icons.light,
        active: Icons.dark,
        title: "明亮模式",
        type: 'theme'
    }, {
        base: Icons.github,
        title: '跳转到GitHub',
        link: '',
        type: 'link'
    }
])

const getIcon = (item: Icon) => {
    if (item.type === 'theme' && themeStore.isDark()) {
        return item.active || item.base
    }
    return item.base
}

const getTitle = (item: Icon) => {
    if (item.type === 'theme') {
        return themeStore.isDark() ? '暗黑模式' : '明亮模式'
    }
    return item.title
}

const onClick = (item: Icon) => {
    if (item.type === 'theme') {
        themeStore.toggleTheme()
    } else if (item.link) {
        window.open(item.link, '_blank')
    }
}

</script>
