/* prettier-ignore */
/**
* @Author: hyy
* @Date: 2026-07-21 14:42:52
* @Description: 文件路径
*/

<template>

    <EmptyData :icon-name="Icons.folder" size="small" v-if="!folderPathList.length">
        暂无文件路径，请添加 NAS 音乐目录
    </EmptyData>

    <ul>
        <li class="flex items-center  gap-1" v-for="item in folderPathList">
            <Icon :name="Icons.folder" class="text-brand"></Icon>
            <span class="flex-1 text-sub pl-4">{{ item.path }}</span>
            <ButtonIcon :icon-name="Icons.sideMyPlayListEdit" @click="onEditPath(item)"></ButtonIcon>
            <ButtonIcon :icon-name="Icons.delete" @click="delFolderPathListItem(item.id)"></ButtonIcon>

        </li>
    </ul>

    <div class="flex items-center gap-2.5 mt-3 " v-for="(item, index) in folderPaths">
        <input v-model="item.path" placeholder="输入文件路径，如 /music/nas/"
            class=" focus:border-brand bg-input h-9 w-full  pt-0 pr-10  pb-0 pl-4 border border-line outline-none rounded-3xl placeholder:text-sm placeholder:text-muted" />

        <ButtonIcon :icon-name="Icons.add" is-full button-title="添加" @click="addPath">

        </ButtonIcon>
        <ButtonIcon :icon-name="Icons.delete" is-full button-title="删除" @click="delPath(index)">

        </ButtonIcon>
    </div>


    <div class=" justify-center mt-2 gap-2 flex">
        <BasicButton @click="updatePath">更新</BasicButton>
        <BasicButton @click="scanMusic" type="danger">扫描音乐</BasicButton>


    </div>

</template>

<script setup lang="ts">
import SettingApi from '@/api/setting';
import { Icons } from '@/config/icons';
import { onBeforeMount, ref } from 'vue';
import type { SettingFolderPathResponse, SettingFolderPathWrite } from '../types';
import { useToast } from 'vue-toastification';

const folderPaths = ref<SettingFolderPathWrite[]>([{
    path: "",
    id: null
}])
const folderPathList = ref<SettingFolderPathResponse[]>([])

const toast = useToast()

const getFolderPathList = async () => {
    const data = await SettingApi.getFolderPathList()
    if (data && data.length) {

        folderPathList.value = data
    }
}

const delFolderPathListItem = async (id: number) => {
    try {

        await SettingApi.delFolderPathItem(id)
        await getFolderPathList()
    } catch (err) {
        toast.error(err)
    }

}
const delPath = (idx: number) => {

    folderPaths.value.splice(idx, 1)

}

const onEditPath = (item: SettingFolderPathResponse) => {
    if (folderPaths.value.some(f => f.id === item.id)) {
        return
    }

    let v = JSON.parse(JSON.stringify(item))

    folderPaths.value.push(v)
}

const addPath = () => {
    folderPaths.value.push({
        path: "",
        id: null
    })

}

const updatePath = async () => {

    try {
        clearInterval(timer)

        await SettingApi.addFolderPath(folderPaths.value)
        toast.success("添加成功")
        await getFolderPathList()

    } catch (err) {
        toast.error(err)

    }
}

let timer: number | undefined = undefined


// 扫描音乐
const scanMusic = async () => {

    try {

        await SettingApi.scanMusic()

        timer = setInterval(() => {
            getScanMusicStatus()
        }, 2000)

    } catch (err) {
        toast.error(err)
    }


}

// 获取扫描状态
const getScanMusicStatus = async () => {
    let status = await SettingApi.getScanMusicStatus()
    if (!status) {
        toast.success("音乐库更新成功")
        clearInterval(timer)
    }
}

onBeforeMount(() => {

    clearInterval(timer)

})

getFolderPathList()

</script>