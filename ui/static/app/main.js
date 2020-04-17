const routes = [
  {
    path: '/login',
    name: 'login',
    component: Login
  },
  {
    path: '/',
    name: 'home',
    component: Home,
    beforeEnter: (to, from, next) => {
      let voterCode = localStorage.getItem('voter-code')
      if (!voterCode) {
        next({ path: '/login' })
      } else {
        next()
      }
    }
  }
]

const router = new VueRouter({
  mode: 'history',
  base: '/app',
  routes: routes
})

const app = new Vue({
  router
}).$mount('#app')
