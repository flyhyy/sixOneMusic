<template>
  <IconifyIcon :icon="IconName" :class="className"></IconifyIcon>
</template>
<script lang="ts" setup>
import type { IconName } from "@/config/icons";
import { Icon as IconifyIcon } from "@iconify/vue";
import { computed } from "vue";

interface Prop {
  name: IconName | string;
  perifex?: string;
  size?: SizeType
}

const { name, perifex, size = "default" } = defineProps<Prop>();

const IconName = computed(() => {
  if (/(:)/.test(name)) {
    return name;
  }

  return `${perifex}${name}`;
});

const Size = {
  "default": "w-4 h-4",
  "large": "w-5 h-5",
  "small": "w-2 h-2"
}
const className = computed(() => {

  return Size[size]

})
</script>
