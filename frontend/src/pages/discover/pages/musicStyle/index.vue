<template>
    <div class="text-lg flex items-center gap-2 font-extrabold mt-6 text-ink">
        <Icon name="material-symbols:filter-list-rounded"></Icon>
        <span class="">
            古典 · 0 首
        </span>
    </div>
    <section>
        <EmptyData :icon-name="Icons.discoverStyleFilter">

            没有匹配的歌曲
        </EmptyData>

    </section>

</template>
<script lang="ts" setup>
import { Icons } from '@/config/icons';


</script>