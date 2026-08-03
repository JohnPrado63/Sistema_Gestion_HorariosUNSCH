<template>
  <div>
    <header class="topbar">
      <h1 class="page-title">Carga Académica</h1>
      <div class="page-actions">
        <select v-model="selectedPeriodo" class="form-input" style="width: 140px" @change="loadCargas">
          <option value="">Periodo...</option>
          <option v-for="p in periodos" :key="p.id_periodo" :value="p.id_periodo">
            {{ p.codigo }}
          </option>
        </select>
        <select v-model="selectedEscuela" class="form-input" style="width: 200px" @change="loadCargas">
          <option value="">Todas las escuelas</option>
          <option v-for="e in escuelas" :key="e.id_escuela" :value="e.id_escuela">
            {{ e.nombre }}
          </option>
        </select>
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
        <button class="tab" :class="{ active: activeTab === 'cargas' }" @click="activeTab = 'cargas'">
          <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" style="margin-right: 6px">
            <path d="M4 19.5A2.5 2.5 0 0 1 6.5 17H20"/>
            <path d="M6.5 2H20v20H6.5A2.5 2.5 0 0 1 4 19.5v-15A2.5 2.5 0 0 1 6.5 2z"/>
          </svg>
          Cargas ({{ cargas.length }})
        </button>
        <button class="tab" :class="{ active: activeTab === 'docentes' }" @click="activeTab = 'docentes'; loadResumenDocentes()">
          <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" style="margin-right: 6px">
            <path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"/>
            <circle cx="9" cy="7" r="4"/>
            <path d="M23 21v-2a4 4 0 0 0-3-3.87"/>
            <path d="M16 3.13a4 4 0 0 1 0 7.75"/>
          </svg>
          Resumen Docentes
        </button>
      </div>

      <!-- Loading -->
      <div v-if="loading" class="loading">
        <div class="spinner"></div>
        <span style="margin-left: 12px">Cargando...</span>
      </div>

      <!-- Error -->
      <div v-else-if="error" class="error-alert">
        <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <circle cx="12" cy="12" r="10"/>
          <line x1="12" y1="8" x2="12" y2="12"/>
          <line x1="12" y1="16" x2="12.01" y2="16"/>
        </svg>
        {{ error }}
      </div>

      <!-- Tab: Cargas -->
      <template v-else-if="activeTab === 'cargas'">
        <div v-if="cargas.length === 0" class="empty-state">
          <svg width="64" height="64" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5">
            <path d="M4 19.5A2.5 2.5 0 0 1 6.5 17H20"/>
            <path d="M6.5 2H20v20H6.5A2.5 2.5 0 0 1 4 19.5v-15A2.5 2.5 0 0 1 6.5 2z"/>
          </svg>
          <p>No hay cargas académicas registradas</p>
          <p style="font-size: 0.875rem">Selecciona un periodo y escuela para comenzar</p>
        </div>

        <div v-else class="cargas-grid">
          <div v-for="carga in cargas" :key="carga.id_carga" class="carga-card" :class="{ 'autorizado': carga.estado === 'AUTORIZADO' }">
            <div class="carga-header">
              <div>
                <span class="badge" :class="carga.estado === 'AUTORIZADO' ? 'badge-success' : 'badge-warning'">
                  {{ carga.estado }}
                </span>
                <span v-if="carga.escuela" class="escuela-tag">{{ carga.escuela }}</span>
              </div>
              <button v-if="carga.estado === 'BORRADOR'" class="btn btn-primary btn-sm" @click="openAgregarGrupo(carga)">
                + Grupo
              </button>
            </div>
            
            <div class="carga-curso">
              <span class="curso-codigo">{{ carga.curso?.codigo }}</span>
              <h3 class="curso-nombre">{{ carga.curso?.nombre }}</h3>
              <div class="curso-meta">
                <span>{{ carga.curso?.horas_teoria }}h teoría</span>
                <span>{{ carga.curso?.horas_practica }}h práctica</span>
                <span class="creditos">{{ carga.curso?.creditos }} créditos</span>
              </div>
            </div>

            <div class="grupos-list">
              <div v-if="!carga.grupos || carga.grupos.length === 0" class="no-grupos">
                Sin grupos asignados
              </div>
              <div v-for="grupo in carga.grupos" :key="grupo.id_grupo" class="grupo-item">
                <div class="grupo-info">
                  <span class="grupo-codigo">{{ grupo.codigo_grupo }}</span>
                  <span class="badge badge-gray" style="font-size: 0.7rem">{{ grupo.tipo_componente }}</span>
                  <span v-if="grupo.es_nueva_necesidad" class="badge badge-danger" style="font-size: 0.7rem">Nueva Necesidad</span>
                </div>
                <div class="grupo-docente">
                  <span v-if="grupo.docente" class="docente-nombre">{{ grupo.docente }}</span>
                  <span v-else class="docente-vacante">Sin docente asignado</span>
                </div>
                <div class="grupo-acciones">
                  <button class="btn btn-secondary btn-sm" @click="editarGrupo(grupo)">
                    <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                      <path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"/>
                      <path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z"/>
                    </svg>
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </template>

      <!-- Tab: Resumen Docentes -->
      <template v-else-if="activeTab === 'docentes'">
        <div v-if="resumenDocentes.length === 0" class="empty-state">
          <svg width="64" height="64" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5">
            <path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"/>
            <circle cx="9" cy="7" r="4"/>
            <path d="M23 21v-2a4 4 0 0 0-3-3.87"/>
            <path d="M16 3.13a4 4 0 0 1 0 7.75"/>
          </svg>
          <p>No hay docentes con carga asignada</p>
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
                ✓ Carga completa
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

    <!-- Modal Agregar/Editar Grupo -->
    <div v-if="showModal" class="modal-overlay" @click.self="closeModal">
      <div class="modal">
        <div class="modal-header">
          <h2>{{ editingGrupo ? 'Editar Grupo' : 'Agregar Grupo' }}</h2>
          <button class="btn btn-icon btn-secondary" @click="closeModal">
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

          <div class="form-group">
            <label class="form-label">Tipo de Componente</label>
            <select v-model="grupoForm.tipo_componente" class="form-input">
              <option value="TEORIA">Teoría</option>
              <option value="PRACTICA">Práctica</option>
            </select>
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
            <label class="form-label">
              <input v-model="grupoForm.es_nueva_necesidad" type="checkbox">
              Marcar como "Nueva Necesidad"
            </label>
          </div>

          <div class="form-group">
            <label class="form-label">Matriculados Proyectados</label>
            <input v-model.number="grupoForm.matriculados_proyectados" type="number" class="form-input" min="0">
          </div>
        </div>
        <div class="modal-footer">
          <button class="btn btn-secondary" @click="closeModal">Cancelar</button>
          <button class="btn btn-primary" @click="saveGrupo" :disabled="saving">
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
.cargas-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(360px, 1fr));
  gap: 24px;
}

