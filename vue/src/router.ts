import { createRouter, createWebHistory, RouteLocationNormalized, RouteRecordRaw, useRoute } from 'vue-router'
import index from './views/index.vue'
import table from './views/table.vue'
import find from './views/find.vue'

declare module 'vue-router' {
  interface RouteMeta {
    scrollToTop?: boolean
  }
}

function dynamicPropsFn(route: RouteLocationNormalized) {
  let t: string = ""
  if (typeof route.params.table == "string") {
    t = route.params.table
  } else {
    t = route.params.table[0]
  }
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
    meta: { scrollToTop: true }
  },
  {
    path: '/list',
    component: () => import('./views/user.vue'),
    props: {
      uid: 0
    },
    meta: { scrollToTop: true }
  },
  {
    path: '/user',
    component: find,
  }
]

const router = createRouter({
  history: createWebHistory(),
  scrollBehavior: (to, from, savedPosition) => {
    if (savedPosition) {
      return savedPosition
    }
    if (to.hash) {
      return {
        el: to.hash,
        behavior: 'smooth',
      }
    }
    let position: { left: number, top: number } = {
      left: 0,
      top: 0,
    }
    if (to.meta.scrollToTop) {
      position.left = 0
      position.top = 0
    } else {
      return false
    }
    return position
  },
  routes
})

export default router
