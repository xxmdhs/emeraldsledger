<template>
    <el-container>
        <el-header class="header">
            <el-menu mode="horizontal" router>
                <el-menu-item index="1" :route="{path:'/'}">Home</el-menu-item>
                <el-menu-item index="2" :route="{path:'/user'}">搜索</el-menu-item>
            </el-menu>
        </el-header>
        <el-main class="main">
            <router-view></router-view>
        </el-main>
    </el-container>
</template>

<script setup lang="ts">
import { RouterView } from 'vue-router';
import { ElContainer } from 'element-plus';

</script>

<style scoped>
.main {
    font-size: 16px;
    font-family: "Helvetica Neue", Helvetica, "PingFang SC", "Hiragino Sans GB",
        "Microsoft YaHei", "微软雅黑", Arial, sans-serif;
    padding-right: 30px;
    padding-left: 30px;
}
.header {
    background-color: #fff;
    font-size: 16px;
    font-family: "Helvetica Neue", Helvetica, "PingFang SC", "Hiragino Sans GB",
        "Microsoft YaHei", "微软雅黑", Arial, sans-serif;
    border-bottom: 1px solid #dcdfe6;
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