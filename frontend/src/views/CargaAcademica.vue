<template>
  <div class="carga-page">
    <header class="page-header">
      <div class="header-left">
        <h1 class="page-title">
          <svg width="28" height="28" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/>
            <polyline points="14 2 14 8 20 8"/>
            <line x1="12" y1="18" x2="12" y2="12"/>
            <line x1="9" y1="15" x2="15" y2="15"/>
          </svg>
          Carga Académica
        </h1>
        <p class="page-subtitle">Gestión de cursos, grupos y asignación docente</p>
      </div>
      <div class="header-actions">
        <div class="filter-group">
          <select v-model="selectedEscuela" class="filter-select" @change="loadCargas">
            <option value="">Todas las escuelas</option>
            <option v-for="e in escuelas" :key="e.id_escuela" :value="e.id_escuela">
              {{ e.nombre }}
            </option>
          </select>
        </div>
        <select v-model="selectedPeriodo" class="filter-select" @change="loadCargas">
          <option value="">Periodo...</option>
          <option v-for="p in periodos" :key="p.id_periodo" :value="p.id_periodo">
            {{ p.codigo }}
          </option>
        </select>
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
          <button class="tab" :class="{ active: activeTab === 'cargas' }" @click="activeTab = 'cargas'">
            <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M4 19.5A2.5 2.5 0 0 1 6.5 17H20"/>
              <path d="M6.5 2H20v20H6.5A2.5 2.5 0 0 1 4 19.5v-15A2.5 2.5 0 0 1 6.5 2z"/>
            </svg>
            Cargas
            <span class="tab-count">{{ cargas.length }}</span>
          </button>
          <button class="tab" :class="{ active: activeTab === 'docentes' }" @click="activeTab = 'docentes'; loadResumenDocentes()">
            <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"/>
              <circle cx="9" cy="7" r="4"/>
              <path d="M23 21v-2a4 4 0 0 0-3-3.87"/>
              <path d="M16 3.13a4 4 0 0 1 0 7.75"/>
            </svg>
            Resumen Docentes
          </button>
        </div>
      </div>

      <div v-if="loading" class="loading-overlay">
        <div class="spinner-lg"></div>
        <p>Cargando...</p>
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

      <template v-else-if="activeTab === 'cargas'">
        <div v-if="cargas.length === 0" class="empty-state">
          <div class="empty-icon">
            <svg width="80" height="80" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1">
              <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/>
              <polyline points="14 2 14 8 20 8"/>
            </svg>
          </div>
          <h2>No hay cargas académicas</h2>
          <p>Selecciona un periodo y escuela para comenzar</p>
        </div>

        <div v-else class="cargas-grid">
          <div v-for="carga in cargas" :key="carga.id_carga" class="carga-card" :class="{ 'autorizado': carga.estado === 'AUTORIZADO' }">
            <div class="carga-header">
              <div class="carga-badges">
                <span class="estado-badge" :class="carga.estado === 'AUTORIZADO' ? 'autorizado' : 'borrador'">
                  {{ carga.estado }}
                </span>
                <span v-if="carga.escuela" class="escuela-tag">{{ carga.escuela }}</span>
              </div>
              <button v-if="carga.estado === 'BORRADOR'" class="btn btn-success btn-sm" @click="openAgregarGrupo(carga)">
                <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <line x1="12" y1="5" x2="12" y2="19"/>
                  <line x1="5" y1="12" x2="19" y2="12"/>
                </svg>
                + Grupo
              </button>
            </div>

            <div class="carga-curso">
              <span class="curso-codigo">{{ carga.curso?.codigo }}</span>
              <h3 class="curso-nombre">{{ carga.curso?.nombre }}</h3>
              <div class="curso-meta">
                <span class="meta-item">
                  <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <circle cx="12" cy="12" r="10"/>
                    <polyline points="12 6 12 12 16 14"/>
                  </svg>
                  {{ carga.curso?.horas_teoria }}h teoría
                </span>
                <span class="meta-item">
                  <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <path d="M14.5 10c-.83 0-1.5-.67-1.5-1.5v-5c0-.83.67-1.5 1.5-1.5s1.5.67 1.5 1.5v5c0 .83-.67 1.5-1.5 1.5z"/>
                    <path d="M20.5 10H19V8.5c0-.83.67-1.5 1.5-1.5s1.5.67 1.5 1.5-.67 1.5-1.5 1.5z"/>
                  </svg>
                  {{ carga.curso?.horas_practica }}h práctica
                </span>
                <span class="creditos">{{ carga.curso?.creditos }} créditos</span>
              </div>
            </div>

            <div class="grupos-list">
              <div v-if="!carga.grupos || carga.grupos.length === 0" class="no-grupos">
                <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <circle cx="12" cy="12" r="10"/>
                  <line x1="12" y1="8" x2="12" y2="12"/>
                  <line x1="12" y1="16" x2="12.01" y2="16"/>
                </svg>
                Sin grupos asignados
              </div>
              <div v-for="grupo in carga.grupos" :key="grupo.id_grupo" class="grupo-item">
                <div class="grupo-info">
                  <span class="grupo-codigo">{{ grupo.codigo_grupo }}</span>
                  <span class="tipo-badge" :class="grupo.tipo_componente">{{ grupo.tipo_componente }}</span>
                  <span v-if="grupo.es_nueva_necesidad" class="nueva-necesidad-badge">Nueva Necesidad</span>
                </div>
                <div class="grupo-docente">
                  <span v-if="grupo.docente" class="docente-nombre">
                    <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                      <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"/>
                      <circle cx="12" cy="7" r="4"/>
                    </svg>
                    {{ grupo.docente }}
                  </span>
                  <span v-else class="docente-vacante">
                    <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                      <circle cx="12" cy="12" r="10"/>
                      <line x1="12" y1="8" x2="12" y2="12"/>
                      <line x1="12" y1="16" x2="12.01" y2="16"/>
                    </svg>
                    Sin docente asignado
                  </span>
                </div>
                <button class="btn btn-icon-sm" @click="editarGrupo(grupo)">
                  <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"/>
                    <path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z"/>
                  </svg>
                </button>
              </div>
            </div>
          </div>
        </div>
      </template>

      <template v-else-if="activeTab === 'docentes'">
        <div v-if="resumenDocentes.length === 0" class="empty-state">
          <div class="empty-icon">
            <svg width="80" height="80" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1">
              <path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"/>
              <circle cx="9" cy="7" r="4"/>
              <path d="M23 21v-2a4 4 0 0 0-3-3.87"/>
              <path d="M16 3.13a4 4 0 0 1 0 7.75"/>
            </svg>
          </div>
          <h2>No hay docentes con carga</h2>
          <p>Selecciona una escuela y periodo para ver el resumen</p>
        </div>

        <div v-else class="docentes-grid">
          <div v-for="doc in resumenDocentes" :key="doc.id_docente" class="docente-card">
            <div class="docente-header">
              <div class="docente-avatar">
                {{ doc.nombre.charAt(0) }}
              </div>
              <div class="docente-info">
                <h3>{{ doc.nombre }}</h3>
                <span class="docente-plaza">{{ doc.codigo_plaza }}</span>
              </div>
            </div>

            <div class="horas-progress">
              <div class="horas-info">
                <span class="horas-label">Horas Asignadas</span>
                <span class="horas-value" :class="{ 'overload': doc.horas_asignadas > 16 }">
                  {{ doc.horas_asignadas }}/16h
                </span>
              </div>
              <div class="progress-bar">
                <div class="progress-fill"
                  :class="{ 'warning': doc.horas_asignadas >= 12 && doc.horas_asignadas < 16, 'danger': doc.horas_asignadas >= 16 }"
                  :style="{ width: Math.min(doc.horas_asignadas / 16 * 100, 100) + '%' }">
                </div>
              </div>
              <span v-if="doc.horas_restantes > 0" class="horas-restantes">
                {{ doc.horas_restantes }}h disponibles
              </span>
              <span v-else-if="doc.horas_asignadas >= 16" class="horas-full">
                Carga completa
              </span>
            </div>

            <div class="cursos-list">
              <h4>Cursos Asignados</h4>
              <div v-for="curso in doc.cursos" :key="curso.id_carga" class="curso-item">
                <span class="curso-item-codigo">{{ curso.codigo_curso }}</span>
                <span class="curso-item-nombre">{{ curso.nombre_curso }}</span>
                <span class="curso-item-horas">{{ curso.horas_semanales }}h</span>
              </div>
            </div>
          </div>
        </div>
      </template>
    </div>

    <div v-if="showModal" class="modal-overlay" @click.self="closeModal">
      <div class="modal modal-lg">
        <div class="modal-header">
          <h2>
            <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"/>
              <path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z"/>
            </svg>
            {{ editingGrupo ? 'Editar' : 'Agregar' }} Grupo
          </h2>
          <button class="btn btn-icon" @click="closeModal">
            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <line x1="18" y1="6" x2="6" y2="18"/>
              <line x1="6" y1="6" x2="18" y2="18"/>
            </svg>
          </button>
        </div>
        <div class="modal-body">
          <div v-if="selectedCarga" class="selected-curso">
            <span class="curso-codigo">{{ selectedCarga.curso?.codigo }}</span>
            <span>{{ selectedCarga.curso?.nombre }}</span>
          </div>

          <div class="form-group">
            <label class="form-label">Código de Grupo</label>
            <input v-model="grupoForm.codigo_grupo" type="text" class="form-input" placeholder="Ej: Grupo A">
          </div>

          <div class="form-grid form-grid-2">
            <div class="form-group">
              <label class="form-label">Tipo de Componente</label>
              <select v-model="grupoForm.tipo_componente" class="form-input">
                <option value="TEORIA">Teoría</option>
                <option value="PRACTICA">Práctica</option>
              </select>
            </div>
            <div class="form-group">
              <label class="form-label">Matriculados Proyectados</label>
              <input v-model.number="grupoForm.matriculados_proyectados" type="number" class="form-input" min="0">
            </div>
          </div>

          <div class="form-group">
            <label class="form-label">Docente</label>
            <select v-model="grupoForm.id_docente" class="form-input">
              <option :value="null">-- Sin asignar --</option>
              <option v-for="d in docentes" :key="d.id_docente" :value="d.id_docente">
                {{ d.nombres }} {{ d.apellidos }} ({{ d.codigo_plaza }})
              </option>
            </select>
            <small v-if="docenteHoras" class="form-hint" :class="{ 'text-danger': docenteHoras.horas_asignadas >= 16 }">
              Horas actuales: {{ docenteHoras.horas_asignadas }}/16
              <span v-if="docenteHoras.horas_asignadas >= 16"> — Límite alcanzado</span>
            </small>
          </div>

          <div v-if="docenteBloques.length > 0" class="form-group">
            <label class="form-label">Horarios Asignados al Docente</label>
            <div class="bloques-info">
              <div v-for="bloque in docenteBloques" :key="bloque.id_bloque" class="bloque-item">
                <span class="bloque-curso">{{ bloque.curso_codigo }}</span>
                <span class="bloque-escuela">{{ bloque.escuela_nombre }}</span>
                <span class="bloque-dia">{{ formatDia(bloque.dia_semana) }}</span>
                <span class="bloque-hora">{{ formatSlot(bloque.slot_inicio) }}-{{ formatSlot(bloque.slot_fin) }}</span>
              </div>
            </div>
          </div>

          <div class="form-group">
            <label class="checkbox-label">
              <input v-model="grupoForm.es_nueva_necesidad" type="checkbox">
              <span class="checkbox-custom"></span>
              Marcar como "Nueva Necesidad"
            </label>
          </div>
        </div>
        <div class="modal-footer">
          <button class="btn btn-secondary" @click="closeModal">Cancelar</button>
          <button class="btn btn-success" @click="saveGrupo" :disabled="saving">
            <svg v-if="saving" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" class="spin">
              <line x1="12" y1="2" x2="12" y2="6"/>
              <line x1="12" y1="18" x2="12" y2="22"/>
              <line x1="4.93" y1="4.93" x2="7.76" y2="7.76"/>
              <line x1="16.24" y1="16.24" x2="19.07" y2="19.07"/>
              <line x1="2" y1="12" x2="6" y2="12"/>
              <line x1="18" y1="12" x2="22" y2="12"/>
              <line x1="4.93" y1="19.07" x2="7.76" y2="16.24"/>
              <line x1="16.24" y1="7.76" x2="19.07" y2="4.93"/>
            </svg>
            {{ saving ? 'Guardando...' : 'Guardar' }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
import api from '../services/api'

const activeTab = ref('cargas')
const loading = ref(false)
const error = ref('')
const cargas = ref([])
const periodos = ref([])
const escuelas = ref([])
const docentes = ref([])
const resumenDocentes = ref([])

const selectedPeriodo = ref('')
const selectedEscuela = ref('')

const showModal = ref(false)
const saving = ref(false)
const selectedCarga = ref(null)
const editingGrupo = ref(null)
const docenteHoras = ref(null)
const docenteBloques = ref([])

const grupoForm = ref({
  codigo_grupo: '',
  tipo_componente: 'TEORIA',
  id_docente: null,
  es_nueva_necesidad: false,
  matriculados_proyectados: 0
})

async function loadCatalogos() {
  try {
    const [p, e, d] = await Promise.all([
      api.periodos.list(),
      api.escuelas.list(),
      api.docentes.list()
    ])
    periodos.value = p
    escuelas.value = e
    docentes.value = d

    const activo = p.find(x => x.activo)
    if (activo) {
      selectedPeriodo.value = activo.id_periodo
    }
  } catch (e) {
    error.value = 'Error cargando catálogos: ' + e.message
  }
}

async function loadCargas() {
  if (!selectedPeriodo.value) {
    cargas.value = []
    return
  }

  loading.value = true
  error.value = ''
  try {
    const params = new URLSearchParams()
    params.append('periodo', selectedPeriodo.value)
    if (selectedEscuela.value) {
      params.append('escuela', selectedEscuela.value)
    }

    cargas.value = await api.get(`/carga-academica?${params.toString()}`)
  } catch (e) {
    error.value = e.message
  } finally {
    loading.value = false
  }
}

async function loadResumenDocentes() {
  if (!selectedEscuela.value || !selectedPeriodo.value) {
    resumenDocentes.value = []
    return
  }

  loading.value = true
  error.value = ''
  try {
    resumenDocentes.value = await api.get(`/carga-academica/resumen-docentes?escuela=${selectedEscuela.value}&periodo=${selectedPeriodo.value}`)
  } catch (e) {
    error.value = e.message
  } finally {
    loading.value = false
  }
}

function openAgregarGrupo(carga) {
  selectedCarga.value = carga
  editingGrupo.value = null
  grupoForm.value = {
    codigo_grupo: '',
    tipo_componente: 'TEORIA',
    id_docente: null,
    es_nueva_necesidad: false,
    matriculados_proyectados: carga.curso?.creditos * 2 || 30
  }
  docenteHoras.value = null
  docenteBloques.value = []
  showModal.value = true
}

function editarGrupo(grupo) {
  editingGrupo.value = grupo
  selectedCarga.value = cargas.value.find(c => c.grupos?.some(g => g.id_grupo === grupo.id_grupo))
  grupoForm.value = {
    codigo_grupo: grupo.codigo_grupo,
    tipo_componente: grupo.tipo_componente,
    id_docente: grupo.id_docente,
    es_nueva_necesidad: grupo.es_nueva_necesidad,
    matriculados_proyectados: grupo.matriculados_proyectados
  }
  showModal.value = true
}

async function saveGrupo() {
  if (!selectedCarga.value) return

  saving.value = true
  try {
    if (editingGrupo.value) {
      await api.post(`/carga-academica/grupos/${editingGrupo.value.id_grupo}`, {
        id_docente: grupoForm.value.id_docente,
        codigo_grupo: grupoForm.value.codigo_grupo,
        es_nueva_necesidad: grupoForm.value.es_nueva_necesidad,
        matriculados_proyectados: grupoForm.value.matriculados_proyectados
      })
    } else {
      await api.post(`/carga-academica/${selectedCarga.value.id_carga}/grupos`, {
        codigo_grupo: grupoForm.value.codigo_grupo,
        tipo_componente: grupoForm.value.tipo_componente,
        id_docente: grupoForm.value.id_docente,
        es_nueva_necesidad: grupoForm.value.es_nueva_necesidad,
        matriculados_proyectados: grupoForm.value.matriculados_proyectados
      })
    }
    closeModal()
    loadCargas()
  } catch (e) {
    alert('Error: ' + e.message)
  } finally {
    saving.value = false
  }
}

function closeModal() {
  showModal.value = false
  selectedCarga.value = null
  editingGrupo.value = null
  docenteBloques.value = []
}

function formatDia(dia) {
  const dias = ['', 'Lunes', 'Martes', 'Miércoles', 'Jueves', 'Viernes', 'Sábado']
  return dias[dia] || `Día ${dia}`
}

function formatSlot(slot) {
  const horas = {
    1: '07:00', 2: '08:00', 3: '09:00', 4: '10:00', 5: '11:00', 6: '12:00',
    7: '13:00', 8: '14:00', 9: '15:00', 10: '16:00', 11: '17:00', 12: '18:00',
    13: '19:00', 14: '20:00'
  }
  return horas[slot] || `${slot}:00`
}

watch(() => grupoForm.value.id_docente, async (newVal) => {
  if (newVal) {
    try {
      const [horas, bloques] = await Promise.all([
        api.get(`/carga-academica/docente/${newVal}/horas?periodo=${selectedPeriodo.value}`),
        api.get(`/disponibilidad/docente/${newVal}?periodo=${selectedPeriodo.value}`)
      ])
      docenteHoras.value = horas
      docenteBloques.value = bloques.bloques_asignados || []
    } catch {
      docenteHoras.value = null
      docenteBloques.value = []
    }
  } else {
    docenteHoras.value = null
    docenteBloques.value = []
  }
})

function refresh() {
  if (activeTab.value === 'cargas') {
    loadCargas()
  } else {
    loadResumenDocentes()
  }
}

onMounted(() => {
  loadCatalogos().then(() => {
    if (selectedPeriodo.value) {
      loadCargas()
    }
  })
})
</script>

<style scoped>
.carga-page {
  min-height: 100vh;
  background: linear-gradient(135deg, #f0f4ff 0%, #fdf2f8 50%, #f0fdf4 100%);
}

.page-header {
  background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%);
  color: white;
  padding: 24px 32px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  box-shadow: 0 8px 32px rgba(240, 147, 251, 0.3);
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

.header-actions {
  display: flex;
  gap: 12px;
  align-items: center;
}

.filter-select {
  padding: 10px 16px;
  border: 2px solid rgba(255,255,255,0.3);
  border-radius: 12px;
  background: rgba(255,255,255,0.15);
  color: white;
  font-size: 0.9rem;
  cursor: pointer;
  min-width: 180px;
  transition: all 0.3s;
}

.filter-select option {
  color: #1e293b;
  background: white;
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

.btn-success {
  background: linear-gradient(135deg, #11998e 0%, #38ef7d 100%);
  color: white;
}

.btn-success:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(17, 153, 142, 0.4);
}

.btn-sm {
  padding: 6px 12px;
  font-size: 0.8rem;
}

.btn-icon {
  background: transparent;
  border: none;
  color: #64748b;
  padding: 8px;
  border-radius: 8px;
}

.btn-icon:hover {
  background: #f1f5f9;
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
}

.tab:hover {
  background: #f0f4ff;
  color: #f5576c;
}

.tab.active {
  background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%);
  color: white;
  box-shadow: 0 4px 12px rgba(245, 87, 108, 0.3);
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
  border: 4px solid rgba(245, 87, 108, 0.2);
  border-top-color: #f5576c;
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

.empty-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 80px 20px;
  text-align: center;
  background: white;
  border-radius: 24px;
  box-shadow: 0 8px 32px rgba(0,0,0,0.08);
}

.empty-icon {
  width: 120px;
  height: 120px;
  background: linear-gradient(135deg, #f0f4ff 0%, #fdf2f8 100%);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 24px;
  color: #f5576c;
}

.empty-state h2 {
  font-size: 1.5rem;
  color: #1e293b;
  margin: 0 0 8px;
}

.empty-state p {
  color: #64748b;
  margin: 0;
}

.cargas-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(380px, 1fr));
  gap: 24px;
}

.carga-card {
  background: white;
  border-radius: 20px;
  box-shadow: 0 8px 32px rgba(0,0,0,0.08);
  overflow: hidden;
  transition: all 0.3s;
  border: 2px solid transparent;
}

.carga-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 12px 40px rgba(0,0,0,0.12);
}

