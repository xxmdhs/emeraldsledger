<template>
    <el-container>
        <el-header class="header">
            <div class="main">
                <el-menu mode="horizontal" router :default-active="'1'">
                    <el-menu-item index="1" :route="{ path: '/' }">Home</el-menu-item>
                    <el-menu-item index="2" :route="{ path: '/user' }">搜索</el-menu-item>
                </el-menu>
            </div>
        </el-header>
        <el-main>
            <div class="main">
                <router-view></router-view>
            </div>
        </el-main>
    </el-container>
</template>

<script setup lang="ts">
import { RouterView } from 'vue-router';
import { ElContainer } from 'element-plus';

</script>

<style scoped>
.main {
    font-family: "Helvetica Neue", Helvetica, "PingFang SC", "Hiragino Sans GB",
        "Microsoft YaHei", "微软雅黑", Arial, sans-serif;
    max-width: 1024px;
    margin-right: auto;
    margin-left: auto;
}
.header {
    background-color: #fff;
    font-family: "Helvetica Neue", Helvetica, "PingFang SC", "Hiragino Sans GB",
        "Microsoft YaHei", "微软雅黑", Arial, sans-serif;
}
</style>

<style>
a {
    color: #0366d6;
    text-decoration: none;
}

a:hover {
    text-decoration: underline;
}
</style>