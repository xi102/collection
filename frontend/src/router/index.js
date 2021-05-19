import Vue from 'vue'
import VueRouter from 'vue-router'
//import { component } from 'vue/types/umd'
import Home from '../views/Home.vue'

Vue.use(VueRouter)

const originalPush = VueRouter.prototype.push
VueRouter.prototype.push = function push(location) {
  return originalPush.call(this, location).catch(err => err)
}

const routes = [
  {
    path: '/',
    name: 'Home',
    redirect: '/home',
    component: () => import('../Layout/layout.vue'),
    children: [
      {
        path:'/home',
        component: Home
      }
    ]
  },
  {
    path: '/about',
    name: 'About',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import('../Layout/layout.vue'),
    children: [
      {
        path:'/about',
        component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
      }
    ]
  },
  {
    path: '/testfiles',
    name: 'About',
    component: () => import('../Layout/layout.vue'),
    children: [
      {
        path:'/testfiles',
        component: () => import(/* webpackChunkName: "about" */ '../views/UploadFile.vue')
      }
    ]
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import(/* webpackChunkName: "about" */ '../views/Login.vue')
  },
  {
    path: '/files_show',
    name: 'Files',
    component: () => import(/* webpackChunkName: "about" */ '../views/Files.vue')
  },
  {
    path: '/register',
    name: 'Register',
    component: () => import(/* webpackChunkName: "about" */ '../views/Register.vue')
  }
]

const router = new VueRouter({
  mode: 'history',
  routes
})

export default router