.carga-card.autorizado {
  border-color: #11998e;
}

.carga-header {
  padding: 16px 20px;
  background: linear-gradient(135deg, #f8fafc 0%, #f0f4ff 100%);
  border-bottom: 1px solid #e2e8f0;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.carga-badges {
  display: flex;
  gap: 8px;
  align-items: center;
}

.estado-badge {
  padding: 4px 12px;
  border-radius: 20px;
  font-size: 0.75rem;
  font-weight: 600;
}

.estado-badge.autorizado {
  background: linear-gradient(135deg, #11998e20, #38ef7d20);
  color: #11998e;
}

.estado-badge.borrador {
  background: #fef3c7;
  color: #d97706;
}

.escuela-tag {
  font-size: 0.75rem;
  color: #64748b;
}

.carga-curso {
  padding: 20px;
}

.curso-codigo {
  display: inline-block;
  background: linear-gradient(135deg, #667eea, #764ba2);
  color: white;
  padding: 4px 12px;
  border-radius: 8px;
  font-size: 0.8rem;
  font-weight: 700;
  margin-bottom: 10px;
}

.curso-nombre {
  font-size: 1.1rem;
  font-weight: 700;
  color: #0f172a;
  margin: 0 0 10px;
}

.curso-meta {
  display: flex;
  gap: 16px;
  font-size: 0.85rem;
  color: #64748b;
}

.meta-item {
  display: flex;
  align-items: center;
  gap: 5px;
}

.creditos {
  font-weight: 600;
  color: #f5576c;
}

.grupos-list {
  padding: 0 20px 20px;
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.no-grupos {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  padding: 24px;
  color: #94a3b8;
  font-style: italic;
  background: #f8fafc;
  border-radius: 12px;
}

.grupo-item {
  background: linear-gradient(135deg, #f8fafc 0%, #f0f4ff 100%);
  border-radius: 12px;
  padding: 14px 16px;
  display: flex;
  align-items: center;
  gap: 12px;
  border: 1px solid #e2e8f0;
}

.grupo-info {
  display: flex;
  align-items: center;
  gap: 8px;
  flex: 1;
}

.grupo-codigo {
  font-weight: 700;
  color: #1e293b;
}

.tipo-badge {
  padding: 2px 8px;
  border-radius: 6px;
  font-size: 0.7rem;
  font-weight: 600;
  text-transform: uppercase;
}

.tipo-badge.TEORIA {
  background: linear-gradient(135deg, #667eea20, #764ba220);
  color: #667eea;
}

.tipo-badge.PRACTICA {
  background: linear-gradient(135deg, #11998e20, #38ef7d20);
  color: #11998e;
}

.nueva-necesidad-badge {
  background: #fef3c7;
  color: #d97706;
  padding: 2px 8px;
  border-radius: 6px;
  font-size: 0.65rem;
  font-weight: 600;
}

.grupo-docente {
  flex: 1;
}

.docente-nombre {
  display: flex;
  align-items: center;
  gap: 5px;
  color: #334155;
  font-size: 0.85rem;
}

.docente-vacante {
  display: flex;
  align-items: center;
  gap: 5px;
  color: #ef4444;
  font-size: 0.85rem;
  font-style: italic;
}

.docentes-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(340px, 1fr));
  gap: 20px;
}

.docente-card {
  background: white;
  border-radius: 20px;
  box-shadow: 0 8px 32px rgba(0,0,0,0.08);
  padding: 20px;
  transition: all 0.3s;
}

.docente-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 12px 40px rgba(0,0,0,0.12);
}

.docente-header {
  display: flex;
  align-items: center;
  gap: 14px;
  margin-bottom: 16px;
}

.docente-avatar {
  width: 52px;
  height: 52px;
  background: linear-gradient(135deg, #667eea, #764ba2);
  color: white;
  border-radius: 14px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 1.3rem;
  font-weight: 700;
}

.docente-info h3 {
  font-size: 1rem;
  font-weight: 700;
  color: #0f172a;
  margin: 0 0 2px;
}

.docente-plaza {
  font-size: 0.75rem;
  color: #64748b;
  font-family: monospace;
}

.horas-progress {
  margin-bottom: 16px;
}

.horas-info {
  display: flex;
  justify-content: space-between;
  margin-bottom: 8px;
}

.horas-label {
  font-size: 0.8rem;
  color: #64748b;
}

.horas-value {
  font-size: 0.9rem;
  font-weight: 700;
  color: #11998e;
}

.horas-value.overload {
  color: #ef4444;
}

.progress-bar {
  height: 10px;
  background: #e2e8f0;
  border-radius: 5px;
  overflow: hidden;
}

.progress-fill {
  height: 100%;
  background: linear-gradient(135deg, #11998e, #38ef7d);
  border-radius: 5px;
  transition: width 0.3s;
}

.progress-fill.warning {
  background: linear-gradient(135deg, #f59e0b, #fbbf24);
}

.progress-fill.danger {
  background: linear-gradient(135deg, #ef4444, #f87171);
}

.horas-restantes {
  font-size: 0.75rem;
  color: #11998e;
  margin-top: 6px;
  display: block;
}

.horas-full {
  font-size: 0.75rem;
  color: #11998e;
  font-weight: 500;
}

.cursos-list {
  border-top: 1px solid #f1f5f9;
  padding-top: 14px;
}

.cursos-list h4 {
  font-size: 0.75rem;
  font-weight: 700;
  color: #64748b;
  text-transform: uppercase;
  letter-spacing: 0.5px;
  margin: 0 0 10px;
}

.curso-item {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 8px 0;
  font-size: 0.85rem;
}

.curso-item-codigo {
  background: linear-gradient(135deg, #667eea20, #764ba220);
  color: #667eea;
  padding: 2px 8px;
  border-radius: 6px;
  font-size: 0.75rem;
  font-weight: 600;
  font-family: monospace;
}

.curso-item-nombre {
  flex: 1;
  color: #334155;
}

.curso-item-horas {
  font-weight: 700;
  color: #f5576c;
}

.modal-overlay {
  position: fixed;
  inset: 0;
  background: rgba(0,0,0,0.4);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
  backdrop-filter: blur(8px);
}

.modal {
  background: white;
  border-radius: 24px;
  width: 100%;
  max-width: 560px;
  max-height: 90vh;
  overflow: hidden;
  box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.25);
}

.modal-lg {
  max-width: 600px;
}

.modal-header {
  padding: 24px;
  border-bottom: 1px solid #f1f5f9;
  display: flex;
  justify-content: space-between;
  align-items: center;
  background: linear-gradient(135deg, #f8fafc 0%, #f0f4ff 100%);
}

.modal-header h2 {
  display: flex;
  align-items: center;
  gap: 12px;
  margin: 0;
  font-size: 1.25rem;
  color: #1e293b;
}

.modal-body {
  padding: 24px;
  max-height: calc(90vh - 180px);
  overflow-y: auto;
}

.modal-footer {
  padding: 20px 24px;
  border-top: 1px solid #f1f5f9;
  display: flex;
  justify-content: flex-end;
  gap: 12px;
  background: linear-gradient(135deg, #f8fafc 0%, #f0f4ff 100%);
}

.selected-curso {
  background: linear-gradient(135deg, #f8fafc 0%, #f0f4ff 100%);
  padding: 14px 18px;
  border-radius: 12px;
  margin-bottom: 20px;
  display: flex;
  gap: 12px;
  align-items: center;
  font-size: 0.9rem;
}

.selected-curso .curso-codigo {
  margin-bottom: 0;
}

.form-grid {
  display: grid;
  gap: 16px;
}

.form-grid-2 {
  grid-template-columns: repeat(2, 1fr);
}

.form-group {
  margin-bottom: 16px;
}

.form-label {
  display: block;
  font-size: 0.875rem;
  font-weight: 600;
  color: #374151;
  margin-bottom: 6px;
}

.form-input {
  width: 100%;
  padding: 12px 16px;
  border: 2px solid #e2e8f0;
  border-radius: 12px;
  font-size: 0.95rem;
  transition: all 0.3s;
  box-sizing: border-box;
}

.form-input:focus {
  outline: none;
  border-color: #f5576c;
  box-shadow: 0 0 0 4px rgba(245, 87, 108, 0.1);
}

.form-hint {
  display: block;
  margin-top: 6px;
  font-size: 0.75rem;
  color: #64748b;
}

.text-danger {
  color: #ef4444 !important;
}

.checkbox-label {
  display: flex;
  align-items: center;
  gap: 10px;
  cursor: pointer;
  font-size: 0.9rem;
  color: #374151;
}

.checkbox-label input {
  display: none;
}

.checkbox-custom {
  width: 20px;
  height: 20px;
  border: 2px solid #d1d5db;
  border-radius: 6px;
  transition: all 0.2s;
  position: relative;
}

.checkbox-label input:checked + .checkbox-custom {
  background: linear-gradient(135deg, #f093fb, #f5576c);
  border-color: #f5576c;
}

.checkbox-label input:checked + .checkbox-custom::after {
  content: '✓';
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  color: white;
  font-size: 12px;
  font-weight: 700;
}

.bloques-info {
  background: linear-gradient(135deg, #f8fafc 0%, #f0f4ff 100%);
  border-radius: 12px;
  padding: 14px;
  max-height: 160px;
  overflow-y: auto;
}

.bloque-item {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 10px;
  background: white;
  border-radius: 8px;
  margin-bottom: 8px;
  font-size: 0.85rem;
}

.bloque-item:last-child {
  margin-bottom: 0;
}

.bloque-curso {
  background: linear-gradient(135deg, #667eea, #764ba2);
  color: white;
  padding: 2px 8px;
  border-radius: 6px;
  font-size: 0.7rem;
  font-weight: 600;
}

.bloque-escuela {
  flex: 1;
  color: #64748b;
  font-size: 0.75rem;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.bloque-dia {
  color: #1e293b;
  font-weight: 500;
}

.bloque-hora {
  background: #e2e8f0;
  padding: 2px 8px;
  border-radius: 6px;
  font-size: 0.75rem;
  color: #475569;
}

.spin {
  animation: spin 1s linear infinite;
}

@media (max-width: 768px) {
  .page-header {
    flex-direction: column;
    gap: 16px;
  }
  .header-actions {
    flex-wrap: wrap;
    justify-content: center;
  }
  .cargas-grid {
    grid-template-columns: 1fr;
  }
  .form-grid-2 {
    grid-template-columns: 1fr;
  }
}
</style>