.carga-card {
  background: white;
  border-radius: 12px;
  box-shadow: 0 1px 3px rgba(0,0,0,0.1);
  overflow: hidden;
  transition: all 0.2s ease;
  border: 2px solid transparent;
}

.carga-card:hover {
  box-shadow: 0 4px 12px rgba(0,0,0,0.15);
  transform: translateY(-2px);
}

.carga-card.autorizado {
  border-color: #10b981;
}

.carga-header {
  padding: 16px 20px;
  background: #f8fafc;
  border-bottom: 1px solid #e2e8f0;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.escuela-tag {
  margin-left: 10px;
  font-size: 0.75rem;
  color: #64748b;
}

.carga-curso {
  padding: 20px;
}

.curso-codigo {
  display: inline-block;
  background: #1e40af;
  color: white;
  padding: 2px 8px;
  border-radius: 4px;
  font-size: 0.75rem;
  font-weight: 600;
  margin-bottom: 8px;
}

.curso-nombre {
  font-size: 1.125rem;
  font-weight: 600;
  color: #0f172a;
  margin-bottom: 8px;
}

.curso-meta {
  display: flex;
  gap: 16px;
  font-size: 0.875rem;
  color: #64748b;
}

.creditos {
  font-weight: 500;
  color: #1e40af;
}

.grupos-list {
  padding: 0 20px 20px;
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.no-grupos {
  text-align: center;
  padding: 20px;
  color: #94a3b8;
  font-style: italic;
}

.grupo-item {
  background: #f8fafc;
  border-radius: 8px;
  padding: 12px 16px;
  display: flex;
  align-items: center;
  gap: 12px;
}

.grupo-info {
  display: flex;
  align-items: center;
  gap: 8px;
  flex: 1;
}

.grupo-codigo {
  font-weight: 600;
  color: #0f172a;
}

.docente-nombre {
  color: #334155;
  font-size: 0.875rem;
}

.docente-vacante {
  color: #ef4444;
  font-size: 0.875rem;
  font-style: italic;
}

.docentes-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(340px, 1fr));
  gap: 20px;
}

