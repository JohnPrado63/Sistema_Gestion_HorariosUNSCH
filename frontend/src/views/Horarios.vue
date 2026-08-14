<template>
  <div class="horarios-page">
    <header class="page-header">
      <div class="header-left">
        <h1 class="page-title">
          <svg width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <rect x="3" y="4" width="18" height="18" rx="2" ry="2"/>
            <line x1="16" y1="2" x2="16" y2="6"/>
            <line x1="8" y1="2" x2="8" y2="6"/>
            <line x1="3" y1="10" x2="21" y2="10"/>
          </svg>
          Gestión de Horarios
        </h1>
        <p class="page-subtitle">Crea y administra los horarios académicos</p>
      </div>
      <div class="header-right">
        <select v-model="selectedPeriodo" class="filter-select" @change="loadHorarios">
          <option value="">Periodo...</option>
          <option v-for="p in periodos" :key="p.id_periodo" :value="p.id_periodo">
            {{ p.codigo }}
          </option>
        </select>
        <select v-model="selectedEscuela" class="filter-select" @change="loadHorarios">
          <option value="">Todas las escuelas</option>
          <option v-for="e in escuelas" :key="e.id_escuela" :value="e.id_escuela">
            {{ e.nombre }}
          </option>
        </select>
        <button class="btn btn-primary" @click="openCrearHorario">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <line x1="12" y1="5" x2="12" y2="19"/>
            <line x1="5" y1="12" x2="19" y2="12"/>
          </svg>
          Nuevo Horario
        </button>
      </div>
    </header>

    <div class="page-content">
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
        <div class="card">
          <div class="card-header">
            <h2 class="card-title">Horarios Creados</h2>
            <button class="btn btn-secondary btn-sm" @click="loadHorarios">Actualizar</button>
          </div>

          <div v-if="horarios.length === 0" class="empty-state">
            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5">
              <rect x="3" y="4" width="18" height="18" rx="2" ry="2"/>
              <line x1="16" y1="2" x2="16" y2="6"/>
              <line x1="8" y1="2" x2="8" y2="6"/>
              <line x1="3" y1="10" x2="21" y2="10"/>
            </svg>
            <p>No hay horarios creados todavía</p>
            <button class="btn btn-primary" style="margin-top: 16px;" @click="openCrearHorario">
              Crear Primer Horario
            </button>
          </div>

          <div v-else class="horarios-grid">
            <div v-for="h in horarios" :key="h.id_horario" class="horario-card" :class="{ 'selected': selectedHorario?.id_horario === h.id_horario }" @click="selectHorario(h)">
              <div class="horario-header">
                <span class="badge" :class="getEstadoClass(h.estado)">{{ h.estado }}</span>
                <span class="horario-version">v{{ h.version_reajuste }}</span>
              </div>
              <div class="horario-info">
                <h3>{{ getEscuelaNombre(h.id_escuela) }}</h3>
                <p>{{ getPeriodoCodigo(h.id_periodo) }} - {{ getSerieDescripcion(getSerieNumero(h.id_serie)) }} - Semestre {{ h.semestre }}</p>
              </div>
              <div class="horario-footer" @click.stop>
                <span class="horario-bloques">{{ getBloquesCount(h.id_horario) }} bloques</span>
                <div class="horario-actions">
                  <button class="btn btn-danger btn-sm" @click.stop="eliminarHorario(h)" title="Eliminar horario">
                    <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                      <polyline points="3 6 5 6 21 6"/>
                      <path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"/>
                    </svg>
                  </button>
                  <button class="btn btn-primary btn-sm" @click.stop="openAgregarBloque(h, 1, 1)">+ Bloque</button>
                </div>
              </div>
            </div>
          </div>
        </div>

        <div v-if="selectedHorario" class="card grilla-card" style="margin-top: 24px;">
          <div class="card-header">
            <div class="grilla-header-left">
              <h2 class="card-title">Grilla Horaria</h2>
              <span class="escuela-badge">{{ getEscuelaNombre(selectedHorario.id_escuela) }}</span>
              <span class="serie-badge">{{ getSerieDescripcion(getSerieNumero(selectedHorario.id_serie)) }} - Semestre {{ selectedHorario.semestre }}</span>
              <span class="estado-badge" :class="getEstadoClass(selectedHorario.estado)">{{ selectedHorario.estado }}</span>
            </div>
            <div class="grilla-header-right">
              <select v-model="filtroAula" class="form-input filter-aula">
                <option value="">Todas las aulas</option>
                <option v-for="a in aulas" :key="a.id_aula" :value="a.id_aula">
                  {{ a.codigo }} - {{ a.nombre }}
                </option>
              </select>
              <button class="btn btn-secondary btn-sm" @click="verificarConflictos">
                <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z"/>
                </svg>
                Verificar
              </button>
            </div>
          </div>

          <div class="grilla-wrapper">
            <table class="grilla-table">
              <thead>
                <tr>
                  <th class="hora-header">Hora</th>
                  <th v-for="(dia, idx) in diasSemana" :key="idx" class="dia-header">
                    {{ dia }}
                  </th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="(hora, horaIdx) in horasLista" :key="horaIdx">
                  <td class="hora-cell">
                    <span class="hora-inicio">{{ hora }}</span>
                    <span class="hora-fin">{{ getHoraFin(horaIdx) }}</span>
                  </td>
                  <template v-for="(dia, diaIdx) in diasSemana" :key="diaIdx">
                    <td v-if="!isCellHidden(diaIdx, horaIdx)"
                        class="dia-cell"
                        :class="{ 'has-event': getEventAt(diaIdx, horaIdx), 'teoria-cell': getEventAt(diaIdx, horaIdx)?.tipo === 'TEORIA', 'practica-cell': getEventAt(diaIdx, horaIdx)?.tipo === 'PRACTICA' }"
                        :rowspan="getEventSpan(diaIdx, horaIdx)"
                        @click="handleCellClick(diaIdx, horaIdx)">
                      <div v-if="getEventCell(diaIdx, horaIdx)" class="event-card" :class="getEventCell(diaIdx, horaIdx)?.tipo">
                        <div class="event-header">
                          <span class="event-codigo">{{ getEventCell(diaIdx, horaIdx)?.codigo_curso }}</span>
                          <span class="event-tipo" :class="getEventCell(diaIdx, horaIdx)?.tipo">{{ getEventCell(diaIdx, horaIdx)?.tipo }}</span>
                          <button class="btn-delete" @click.stop="eliminarBloque(getEventCell(diaIdx, horaIdx))" title="Eliminar bloque">
                            <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                              <polyline points="3 6 5 6 21 6"/>
                              <path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"/>
                            </svg>
                          </button>
                        </div>
                        <div class="event-curso">{{ getEventCell(diaIdx, horaIdx)?.nombre_curso }}</div>
                        <div class="event-footer">
                          <div class="event-grupo">
                            <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                              <path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"/>
                              <circle cx="9" cy="7" r="4"/>
                              <path d="M23 21v-2a4 4 0 0 0-3-3.87"/>
                              <path d="M16 3.13a4 4 0 0 1 0 7.75"/>
                            </svg>
                            {{ getEventCell(diaIdx, horaIdx)?.codigo_grupo }}
                          </div>
                          <div class="event-aula" v-if="getEventCell(diaIdx, horaIdx)?.aula_codigo">
                            <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                              <path d="M3 9l9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"/>
                              <polyline points="9 22 9 12 15 12 15 22"/>
                            </svg>
                            {{ getEventCell(diaIdx, horaIdx)?.aula_codigo }}
                          </div>
                        </div>
                        <div class="event-docente" v-if="getEventCell(diaIdx, horaIdx)?.docente_nombre">
                          <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                            <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"/>
                            <circle cx="12" cy="7" r="4"/>
                          </svg>
                          {{ getEventCell(diaIdx, horaIdx)?.docente_nombre }}
                        </div>
                      </div>
                      <div v-else-if="canAddBlock(diaIdx, horaIdx)" class="add-block-hint" @click.stop="openAgregarBloque(selectedHorario, diaIdx + 1, horaIdx + 1)">
                        <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                          <line x1="12" y1="5" x2="12" y2="19"/>
                          <line x1="5" y1="12" x2="19" y2="12"/>
                        </svg>
                      </div>
                    </td>
                  </template>
                </tr>
              </tbody>
            </table>
          </div>

          <div class="grilla-legend">
            <div class="legend-item">
              <span class="legend-color teoria"></span>
              <span>Teoría</span>
            </div>
            <div class="legend-item">
              <span class="legend-color practica"></span>
              <span>Práctica</span>
            </div>
          </div>
        </div>
      </template>
    </div>

    <div v-if="showModalHorario" class="modal-overlay" @click.self="showModalHorario = false">
      <div class="modal">
        <div class="modal-header">
          <h2>Crear Horario</h2>
          <button class="btn btn-icon btn-secondary" @click="showModalHorario = false">
            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <line x1="18" y1="6" x2="6" y2="18"/>
              <line x1="6" y1="6" x2="18" y2="18"/>
            </svg>
          </button>
        </div>
        <div class="modal-body">
          <div class="form-group">
            <label class="form-label">Escuela</label>
            <select v-model="horarioForm.id_escuela" class="form-input">
              <option value="">-- Seleccionar --</option>
              <option v-for="e in escuelas" :key="e.id_escuela" :value="e.id_escuela">
                {{ e.nombre }}
              </option>
            </select>
          </div>
          <div class="form-group">
            <label class="form-label">Periodo</label>
            <select v-model="horarioForm.id_periodo" class="form-input">
              <option value="">-- Seleccionar --</option>
              <option v-for="p in periodos" :key="p.id_periodo" :value="p.id_periodo">
                {{ p.codigo }}
              </option>
            </select>
          </div>
          <div class="form-group">
            <label class="form-label">Serie</label>
            <select v-model="horarioForm.id_serie" class="form-input">
              <option value="">-- Seleccionar --</option>
              <option v-for="s in series" :key="s.id_serie" :value="s.id_serie">
                {{ s.numero_ciclo }} - {{ getSerieDescripcion(s.numero_ciclo) }}
              </option>
            </select>
          </div>
          <div class="form-group">
            <label class="form-label">Semestre</label>
            <select v-model="horarioForm.semestre" class="form-input">
              <option value="">-- Seleccionar --</option>
              <option value="I">Impar (I)</option>
              <option value="II">Par (II)</option>
            </select>
          </div>
        </div>
        <div class="modal-footer">
          <button class="btn btn-secondary" @click="showModalHorario = false">Cancelar</button>
          <button class="btn btn-primary" @click="crearHorario" :disabled="saving">
            {{ saving ? 'Creando...' : 'Crear Horario' }}
          </button>
        </div>
      </div>
    </div>

    <div v-if="showModalBloque" class="modal-overlay" @click.self="showModalBloque = false">
      <div class="modal modal-lg">
        <div class="modal-header">
          <h2>Agregar Bloque de Horario</h2>
          <button class="btn btn-icon btn-secondary" @click="showModalBloque = false">
            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <line x1="18" y1="6" x2="6" y2="18"/>
              <line x1="6" y1="6" x2="18" y2="18"/>
            </svg>
          </button>
        </div>
        <div class="modal-body">
          <div v-if="conflictoError && conflictoError.length > 0" class="validation-results" style="margin-bottom: 16px;">
            <div v-for="(c, i) in conflictoError" :key="i" class="validation-item" :class="c.severity || 'blocker'">
              <div class="validation-icon">
                <svg v-if="(c.severity || 'blocker') === 'blocker'" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <circle cx="12" cy="12" r="10"/>
                  <line x1="15" y1="9" x2="9" y2="15"/>
                  <line x1="9" y1="9" x2="15" y2="15"/>
                </svg>
                <svg v-else-if="c.severity === 'warning'" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <path d="M10.29 3.86L1.82 18a2 2 0 0 0 1.71 3h16.94a2 2 0 0 0 1.71-3L13.71 3.86a2 2 0 0 0-3.42 0z"/>
                  <line x1="12" y1="9" x2="12" y2="13"/>
                  <line x1="12" y1="17" x2="12.01" y2="17"/>
                </svg>
                <svg v-else width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <circle cx="12" cy="12" r="10"/>
                  <line x1="12" y1="16" x2="12" y2="12"/>
                  <line x1="12" y1="8" x2="12.01" y2="8"/>
                </svg>
              </div>
              <div class="validation-content">
                <span class="validation-rule">{{ c.tipo }}</span>
                <span class="validation-message">{{ c.mensaje }}</span>
              </div>
            </div>
          </div>

          <div class="selected-info" v-if="selectedHorario">
            <span class="info-label">Horario:</span>
            <span class="info-value">{{ getEscuelaNombre(selectedHorario.id_escuela) }}</span>
            <span class="info-divider">|</span>
            <span class="info-label">Día:</span>
            <span class="info-value">{{ diasSemana[bloqueForm.dia_semana - 1] }}</span>
            <span class="info-divider">|</span>
            <span class="info-label">Hora:</span>
            <span class="info-value">{{ formatHora(bloqueForm.slot_inicio) }} - {{ formatHora(bloqueForm.slot_fin) }}</span>
          </div>

          <div class="form-group">
            <label class="form-label">Grupo</label>
            <select v-model="bloqueForm.id_grupo" class="form-input" @change="onGrupoChange">
              <option value="">-- Seleccionar --</option>
              <option v-for="g in gruposDisponibles" :key="g.id_grupo" :value="g.id_grupo">
                {{ g.codigo_curso }} - {{ g.nombre_curso }} ({{ g.codigo_grupo }}) - {{ g.docente_nombre || 'Sin docente' }}
              </option>
            </select>
          </div>

          <div class="form-row">
            <div class="form-group">
              <label class="form-label">Día</label>
              <select v-model="bloqueForm.dia_semana" class="form-input">
                <option v-for="(d, i) in diasSemana" :key="i" :value="i + 1">{{ d }}</option>
              </select>
            </div>
            <div class="form-group">
              <label class="form-label">Hora Inicio</label>
              <select v-model="bloqueForm.slot_inicio" class="form-input">
                <option v-for="(h, i) in horasLista" :key="i" :value="i + 1">{{ h }}</option>
              </select>
            </div>
            <div class="form-group">
              <label class="form-label">Hora Fin</label>
              <select v-model="bloqueForm.slot_fin" class="form-input">
                <option v-for="(h, i) in horasLista" :key="i" :value="i + 1">{{ h }}</option>
              </select>
            </div>
          </div>

          <div class="form-group">
            <label class="form-label">Aula</label>
            <select v-model="bloqueForm.id_aula" class="form-input">
              <option value="">-- Seleccionar --</option>
              <option v-for="a in aulasDisponibles" :key="a.id_aula" :value="a.id_aula">
                {{ a.codigo }} - {{ a.nombre }} ({{ a.tipo }}) - Cap: {{ a.aforo }}
              </option>
            </select>
          </div>

          <div class="form-group checkbox-group">
            <label class="checkbox-label">
              <input v-model="bloqueForm.verificar_solo" type="checkbox">
              <span class="checkbox-custom"></span>
              Solo verificar sin crear
            </label>
          </div>
        </div>
        <div class="modal-footer">
          <button class="btn btn-secondary" @click="showModalBloque = false">Cancelar</button>
          <button class="btn btn-outline" @click="verificarBloque">
            <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z"/>
            </svg>
            Verificar
          </button>
          <button class="btn btn-primary" @click="crearBloque" :disabled="saving">
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
            {{ saving ? 'Guardando...' : 'Crear Bloque' }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue'
import api from '../services/api'

const loading = ref(false)
const error = ref('')
const horarios = ref([])
const bloquesRaw = ref([])
const periodos = ref([])
const escuelas = ref([])
const series = ref([])
const aulas = ref([])
const gruposDisponibles = ref([])

const selectedPeriodo = ref('')
const selectedEscuela = ref('')
const selectedHorario = ref(null)
const filtroAula = ref('')

const showModalHorario = ref(false)
const showModalBloque = ref(false)
const saving = ref(false)
const conflictoError = ref(null)

const horarioForm = ref({
  id_escuela: '',
  id_periodo: '',
  id_serie: '',
  semestre: ''
})

const bloqueForm = ref({
  id_grupo: '',
  dia_semana: 1,
  slot_inicio: 1,
  slot_fin: 2,
  id_aula: '',
  verificar_solo: false
})

const diasSemana = ['Lunes', 'Martes', 'Miércoles', 'Jueves', 'Viernes', 'Sábado']
const horasLista = ['07:00', '08:00', '09:00', '10:00', '11:00', '12:00', '13:00', '14:00', '15:00', '16:00', '17:00', '18:00', '19:00', '20:00']

const aulasDisponibles = computed(() => {
  if (!filtroAula.value) return aulas.value
  return aulas.value.filter(a => a.id_aula === parseInt(filtroAula.value))
})

const serieAulaMap = {
  100: 'H-202',
  200: 'H-203',
  300: 'H-205',
  400: 'H-206',
  500: 'H-208'
}

function getAulaSugerida(serieId) {
  if (!serieId) return null
  const serie = series.value.find(s => s.id_serie === serieId)
  if (!serie) return null
  const codigoAula = serieAulaMap[serie.numero_ciclo]
  if (!codigoAula) return null
  return aulas.value.find(a => a.codigo === codigoAula)
}

function getHoraFin(horaIdx) {
  const hora = 7 + horaIdx + 1
  return `${hora.toString().padStart(2, '0')}:00`
}

function formatHora(slot) {
  const hora = 7 + (slot - 1)
  return `${hora.toString().padStart(2, '0')}:00`
}

async function loadCatalogos() {
  try {
    const [p, e, s, a] = await Promise.all([
      api.periodos.list(),
      api.escuelas.list(),
      api.series.list(),
      api.aulas.list()
    ])
    periodos.value = p
    escuelas.value = e
    series.value = s
    aulas.value = a

    const activo = p.find(x => x.activo)
    if (activo) {
      selectedPeriodo.value = activo.id_periodo
    }
  } catch (e) {
    error.value = 'Error cargando catálogos: ' + e.message
  }
}

async function loadHorarios() {
  if (!selectedPeriodo.value) {
    horarios.value = []
    bloquesRaw.value = []
    selectedHorario.value = null
    return
  }

  loading.value = true
  error.value = ''
  try {
    const h = await api.horarios.list()
    horarios.value = h.filter(x => x.id_periodo === selectedPeriodo.value)

    if (selectedEscuela.value) {
      horarios.value = horarios.value.filter(x => x.id_escuela === selectedEscuela.value)
    }

    if (selectedHorario.value) {
      loadBloquesHorario(selectedHorario.value.id_horario)
    } else {
      bloquesRaw.value = []
    }
  } catch (e) {
    error.value = e.message
  } finally {
    loading.value = false
  }
}

function selectHorario(h) {
  selectedHorario.value = h
  loadBloquesHorario(h.id_horario)
}

async function loadBloquesHorario(idHorario) {
  try {
    const b = await api.horarios.bloques(idHorario)
    bloquesRaw.value = b || []
  } catch (e) {
    console.error('Error loading bloques:', e)
    bloquesRaw.value = []
  }
}

function getEscuelaNombre(id) {
  const e = escuelas.value.find(x => x.id_escuela === id)
  return e ? e.nombre : `Escuela ${id}`
}

function getPeriodoCodigo(id) {
  const p = periodos.value.find(x => x.id_periodo === id)
  return p ? p.codigo : `Periodo ${id}`
}

function getSerieDescripcion(numeroCiclo) {
  const map = {
    100: 'Ciclos I-II',
    200: 'Ciclos III-IV',
    300: 'Ciclos V-VI',
    400: 'Ciclos VII-VIII',
    500: 'Ciclos IX-X'
  }
  return map[numeroCiclo] || `Ciclos ${numeroCiclo}`
}

function getSerieNumero(idSerie) {
  if (!idSerie) return 0
  const s = series.value.find(x => x.id_serie === idSerie)
  return s ? s.numero_ciclo : 0
}

function getBloquesCount(idHorario) {
  if (!bloquesRaw.value) return 0
  return bloquesRaw.value.filter(x => x.id_horario === idHorario).length
}

function getBloquesHorario() {
  if (!selectedHorario.value) return []
  if (!bloquesRaw.value) return []
  let bloques = bloquesRaw.value.filter(x => x.id_horario === selectedHorario.value.id_horario)
  if (filtroAula.value) {
    bloques = bloques.filter(b => b.id_aula === parseInt(filtroAula.value))
  }
  return bloques
}

function getEventCell(dia, horaIdx) {
  const bloques = getBloquesHorario()
  const slot = horaIdx + 1
  const diaNum = dia + 1

  for (const bloque of bloques) {
    if (bloque.dia_semana === diaNum &&
        bloque.slot_inicio === slot &&
        bloque.slot_inicio < bloque.slot_fin) {
      const aula = aulas.value.find(a => a.id_aula === bloque.id_aula)
      return {
        ...bloque,
        tipo: bloque.tipo_componente,
        aula_codigo: aula?.codigo || bloque.codigo_aula || '—',
        aula_nombre: aula?.nombre || '',
        docente_nombre: bloque.nombre_docente || bloque.docente || 'Sin docente'
      }
    }
  }
  return null
}

function getEventAt(dia, horaIdx) {
  const bloques = getBloquesHorario()
  const slot = horaIdx + 1
  const diaNum = dia + 1

  return bloques.find(b =>
    b.dia_semana === diaNum &&
    slot >= b.slot_inicio &&
    slot < b.slot_fin
  )
}

function getEventSpan(dia, horaIdx) {
  const event = getEventCell(dia, horaIdx)
  if (!event) {
    const hasEvent = getEventAt(dia, horaIdx)
    if (hasEvent) return 0
    return 1
  }
  return event.slot_fin - event.slot_inicio
}

function isCellHidden(dia, horaIdx) {
  const event = getEventCell(dia, horaIdx)
  if (event) return false
  const hasEvent = getEventAt(dia, horaIdx)
  return !!hasEvent
}

function getEventStyle(dia, horaIdx) {
  const event = getEventCell(dia, horaIdx)
  if (!event) return {}
  return {
    '--event-span': event.slot_fin - event.slot_inicio
  }
}

function canAddBlock(dia, horaIdx) {
  return !getEventAt(dia, horaIdx)
}

function handleCellClick(dia, horaIdx) {
  if (canAddBlock(dia, horaIdx)) {
    openAgregarBloque(selectedHorario.value, dia + 1, horaIdx + 1)
  }
}

function getEstadoClass(estado) {
  const map = {
    'BORRADOR': 'badge-gray',
    'PRELIMINAR': 'badge-primary',
    'EN_REAJUSTE': 'badge-warning',
    'OFICIAL': 'badge-success',
    'REAJUSTADO': 'badge-success'
  }
  return map[estado] || 'badge-gray'
}

function openCrearHorario() {
  horarioForm.value = {
    id_escuela: selectedEscuela.value || '',
    id_periodo: selectedPeriodo.value || '',
    id_serie: '',
    semestre: ''
  }
  showModalHorario.value = true
}

async function crearHorario() {
  if (!horarioForm.value.id_escuela || !horarioForm.value.id_periodo || !horarioForm.value.id_serie || !horarioForm.value.semestre) {
    alert('Selecciona escuela, periodo, serie y semestre')
    return
  }

  const data = {
    id_escuela: parseInt(horarioForm.value.id_escuela),
    id_periodo: parseInt(horarioForm.value.id_periodo),
    id_serie: parseInt(horarioForm.value.id_serie),
    semestre: horarioForm.value.semestre
  }
  console.log('Creando horario con:', data)
  saving.value = true
  try {
    await api.horarios.create(data)
    showModalHorario.value = false
    loadHorarios()
  } catch (e) {
    console.error('Error creando horario:', e)
    alert('Error: ' + e.message)
  } finally {
    saving.value = false
  }
}

async function eliminarHorario(horario) {
  if (!confirm(`¿Eliminar el horario de ${getEscuelaNombre(horario.id_escuela)} - ${getPeriodoCodigo(horario.id_periodo)}?`)) {
    return
  }

  try {
    await api.horarios.delete(horario.id_horario)
    if (selectedHorario.value?.id_horario === horario.id_horario) {
      selectedHorario.value = null
    }
    loadHorarios()
  } catch (e) {
    alert('Error al eliminar: ' + e.message)
  }
}

async function openAgregarBloque(horario, dia = 1, slot = 1) {
  selectedHorario.value = horario

  const aulaSugerida = getAulaSugerida(horario.id_serie)

  bloqueForm.value = {
    id_grupo: '',
    dia_semana: dia,
    slot_inicio: slot,
    slot_fin: Math.min(slot + 1, 14),
    id_aula: aulaSugerida ? aulaSugerida.id_aula.toString() : '',
    verificar_solo: false
  }
  conflictoError.value = null
  showModalBloque.value = true

  try {
    gruposDisponibles.value = await api.gruposHorario.list(
      horario.id_escuela,
      horario.id_periodo,
      horario.id_serie,
      horario.semestre
    )
  } catch (e) {
    gruposDisponibles.value = []
  }
}

function onGrupoChange() {
  conflictoError.value = null
}

async function verificarBloque() {
  if (!bloqueForm.value.id_grupo || !bloqueForm.value.id_aula) {
    alert('Selecciona grupo y aula')
    return
  }

  saving.value = true
  conflictoError.value = null

  try {
    const grupo = gruposDisponibles.value.find(g => g.id_grupo === parseInt(bloqueForm.value.id_grupo))
    const aula = aulas.value.find(a => a.id_aula === parseInt(bloqueForm.value.id_aula))

    const existingBloques = getBloquesHorario().map(b => ({
      id: b.id_bloque,
      teacher_id: b.id_docente,
      day: b.dia_semana,
      start_slot: b.slot_inicio,
      end_slot: b.slot_fin,
      room_id: b.id_aula,
      school_id: selectedHorario.value.id_escuela,
      pavilion_id: b.id_pabellon || 0,
      series_id: 0,
      course_id: b.id_curso || 0,
      department_id: b.id_departamento || 0,
      room_capacity: aula?.aforo || 0,
      enrollment: grupo?.matriculados_proyectados || 0,
      room_shared: aula?.es_compartida || false,
      group_id: b.id_grupo,
      component_type: b.tipo_componente || '',
      course_hours_teoria: grupo?.horas_teoria || 0,
      course_hours_practica: grupo?.horas_practica || 0
    }))

    const proposed = {
      id: 0,
      teacher_id: grupo?.id_docente || 0,
      day: bloqueForm.value.dia_semana,
      start_slot: bloqueForm.value.slot_inicio,
      end_slot: bloqueForm.value.slot_fin,
      room_id: parseInt(bloqueForm.value.id_aula),
      school_id: selectedHorario.value.id_escuela,
      pavilion_id: aula?.id_pabellon || 0,
      series_id: 0,
      course_id: grupo?.id_curso || 0,
      department_id: grupo?.id_departamento || 0,
      room_capacity: aula?.aforo || 0,
      enrollment: grupo?.matriculados_proyectados || 0,
      room_shared: aula?.es_compartida || false,
      group_id: parseInt(bloqueForm.value.id_grupo),
      component_type: grupo?.tipo_componente || '',
      course_hours_teoria: grupo?.horas_teoria || 0,
      course_hours_practica: grupo?.horas_practica || 0
    }

    const result = await api.validaciones.placement({
      proposed,
      existing: existingBloques,
      state: selectedHorario.value.estado
    })

    if (result.findings && result.findings.length > 0) {
      conflictoError.value = result.findings.map(f => ({
        tipo: f.rule,
        mensaje: f.message,
        severity: f.severity.toLowerCase()
      }))
    } else {
      alert('No hay conflictos. El bloque puede crearse.')
    }
  } catch (e) {
    conflictoError.value = [{ tipo: 'ERROR', mensaje: e.message, severity: 'blocker' }]
  } finally {
    saving.value = false
  }
}

async function crearBloque() {
  if (!bloqueForm.value.id_grupo || !bloqueForm.value.id_aula) {
    alert('Selecciona grupo y aula')
    return
  }

  saving.value = true
  conflictoError.value = null

  try {
    const grupo = gruposDisponibles.value.find(g => g.id_grupo === parseInt(bloqueForm.value.id_grupo))
    const aula = aulas.value.find(a => a.id_aula === parseInt(bloqueForm.value.id_aula))

    const existingBloques = getBloquesHorario().map(b => ({
      id: b.id_bloque,
      teacher_id: b.id_docente || 0,
      day: b.dia_semana,
      start_slot: b.slot_inicio,
      end_slot: b.slot_fin,
      room_id: b.id_aula,
      school_id: selectedHorario.value.id_escuela,
      pavilion_id: b.id_aula_pabellon || 0,
      series_id: 0,
      course_id: b.id_curso || 0,
      department_id: b.id_docente_departamento || 0,
      room_capacity: 0,
      enrollment: 0,
      room_shared: false,
      group_id: b.id_grupo,
      component_type: b.tipo_componente || '',
      course_hours_teoria: b.horas_teoria || 0,
      course_hours_practica: b.horas_practica || 0
    }))

    const proposed = {
      id: 0,
      teacher_id: grupo?.id_docente || 0,
      day: bloqueForm.value.dia_semana,
      start_slot: bloqueForm.value.slot_inicio,
      end_slot: bloqueForm.value.slot_fin,
      room_id: parseInt(bloqueForm.value.id_aula),
      school_id: selectedHorario.value.id_escuela,
      pavilion_id: aula?.id_pabellon || 0,
      series_id: 0,
      course_id: grupo?.id_curso || 0,
      department_id: grupo?.id_departamento || 0,
      room_capacity: aula?.aforo || 0,
      enrollment: grupo?.matriculados_proyectados || 0,
      room_shared: aula?.es_compartida || false,
      group_id: parseInt(bloqueForm.value.id_grupo),
      component_type: grupo?.tipo_componente || '',
      course_hours_teoria: grupo?.horas_teoria || 0,
      course_hours_practica: grupo?.horas_practica || 0
    }

    const validationResult = await api.validaciones.placement({
      proposed,
      existing: existingBloques,
      state: selectedHorario.value.estado
    })

    console.log('=== DEBUG crearBloque ===')
    console.log('selectedHorario.id_horario:', selectedHorario.value.id_horario)
    console.log('bloquesRaw.length:', bloquesRaw.value.length)
    console.log('getBloquesHorario().length:', getBloquesHorario().length)
    console.log('proposed:', proposed)
    console.log('existingBloques:', existingBloques)
    console.log('validationResult:', validationResult)

    const blockers = validationResult.findings?.filter(f => f.severity === 'BLOCKER') || []
    if (blockers.length > 0) {
      conflictoError.value = blockers.map(f => ({
        tipo: f.rule,
        mensaje: f.message,
        severity: f.severity.toLowerCase()
      }))
      saving.value = false
      return
    }

    if (validationResult.findings && validationResult.findings.length > 0) {
      conflictoError.value = validationResult.findings.map(f => ({
        tipo: f.rule,
        mensaje: f.message,
        severity: f.severity.toLowerCase()
      }))
    }

    await api.bloques.create({
      id_horario: selectedHorario.value.id_horario,
      id_grupo: parseInt(bloqueForm.value.id_grupo),
      id_docente: grupo?.id_docente || null,
      id_aula: parseInt(bloqueForm.value.id_aula),
      dia_semana: bloqueForm.value.dia_semana,
      slot_inicio: bloqueForm.value.slot_inicio,
      slot_fin: bloqueForm.value.slot_fin
    })

    showModalBloque.value = false
    loadBloquesHorario(selectedHorario.value.id_horario)
  } catch (e) {
    conflictoError.value = [{ tipo: 'ERROR', mensaje: e.message, severity: 'blocker' }]
  } finally {
    saving.value = false
  }
}

async function verificarConflictos() {
  if (!selectedHorario.value) return
  alert('Verificación de conflictos iniciada para el horario ' + selectedHorario.value.id_horario)
}

async function eliminarBloque(bloque) {
  if (!confirm(`¿Eliminar el bloque de ${bloque.nombre_curso || 'este curso'} (${bloque.codigo_grupo})?`)) {
    return
  }

  try {
    console.log('Eliminando bloque con id:', bloque.id_bloque)
    const response = await api.bloques.delete(bloque.id_bloque)
    console.log('Delete response:', response)
    await loadBloquesHorario(selectedHorario.value.id_horario)
    console.log('Bloques después de eliminar:', bloquesRaw.value.length)
  } catch (e) {
    console.error('Error al eliminar:', e)
    alert('Error al eliminar: ' + e.message)
  }
}

onMounted(() => {
  loadCatalogos().then(() => {
    if (selectedPeriodo.value) {
      loadHorarios()
    }
  })
})
</script>

<style scoped>
.horarios-page {
  min-height: 100vh;
  background: linear-gradient(135deg, #f0f4ff 0%, #fdf2f8 50%, #f0fdf4 100%);
}

.page-header {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  padding: 28px 32px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  box-shadow: 0 8px 32px rgba(102, 126, 234, 0.3);
}

.header-left .page-title {
  display: flex;
  align-items: center;
  gap: 14px;
  font-size: 1.75rem;
  font-weight: 700;
  margin: 0;
}

.page-subtitle {
  margin: 6px 0 0 46px;
  opacity: 0.85;
  font-size: 0.95rem;
}

.header-right {
  display: flex;
  align-items: center;
  gap: 12px;
}

.filter-select {
  padding: 10px 16px;
  border: 2px solid rgba(255,255,255,0.3);
  border-radius: 10px;
  background: rgba(255,255,255,0.15);
  color: white;
  font-size: 0.9rem;
  cursor: pointer;
  outline: none;
}

.filter-select option {
  color: #1e293b;
  background: white;
}

.page-content {
  padding: 28px;
}

.loading, .error-alert {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 60px;
  text-align: center;
}

.spinner {
  width: 40px;
  height: 40px;
  border: 4px solid #e2e8f0;
  border-top-color: #667eea;
  border-radius: 50%;
  animation: spin 1s linear infinite;
}

@keyframes spin {
  to { transform: rotate(360deg); }
}

.error-alert {
  color: #dc2626;
}

.error-alert svg {
  margin-bottom: 12px;
}

.horarios-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: 16px;
  padding: 20px;
}

.horario-card {
  background: linear-gradient(135deg, #f8fafc 0%, #f0f4ff 100%);
  border: 2px solid transparent;
  border-radius: 14px;
  padding: 16px;
  cursor: pointer;
  transition: all 0.2s;
}

.horario-card:hover {
  border-color: #667eea40;
  transform: translateY(-2px);
  box-shadow: 0 4px 16px rgba(102, 126, 234, 0.15);
}

.horario-card.selected {
  border-color: #667eea;
  background: linear-gradient(135deg, #667eea10, #764ba210);
}

.horario-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
}

.horario-version {
  font-size: 0.75rem;
  color: #64748b;
  background: #e2e8f0;
  padding: 2px 8px;
  border-radius: 6px;
}

.badge {
  padding: 4px 10px;
  border-radius: 20px;
  font-size: 0.7rem;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.badge-gray { background: #f1f5f9; color: #64748b; }
.badge-primary { background: #dbeafe; color: #1d4ed8; }
.badge-warning { background: #fef3c7; color: #b45309; }
.badge-success { background: #dcfce7; color: #15803d; }

.horario-info h3 {
  font-size: 1rem;
  font-weight: 600;
  color: #0f172a;
  margin-bottom: 4px;
}

.horario-info p {
  font-size: 0.875rem;
  color: #64748b;
}

.horario-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-top: 12px;
  padding-top: 12px;
  border-top: 1px solid #e2e8f0;
}

.horario-bloques {
  font-size: 0.75rem;
  color: #64748b;
}

.horario-actions {
  display: flex;
  gap: 8px;
}

.grilla-card {
  overflow: hidden;
}

.card-header {
  padding: 20px 24px;
  border-bottom: 1px solid #f1f5f9;
  display: flex;
  justify-content: space-between;
  align-items: center;
  background: linear-gradient(135deg, #f8fafc 0%, #f0f4ff 100%);
}

.grilla-header-left {
  display: flex;
  align-items: center;
  gap: 12px;
}

.grilla-header-left .card-title {
  margin: 0;
  font-size: 1.1rem;
  color: #1e293b;
}

.escuela-badge {
  background: linear-gradient(135deg, #667eea, #764ba2);
  color: white;
  padding: 4px 12px;
  border-radius: 20px;
  font-size: 0.8rem;
  font-weight: 600;
}

.serie-badge {
  background: linear-gradient(135deg, #f093fb, #f5576c);
  color: white;
  padding: 4px 12px;
  border-radius: 20px;
  font-size: 0.8rem;
  font-weight: 600;
}

.grilla-header-right {
  display: flex;
  align-items: center;
  gap: 12px;
}

.filter-aula {
  width: 180px;
  padding: 8px 12px;
  border: 2px solid #e2e8f0;
  border-radius: 8px;
  font-size: 0.85rem;
  background: white;
}

.grilla-wrapper {
  overflow-x: auto;
  padding: 0;
}

.grilla-table {
  width: 100%;
  border-collapse: collapse;
  min-width: 900px;
}

.hora-header {
  width: 80px;
  background: #f1f5f9;
  padding: 12px 8px;
  font-size: 0.75rem;
  font-weight: 700;
  color: #64748b;
  text-transform: uppercase;
  letter-spacing: 0.5px;
  border-bottom: 2px solid #e2e8f0;
  text-align: center;
}

.dia-header {
  padding: 14px 10px;
  font-size: 0.9rem;
  font-weight: 700;
  text-transform: uppercase;
  letter-spacing: 0.5px;
  border-bottom: 2px solid #e2e8f0;
  text-align: center;
  background: linear-gradient(135deg, #667eea, #764ba2);
  color: white;
  width: 16.66%;
  min-width: 150px;
}

.hora-cell {
  padding: 0;
  background: #f8fafc;
  border-bottom: 1px solid #e2e8f0;
  border-right: 1px solid #e2e8f0;
  vertical-align: top;
}

.hora-cell .hora-inicio {
  display: block;
  padding: 8px 8px 2px;
  font-size: 0.85rem;
  font-weight: 700;
  color: #1e293b;
}

.hora-cell .hora-fin {
  display: block;
  padding: 0 8px 8px;
  font-size: 0.75rem;
  color: #94a3b8;
}

.dia-cell {
  height: 70px;
  width: 16.66%;
  min-width: 150px;
  border-bottom: 1px solid #e2e8f0;
  border-right: 1px solid #f1f5f9;
  padding: 0;
  position: relative;
  vertical-align: top;
  overflow: hidden;
}

.dia-cell.has-event {
  padding: 0;
}

.dia-cell:not(.has-event):hover {
  background: #f0f9ff;
  cursor: pointer;
}

.event-card {
  height: 100%;
  padding: 8px 10px;
  border-radius: 0;
  display: flex;
  flex-direction: column;
  gap: 4px;
  overflow: hidden;
  border-left: 4px solid;
  box-sizing: border-box;
}

.event-curso {
  font-size: 0.75rem;
  font-weight: 600;
  color: #1e293b;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.event-grupo, .event-aula, .event-docente {
  font-size: 0.7rem;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.event-card.TEORIA {
  background: linear-gradient(135deg, #dbeafe, #bfdbfe);
  border-left-color: #3b82f6;
}

.event-card.PRACTICA {
  background: linear-gradient(135deg, #dcfce7, #bbf7d0);
  border-left-color: #22c55e;
}

.event-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 6px;
  position: relative;
}

.btn-delete {
  background: rgba(239, 68, 68, 0.9);
  border: none;
  border-radius: 4px;
  color: white;
  cursor: pointer;
  padding: 3px 6px;
  display: flex;
  align-items: center;
  opacity: 0;
  transition: opacity 0.2s;
}

.event-card:hover .btn-delete {
  opacity: 1;
}

.btn-delete:hover {
  background: #dc2626;
}

.event-codigo {
  font-size: 0.8rem;
  font-weight: 800;
  color: #1e40af;
}

.event-tipo {
  font-size: 0.6rem;
  font-weight: 700;
  padding: 2px 6px;
  border-radius: 4px;
  text-transform: uppercase;
}

.event-tipo.TEORIA {
  background: #3b82f6;
  color: white;
}

.event-tipo.PRACTICA {
  background: #22c55e;
  color: white;
}

.event-curso {
  font-size: 0.8rem;
  font-weight: 600;
  color: #1e3a8a;
  line-height: 1.2;
  flex: 1;
  overflow: hidden;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
}

.event-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 8px;
  margin-top: auto;
}

.event-grupo, .event-aula {
  display: flex;
  align-items: center;
  gap: 4px;
  font-size: 0.7rem;
  color: #64748b;
  background: rgba(255,255,255,0.6);
  padding: 2px 6px;
  border-radius: 4px;
}

.event-docente {
  display: flex;
  align-items: center;
  gap: 4px;
  font-size: 0.7rem;
  color: #475569;
  background: rgba(255,255,255,0.5);
  padding: 2px 6px;
  border-radius: 4px;
  width: 100%;
  margin-top: 4px;
}

.add-block-hint {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100%;
  color: #cbd5e1;
  opacity: 0;
  transition: opacity 0.2s;
}

.dia-cell:hover .add-block-hint {
  opacity: 1;
}

.grilla-legend {
  display: flex;
  justify-content: center;
  gap: 24px;
  padding: 16px;
  background: #f8fafc;
  border-top: 1px solid #e2e8f0;
}

.legend-item {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 0.85rem;
  color: #64748b;
}

.legend-color {
  width: 20px;
  height: 20px;
  border-radius: 4px;
  border-left: 4px solid;
}

.legend-color.teoria {
  background: linear-gradient(135deg, #dbeafe, #bfdbfe);
  border-left-color: #3b82f6;
}

.legend-color.practica {
  background: linear-gradient(135deg, #dcfce7, #bbf7d0);
  border-left-color: #22c55e;
}

.form-row {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 16px;
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
  padding: 12px 14px;
  border: 2px solid #e2e8f0;
  border-radius: 10px;
  font-size: 0.9rem;
  transition: all 0.2s;
  background: white;
  box-sizing: border-box;
}

.form-input:focus {
  outline: none;
  border-color: #667eea;
  box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.15);
}

.checkbox-group {
  margin-top: 8px;
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
  background: linear-gradient(135deg, #667eea, #764ba2);
  border-color: #667eea;
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

.selected-info {
  background: linear-gradient(135deg, #f8fafc 0%, #f0f4ff 100%);
  padding: 12px 16px;
  border-radius: 10px;
  margin-bottom: 16px;
  display: flex;
  align-items: center;
  gap: 8px;
  flex-wrap: wrap;
  font-size: 0.9rem;
}

.info-label {
  font-weight: 600;
  color: #64748b;
}

.info-value {
  color: #1e293b;
  font-weight: 500;
}

.info-divider {
  color: #cbd5e1;
}

.modal-overlay {
  position: fixed;
  inset: 0;
  background: rgba(0,0,0,0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
  backdrop-filter: blur(4px);
}

.modal {
  background: white;
  border-radius: 20px;
  width: 100%;
  max-width: 500px;
  max-height: 90vh;
  overflow: auto;
  box-shadow: 0 25px 50px -12px rgba(0,0,0,0.25);
  animation: modalIn 0.2s ease-out;
}

.modal-lg {
  max-width: 600px;
}

@keyframes modalIn {
  from {
    opacity: 0;
    transform: scale(0.95) translateY(10px);
  }
  to {
    opacity: 1;
    transform: scale(1) translateY(0);
  }
}

.modal-header {
  padding: 24px;
  border-bottom: 1px solid #f1f5f9;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.modal-header h2 {
  font-size: 1.25rem;
  font-weight: 700;
  color: #1e293b;
  margin: 0;
}

.modal-body {
  padding: 24px;
}

.modal-footer {
  padding: 20px 24px;
  border-top: 1px solid #f1f5f9;
  display: flex;
  justify-content: flex-end;
  gap: 12px;
  background: #f8fafc;
}

.btn {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  padding: 10px 18px;
  border-radius: 10px;
  font-size: 0.9rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s;
  border: none;
}

.btn-primary {
  background: linear-gradient(135deg, #667eea, #764ba2);
  color: white;
  box-shadow: 0 4px 12px rgba(102, 126, 234, 0.3);
}

.btn-primary:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(102, 126, 234, 0.4);
}

.btn-primary:disabled {
  opacity: 0.6;
  cursor: not-allowed;
  transform: none;
}

.btn-secondary {
  background: #f1f5f9;
  color: #475569;
  border: 1px solid #e2e8f0;
}

.btn-secondary:hover {
  background: #e2e8f0;
}

.btn-danger {
  background: #ef4444;
  color: white;
}

.btn-danger:hover {
  background: #dc2626;
}

.btn-outline {
  background: white;
  color: #667eea;
  border: 2px solid #667eea;
}

.btn-outline:hover {
  background: #667eea10;
}

.btn-sm {
  padding: 6px 12px;
  font-size: 0.8rem;
}

.btn-icon {
  padding: 8px;
  border-radius: 8px;
  background: transparent;
  border: none;
}

.btn-icon:hover {
  background: #f1f5f9;
}

.validation-results {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.validation-item {
  display: flex;
  gap: 12px;
  padding: 12px 16px;
  border-radius: 10px;
  animation: slideIn 0.2s ease-out;
}

@keyframes slideIn {
  from {
    opacity: 0;
    transform: translateX(-10px);
  }
  to {
    opacity: 1;
    transform: translateX(0);
  }
}

.validation-item.blocker {
  background: linear-gradient(135deg, #fef2f2 0%, #fee2e2 100%);
  border-left: 4px solid #ef4444;
}

.validation-item.warning {
  background: linear-gradient(135deg, #fffbeb 0%, #fef3c7 100%);
  border-left: 4px solid #f59e0b;
}

.validation-item.info {
  background: linear-gradient(135deg, #eff6ff 0%, #dbeafe 100%);
  border-left: 4px solid #3b82f6;
}

.validation-icon {
  flex-shrink: 0;
}

.validation-item.blocker .validation-icon { color: #ef4444; }
.validation-item.warning .validation-icon { color: #f59e0b; }
.validation-item.info .validation-icon { color: #3b82f6; }

.validation-content {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.validation-rule {
  font-size: 0.75rem;
  font-weight: 700;
  color: #64748b;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.validation-message {
  font-size: 0.9rem;
  color: #1e293b;
}

.spin {
  animation: spin 1s linear infinite;
}

@media (max-width: 768px) {
  .page-header {
    flex-direction: column;
    gap: 16px;
  }
  .header-right {
    flex-wrap: wrap;
    justify-content: center;
  }
  .grilla-header-left {
    flex-wrap: wrap;
  }
  .form-row {
    grid-template-columns: 1fr;
  }
}
</style>
