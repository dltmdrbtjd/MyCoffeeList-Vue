import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: () => import(/* webpackChunkName: "home" */ '../views/Home.vue')
  },
  {
    path: '/detail/:id',
    name: 'detail',
    component: () =>
      import(/* webpackChunkName: "detail" */ '../views/Detail.vue')
  },
  {
    path: '/create',
    name: 'Create',
    component: () =>
      import(/* webpackChunckName: 'create' */ '../views/Create.vue')
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
