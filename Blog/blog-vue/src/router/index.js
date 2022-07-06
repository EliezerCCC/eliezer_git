import Vue from 'vue'
import VueRouter from 'vue-router'
import home from '../views/Home.vue'
import login from '../views/Login.vue'
import register from '../views/Register.vue'
import add from '../views/Add.vue'
import mine from '../views/Mine.vue'
import article from '../views/Article.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'login',
    component: login
  },
  {
    path: '/register',
    name: 'register',
    component: register
  },
  {
    path: '/home',
    name: 'home',
    component: home
  },
  {
    path: '/add',
    name: 'add',
    component: add
  },
  {
    path: '/mine',
    name: 'mine',
    component: mine
  },
  {
    path: '/article',
    name: 'article',
    component: article
  }
]

const router = new VueRouter({
  routes
})

export default router
