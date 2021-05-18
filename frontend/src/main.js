import Vue from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify'
import router from './router'
import Axios from 'axios'
import Qs from 'qs'

Vue.config.productionTip = false
Vue.prototype.$axios = Axios
Vue.prototype.$qs = Qs

Axios.defaults.baseURL = 'http://localhost:8080'
Axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded'
Axios.defaults.headers.common['Access-Control-Allow-Origin'] = '*';

new Vue({
  vuetify,
  router,
  render: h => h(App)
}).$mount('#app')
