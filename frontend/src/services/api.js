const API_BASE = '/api/v1'

const TOKEN_KEY = 'auth_token'

function getAuthHeaders() {
  const token = localStorage.getItem(TOKEN_KEY)
  return token ? { 'Authorization': `Bearer ${token}` } : {}
}

async function request(path, options = {}) {
  const url = `${API_BASE}${path}`

  const headers = {
    'Content-Type': 'application/json',
    ...getAuthHeaders(),
    ...options.headers
  }

  const response = await fetch(url, {
    headers,
    ...options
  })

  if (response.status === 401) {
    localStorage.removeItem(TOKEN_KEY)
    localStorage.removeItem('auth_user')
    window.location.href = '/app/login'
    throw new Error('Sesión expirada')
  }

  if (!response.ok) {
    const error = await response.json().catch(() => ({ error: response.statusText }))
    throw new Error(error.error || `HTTP ${response.status}`)
  }

  return response.json()
}

export const api = {
  get: (path) => request(path),
  post: (path, data) => request(path, { method: 'POST', body: JSON.stringify(data) }),
  put: (path, data) => request(path, { method: 'PUT', body: JSON.stringify(data) }),
  delete: (path) => request(path, { method: 'DELETE' }),

  auth: {
    login: (email, password) => request('/login', { method: 'POST', body: JSON.stringify({ email, password }) }),
    logout: () => request('/logout', { method: 'POST' }),
    me: () => request('/me'),

    listUsers: () => request('/usuarios'),
    createUser: (data) => request('/usuarios', { method: 'POST', body: JSON.stringify(data) }),
    updateUser: (id, data) => request(`/usuarios/${id}`, { method: 'PUT', body: JSON.stringify(data) }),
    deleteUser: (id) => request(`/usuarios/${id}`, { method: 'DELETE' })
  },

  facultades: {
    list: () => api.get('/facultades'),
    get: (id) => api.get(`/facultades/${id}`),
    departamentos: (id) => api.get(`/facultades/${id}/departamentos`),
    escuelas: (id) => api.get(`/facultades/${id}/escuelas`)
  },

  departamentos: {
    list: () => api.get('/departamentos'),
    get: (id) => api.get(`/departamentos/${id}`),
    docentes: (id) => api.get(`/departamentos/${id}/docentes`),
    escuelas: (id) => api.get(`/departamentos/${id}/escuelas`)
  },

  escuelas: {
    list: () => api.get('/escuelas'),
    get: (id) => api.get(`/escuelas/${id}`),
    docentes: (id) => api.get(`/escuelas/${id}/docentes`),
    cursos: (id) => api.get(`/escuelas/${id}/cursos`)
  },

  docentes: {
    list: () => api.get('/docentes'),
    get: (id) => api.get(`/docentes/${id}`),
    cursos: (id) => api.get(`/docentes/${id}/cursos`)
  },

  cursos: {
    list: () => api.get('/cursos'),
    get: (id) => api.get(`/cursos/${id}`)
  },

  periodos: {
    list: () => api.get('/periodos')
  },

  aulas: {
    list: () => api.get('/aulas')
  },

  horarios: {
    list: () => api.get('/horarios'),
    get: (id) => api.get(`/horarios/${id}`),
    create: (data) => api.post('/horarios', data),
    bloques: (id) => api.get(`/horarios/${id}/bloques`)
  },

  bloques: {
    list: () => api.get('/bloques'),
    create: (data) => api.post('/bloques', data),
    verificar: (data) => api.post('/bloques/verificar', data)
  },

  gruposHorario: {
    list: (escuela, periodo) => api.get(`/grupos-horario?escuela=${escuela}&periodo=${periodo}`)
  },

  cargas: {
    list: () => api.get('/cargas-academicas')
  },

  validaciones: {
    placement: (data) => api.post('/validaciones/placement', data),
    audit: (data) => api.post('/validaciones/audit', data),
    carga: (data) => api.post('/validaciones/carga', data),
    escenarios: () => api.get('/validaciones/escenarios')
  }
}

export default api
