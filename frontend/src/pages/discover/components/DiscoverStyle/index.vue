/* prettier-ignore */
/**
* @Author: hyy
* @Date: 2026-06-12 13:26:14
* @Description: 发现音乐-风格
*/
<template>
    <div class="flex items-center text-[11px] text-muted tracking-[1px] gap-0.5 mt-6 mb-2.5">

        <Icon :name="Icons.discvoerStyle"></Icon>
        <span>
            风格
        </span>
    </div>
    <!-- 曲风数据 -->
    <div>
        <StyleItemButton v-for="item in styleList" :icon-name="StyleItemsConfig.Classical.iconName"
            @click="onToStyleData">
            {{ item.name }}
        </StyleItemButton>
    </div>



    <!-- 曲风歌曲列表展示 -->
    <section>
        <RouterView></RouterView>
    </section>



</template>
<script lang="ts" setup>

import { Icons } from '@/config/icons';
import StyleItemButton from './components/styleItemButton.vue'
import { useRouter } from 'vue-router';
import type { StyleListResponse } from '../../pages/types.ts';
import { DiscoverApi } from '@/api/discover/index.ts';
import { ref } from 'vue';

const router = useRouter()

const StyleItemsConfig = {

    // 古典音乐
    Classical: { label: '古典', iconName: Icons.discvoerStyleClassical },
    // 摇滚乐
    Rock: { label: '摇滚', iconName: Icons.discvoerStyleRock },
    // 民谣音乐
    Folk: { label: '民谣', iconName: Icons.discvoerStyleFolk },
    //电子音乐
    EDM: { label: '电子', iconName: Icons.discvoerStyleEDM },
    // 嘻哈/说唱
    Hip_Hop: { label: '嘻哈/说唱', iconName: Icons.discvoerStyleHip_Hop },
    // 爵士乐
    Jazz: { label: '爵士乐', iconName: Icons.discvoerStyleJazz },
    // 打碟/混音
    DJ_Remix: { label: '打碟/混音', iconName: Icons.discvoerStyleDJ_Remix },
    // 流行乐
    Pop: { label: '流行乐', iconName: Icons.discvoerStylePop },
    // 轻音乐
    Lo_Fi: { label: '轻音乐', iconName: Icons.discvoerStyleLo_Fi },
    // 复古/经典
    Retro: { label: '复古/经典', iconName: Icons.discvoerStyleRetro },


}

const styleList = ref<StyleListResponse[]>([])

//  曲风页面跳转
const onToStyleData = () => {

    router.push({

    })

}

const getStyleList = async () => {
    const data = await DiscoverApi.GetStyleList()
    styleList.value = data
}
getStyleList()

</script>
