<!-- components/Switch.vue -->
<template>
    <label class="relative inline-flex items-center cursor-pointer">
        <input type="checkbox" class="sr-only peer" v-model="model">
        <div class="w-12 h-6.5 bg-line rounded-full peer-checked:bg-brand transition-colors duration-300"></div>
        <div
            class="absolute left-0.5   top-0.75 bg-white w-5 h-5 rounded-full shadow-sm transition-transform duration-300 peer-checked:translate-x-6">
        </div>


    </label>
</template>

<script lang="ts" setup>


const model = defineModel<boolean>()


</script>

<!-- **Tailwind 原理核心点（敲黑板）：**
* `sr-only`: 这个类会把 input 在视觉上彻底隐藏，但屏幕阅读器依然能读到它，而且用键盘 `Tab` 键能正常选中它。
* `peer`: 给 input 打上标记。
* `peer-checked:xxx`: 写在**跟 input 同级的元素**（兄弟元素）或者它的**子元素**上。当被标记为 peer 的 input 选中时，触发这些样式。 -->