<template>
  <div class="catalogos-page">
    <header class="page-header">
      <div class="header-left">
        <h1 class="page-title">
          <svg width="28" height="28" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M4 19.5A2.5 2.5 0 0 1 6.5 17H20"/>
            <path d="M6.5 2H20v20H6.5A2.5 2.5 0 0 1 4 19.5v-15A2.5 2.5 0 0 1 6.5 2z"/>
          </svg>
          Catálogos Institucionales
        </h1>
        <p class="page-subtitle">Gestión de estructura académica y administrativa</p>
      </div>
      <div class="header-actions">
        <button class="btn btn-secondary" @click="refresh">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M23 4v6h-6"/>
            <path d="M1 20v-6h6"/>
            <path d="M3.51 9a9 9 0 0 1 14.85-3.36L23 10M1 14l4.64 4.36A9 9 0 0 0 20.49 15"/>
          </svg>
          Actualizar
        </button>
      </div>
    </header>

    <div class="page-content">
      <div class="tabs-container">
        <div class="tabs">
          <button class="tab" :class="{ active: activeTab === 'facultades' }" @click="activeTab = 'facultades'">
            <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M3 21h18"/>
              <path d="M5 21V7l8-4v18"/>
              <path d="M19 21V11l-6-4"/>
              <path d="M9 9v.01"/>
              <path d="M9 12v.01"/>
              <path d="M9 15v.01"/>
              <path d="M9 18v.01"/>
            </svg>
            Facultades
            <span class="tab-count">{{ catalogos.facultades.length }}</span>
          </button>
          <button class="tab" :class="{ active: activeTab === 'departamentos' }" @click="activeTab = 'departamentos'">
            <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M3 9l9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"/>
              <polyline points="9 22 9 12 15 12 15 22"/>
            </svg>
            Departamentos
            <span class="tab-count">{{ catalogos.departamentos.length }}</span>
          </button>
          <button class="tab" :class="{ active: activeTab === 'escuelas' }" @click="activeTab = 'escuelas'">
            <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M22 10v6M2 10l10-5 10 5-10 5z"/>
              <path d="M6 12v5c0 2 2 3 6 3s6-1 6-3v-5"/>
            </svg>
            Escuelas
            <span class="tab-count">{{ catalogos.escuelas.length }}</span>
          </button>
          <button class="tab" :class="{ active: activeTab === 'docentes' }" @click="activeTab = 'docentes'">
            <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"/>
              <circle cx="9" cy="7" r="4"/>
              <path d="M23 21v-2a4 4 0 0 0-3-3.87"/>
              <path d="M16 3.13a4 4 0 0 1 0 7.75"/>
            </svg>
            Docentes
            <span class="tab-count">{{ catalogos.docentes.length }}</span>
          </button>
          <button class="tab" :class="{ active: activeTab === 'cursos' }" @click="activeTab = 'cursos'">
            <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M4 19.5A2.5 2.5 0 0 1 6.5 17H20"/>
              <path d="M6.5 2H20v20H6.5A2.5 2.5 0 0 1 4 19.5v-15A2.5 2.5 0 0 1 6.5 2z"/>
            </svg>
            Cursos
            <span class="tab-count">{{ catalogos.cursos.length }}</span>
          </button>
        </div>
      </div>

      <div v-if="loading" class="loading-overlay">
        <div class="spinner-lg"></div>
        <p>Cargando catálogos...</p>
      </div>

      <div v-else-if="error" class="error-banner">
        <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <circle cx="12" cy="12" r="10"/>
          <line x1="12" y1="8" x2="12" y2="12"/>
          <line x1="12" y1="16" x2="12.01" y2="16"/>
        </svg>
        <span>{{ error }}</span>
        <button @click="refresh" class="btn btn-sm">Reintentar</button>
      </div>

      <template v-else>
        <div v-show="activeTab === 'facultades'" class="tab-content">
          <div class="content-header">
            <h2>Facultades</h2>
            <p>Listado de facultades registradas en el sistema</p>
          </div>
          <div class="table-container">
            <table class="data-table">
              <thead>
                <tr>
                  <th>ID</th>
                  <th>Nombre</th>
                  <th>Acciones</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="item in catalogos.facultades" :key="item.id_facultad">
                  <td class="id-cell">{{ item.id_facultad }}</td>
                  <td class="name-cell">
                    <div class="name-content">
                      <span class="entity-icon facultad-icon">
                        <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                          <path d="M3 21h18"/>
                          <path d="M5 21V7l8-4v18"/>
                        </svg>
                      </span>
                      {{ item.nombre }}
                    </div>
                  </td>
                  <td class="actions-cell">
                    <button class="btn btn-icon-sm" title="Ver detalle">
                      <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"/>
                        <circle cx="12" cy="12" r="3"/>
                      </svg>
                    </button>
                    <button class="btn btn-icon-sm" title="Editar">
                      <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"/>
                        <path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z"/>
                      </svg>
                    </button>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>

        <div v-show="activeTab === 'departamentos'" class="tab-content">
          <div class="content-header">
            <h2>Departamentos Académicos</h2>
            <p>Departamentos organizados por facultad</p>
          </div>
          <div class="table-container">
            <table class="data-table">
              <thead>
                <tr>
                  <th>ID</th>
                  <th>Facultad</th>
                  <th>Nombre</th>
                  <th>Acciones</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="item in catalogos.departamentos" :key="item.id_departamento">
                  <td class="id-cell">{{ item.id_departamento }}</td>
                  <td>
                    <span class="facultad-badge">{{ getFacultadNombre(item.id_facultad) }}</span>
                  </td>
                  <td class="name-cell">
                    <div class="name-content">
                      <span class="entity-icon depto-icon">
                        <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                          <path d="M3 9l9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"/>
                        </svg>
                      </span>
                      {{ item.nombre }}
                    </div>
                  </td>
                  <td class="actions-cell">
                    <button class="btn btn-icon-sm"><svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"/><circle cx="12" cy="12" r="3"/></svg></button>
                    <button class="btn btn-icon-sm"><svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"/><path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z"/></svg></button>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>

        <div v-show="activeTab === 'escuelas'" class="tab-content">
          <div class="content-header">
            <h2>Escuelas Profesionales</h2>
            <p>Escuelas organizadas por departamento y facultad</p>
          </div>
          <div class="table-container">
            <table class="data-table">
              <thead>
                <tr>
                  <th>ID</th>
                  <th>Facultad</th>
                  <th>Departamento</th>
                  <th>Nombre</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="item in catalogos.escuelas" :key="item.id_escuela">
                  <td class="id-cell">{{ item.id_escuela }}</td>
                  <td><span class="facultad-badge">{{ getFacultadNombre(item.id_facultad) }}</span></td>
                  <td><span class="depto-badge">{{ getDepartamentoNombre(item.id_departamento) }}</span></td>
                  <td class="name-cell">
                    <div class="name-content">
                      <span class="entity-icon escuela-icon">
                        <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                          <path d="M22 10v6M2 10l10-5 10 5z"/>
                          <path d="M6 12v5c0 2 2 3 6 3s6-1 6-3v-5"/>
                        </svg>
                      </span>
                      {{ item.nombre }}
                    </div>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>

        <div v-show="activeTab === 'docentes'" class="tab-content">
          <div class="content-header">
            <h2>Docentes</h2>
            <p>Personal docente registrado en el sistema</p>
          </div>
          <div class="table-container">
            <table class="data-table">
              <thead>
                <tr>
                  <th>ID</th>
                  <th>Departamento</th>
                  <th>Código Plaza</th>
                  <th>Nombres</th>
                  <th>Apellidos</th>
                  <th>Email</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="item in catalogos.docentes" :key="item.id_docente">
                  <td class="id-cell">{{ item.id_docente }}</td>
                  <td><span class="depto-badge">{{ getDepartamentoNombre(item.id_departamento) }}</span></td>
                  <td><span class="codigo-badge">{{ item.codigo_plaza }}</span></td>
                  <td>{{ item.nombres }}</td>
                  <td class="fw-600">{{ item.apellidos }}</td>
                  <td class="email-cell">{{ item.email || '—' }}</td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>

        <div v-show="activeTab === 'cursos'" class="tab-content">
          <div class="content-header">
            <h2>Cursos</h2>
            <p>Cursos registrados en el sistema académico</p>
          </div>
          <div class="table-container">
            <table class="data-table">
              <thead>
                <tr>
                  <th>ID</th>
                  <th>Código</th>
                  <th>Nombre</th>
                  <th>Créditos</th>
                  <th>Hrs. Teoría</th>
                  <th>Hrs. Práctica</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="item in catalogos.cursos" :key="item.id_curso">
                  <td class="id-cell">{{ item.id_curso }}</td>
                  <td><span class="codigo-badge curso-codigo">{{ item.codigo }}</span></td>
                  <td class="fw-600">{{ item.nombre }}</td>
                  <td><span class="creditos-badge">{{ item.creditos }}</span></td>
                  <td>{{ item.horas_teoria }}</td>
                  <td>{{ item.horas_practica }}</td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </template>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import api from '../services/api'

