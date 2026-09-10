
export interface LyricLine {
    time: number,
    text: string
}


export const ParseLrc = (lrcString: string): LyricLine[] => {
    // 按照换行符将纯文本分割成数组
    const lines = lrcString.split('\n')
    const result: LyricLine[] = []

    // 正则表达式匹配 [mm:ss.xx] 或 [mm:ss.xxx] 格式的时间戳
    const timeRegExp = /\[(\d{2,}):(\d{2})\.(\d{2,3})\]/

    for (const line of lines) {
        const match = timeRegExp.exec(line)
        if (!match) continue // 跳过没有时间戳的行（比如 [ti:歌曲名] 等元数据）

        // 提取分钟、秒、毫秒
        const min = parseInt(match[1], 10)
        const sec = parseInt(match[2], 10)
        // 兼容两位或三位数的毫秒
        const ms = match[3].length === 2 ? parseInt(match[3], 10) * 10 : parseInt(match[3], 10)

        // 计算当前行歌词对应的总秒数
        const time = min * 60 + sec + ms / 1000

        // 剔除时间戳，保留纯文本并去除首尾空格
        const text = line.replace(timeRegExp, '').trim()

        // 可选：过滤掉空行
        if (text) {
            result.push({ time, text })
        }
    }

    return result
}