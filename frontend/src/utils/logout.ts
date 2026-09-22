/* prettier-ignore */
/**
 * @Author: hyy
 * @Date: 2026-09-10
 * @Description: 退出登录
 */
import { clearToken } from '@/storage/token'
import router from '@/router'
import { useMusicStore } from '@/store/music'
import { useAudioStore } from '@/store/audio'
import { usePlayListStore } from '@/store/playList'

/**
 * 退出登录：清除本地 token、清空会话数据并跳转回登录页。
 * 使用 vue-router 进行 SPA 跳转（replace 避免退出前的页面残留在历史记录中）。
 */
export function logout() {
    clearToken()
    useMusicStore().reset()
    useAudioStore().reset()
    usePlayListStore().reset()
    router.replace('/')
}
