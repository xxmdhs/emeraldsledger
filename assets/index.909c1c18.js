export function __vite_legacy_guard(){import("data:text/javascript,")}import{d as e,o as t,c as a,w as r,a as o,u as n,R as l,E as s,b as u,e as i,f as d,g as c,h as p,i as f,r as m,j as h,k as y,l as _,t as v,m as b,F as g,n as w,p as E,q as k,s as C,v as N,x,y as L,S as j,z as A,A as I,B as P,C as U,D as V,G as $,H as D,I as O,J as S}from"./vendor.0cabb7b3.js";!function(){const e=document.createElement("link").relList;if(!(e&&e.supports&&e.supports("modulepreload"))){for(const e of document.querySelectorAll('link[rel="modulepreload"]'))t(e);new MutationObserver((e=>{for(const a of e)if("childList"===a.type)for(const e of a.addedNodes)"LINK"===e.tagName&&"modulepreload"===e.rel&&t(e)})).observe(document,{childList:!0,subtree:!0})}function t(e){if(e.ep)return;e.ep=!0;const t=function(e){const t={};return e.integrity&&(t.integrity=e.integrity),e.referrerpolicy&&(t.referrerPolicy=e.referrerpolicy),"use-credentials"===e.crossorigin?t.credentials="include":"anonymous"===e.crossorigin?t.credentials="omit":t.credentials="same-origin",t}(e);fetch(e.href,t)}}();var T=(e,t)=>{for(const[a,r]of t)e[a]=r;return e};const q=u("Home"),z=u("搜索");var F=T(e({setup:e=>(e,u)=>{const f=i,m=d,h=c,y=p;return t(),a(n(s),null,{default:r((()=>[o(h,{class:"header"},{default:r((()=>[o(m,{mode:"horizontal",router:""},{default:r((()=>[o(f,{index:"1",route:{path:"/"}},{default:r((()=>[q])),_:1}),o(f,{index:"2",route:{path:"/user"}},{default:r((()=>[z])),_:1})])),_:1})])),_:1}),o(y,{class:"main"},{default:r((()=>[o(n(l))])),_:1})])),_:1})}}),[["__scopeId","data-v-7b9a7890"]]);const R={},B=function(e,t){return t&&0!==t.length?Promise.all(t.map((e=>{if((e=`/${e}`)in R)return;R[e]=!0;const t=e.endsWith(".css"),a=t?'[rel="stylesheet"]':"";if(document.querySelector(`link[href="${e}"]${a}`))return;const r=document.createElement("link");return r.rel=t?"stylesheet":"modulepreload",t||(r.as="script",r.crossOrigin=""),r.href=e,document.head.appendChild(r),t?new Promise(((e,t)=>{r.addEventListener("load",e),r.addEventListener("error",t)})):void 0}))).then((()=>e())):e()};async function H(){if(G.length>0)return G;let e;try{let t=await fetch("/data.json");e=await t.json()}catch(a){return console.warn(a),f({title:"Error",message:"Could not load data",type:"error",onClick:()=>{location.reload()},duration:0,showClose:!1}),[]}let t=[];for(const r in e)t.push(e[r]);return G=t,t}let G=[];const J=u("一月内绿宝石使用排行"),K=u("近三月绿宝石使用排行"),M=u("近一年绿宝石使用排行"),W=u("总绿宝石使用排行"),Q=u("绿宝石使用列表"),X=u("某个用户绿宝石使用详单"),Y=_("p",null,[_("a",{href:"https://greasyfork.org/zh-CN/scripts/424437-%E6%9F%A5%E7%9C%8B%E8%AF%84%E5%88%86"},"某个用户被评分列表")],-1),Z=e({setup(e){let a=m(0);return h((async()=>{let e=await H(),t=0;for(const a of e)t+=a.Count;a.value=t,document.title="绿宝石"})),(e,l)=>(t(),y(g,null,[_("p",null,"总绿宝石使用数："+v(n(a)),1),_("p",null,[o(n(b),{to:"/table30"},{default:r((()=>[J])),_:1})]),_("p",null,[o(n(b),{to:"/table90"},{default:r((()=>[K])),_:1})]),_("p",null,[o(n(b),{to:"/table365"},{default:r((()=>[M])),_:1})]),_("p",null,[o(n(b),{to:"/all"},{default:r((()=>[W])),_:1})]),_("p",null,[o(n(b),{to:"/list"},{default:r((()=>[Q])),_:1})]),_("p",null,[o(n(b),{to:"/user"},{default:r((()=>[X])),_:1})]),Y],64))}}),ee={key:1},te=e({props:{day:{type:Number,default:30}},setup(e){let l=w(e),s=m(""),i=m([]);function d(e){return e+1}return h((()=>{E((async()=>{0==l.day.value?(s.value="总绿宝石使用排行",document.title=s.value):(s.value=`${l.day.value} 天内绿宝石使用排行`,document.title=s.value),i.value=[];let e=await H(),t=(new Date).getTime()/1e3,a=24*l.day.value*3600,r=[],o={};for(const n of e)if(t-Number(n.Time)<a||0==l.day.value){let e=String(n.Uid);o[e]?o[e].count+=Number(n.Count):o[e]={uid:Number(n.Uid),name:n.Username,count:Number(n.Count),href:`/user/${n.Uid}`,v:"详情"}}for(const n in o)r.push(o[n]);i.value=r}))})),(e,l)=>{const c=k("router-link");return t(),y(g,null,[_("h1",null,v(n(s)),1),o(n(N),{data:n(i),"default-sort":{prop:"count",order:"ascending"}},{default:r((()=>[o(n(C),{type:"index",label:"排名",index:d}),o(n(C),{prop:"uid",label:"uid"}),o(n(C),{prop:"name",label:"用户名"}),o(n(C),{prop:"count",label:"总数",sortable:""}),o(n(C),{label:"详情"},{default:r((e=>[e.row.href?(t(),a(c,{key:0,to:e.row.href},{default:r((()=>[u(v(e.row.v),1)])),_:2},1032,["to"])):(t(),y("span",ee,v(e.row.v),1))])),_:1})])),_:1},8,["data"])],64)}}});const ae=(e=>(A("data-v-cefa5782"),e=e(),I(),e))((()=>_("h1",null,"查询某个用户绿宝石使用详单",-1))),re={class:"c"};const oe=[{path:"/",component:Z},{path:"/:table(table\\d+)",component:te,props:function(e){let t=e.params.table.match(/\d+/g);return null!=t&&t.length>0?{day:Number(t[0])}:{}}},{path:"/all",component:te,props:{day:"0"}},{path:"/user/:uid",component:()=>B((()=>import("./user.3030ecfe.js")),["assets/user.3030ecfe.js","assets/user.ad7c338f.css","assets/vendor.0cabb7b3.js","assets/vendor.cdde2078.css"]),props:e=>({uid:Number(e.params.uid)})},{path:"/list",component:()=>B((()=>import("./user.3030ecfe.js")),["assets/user.3030ecfe.js","assets/user.ad7c338f.css","assets/vendor.0cabb7b3.js","assets/vendor.cdde2078.css"]),props:{uid:0}},{path:"/user",component:T(e({setup(e){let a=m("");const r=x();function l(){r.push({path:"/user/"+a.value})}return(e,r)=>{const s=P,u=U,i=V;return t(),y(g,null,[ae,o(s),_("div",re,[o(u,{modelValue:n(a),"onUpdate:modelValue":r[0]||(r[0]=e=>L(a)?a.value=e:a=e),type:"tel",placeholder:"Please input"},null,8,["modelValue"]),o(i,{icon:n(j),circle:"",onClick:l},null,8,["icon"])])],64)}}}),[["__scopeId","data-v-cefa5782"]])}],ne=$({history:D(),routes:oe});let le=O(F);le.use(S),le.use(ne),le.mount("#app");export{H as g};
//# sourceMappingURL=index.909c1c18.js.map
