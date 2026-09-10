<template>

    <div class=" font-extrabold flex items-center gap-2 text-ink" :class="sizeCtr">
        <Icon :name="iconName"></Icon>
        <span>

            <slot> </slot>
        </span>
    </div>

</template>

<script lang="ts" setup>
import type { IconName } from '@/config/icons';
import { computed } from 'vue';

export type SizeType = 'large' | 'default' | 'small'

interface ItemIndexProps {
    iconName: IconName
    size?: SizeType
}
const { iconName, size = 'default' } = defineProps<ItemIndexProps>()

const sizeCtr = computed(() => {

    switch (size) {

        case 'large':
            return 'text-[28px]';
        case 'small':
            return 'text-[14px]';
        default:
            return 'text-[18px]'

    }

})


</script>