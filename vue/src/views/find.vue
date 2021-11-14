<template>
    <h1>查询某个用户绿宝石使用详单</h1>
    <el-divider></el-divider>
    <div class="c">
        <el-input v-model="input" type="tel" placeholder="Please input" />
        <el-button :icon="Search" circle @click="click"></el-button>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { Search } from '@element-plus/icons'
import { useRouter } from 'vue-router'

let input = ref('');

const router = useRouter();
function click() {
    router.push({
        path: '/user/' + input.value
    })
}
</script>

<style scoped>
.c {
    display: flex;
    justify-content: center;
    align-items: center;
    max-width: max-content;
}
</style>