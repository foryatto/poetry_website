import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import Poets from '../views/Poets.vue'
import Poems from '../views/Poems.vue'
import PoemDetail from '../views/PoemDetail.vue'
import SearchResult from '../views/SearchResult.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/poets',
    name: 'Poets',
    component: Poets
  },
  {
    path: '/poems',
    name: 'Poems',
    component: Poems
  },
  {
    path: '/poems/detail',
    name: 'PoemDetail',
    component: PoemDetail
  },
  {
    path: '/search_result',
    name: 'SearchResult',
    component: SearchResult
  },
  
  {
    path: '/about',
    name: 'About',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
  }
]

const router = new VueRouter({
  routes
})

export default router
