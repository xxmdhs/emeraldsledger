import { createRouter, createWebHashHistory, RouteRecordRaw, useRoute } from 'vue-router'
import index from './views/index.vue'
import table from './views/table.vue'
import find from './views/find.vue'


function dynamicPropsFn(route: any) {
  let t: string = route.params.table
  let l = t.match(/\d+/g)
  if (l != null && l.length > 0) {
    return {
      day: Number(l[0]),
    }
  }
  return {}
}
const routes: RouteRecordRaw[] = [
  {
    path: '/',
    component: index,
  },
  {
    path: '/:table(table\\d+)',
    component: table,
    props: dynamicPropsFn
  },
  {
    path: '/all',
    component: table,
    props: {
      day: '0'
    }
  },
  {
    path: '/user/:uid',
    component: () => import('./views/user.vue'),
    props: route => ({
      uid: Number(route.params.uid),
    }),
  },
  {
    path: '/list',
    component: () => import('./views/user.vue'),
    props: {
      uid: 0
    },
  },
  {
    path: '/user',
    component: find,
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
