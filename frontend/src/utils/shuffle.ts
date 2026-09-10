// 经典的 Fisher-Yates 洗牌算法，返回一个新的打乱数组
// 洗牌算法 
export function shuffleArray<T>(arr: T[]): T[] {

    const result = [...arr]

    for (let i = result.length - 1; i > 0; i--) {
        const j = Math.floor(Math.random() * (i + 1));
        [result[i], result[j]] = [result[j], result[i]]; // 交换位置
    }
    return result
}

