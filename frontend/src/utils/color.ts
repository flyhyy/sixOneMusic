/**
 * 将字符串转换为一段美观的 CSS 线性渐变代码
 * @param str 输入的字符串 (如用户名、标签名)
 * @param s 饱和度 (默认 75%，推荐 60%-80% 保证色彩鲜艳)
 * @param l 亮度 (默认 60%，推荐 50%-70% 保证文字易读)
 * @returns 完整的 CSS background-image 字符串
 */
export const StringToGradient = (str: string, s: number = 75, l: number = 60): string => {
    if (!str) return 'linear-gradient(135deg, #e2e8f0, #cbd5e1)'; // 空字符串兜底色

    // 1. 简易且高效的 DJB2 Hash 算法
    let hash = 0;
    for (let i = 0; i < str.length; i++) {
        // hash * 33 + charCode
        hash = str.charCodeAt(i) + ((hash << 5) - hash);
    }

    // 防止 hash 为负数
    hash = Math.abs(hash);

    // 2. 将 Hash 值映射为 0-360 的色相 (Hue)
    const h1 = hash % 360;

    // 3. 计算渐变的第二个色相，偏移 50 度左右通常视觉上最和谐 (相邻色)
    const h2 = (h1 + 50) % 360;

    // 4. 组装 HSL 颜色
    const color1 = `hsl(${h1}, ${s}%, ${l}%)`;
    const color2 = `hsl(${h2}, ${s}%, ${l}%)`;

    // 5. 返回渐变 CSS
    return `linear-gradient(135deg, ${color1}, ${color2})`;
};