import {
    createApp
} from 'vue'
import App from './App.vue'
import router from './router'

// import './assets/main.css'
import 'ant-design-vue/dist/antd.css';

const app = createApp(App)

app.use(router)

// "http://localhost:8080/api" => "/api"
app.provide("appConfig", {
        "apiBaseUrl": "http://localhost:8080/api",
    }
)

app.mount('#app')