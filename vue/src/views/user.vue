<template>
    <p v-if="uid != 0">
        {{ username }}({{ uid }})
        <br />
        总数： {{ count }}
    </p>
    <el-table :data="list">
        <el-table-column v-if="uid == 0" prop="Uid" label="uid">
            <template #default="{ row }">
                <router-link :to="'/user/' + row.Uid">{{ row.Uid }}</router-link>
            </template>
        </el-table-column>
        <el-table-column v-if="uid == 0" prop="Username" label="用户名" />
        <el-table-column prop="Count" label="绿宝石数" sortable />
        <el-table-column prop="time" label="时间" sortable />
        <el-table-column prop="Cause" label="原因">
            <template #default="{ row }">
                <span v-html="row.Cause"></span>
            </template>
        </el-table-column>
        <el-table-column label="链接">
            <template #default="props">
                <a
                    v-if="props.row.Link"
                    :href="props.row.Link"
                    target="_blank"
                    referrerpolicy="no-referrer"
                >{{ props.row.v }}</a>
                <span v-else>{{ props.row.v }}</span>
            </template>
        </el-table-column>
    </el-table>
</template>

<script setup lang="ts">
import { onMounted, ref, watchEffect } from 'vue';
import { getData, emdata } from '../data';
import { RouterLink } from 'vue-router';

let username = ref('');
let count = ref(0);

interface td extends emdata {
    v: string;
    time: string;
}

let list = ref([] as td[]);

const props = defineProps({
    uid: {
        type: Number,
        required: true,
    }
});

onMounted(() => {
    watchEffect(() => {
        count.value = 0
        getUser(props.uid);
    })
})

async function getUser(uid: number) {
    let d = await getData()
    let l = [] as emdata[]
    list.value = []

    for (const v in d) {
        if (d[v].Uid == String(uid) || uid == 0) {
            username.value = d[v].Username;
            count.value += d[v].Count;
            l.push(d[v])
        }
    }
    if (uid == 0) {
        document.title = '绿宝石列表'
    } else {
        document.title = `${username.value} 的绿宝石使用列表`
    }

    l.sort((a, b) => {
        return b.Time - a.Time
    })
    for (const v of l) {
        let vv = v as td
        if (vv.Link == "" && vv.Type == "mcbbsAd") {
            vv.v = "none（宣传栏）"
        } else {
            vv.v = "link"
        }
        vv.time = transformTime(vv.Time)
        list.value.push(vv)
    }
}

function transformTime(timestamp: number) {
    var time = new Date(timestamp * 1000);
    var y = time.getFullYear();
    var M = time.getMonth() + 1;
    var d = time.getDate();
    var h = time.getHours();
    var m = time.getMinutes();
    return y + '-' + addZero(M) + '-' + addZero(d) + ' ' + addZero(h) + ':' + addZero(m)
}
function addZero(m: number): string {
    return m < 10 ? '0' + m : String(m);
}
</script>