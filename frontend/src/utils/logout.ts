/* prettier-ignore */
/**
 * @Author: hyy
 * @Date: 2026-09-10
 * @Description: 退出登录
 */
import { clearToken } from '@/storage/token'

/**
 * 退出登录：清除本地 token 并跳转回登录页。
 * 使用整页跳转会同时清空内存中的 pinia 状态（音乐/播放列表/播放器等）。
 */
export function logout() {
    clearToken()
    window.location.replace('/')
}
