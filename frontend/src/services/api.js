const API_BASE = '/api/v1'

async function request(path, options = {}) {
  const url = `${API_BASE}${path}`
  const response = await fetch(url, {
    headers: {
      'Content-Type': 'application/json',
      ...options.headers
    },
    ...options
  })
  
  if (!response.ok) {
    const error = await response.json().catch(() => ({ error: response.statusText }))
    throw new Error(error.error || `HTTP ${response.status}`)
  }
  
  return response.json()
}

export const api = {
  get: (path) => request(path),
  post: (path, data) => request(path, { method: 'POST', body: JSON.stringify(data) }),
  
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
    get: (id) => api.get(`/horarios/${id}`)
  },
  
  bloques: {
    list: () => api.get('/bloques')
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
