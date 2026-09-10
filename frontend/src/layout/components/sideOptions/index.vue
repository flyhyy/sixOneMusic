/** * @Author: hyy * @Date: 2026-06-09 20:01:47 * @Description:
侧边栏选项控制列表 */

<template>
  <ul class="flex flex-col gap-1">
    <li :class="[itemBaseClassName, currentPath === item.routerLink ? itemActiveClassName : itemNormalClassName]"
      v-for="item in optionsConfigList" @click="onItemClick(item)">
      <Icon :name="item.icon" class=""></Icon>

      <span>
        {{ item.text }}
      </span>

      <span v-if="item.num != null" class="ml-auto min-w-6 rounded-xl px-2 py-0.5 text-center text-sm transition-colors"
        :class="[badgeBaseClassName, currentPath === item.routerLink ? badgeActiveClassName : badgeNormalClassName]">
        {{ item.num }}
      </span>
    </li>
  </ul>
</template>

<script setup lang="ts">
import { Icons, type IconName } from "@/config/icons";
import { ref, watch } from "vue";
import { useRoute, useRouter } from "vue-router";

// 基础类名
const itemBaseClassName =
  "whitespace-normal flex items-center py-3 px-4 gap-3 hover:cursor-pointer text-base rounded-lg transition duration-200 ease-in";
// 默认样式
const itemNormalClassName =
  "hover:bg-hover hover:text-ink text-sub ";
// 激活样式
const itemActiveClassName = "bg-brand/12 text-brand ";

//基础类名
const badgeBaseClassName = " ml-auto min-w-6 text-sm ";
// 默认样式
const badgeNormalClassName = "bg-hover text-muted";
// 激活样式
const badgeActiveClassName = "bg-brand text-white ";


const route = useRoute()
const router = useRouter()

interface OptionConfigList {
  icon: IconName;
  text: string;
  routerLink: string;
  routerQuery?: Record<string, any>
  num?: string | null | number;
}



const currentPath = ref(route.path)

const onItemClick = (item: OptionConfigList) => {
  router.push({
    path: item.routerLink,
    query: item.routerQuery || {}
  })
}
watch(() => route.path, (nv) => {
  currentPath.value = nv
})


// 侧边栏选项控制列表
const optionsConfigList: OptionConfigList[] = [

  {
    icon: Icons.sideHome,
    text: "发现音乐",
    routerLink: "/music/discover",
    num: null,
  },
  {
    icon: Icons.sidePlayQueue,
    text: "播放队列",
    routerLink: "/music/playQueue",
    routerQuery: {
      type: "PlayQueue"
    },
    num: null,
  },
  {
    icon: Icons.sideHome,
    text: "全部音乐",
    routerLink: "/music/allMusic",
    routerQuery: {
      type: "AllMusic"
    },
    num: null,
  },
  {
    icon: Icons.sideMyCollection,
    text: "我的收藏",
    routerLink: "/music/myCollection",
    num: 0,
    routerQuery: {
      type: "MyCollect"
    },
  },
  // {
  //   icon: Icons.sidePersonalCenter,
  //   text: "个人中心",
  //   routerLink: "/music/personalCenter",
  //   num: 0,
  // },
  {
    icon: Icons.sideSetting,
    text: "设置",
    routerLink: "/music/setting",
    num: null,
  },
];
</script>
