import { ElNotification } from 'element-plus'

export async function getData(): Promise<emdata[]> {
    if (data.length > 0) {
        return data;
    }
    let d: r
    try {
        let r = await fetch("https://cdn.jsdelivr.net/gh/xxmdhs/emeraldsledger@master/data.json")
        d = await r.json();
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
        return [];
    }
    let l: emdata[] = []
    for (const k in d) {
        l.push(d[k]);
    }
    data = l;
    return l;
}

type r = { [key: string]: emdata }

let data: emdata[] = []

export interface emdata {
    Uid: string;
    Username: string;
    Count: number;
    Time: number;
    Cause: string;
    Type: string;
    Link: string;
}