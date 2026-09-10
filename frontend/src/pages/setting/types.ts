
export interface SettingFolderPathResponse {
    id: number
    path: string

}

export interface SettingFolderPathWrite {
    path: string,
    id: number | null
}