.docente-card {
  background: white;
  border-radius: 12px;
  box-shadow: 0 1px 3px rgba(0,0,0,0.1);
  padding: 20px;
}

.docente-header {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 16px;
}

.docente-avatar {
  width: 48px;
  height: 48px;
  background: linear-gradient(135deg, #1e40af, #3b82f6);
  color: white;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 1.25rem;
  font-weight: 600;
}

.docente-info h3 {
  font-size: 1rem;
  font-weight: 600;
  color: #0f172a;
  margin-bottom: 2px;
}

.docente-plaza {
  font-size: 0.75rem;
  color: #64748b;
}

.horas-progress {
  margin-bottom: 16px;
}

.horas-info {
  display: flex;
  justify-content: space-between;
  margin-bottom: 6px;
}

.horas-label {
  font-size: 0.75rem;
  color: #64748b;
}

.horas-value {
  font-size: 0.875rem;
  font-weight: 600;
  color: #10b981;
}

.horas-value.overload {
  color: #ef4444;
}

.progress-bar {
  height: 8px;
  background: #e2e8f0;
  border-radius: 4px;
  overflow: hidden;
}

.progress-fill {
  height: 100%;
  background: #10b981;
  border-radius: 4px;
  transition: width 0.3s ease;
}

.progress-fill.warning {
  background: #f59e0b;
}

.progress-fill.danger {
  background: #ef4444;
}

.horas-restantes {
  font-size: 0.75rem;
  color: #10b981;
  margin-top: 4px;
  display: block;
}

.horas-full {
  font-size: 0.75rem;
  color: #10b981;
  font-weight: 500;
}

.cursos-list {
  border-top: 1px solid #e2e8f0;
  padding-top: 12px;
}

.cursos-list h4 {
  font-size: 0.75rem;
  font-weight: 600;
  color: #64748b;
  text-transform: uppercase;
  letter-spacing: 0.05em;
  margin-bottom: 8px;
}

.curso-item {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 6px 0;
  font-size: 0.875rem;
}

.curso-item-codigo {
  background: #f1f5f9;
  padding: 2px 6px;
  border-radius: 4px;
  font-size: 0.75rem;
  font-weight: 500;
  color: #475569;
}

.curso-item-nombre {
  flex: 1;
  color: #334155;
}

.curso-item-horas {
  font-weight: 600;
  color: #1e40af;
}

/* Modal */
.modal-overlay {
  position: fixed;
  inset: 0;
  background: rgba(0,0,0,0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}

.modal {
  background: white;
  border-radius: 16px;
  width: 100%;
  max-width: 500px;
  max-height: 90vh;
  overflow: auto;
  box-shadow: 0 20px 25px -5px rgba(0,0,0,0.1);
}

.modal-header {
  padding: 20px 24px;
  border-bottom: 1px solid #e2e8f0;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.modal-header h2 {
  font-size: 1.25rem;
  font-weight: 600;
}

.modal-body {
  padding: 24px;
}

.modal-footer {
  padding: 16px 24px;
  border-top: 1px solid #e2e8f0;
  display: flex;
  justify-content: flex-end;
  gap: 12px;
}

.selected-curso {
  background: #f8fafc;
  padding: 12px 16px;
  border-radius: 8px;
  margin-bottom: 20px;
  display: flex;
  gap: 12px;
  align-items: center;
  font-size: 0.875rem;
}

.form-hint {
  display: block;
  margin-top: 4px;
  font-size: 0.75rem;
  color: #64748b;
}

.text-danger {
  color: #ef4444 !important;
}

.bloques-info {
  background: #f8fafc;
  border-radius: 8px;
  padding: 12px;
  max-height: 150px;
  overflow-y: auto;
}

.bloque-item {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 8px;
  background: white;
  border-radius: 6px;
  margin-bottom: 6px;
  font-size: 0.875rem;
}

.bloque-item:last-child {
  margin-bottom: 0;
}

.bloque-curso {
  background: #1e40af;
  color: white;
  padding: 2px 6px;
  border-radius: 4px;
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
  color: #0f172a;
  font-weight: 500;
}

.bloque-hora {
  background: #e2e8f0;
  padding: 2px 8px;
  border-radius: 4px;
  font-size: 0.75rem;
  color: #475569;
}
</style>
