<template>
    <p>总绿宝石使用数：{{ all }}</p>
    <p>
        <router-link to="/table30">一月内绿宝石使用排行</router-link>
    </p>
    <p>
        <router-link to="/table90">近三月绿宝石使用排行</router-link>
    </p>
    <p>
        <router-link to="/table365">近一年绿宝石使用排行</router-link>
    </p>
    <p>
        <router-link to="/all">总绿宝石使用排行</router-link>
    </p>
    <p>
        <router-link to="/list">绿宝石使用列表</router-link>
    </p>
    <p>
        <router-link to="/user">某个用户绿宝石使用详单</router-link>
    </p>
    <p>
        <a
            href="https://greasyfork.org/zh-CN/scripts/424437-%E6%9F%A5%E7%9C%8B%E8%AF%84%E5%88%86"
        >某个用户被评分列表</a>
    </p>
</template>

<script setup lang="ts">
import { ElNotification } from 'element-plus';
import { onMounted, ref } from 'vue';
import { RouterLink } from 'vue-router';
import { emdata, getData } from '../data';

let all = ref(0)

onMounted(async () => {
    let d: emdata[]
    try {
        d = await getData()
    } catch (e) {
        console.warn(e);
        ElNotification({
            title: 'Error',
            message: 'Could not load data',
            type: 'error',
            onClick: () => {
                location.reload();
            },
            duration: 0,
            showClose: false
        })
        return
    }
    let All = 0
    for (const v of d) {
        All += v.Count
    }
    all.value = All
    document.title = `绿宝石`
})

</script>

<style>
</style>
