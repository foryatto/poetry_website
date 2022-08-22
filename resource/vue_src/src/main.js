import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'

Vue.config.productionTip = false

import VueAxios from 'vue-axios'
import axios from 'axios'
Vue.use(VueAxios, axios)
axios.defaults.baseURL="http://localhost:8080/api/";
// axios.defaults.baseURL="/api/";

import Antd from 'ant-design-vue';
import 'ant-design-vue/dist/antd.css'
Vue.use(Antd);

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
