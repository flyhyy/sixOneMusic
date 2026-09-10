import http from "@/http";
import type { MusicBase, MusicListAllResponse } from "@/pages/publicSongList/index.vue";


const MusicApi = {

    getMusicListAll: (pageSize: number, page: number) => http.get<MusicListAllResponse>("/music/all", {
        pageSize,
        page
    }),
    musicPath: (id: number) => `/music/play/${id}`,

    collect: (song_id: number, is_collect: boolean) => http.put<number | string>("/music/collect", {
        song_id,
        is_collect
    }),

    getCollectSongs: () => http.get<MusicBase[]>("/music/collectSongs")

}

export default MusicApi