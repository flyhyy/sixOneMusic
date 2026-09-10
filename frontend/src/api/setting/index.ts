import http from "@/http";
import type { SettingFolderPathResponse, SettingFolderPathWrite } from "@/pages/setting/types";

const SettingApi = {

    /**
     * 获取文件路径列表
     * @returns 
     */
    getFolderPathList: () => http.get<SettingFolderPathResponse[]>('/folder/query'),
    /**
     * 删除某个文件路径
     * @param id 
     * @returns 
     */
    delFolderPathItem: (id: number) => http.delete('/folder/del', { id }),

    /**
     * 添加文件路径
     * @param paths 
     * @returns 
     */
    addFolderPath: (paths: SettingFolderPathWrite[]) => http.post("/folder/write", paths),

    /**
     * 扫描音乐
     * @returns 
     */
    scanMusic: () => http.get("/scan/handler"),

    /**
     * 获取扫描音乐状态
     * @returns 
     */
    getScanMusicStatus: () => http.get<boolean>("/scan/status")
}

export default SettingApi