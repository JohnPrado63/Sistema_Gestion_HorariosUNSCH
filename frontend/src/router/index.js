import { createRouter, createWebHistory } from 'vue-router'
import Dashboard from '../views/Dashboard.vue'
import Catalogos from '../views/Catalogos.vue'
import Validaciones from '../views/Validaciones.vue'
import Horarios from '../views/Horarios.vue'
import Siige from '../views/Siige.vue'

const routes = [
  { path: '/', redirect: '/dashboard' },
  { path: '/dashboard', name: 'Dashboard', component: Dashboard },
  { path: '/catalogos', name: 'Catalogos', component: Catalogos },
  { path: '/validaciones', name: 'Validaciones', component: Validaciones },
  { path: '/horarios', name: 'Horarios', component: Horarios },
  { path: '/siige', name: 'Siige', component: Siige },
]

const router = createRouter({
  history: createWebHistory('/app/'),
  routes
})

export default router
