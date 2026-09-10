<template>
    <div class="flex items-center justify-between py-3  px-0   border-b border-b-line gap-4">
        <div class="text-sm text-sub flex items-center gap-2">
            <template v-if="iconName">
                <Icon :name="iconName"></Icon>
            </template>
            {{ label }}
        </div>

        <div>
            <slot></slot>
        </div>
    </div>

</template>

<script setup lang="ts">
import type { IconName } from '@/config/icons';


interface SettingRowProps {
    iconName?: IconName
    label: string

}
const { iconName, label } = defineProps<SettingRowProps>()
</script>