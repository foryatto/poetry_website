import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/Home.vue'
import PoetsView from '../views/Poets.vue'
import PoemsView from '../views/Poems.vue'
import PoemView from '../views/Poem.vue'
import SearchView from '../views/Search.vue'


const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/poets',
      name: 'poets',
      component: PoetsView
    },
    {
      path: '/poems',
      name: 'poems',
      component: PoemsView
    },
    {
      path: '/poem',
      name: 'poem',
      component: PoemView
    },
    {
      path: '/search',
      name: 'search',
      component: SearchView
    },
  ]
})

export default router