const activeTab = ref('facultades')
const loading = ref(true)
const error = ref('')

const catalogos = ref({
  facultades: [],
  departamentos: [],
  escuelas: [],
  docentes: [],
  cursos: []
})

async function loadAll() {
  loading.value = true
  error.value = ''
  try {
    const [facultades, departamentos, escuelas, docentes, cursos] = await Promise.all([
      api.facultades.list(),
      api.departamentos.list(),
      api.escuelas.list(),
      api.docentes.list(),
      api.cursos.list()
    ])

    catalogos.value = { facultades, departamentos, escuelas, docentes, cursos }
  } catch (e) {
    error.value = e.message
  } finally {
    loading.value = false
  }
}

function refresh() {
  loadAll()
}

function getFacultadNombre(id) {
  const f = catalogos.value.facultades.find(x => x.id_facultad === id)
  return f ? f.nombre : `Facultad #${id}`
}

function getDepartamentoNombre(id) {
  const d = catalogos.value.departamentos.find(x => x.id_departamento === id)
  return d ? d.nombre : `Depto #${id}`
}

onMounted(loadAll)
</script>

<style scoped>
.catalogos-page {
  min-height: 100vh;
  background: linear-gradient(135deg, #f0f4ff 0%, #fdf2f8 50%, #f0fdf4 100%);
}

.page-header {
  background: linear-gradient(135deg, #11998e 0%, #38ef7d 100%);
  color: white;
  padding: 24px 32px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  box-shadow: 0 8px 32px rgba(17, 153, 142, 0.3);
}

.header-left .page-title {
  display: flex;
  align-items: center;
  gap: 12px;
  font-size: 1.75rem;
  font-weight: 700;
  margin: 0;
}

.page-subtitle {
  margin: 4px 0 0 40px;
  opacity: 0.85;
  font-size: 0.95rem;
}

.btn {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  padding: 10px 20px;
  border-radius: 12px;
  font-size: 0.9rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s;
  border: none;
  box-shadow: 0 4px 12px rgba(0,0,0,0.1);
}

.btn-secondary {
  background: rgba(255,255,255,0.2);
  color: white;
  border: 2px solid rgba(255,255,255,0.3);
}

.btn-secondary:hover {
  background: rgba(255,255,255,0.3);
  transform: translateY(-2px);
}

.btn-sm {
  padding: 6px 12px;
  font-size: 0.8rem;
}

.btn-icon-sm {
  background: #f1f5f9;
  border: none;
  padding: 8px;
  border-radius: 8px;
  cursor: pointer;
  color: #64748b;
  transition: all 0.2s;
}

.btn-icon-sm:hover {
  background: #e2e8f0;
  color: #1e293b;
}

.page-content {
  padding: 24px;
}

.tabs-container {
  margin-bottom: 24px;
}

.tabs {
  display: flex;
  gap: 8px;
  background: white;
  padding: 8px;
  border-radius: 16px;
  box-shadow: 0 4px 16px rgba(0,0,0,0.06);
  overflow-x: auto;
}

.tab {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 12px 20px;
  border: none;
  background: transparent;
  border-radius: 10px;
  font-size: 0.9rem;
  font-weight: 600;
  color: #64748b;
  cursor: pointer;
  transition: all 0.3s;
  white-space: nowrap;
}

.tab:hover {
  background: #f0f4ff;
  color: #11998e;
}

.tab.active {
  background: linear-gradient(135deg, #11998e 0%, #38ef7d 100%);
  color: white;
  box-shadow: 0 4px 12px rgba(17, 153, 142, 0.3);
}

.tab-count {
  background: rgba(0,0,0,0.1);
  padding: 2px 8px;
  border-radius: 10px;
  font-size: 0.75rem;
}

.tab.active .tab-count {
  background: rgba(255,255,255,0.3);
}

.loading-overlay {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 80px 20px;
  color: #64748b;
}

.spinner-lg {
  width: 48px;
  height: 48px;
  border: 4px solid rgba(17, 153, 142, 0.2);
  border-top-color: #11998e;
  border-radius: 50%;
  animation: spin 1s linear infinite;
}

@keyframes spin {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}

.error-banner {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 16px 20px;
  background: white;
  border-radius: 16px;
  box-shadow: 0 4px 20px rgba(239, 68, 68, 0.15);
  border-left: 4px solid #ef4444;
  color: #dc2626;
  margin-bottom: 24px;
}

.tab-content {
  animation: fadeIn 0.3s;
}

@keyframes fadeIn {
  from { opacity: 0; transform: translateY(10px); }
  to { opacity: 1; transform: translateY(0); }
}

.content-header {
  margin-bottom: 20px;
}

.content-header h2 {
  margin: 0 0 4px;
  font-size: 1.25rem;
  color: #1e293b;
}

.content-header p {
  margin: 0;
  font-size: 0.9rem;
  color: #64748b;
}

.table-container {
  background: white;
  border-radius: 20px;
  box-shadow: 0 8px 32px rgba(0,0,0,0.08);
  overflow: hidden;
  border: 1px solid rgba(0,0,0,0.05);
}

.data-table {
  width: 100%;
  border-collapse: collapse;
}

.data-table th {
  text-align: left;
  padding: 16px 20px;
  font-size: 0.75rem;
  font-weight: 700;
  text-transform: uppercase;
  letter-spacing: 0.5px;
  color: #64748b;
  background: linear-gradient(135deg, #f8fafc 0%, #f0f4ff 100%);
  border-bottom: 2px solid #e2e8f0;
}

.data-table td {
  padding: 14px 20px;
  border-bottom: 1px solid #f1f5f9;
  font-size: 0.9rem;
  color: #334155;
}

.data-table tr:last-child td {
  border-bottom: none;
}

.data-table tr:hover td {
  background: linear-gradient(135deg, #f8fafc 0%, #f0f4ff 100%);
}

.id-cell {
  font-weight: 600;
  color: #94a3b8;
  width: 60px;
}

.name-cell {
  font-weight: 500;
}

.name-content {
  display: flex;
  align-items: center;
  gap: 10px;
}

.entity-icon {
  width: 32px;
  height: 32px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
}

.facultad-icon { background: linear-gradient(135deg, #667eea, #764ba2); }
.depto-icon { background: linear-gradient(135deg, #11998e, #38ef7d); }
.escuela-icon { background: linear-gradient(135deg, #f093fb, #f5576c); }

.actions-cell {
  width: 100px;
}

.fw-600 {
  font-weight: 600;
}

.facultad-badge {
  background: linear-gradient(135deg, #667eea20, #764ba220);
  color: #667eea;
  padding: 4px 10px;
  border-radius: 8px;
  font-size: 0.8rem;
  font-weight: 500;
}

.depto-badge {
  background: linear-gradient(135deg, #11998e20, #38ef7d20);
  color: #11998e;
  padding: 4px 10px;
  border-radius: 8px;
  font-size: 0.8rem;
  font-weight: 500;
}

.codigo-badge {
  background: #f1f5f9;
  color: #475569;
  padding: 4px 10px;
  border-radius: 8px;
  font-size: 0.8rem;
  font-weight: 600;
  font-family: monospace;
}

.curso-codigo {
  background: linear-gradient(135deg, #f093fb20, #f5576c20);
  color: #f5576c;
}

.creditos-badge {
  background: linear-gradient(135deg, #fbbf24, #f59e0b);
  color: white;
  padding: 4px 10px;
  border-radius: 8px;
  font-size: 0.8rem;
  font-weight: 700;
}

.email-cell {
  color: #64748b;
  font-style: italic;
}

@media (max-width: 1024px) {
  .tabs {
    flex-wrap: wrap;
  }
  .table-container {
    overflow-x: auto;
  }
  .data-table {
    min-width: 800px;
  }
}

@media (max-width: 768px) {
  .page-header {
    flex-direction: column;
    gap: 16px;
    text-align: center;
  }
  .page-subtitle {
    margin: 4px 0 0 0;
  }
}
</style>
