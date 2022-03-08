import { createRouter, createWebHistory } from 'vue-router'
import Index from '@/views/Index.vue'
import Types from '@/views/Types.vue'
import Relayers from '@/views/Relayers.vue'

const routerHistory = createWebHistory()
const routes = [
  {
    path: '/wallet/',
    component: Index,
  },
  { path: '/wallet/types', component: Types },
  { path: '/wallet/relayers', component: Relayers },
]

const router = createRouter({
  history: routerHistory,
  routes,
})

export default router
