import http from "@/http";
import type { PlayListItemBase } from "@/layout/components/sideMyPlayList/index.vue";
import type { MusicBase } from "@/pages/publicSongList/index.vue";




export const PlayListApi = {

    Add: (name: string) => http.post("/playList/add", { name }),

    list: () => http.get<PlayListItemBase[]>("/playList/queryList"),

    Update: (data: PlayListItemBase) => http.put("/playList/update", data),

    Del: (id: number) => http.delete(`/playList/del/${id}`),

    addSong: (playListId: number, songId: number) => http.post("/playList/addSong", {
        playListId,
        songId
    }),
    getSongs: (playListId: number) => http.get<MusicBase[]>(`/playList/getSongs/${playListId}`)


}