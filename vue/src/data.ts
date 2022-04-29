
export async function getData(): Promise<emdata[]> {
    if (data.length > 0) {
        return data;
    }
    let r = await fetch("https://emerald.xmdhs.top/data.json")
    let d: r = await r.json();
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