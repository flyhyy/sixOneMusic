/* prettier-ignore */
/**
 * @Author: hyy
 * @Date: 2026-09-07
 * @Description: 明暗主题 store
 */
import { defineStore } from 'pinia';
import { ref } from 'vue';

const THEME_KEY = 'theme';

/**
 * 初始化主题：优先读本地记录，否则默认浅色
 */
function getInitialTheme(): ThemeType {
    const stored = localStorage.getItem(THEME_KEY);
    if (stored === 'dark' || stored === 'light') {
        return stored;
    }
    return 'light';
}

export const useThemeStore = defineStore('theme', () => {
    const theme = ref<ThemeType>(getInitialTheme());

    const isDark = () => theme.value === 'dark';

    function apply() {
        document.documentElement.classList.toggle('dark', theme.value === 'dark');
        localStorage.setItem(THEME_KEY, theme.value);
    }

    function toggleTheme() {
        theme.value = theme.value === 'dark' ? 'light' : 'dark';
        apply();
    }

    apply();

    return {
        theme,
        isDark,
        toggleTheme
    };
});
