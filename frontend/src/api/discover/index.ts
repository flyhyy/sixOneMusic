import http from "@/http";
import type { AblumListRespose, SingerListResponse, StyleListResponse } from "@/pages/discover/pages/types";
import type { MusicBase } from "@/pages/publicSongList/index.vue";



export const DiscoverApi = {
    //  获取风格列表
    GetStyleList: () => http.get<StyleListResponse[]>("/style/list"),
    // 获取艺术家
    GetSingerList: () => http.get<SingerListResponse[]>("/singer/list"),
    // 获取专辑列表
    GetAlbumList: () => http.get<AblumListRespose[]>("/album/list"),
    // 获取艺术家对应的歌曲列表
    GetSingerSongList: (singerName: string) => http.get<MusicBase[]>(`/singer/song/${singerName}`),
    // 获取专辑对应的歌曲列表
    GetAlbumSongList: (albumName: string) => http.get<MusicBase[]>(`/album/song/${albumName}`)

}

