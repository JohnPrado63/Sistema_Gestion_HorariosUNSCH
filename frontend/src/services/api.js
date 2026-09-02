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
    create: (data) => api.post('/facultades', data),
    update: (id, data) => api.put(`/facultades/${id}`, data),
    delete: (id) => api.delete(`/facultades/${id}`),
    departamentos: (id) => api.get(`/facultades/${id}/departamentos`),
    escuelas: (id) => api.get(`/facultades/${id}/escuelas`)
  },

  departamentos: {
    list: () => api.get('/departamentos'),
    get: (id) => api.get(`/departamentos/${id}`),
    create: (data) => api.post('/departamentos', data),
    update: (id, data) => api.put(`/departamentos/${id}`, data),
    delete: (id) => api.delete(`/departamentos/${id}`),
    docentes: (id) => api.get(`/departamentos/${id}/docentes`),
    escuelas: (id) => api.get(`/departamentos/${id}/escuelas`)
  },

  escuelas: {
    list: () => api.get('/escuelas'),
    get: (id) => api.get(`/escuelas/${id}`),
    create: (data) => api.post('/escuelas', data),
    update: (id, data) => api.put(`/escuelas/${id}`, data),
    delete: (id) => api.delete(`/escuelas/${id}`),
    docentes: (id) => api.get(`/escuelas/${id}/docentes`),
    cursos: (id) => api.get(`/escuelas/${id}/cursos`)
  },

  docentes: {
    list: () => api.get('/docentes'),
    get: (id) => api.get(`/docentes/${id}`),
    create: (data) => api.post('/docentes', data),
    update: (id, data) => api.put(`/docentes/${id}`, data),
    delete: (id) => api.delete(`/docentes/${id}`),
    cursos: (id) => api.get(`/docentes/${id}/cursos`)
  },

  cursos: {
    list: () => api.get('/cursos'),
    get: (id) => api.get(`/cursos/${id}`),
    create: (data) => api.post('/cursos', data),
    update: (id, data) => api.put(`/cursos/${id}`, data),
    delete: (id) => api.delete(`/cursos/${id}`)
  },

  series: {
    list: () => api.get('/series')
  },

  periodos: {
    list: () => api.get('/periodos'),
    create: (data) => api.post('/periodos', data),
    update: (id, data) => api.put(`/periodos/${id}`, data),
    delete: (id) => api.delete(`/periodos/${id}`)
  },

  planesEstudio: {
    list: () => api.get('/planes-estudio'),
    create: (data) => api.post('/planes-estudio', data),
    update: (id, data) => api.put(`/planes-estudio/${id}`, data),
    delete: (id) => api.delete(`/planes-estudio/${id}`)
  },

  aulas: {
    list: () => api.get('/aulas'),
    listAll: () => api.get('/aulas/todas'),
    create: (data) => api.post('/aulas', data),
    update: (id, data) => api.put(`/aulas/${id}`, data),
    delete: (id) => api.delete(`/aulas/${id}`),
    setActivo: (id, activo) => api.put(`/aulas/${id}/activo`, { activo })
  },

  horarios: {
    list: () => api.get('/horarios'),
    get: (id) => api.get(`/horarios/${id}`),
    create: (data) => api.post('/horarios', data),
    generate: (data) => api.post('/horarios/generar', data),
    delete: (id) => api.delete(`/horarios/${id}`),
    bloques: (id) => api.get(`/horarios/${id}/bloques`)
  },

  bloques: {
    list: () => api.get('/bloques'),
    create: (data) => api.post('/bloques', data),
    delete: (id) => api.delete(`/bloques/${id}`),
    verificar: (data) => api.post('/bloques/verificar', data)
  },

  gruposHorario: {
    list: (escuela, periodo, serie, semestre) => {
      let url = `/grupos-horario?escuela=${escuela}&periodo=${periodo}`
      if (serie) url += `&serie=${serie}`
      if (semestre) url += `&semestre=${semestre}`
      return api.get(url)
    }
  },

  cargas: {
    list: () => api.get('/carga-academica'),
    create: (data) => api.post('/carga-academica', data),
    delete: (id) => api.delete(`/carga-academica/${id}`),
    approve: (id, data) => api.post(`/carga-academica/${id}/aprobar`, data)
  },

  validaciones: {
    placement: (data) => api.post('/validaciones/placement', data),
    audit: (data) => api.post('/validaciones/audit', data),
    carga: (data) => api.post('/validaciones/carga', data),
    escenarios: () => api.get('/validaciones/escenarios')
  }
}

export default api
