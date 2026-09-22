<template>
  <div class="flex items-center justify-between px-4 py-1">
    <span class="text-sm tracking-[1px] text-muted"> 我的歌单 </span>
    <button class="rounded text-base text-brand hover:cursor-pointer hover:bg-brand/12" title="新建歌单"
      @click="onCreatePlayList">
      <Icon :name="Icons.add"></Icon>
    </button>
  </div>

  <div class="min-h-0 flex-1 overflow-y-auto">
    <div v-if="showCreatePlayListInput" class="box-border flex h-10 items-center px-4 text-sm">
      <input type="text" placeholder="歌单名称"
        class="h-7 w-full rounded-lg border border-brand bg-canvas px-2 text-ink transition duration-200 outline-none focus:border-transparent focus:ring-1 focus:ring-brand"
        @blur="onBlurCreatePlayListInput" @keyup.enter="onKeyupEnterCreatePlayListInput" v-focus
        v-model.trim="playListName" />
    </div>

    <div
      class="group flex h-10 items-center gap-2 overflow-hidden px-4 text-sm whitespace-nowrap text-sub transition duration-200 ease-in hover:cursor-pointer hover:rounded-lg hover:bg-hover hover:text-ink"
      v-for="(item, index) in playList" :key="index" @click="onPlayListItem(item)">
      <input v-if="item.isEdit" type="text" :placeholder="item.name"
        class="h-7 w-full rounded-lg border border-brand bg-canvas px-2 text-ink transition duration-200 outline-none focus:ring-1 focus:ring-brand"
        @blur="onBlurItemEditInput(index)" v-focus v-model.trim="item.name" />

      <template v-if="!item.isEdit">
        <Icon :name="Icons.sideMyPlayList" class="shrink-0"></Icon>

        <span class="flex-1 truncate">{{ item.name }}</span>

        <div class="ml-auto flex shrink-0 gap-0.5 opacity-0 transition-opacity group-hover:opacity-100">
          <button title="重命名"
            class="flex h-6 w-6 items-center justify-center rounded-[50%] bg-transparent text-sm hover:cursor-pointer hover:bg-brand/6 hover:text-brand"
            @click.stop="onRestName(index)">
            <Icon :name="Icons.sideMyPlayListEdit"></Icon>
          </button>
          <button title="删除"
            class="flex h-6 w-6 items-center justify-center rounded-[50%] bg-transparent text-sm hover:cursor-pointer hover:bg-brand/6 hover:text-brand"
            @click.stop="onDelItemPlayList(item)">
            <Icon :name="Icons.delete"></Icon>
          </button>
        </div>
      </template>
    </div>
  </div>
</template>

<script setup lang="ts">
import { Icons } from "@/config/icons";
import { usePlayListStore } from "@/store/playList";
import { storeToRefs } from "pinia";
import { onMounted, ref } from "vue";
import { useRouter } from "vue-router";

export interface PlayListItemBase {
  name: string;
  id?: number;
}
export interface PlayListItem extends PlayListItemBase {
  isEdit: boolean;

}

const playListStore = usePlayListStore()

const { playList } = storeToRefs(playListStore)

const playListName = ref("");

const router = useRouter()
const showCreatePlayListInput = ref<boolean>(false);

// 每次进入（含退出后重新登录）都重新拉取歌单列表
onMounted(() => {
  playListStore.getPlayList()
})

const onPlayListItem = (item: PlayListItem) => {

  // playListStore.getSongPlayList(item.id || 0)

  router.push({
    path: `/music/songList`,
    query: {
      playListId: item.id || 0,
      type: 'PlayList',
      playListName: item.name
    }
  })
}

// 聚焦事件指令
const vFocus = {
  mounted: (el: HTMLInputElement) => {
    el.focus();
  },
};

// 创建歌单
const onCreatePlayList = async () => {
  showCreatePlayListInput.value = true;
};

// 监听输入框失焦事件
const onBlurCreatePlayListInput = () => {
  resetPlayListInputValue();
};
// 监听键盘enter事件
const onKeyupEnterCreatePlayListInput = () => {
  if (playListName.value) {
    addHandler(playListName.value)
  } else {
    alert("歌单名称不能为空");
  }
};



const addHandler = async (name: string) => {

  playListStore.addPlayList(name)

}





// 置空歌单输入框
const resetPlayListInputValue = () => {
  showCreatePlayListInput.value = false;
  playListName.value = "";
};

//歌单子项重命名
const onRestName = async (idx: number) => {
  playList.value[idx].isEdit = true;
};

// 歌单子项失焦处理
const onBlurItemEditInput = (idx: number) => {
  if (playList.value[idx].name) {
    playList.value[idx].isEdit = false;
    playListStore.updatePlayList(playList.value[idx])


  } else {
    alert("重命名不能为空");
  }
};
// 歌单删除
const onDelItemPlayList = (item: PlayListItem) => {
  if (item.id) {
    playListStore.delPlayList(item.id)
  }
};
</script>

<style></style>
