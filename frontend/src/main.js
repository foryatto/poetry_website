import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'

Vue.config.productionTip = false

import VueAxios from 'vue-axios'
import axios from 'axios'
Vue.use(VueAxios, axios)
// axios.defaults.baseURL="https://shici.foryatto.com/api/v1";
axios.defaults.baseURL="http://127.0.0.1:8000/api/v1/poetry/";

import Antd from 'ant-design-vue';
import 'ant-design-vue/dist/antd.css'
Vue.use(Antd);

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
