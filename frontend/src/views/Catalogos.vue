<template>
  <div>
    <header class="topbar">
      <h1 class="page-title">Catálogos Institucionales</h1>
      <div class="page-actions">
        <button class="btn btn-secondary btn-sm" @click="refresh">
          <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M23 4v6h-6"/>
            <path d="M1 20v-6h6"/>
            <path d="M3.51 9a9 9 0 0 1 14.85-3.36L23 10M1 14l4.64 4.36A9 9 0 0 0 20.49 15"/>
          </svg>
          Actualizar
        </button>
      </div>
    </header>
    
    <div class="page-content">
      <div class="tabs">
        <button class="tab" :class="{ active: activeTab === 'facultades' }" @click="activeTab = 'facultades'">
          Facultades ({{ catalogos.facultades.length }})
        </button>
        <button class="tab" :class="{ active: activeTab === 'departamentos' }" @click="activeTab = 'departamentos'">
          Departamentos ({{ catalogos.departamentos.length }})
        </button>
        <button class="tab" :class="{ active: activeTab === 'escuelas' }" @click="activeTab = 'escuelas'">
          Escuelas ({{ catalogos.escuelas.length }})
        </button>
        <button class="tab" :class="{ active: activeTab === 'docentes' }" @click="activeTab = 'docentes'">
          Docentes ({{ catalogos.docentes.length }})
        </button>
        <button class="tab" :class="{ active: activeTab === 'cursos' }" @click="activeTab = 'cursos'">
          Cursos ({{ catalogos.cursos.length }})
        </button>
      </div>
      
      <div v-if="loading" class="loading">
        <div class="spinner"></div>
      </div>
      
      <div v-else-if="error" class="error-alert">
        <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <circle cx="12" cy="12" r="10"/>
          <line x1="12" y1="8" x2="12" y2="12"/>
          <line x1="12" y1="16" x2="12.01" y2="16"/>
        </svg>
        {{ error }}
      </div>
      
      <template v-else>
        <div v-show="activeTab === 'facultades'" class="card">
          <div class="card-header">
            <h2 class="card-title">Facultades</h2>
          </div>
          <div class="table-wrapper">
            <table>
              <thead>
                <tr>
                  <th>ID</th>
                  <th>Nombre</th>
                  <th>Acciones</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="item in catalogos.facultades" :key="item.id_facultad">
                  <td>{{ item.id_facultad }}</td>
                  <td>{{ item.nombre }}</td>
                  <td>
                    <button class="btn btn-secondary btn-sm" @click="verDetalle('facultad', item.id_facultad)">
                      Ver Detalle
                    </button>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
        
        <div v-show="activeTab === 'departamentos'" class="card">
          <div class="card-header">
            <h2 class="card-title">Departamentos Académicos</h2>
          </div>
          <div class="table-wrapper">
            <table>
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
                  <td>{{ item.id_departamento }}</td>
                  <td>{{ getFacultadNombre(item.id_facultad) }}</td>
                  <td>{{ item.nombre }}</td>
                  <td>
                    <button class="btn btn-secondary btn-sm" @click="verDetalle('departamento', item.id_departamento)">
                      Ver Detalle
                    </button>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
        
        <div v-show="activeTab === 'escuelas'" class="card">
          <div class="card-header">
            <h2 class="card-title">Escuelas Profesionales</h2>
          </div>
          <div class="table-wrapper">
            <table>
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
                  <td>{{ item.id_escuela }}</td>
                  <td>{{ getFacultadNombre(item.id_facultad) }}</td>
                  <td>{{ getDepartamentoNombre(item.id_departamento) }}</td>
                  <td>{{ item.nombre }}</td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
        
        <div v-show="activeTab === 'docentes'" class="card">
          <div class="card-header">
            <h2 class="card-title">Docentes</h2>
          </div>
          <div class="table-wrapper">
            <table>
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
                  <td>{{ item.id_docente }}</td>
                  <td>{{ getDepartamentoNombre(item.id_departamento) }}</td>
                  <td>{{ item.codigo_plaza }}</td>
                  <td>{{ item.nombres }}</td>
                  <td>{{ item.apellidos }}</td>
                  <td>{{ item.email || '—' }}</td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
        
        <div v-show="activeTab === 'cursos'" class="card">
          <div class="card-header">
            <h2 class="card-title">Cursos</h2>
          </div>
          <div class="table-wrapper">
            <table>
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
                  <td>{{ item.id_curso }}</td>
                  <td><span class="badge badge-gray">{{ item.codigo }}</span></td>
                  <td>{{ item.nombre }}</td>
                  <td>{{ item.creditos }}</td>
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

function verDetalle(tipo, id) {
  alert(`Ver detalle de ${tipo} #${id} - Por implementar`)
}

onMounted(loadAll)
</script>
