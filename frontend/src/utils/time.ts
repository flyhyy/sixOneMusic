// 格式化当前秒为 mm:ss
export const formatTime = (time: number) => {
    if (!time || isNaN(time)) return '00:00'
    const mins = Math.floor(time / 60)
    const secs = Math.floor(time % 60)
    return `${mins.toString().padStart(2, '0')}:${secs.toString().padStart(2, '0')}`
}
