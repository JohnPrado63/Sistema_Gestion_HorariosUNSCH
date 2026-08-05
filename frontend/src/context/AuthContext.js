import { ref, readonly } from 'vue'

const TOKEN_KEY = 'auth_token'
const USER_KEY = 'auth_user'

const user = ref(null)
const token = ref(null)
const loading = ref(true)

function loadFromStorage() {
  const savedToken = localStorage.getItem(TOKEN_KEY)
  const savedUser = localStorage.getItem(USER_KEY)

  if (savedToken && savedUser) {
    try {
      token.value = savedToken
      user.value = JSON.parse(savedUser)
    } catch (e) {
      localStorage.removeItem(TOKEN_KEY)
      localStorage.removeItem(USER_KEY)
    }
  }
  loading.value = false
}

loadFromStorage()

const login = async (email, password) => {
  const response = await fetch('/api/v1/login', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({ email, password })
  })

  if (!response.ok) {
    const error = await response.json()
    throw new Error(error.error || 'Error al iniciar sesión')
  }

  const data = await response.json()

  localStorage.setItem(TOKEN_KEY, data.token)
  localStorage.setItem(USER_KEY, JSON.stringify(data.user))

  token.value = data.token
  user.value = data.user

  return data
}

const logout = async () => {
  if (token.value) {
    try {
      await fetch('/api/v1/logout', {
        method: 'POST',
        headers: {
          'Authorization': `Bearer ${token.value}`
        }
      })
    } catch (e) {
      console.warn('Logout request failed:', e)
    }
  }

  localStorage.removeItem(TOKEN_KEY)
  localStorage.removeItem(USER_KEY)

  token.value = null
  user.value = null
}

const hasRole = (roles) => {
  if (!user.value) return false
  if (typeof roles === 'string') {
    return user.value.rol === roles
  }
  return roles.includes(user.value.rol)
}

export const useAuth = () => ({
  user: readonly(user),
  token: readonly(token),
  loading: readonly(loading),
  isAuthenticated: !!user.value,
  login,
  logout,
  hasRole
})

export const getAuthHeaders = () => {
  const savedToken = localStorage.getItem(TOKEN_KEY)
  return savedToken ? { 'Authorization': `Bearer ${savedToken}` } : {}
}

export const AuthProvider = {
  install(app) {
    app.provide('auth', {
      user: readonly(user),
      token: readonly(token),
      loading: readonly(loading),
      isAuthenticated: !!user.value,
      login,
      logout,
      hasRole
    })
  }
}
