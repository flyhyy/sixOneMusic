import { createApp } from 'vue'
import "vue-toastification/dist/index.css";
import './style.css'

import App from './App.vue'
import Router from './router/index.ts'

import { createPinia } from 'pinia'
import Toast, { type PluginOptions } from "vue-toastification";

// 1. 引入 Iconify 的 API 核心方法
import { addCollection, type IconifyJSON } from '@iconify/vue';
// 2. 引入你刚刚下载的离线图标 JSON 数据
import riIcons from '@iconify-json/ri/icons.json';
import mdiIcons from '@iconify-json/mdi/icons.json';
import materialSymbolsIcons from '@iconify-json/material-symbols/icons.json';
import systemUiconsIcons from '@iconify-json/system-uicons/icons.json';
// 3. 将它们注册到本地内存中
addCollection(riIcons);
addCollection(mdiIcons);
addCollection(materialSymbolsIcons as IconifyJSON);
addCollection(systemUiconsIcons);

const pinia = createPinia()
const options: PluginOptions = {

}


createApp(App).use(Router).use(pinia).use(Toast, options).mount('#app')

