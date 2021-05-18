import Vue from 'vue'
import VueRouter from 'vue-router'
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
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
  },
  {
    path: '/testfiles',
    name: 'About',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/UploadFile.vue')
  },
  {
    path: '/login',
    name: 'Login',
    redirct: '/login/ll',
    component: () => import('../Layout/layout.vue'),
    children: [{
      name: 'login2',
      path: '/ll',
      component: () => import(/* webpackChunkName: "about" */ '../views/Login.vue')
    }]

  },
  {
    path: '/files_show',
    name: 'Files',
    component: () => import(/* webpackChunkName: "about" */ '../views/Files.vue')
  }
]

const router = new VueRouter({
  mode: 'history',
  routes
})

export default router
