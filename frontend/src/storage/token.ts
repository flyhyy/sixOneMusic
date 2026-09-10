
/* prettier-ignore */
/**
 * @Author: hyy
 * @Date: 2026-07-20 16:53:49
 * @Description: token 操作
 */
// 提取常量，方便统一维护
const TOKEN_KEY = 'token'

/**
 * 设置（保存）token
 * @param token 
 */
export const setToken = (token: string) => {

    if (token) {

        localStorage.setItem(TOKEN_KEY, token)

    }
}

/**
 * 获取token
 * @returns 
 */
export const getToken = () => {
    let token = localStorage.getItem(TOKEN_KEY)

    return token || null

}

/**
 * 移除token
 */
export const clearToken = () => {

    localStorage.removeItem(TOKEN_KEY)
}

