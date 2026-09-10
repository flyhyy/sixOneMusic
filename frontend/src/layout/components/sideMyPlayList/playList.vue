<template>
    <div class="text-xs text-muted uppercase tracking-wider pt-1.5 px-3 pb-2">添加到歌单</div>
    <div class="flex items-center justify-between px-3 py-2 rounded-lg cursor-pointer text-[13px] text-sub gap-2 transition  hover:bg-hover hover:text-ink"
        v-for="pl in playList" :key="pl.id" @click="addToPlayList(pl?.id || 0)">
        <span class="overflow-hidden text-ellipsis whitespace-nowrap">
            {{ pl.name }}
        </span>

        <!-- <span class="text-[11px] text-[#8E90A6] shrink-0">{{ pl.songs.length }}首</span> -->
    </div>
    <!-- <div class="border-t border-[#E0E0EA] mt-1 pt-1" v-if="!userPlaylists.length">
                    <div class="flex items-center justify-between px-3 py-2 rounded-lg cursor-pointer text-[13px] text-[#E8637B] gap-2 transition hover:bg-[##EDEDF3]"
                        @click="startCreatePlaylist(); addToPlaylistTarget = null">+ 新建歌单</div>
                </div> -->

</template>

<script setup lang="ts">
import type { MusicBase } from "@/pages/publicSongList/index.vue";
import { usePlayListStore } from "@/store/playList";
import { storeToRefs } from "pinia";
import { useToast } from "vue-toastification";



const playListStore = usePlayListStore()
const toast = useToast()
const { id } = defineProps<MusicBase>()


const { playList } = storeToRefs(playListStore)


const addToPlayList = async (playListId: number) => {

    try {

        await playListStore.addSongPlayList(playListId, id)
    } catch (err) {
        toast.error(err)
    }
}


</script>