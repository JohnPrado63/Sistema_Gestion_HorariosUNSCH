import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { AuthProvider } from './context/AuthContext'
import './assets/main.css'

const app = createApp(App)

app.use(router)
app.use(AuthProvider)

app.mount('#app')
