import { createRouter, createWebHistory } from 'vue-router'
import Dashboard from '../views/Dashboard.vue'
import Catalogos from '../views/Catalogos.vue'
import Validaciones from '../views/Validaciones.vue'
import Horarios from '../views/Horarios.vue'
import Siige from '../views/Siige.vue'
import CargaAcademica from '../views/CargaAcademica.vue'
import Login from '../views/Login.vue'

const TOKEN_KEY = 'auth_token'

function isAuthenticated() {
  return !!localStorage.getItem(TOKEN_KEY)
}

function getUserRole() {
  const userStr = localStorage.getItem('auth_user')
  if (!userStr) return null
  try {
    return JSON.parse(userStr).rol
  } catch {
    return null
  }
}

function hasRole(roles) {
  const userRole = getUserRole()
  if (!userRole) return false
  if (typeof roles === 'string') {
    return userRole === roles
  }
  return roles.includes(userRole)
}

const routes = [
  {
    path: '/login',
    name: 'Login',
    component: Login,
    meta: { public: true }
  },
  {
    path: '/',
    redirect: '/dashboard'
  },
  {
    path: '/dashboard',
    name: 'Dashboard',
    component: Dashboard,
    meta: { roles: ['ADMIN_TI', 'DGA', 'DIRECTOR_ESCUELA', 'JEFE_DEPTO', 'COORDINADOR', 'DOCENTE', 'ESTUDIANTE'] }
  },
  {
    path: '/catalogos',
    name: 'Catalogos',
    component: Catalogos,
    meta: { roles: ['ADMIN_TI', 'DGA'] }
  },
  {
    path: '/carga-academica',
    name: 'CargaAcademica',
    component: CargaAcademica,
    meta: { roles: ['ADMIN_TI', 'DGA', 'DIRECTOR_ESCUELA', 'JEFE_DEPTO'] }
  },
  {
    path: '/validaciones',
    name: 'Validaciones',
    component: Validaciones,
    meta: { roles: ['ADMIN_TI', 'DGA'] }
  },
  {
    path: '/horarios',
    name: 'Horarios',
    component: Horarios,
    meta: { roles: ['ADMIN_TI', 'DGA', 'DIRECTOR_ESCUELA', 'JEFE_DEPTO', 'COORDINADOR'] }
  },
  {
    path: '/siige',
    name: 'Siige',
    component: Siige,
    meta: { roles: ['ADMIN_TI', 'DGA'] }
  }
]

const router = createRouter({
  history: createWebHistory('/app/'),
  routes
})

router.beforeEach((to, from, next) => {
  if (to.meta.public) {
    if (isAuthenticated() && to.path === '/login') {
      return next('/dashboard')
    }
    return next()
  }

  if (!isAuthenticated()) {
    return next('/login')
  }

  if (to.meta.roles) {
    if (!hasRole(to.meta.roles)) {
      return next('/dashboard')
    }
  }

  next()
})

export default router
