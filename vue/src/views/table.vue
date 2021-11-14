<template>
    <h1>{{ title }}</h1>
    <el-table :data="list" :default-sort="{ prop: 'count', order: 'ascending' }">
        <el-table-column type="index" label="排名" :index="indexMethod" />
        <el-table-column prop="uid" label="uid" />
        <el-table-column prop="name" label="用户名" />
        <el-table-column prop="count" label="总数" sortable />
        <el-table-column label="详情">
            <template #default="props">
                <router-link v-if="props.row.href" :to="props.row.href">{{ props.row.v }}</router-link>
                <span v-else>{{ props.row.v }}</span>
            </template>
        </el-table-column>
    </el-table>
</template>

<script setup lang="ts">
import { onMounted, ref, toRefs, watchEffect } from 'vue';
import { getData } from '../data';
import { ElTable, ElTableColumn } from 'element-plus';

const props = defineProps({
    day: {
        type: Number,
        default: 30
    }
});
let p = toRefs(props)

let title = ref('');

interface tdItem {
    uid: number;
    name: string;
    count: number;
    href: string;
    v: string;
}


let list = ref([] as tdItem[]);


onMounted(() => {
    watchEffect(async () => {
        if (p.day.value == 0) {
            title.value = `总绿宝石使用排行`;
            document.title = title.value;
        } else {
            title.value = `${p.day.value} 天内绿宝石使用排行`;
            document.title = title.value;
        }
        list.value = []

        let l = await getData()
        let utime = new Date().getTime() / 1000;
        let d = p.day.value * 24 * 3600;
        let tl: tdItem[] = []
        let m: { [key: string]: tdItem } = {}
        for (const v of l) {
            if (utime - Number(v.Time) < d || p.day.value == 0) {
                let uid = String(v.Uid);
                if (m[uid]) {
                    m[uid].count += Number(v.Count);
                } else {
                    m[uid] = {
                        uid: Number(v.Uid),
                        name: v.Username,
                        count: Number(v.Count),
                        href: `/user/${v.Uid}`,
                        v: `详情`
                    }
                }
            }
        }
        for (const k in m) {
            tl.push(m[k])
        }
        list.value = tl;
    })
})

function indexMethod(index: number) {
    return index + 1;
}

</script>

<style>
</style>
