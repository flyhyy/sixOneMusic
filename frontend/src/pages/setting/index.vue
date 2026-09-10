/* prettier-ignore */
/**
* @Author: hyy
* @Date: 2026-06-18 14:57:20
* @Description: 设置
*/

<template>
    <TitleInfo title="设置" :icon-name="Icons.sideSetting" size="large"></TitleInfo>

    <!-- <SettingCard title="播放设置" :icon-name="Icons.sideSetting" class="mt-6">
        <SettingRow label="音频淡入淡出" :icon-name="Icons.speak">
            <Switch v-model="t"></Switch>
        </SettingRow>
    </SettingCard> -->

    <SettingCard title="文件路径" :icon-name="Icons.folderLine" class="mt-6">
        <FolderPathList></FolderPathList>
    </SettingCard>
    <SettingCard title="更新说明" :icon-name="Icons.information">

    </SettingCard>

    <SettingCard title="关于" :icon-name="Icons.information">
        <SettingRow label="版本">
            <Text>{{ version }}</Text>
        </SettingRow>
    </SettingCard>

    <div class="flex justify-center">

        <BasicButton>退出登陆</BasicButton>
    </div>

</template>


<script lang="ts" setup>
import { Icons } from '@/config/icons';
import SettingCard from './components/card.vue'
import SettingRow from './components/row.vue'
import FolderPathList from './components/folderPathList.vue'
import { ref } from 'vue';

const version = ref(__APP_VERSION__)
</script>
