<template>
  <div class="validaciones-page">
    <header class="page-header">
      <div class="header-left">
        <h1 class="page-title">
          <svg width="28" height="28" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z"/>
          </svg>
          Motor de Validaciones
        </h1>
        <p class="page-subtitle">Validar asignaciones de bloques horarios</p>
      </div>
    </header>

    <div class="page-content">
      <div class="validation-grid">
        <!-- Columna Izquierda: Seleccion y Formulario -->
        <div class="validation-left">
          <!-- Seleccion de Horario -->
          <div class="card">
            <div class="card-header">
              <h2 class="card-title">
                <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <rect x="3" y="4" width="18" height="18" rx="2" ry="2"/>
                  <line x1="16" y1="2" x2="16" y2="6"/>
                  <line x1="8" y1="2" x2="8" y2="6"/>
                  <line x1="3" y1="10" x2="21" y2="10"/>
                </svg>
                Seleccionar Horario
              </h2>
            </div>
            <div class="card-body">
              <div class="form-row">
                <div class="form-group">
                  <label class="form-label">Escuela *</label>
                  <select v-model="filtro.escuela" class="form-input" @change="onEscuelaChange">
                    <option value="">-- Seleccionar escuela --</option>
                    <option v-for="e in escuelas" :key="e.id_escuela" :value="e.id_escuela">
                      {{ e.nombre }}
                    </option>
                  </select>
                </div>
                <div class="form-group">
                  <label class="form-label">Periodo *</label>
                  <select v-model="filtro.periodo" class="form-input" @change="loadHorarios">
                    <option value="">-- Seleccionar periodo --</option>
                    <option v-for="p in periodos" :key="p.id_periodo" :value="p.id_periodo">
                      {{ p.codigo }} {{ p.activo ? '(Activo)' : '' }}
                    </option>
                  </select>
                </div>
              </div>
              <div class="form-row">
                <div class="form-group">
                  <label class="form-label">Serie</label>
                  <select v-model="filtro.serie" class="form-input" @change="loadHorarios" :disabled="!filtro.escuela">
                    <option value="">-- Todas las series --</option>
                    <option v-for="s in seriesFiltradas" :key="s.id_serie" :value="s.id_serie">
                      {{ s.numero_ciclo }}-{{ s.subciclo === 1 ? 'I' : 'II' }} ({{ s.codigo_plan }})
                    </option>
                  </select>
                </div>
                <div class="form-group">
                  <label class="form-label">Semestre</label>
                  <select v-model="filtro.semestre" class="form-input" @change="loadHorarios">
                    <option value="">-- Ambos --</option>
                    <option value="I">Impar (I)</option>
                    <option value="II">Par (II)</option>
                  </select>
                </div>
              </div>
            </div>
          </div>

          <!-- Bloques Existentes del Horario -->
          <div class="card" v-if="horarioSeleccionado">
            <div class="card-header">
              <h2 class="card-title">
                <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <path d="M22 11.08V12a10 10 0 1 1-5.93-9.14"/>
                  <polyline points="22 4 12 14.01 9 11.01"/>
                </svg>
                Bloques Actuales del Horario
              </h2>
            </div>
            <div class="card-body">
              <div v-if="bloquesExistentes.length === 0" class="empty-message">
                No hay bloques asignados en este horario
              </div>
              <div v-else class="bloques-list">
                <div v-for="bloque in bloquesExistentes" :key="bloque.id_bloque" class="bloque-item">
                  <div class="bloque-info">
                    <span class="bloque-curso">{{ bloque.codigo_curso }}</span>
                    <span class="bloque-nombre">{{ bloque.nombre_curso }}</span>
                    <span class="bloque-grupo">{{ bloque.codigo_grupo }}</span>
                    <span class="bloque-tipo" :class="'tipo-' + bloque.tipo_componente.toLowerCase()">{{ bloque.tipo_componente }}</span>
                  </div>
                  <div class="bloque-detalles">
                    <span>{{ getDiaNombre(bloque.dia_semana) }}</span>
                    <span>Slot {{ bloque.slot_inicio }} - {{ bloque.slot_fin }}</span>
                    <span>{{ bloque.codigo_aula }}</span>
                    <span>{{ bloque.nombre_docente || 'Sin docente' }}</span>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- Formulario Nueva Propuesta -->
          <div class="card">
            <div class="card-header">
              <h2 class="card-title">
                <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <polygon points="13 2 3 14 12 14 11 22 21 10 12 10 13 2"/>
                </svg>
                Proponer Nuevo Bloque
              </h2>
            </div>
            <div class="card-body">
              <div v-if="!horarioSeleccionado" class="alert-message">
                Selecciona un horario primero para proponer un bloque
              </div>
              <template v-else>
                <div class="form-row">
                  <div class="form-group">
                    <label class="form-label">Grupo *</label>
                    <select v-model="propuesta.id_grupo" class="form-input" @change="onGrupoChange">
                      <option value="">-- Seleccionar grupo --</option>
                      <option v-for="g in gruposDisponibles" :key="g.id_grupo" :value="g.id_grupo">
                        {{ g.codigo_curso }} - {{ g.codigo_grupo }} ({{ g.tipo_componente }})
                      </option>
                    </select>
                  </div>
                  <div class="form-group">
                    <label class="form-label">Aula *</label>
                    <select v-model="propuesta.id_aula" class="form-input" @change="onAulaChange">
                      <option value="">-- Seleccionar aula --</option>
                      <option v-for="a in aulasDisponibles" :key="a.id_aula" :value="a.id_aula">
                        {{ a.codigo }} ({{ a.tipo }}) - Cap: {{ a.aforo }}
                      </option>
                    </select>
                  </div>
                </div>
                <div class="form-row">
                  <div class="form-group">
                    <label class="form-label">Día *</label>
                    <select v-model="propuesta.dia_semana" class="form-input">
                      <option value="">-- Seleccionar día --</option>
                      <option v-for="(d, i) in diasSemana" :key="i" :value="i + 1">{{ d }}</option>
                    </select>
                  </div>
                  <div class="form-group">
                    <label class="form-label">Hora Inicio *</label>
                    <select v-model="propuesta.slot_inicio" class="form-input">
                      <option value="">-- Seleccionar --</option>
                      <option v-for="(h, i) in horasLista" :key="i" :value="i + 1">{{ h }}</option>
                    </select>
                  </div>
                  <div class="form-group">
                    <label class="form-label">Hora Fin *</label>
                    <select v-model="propuesta.slot_fin" class="form-input">
                      <option value="">-- Seleccionar --</option>
                      <option v-for="(h, i) in horasLista" :key="i" :value="i + 1">{{ h }}</option>
                    </select>
                  </div>
                </div>
                <div class="form-actions">
                  <button class="btn btn-primary btn-lg" @click="validarPropuesta" :disabled="!puedeValidar || loading">
                    <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                      <polygon points="5 3 19 12 5 21 5 3"/>
                    </svg>
                    Validar Propuesta
                  </button>
                </div>
              </template>
            </div>
          </div>
        </div>

        <!-- Columna Derecha: Resultados y Reglas -->
        <div class="validation-right">
          <!-- Resultado de Validacion -->
          <div class="card result-card" v-if="resultado !== null">
            <div class="card-header">
              <h2 class="card-title">
                <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z"/>
                </svg>
                Resultado de Validación
              </h2>
              <span class="result-badge" :class="resultado.findings?.length === 0 ? 'success' : 'has-issues'">
                {{ resultado.findings?.length === 0 ? 'Sin problemas' : resultado.findings?.length + ' hallazgo(s)' }}
              </span>
            </div>
            <div class="card-body">
              <div v-if="loading" class="loading-state">
                <div class="spinner-lg"></div>
                <p>Validando...</p>
              </div>

              <div v-else-if="resultado.findings?.length === 0" class="success-state">
                <svg width="64" height="64" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <path d="M22 11.08V12a10 10 0 1 1-5.93-9.14"/>
                  <polyline points="22 4 12 14.01 9 11.01"/>
                </svg>
                <p>Sin observaciones - La asignación puede continuar</p>
              </div>

              <div v-else class="findings-list">
                <div v-for="(finding, idx) in resultado.findings" :key="idx" class="finding-item" :class="'severity-' + finding.severity.toLowerCase()">
                  <div class="finding-header">
                    <span class="finding-rule">{{ finding.rule }}</span>
                    <span class="finding-severity" :class="'severity-' + finding.severity.toLowerCase()">
                      {{ finding.severity }}
                    </span>
                  </div>
                  <p class="finding-message">{{ finding.message }}</p>
                </div>
              </div>
            </div>
          </div>

          <!-- Reglas Disponibles -->
          <div class="card rules-card">
            <div class="card-header">
              <h2 class="card-title">
                <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z"/>
                </svg>
                Reglas Disponibles
              </h2>
            </div>
            <div class="card-body">
              <div class="rules-list">
                <div class="rule-item blocker">
                  <span class="rule-code">RV-01</span>
                  <span>Conflicto de docente</span>
                </div>
                <div class="rule-item blocker">
                  <span class="rule-code">RV-02</span>
                  <span>Conflicto de aula</span>
                </div>
                <div class="rule-item blocker">
                  <span class="rule-code">RV-03</span>
                  <span>Sesión de departamento</span>
                </div>
                <div class="rule-item blocker">
                  <span class="rule-code">RV-04a</span>
                  <span>Sin tiempo de traslado</span>
                </div>
                <div class="rule-item warning">
                  <span class="rule-code">RV-04b</span>
                  <span>Tiempo traslado insuficiente</span>
                </div>
                <div class="rule-item warning">
                  <span class="rule-code">RV-05</span>
                  <span>Carga lectiva > 16h</span>
                </div>
                <div class="rule-item info">
                  <span class="rule-code">RV-06</span>
                  <span>Misma serie mismo horario</span>
                </div>
                <div class="rule-item blocker">
                  <span class="rule-code">RV-07</span>
                  <span>Aula compartida reservada</span>
                </div>
                <div class="rule-item blocker">
                  <span class="rule-code">RV-08</span>
                  <span>Justificación requerida</span>
                </div>
                <div class="rule-item warning">
                  <span class="rule-code">RV-09</span>
                  <span>Matrícula > Capacidad aula</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue'
import api from '../services/api'

const diasSemana = ['Lunes', 'Martes', 'Miércoles', 'Jueves', 'Viernes', 'Sábado']
const horasLista = ['07:00', '08:00', '09:00', '10:00', '11:00', '12:00', '13:00', '14:00', '15:00', '16:00', '17:00', '18:00', '19:00', '20:00']

const escuelas = ref([])
const periodos = ref([])
const series = ref([])
const aulas = ref([])
const horarios = ref([])
const gruposDisponibles = ref([])

const bloquesExistentes = ref([])
const resultado = ref(null)
const loading = ref(false)
const horarioSeleccionado = ref(null)

const filtro = ref({
  escuela: '',
  periodo: '',
  serie: '',
  semestre: ''
})

const propuesta = ref({
  id_grupo: '',
  id_aula: '',
  dia_semana: '',
  slot_inicio: '',
  slot_fin: ''
})

const seriesFiltradas = computed(() => {
  if (!filtro.value.escuela) return []
  return series.value.filter(s => {
    const plan = plansCache.value.find(p => p.id_plan === s.id_plan)
    return plan && plan.id_escuela === parseInt(filtro.value.escuela)
  })
})

const plansCache = ref([])

const aulasDisponibles = computed(() => {
  return aulas.value.filter(a => a.activo)
})

const puedeValidar = computed(() => {
  return propuesta.value.id_grupo &&
    propuesta.value.id_aula &&
    propuesta.value.dia_semana &&
    propuesta.value.slot_inicio &&
    propuesta.value.slot_fin &&
    horarioSeleccionado.value
})

async function loadCatalogos() {
  try {
    const [esc, per, ser, aul, plans] = await Promise.all([
      api.escuelas.list(),
      api.periodos.list(),
      api.series.list(),
      api.aulas.listAll(),
      api.get('/planes-estudio')
    ])
    escuelas.value = esc
    periodos.value = per
    series.value = ser
    aulas.value = aul
    plansCache.value = plans
  } catch (e) {
    console.error('Error cargando catalogos:', e)
  }
}

async function onEscuelaChange() {
  filtro.value.serie = ''
  await loadSeriesPorEscuela()
  loadHorarios()
}

async function loadSeriesPorEscuela() {
  if (!filtro.value.escuela) {
    series.value = []
    return
  }
  try {
    const ser = await api.get(`/series?escuela=${filtro.value.escuela}`)
    series.value = ser
  } catch (e) {
    console.error('Error cargando series:', e)
  }
}

async function loadHorarios() {
  if (!filtro.value.escuela || !filtro.value.periodo) {
    horarios.value = []
    horarioSeleccionado.value = null
    bloquesExistentes.value = []
    return
  }

  try {
    const h = await api.horarios.list()
    horarios.value = h.filter(hor => {
      if (hor.id_escuela !== parseInt(filtro.value.escuela)) return false
      if (hor.id_periodo !== parseInt(filtro.value.periodo)) return false
      if (filtro.value.serie && hor.id_serie !== parseInt(filtro.value.serie)) return false
      if (filtro.value.semestre && hor.semestre !== filtro.value.semestre) return false
      return true
    })

    if (horarios.value.length === 1) {
      horarioSeleccionado.value = horarios.value[0]
      await loadBloquesDelHorario()
      await loadGruposDisponibles()
    } else {
      horarioSeleccionado.value = null
      bloquesExistentes.value = []
    }
  } catch (e) {
    console.error('Error cargando horarios:', e)
  }
}

async function loadBloquesDelHorario() {
  if (!horarioSeleccionado.value) return
  try {
    const bloques = await api.horarios.bloques(horarioSeleccionado.value.id_horario)
    bloquesExistentes.value = bloques
  } catch (e) {
    console.error('Error cargando bloques:', e)
  }
}

async function loadGruposDisponibles() {
  if (!horarioSeleccionado.value) return
  try {
    const grupos = await api.gruposHorario.list(
      horarioSeleccionado.value.id_escuela,
      horarioSeleccionado.value.id_periodo,
      horarioSeleccionado.value.id_serie || '',
      horarioSeleccionado.value.semestre || ''
    )
    gruposDisponibles.value = grupos
  } catch (e) {
    console.error('Error cargando grupos:', e)
  }
}

function onGrupoChange() {
  const grupo = gruposDisponibles.value.find(g => g.id_grupo === parseInt(propuesta.value.id_grupo))
  if (grupo) {
    propuesta.value.grupo_data = grupo
  }
}

function onAulaChange() {
  const aula = aulas.value.find(a => a.id_aula === parseInt(propuesta.value.id_aula))
  if (aula) {
    propuesta.value.aula_data = aula
  }
}

async function validarPropuesta() {
  if (!puedeValidar.value) return

  loading.value = true
  resultado.value = null

  try {
    const grupo = gruposDisponibles.value.find(g => g.id_grupo === parseInt(propuesta.value.id_grupo))
    const aula = aulas.value.find(a => a.id_aula === parseInt(propuesta.value.id_aula))

    const existingBlocks = bloquesExistentes.value.map(b => ({
      id: b.id_bloque,
      schedule_id: horarioSeleccionado.value.id_horario,
      school_id: horarioSeleccionado.value.id_escuela,
      group_id: b.id_grupo,
      course_id: b.id_curso,
      series_id: 0,
      teacher_id: b.id_docente || 0,
      department_id: b.id_docente_departamento || 0,
      room_id: b.id_aula,
      room_shared: aula?.es_compartida || false,
      pavilion_id: b.id_aula_pabellon || 0,
      day: b.dia_semana,
      start_slot: b.slot_inicio,
      end_slot: b.slot_fin,
      enrollment: grupo?.matriculados_proyectados || 0,
      room_capacity: aula?.aforo || 0,
      component_type: b.tipo_componente || '',
      course_hours_teoria: b.horas_teoria || 0,
      course_hours_practica: b.horas_practica || 0
    }))

    const proposedBlock = {
      id: 0,
      schedule_id: horarioSeleccionado.value.id_horario,
      school_id: horarioSeleccionado.value.id_escuela,
      group_id: parseInt(propuesta.value.id_grupo),
      course_id: grupo?.id_curso || 0,
      series_id: 0,
      teacher_id: grupo?.id_docente || 0,
      department_id: 0,
      room_id: parseInt(propuesta.value.id_aula),
      room_shared: aula?.es_compartida || false,
      pavilion_id: aula?.id_pabellon || 0,
      day: parseInt(propuesta.value.dia_semana),
      start_slot: parseInt(propuesta.value.slot_inicio),
      end_slot: parseInt(propuesta.value.slot_fin),
      enrollment: grupo?.matriculados_proyectados || 0,
      room_capacity: aula?.aforo || 0,
      component_type: grupo?.tipo_componente || '',
      course_hours_teoria: grupo?.horas_teoria || 0,
      course_hours_practica: grupo?.horas_practica || 0
    }

    const payload = {
      proposed: proposedBlock,
      existing: existingBlocks,
      state: horarioSeleccionado.value.estado || 'BORRADOR',
      department_sessions: [],
      distances: []
    }

    resultado.value = await api.validaciones.placement(payload)
  } catch (e) {
    alert('Error validando: ' + e.message)
  } finally {
    loading.value = false
  }
}

function getDiaNombre(dia) {
  return diasSemana[dia - 1] || 'Día ' + dia
}

watch(() => filtro.value.escuela, () => {
  filtro.value.periodo = ''
  filtro.value.serie = ''
  filtro.value.semestre = ''
  horarios.value = []
  horarioSeleccionado.value = null
  bloquesExistentes.value = []
  gruposDisponibles.value = []
  resultado.value = null
})

watch(() => filtro.value.periodo, () => {
  loadHorarios()
})

onMounted(loadCatalogos)
</script>

<style scoped>
.validaciones-page {
  min-height: 100vh;
  background: linear-gradient(135deg, #f0f4ff 0%, #fdf2f8 50%, #f0fdf4 100%);
}

.page-header {
  background: linear-gradient(135deg, #5C0000 0%, #8B0000 100%);
  color: white;
  padding: 24px 32px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  box-shadow: 0 8px 32px rgba(139, 0, 0, 0.4);
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

.page-content {
  padding: 28px;
}

.validation-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 24px;
}

.validation-left,
.validation-right {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.card {
  background: white;
  border-radius: 20px;
  box-shadow: 0 8px 32px rgba(0,0,0,0.08);
  overflow: hidden;
}

.card-header {
  padding: 16px 20px;
  border-bottom: 1px solid #f1f5f9;
  background: linear-gradient(135deg, #f8fafc 0%, #f0f4ff 100%);
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.card-title {
  display: flex;
  align-items: center;
  gap: 10px;
  margin: 0;
  font-size: 1rem;
  color: #1e293b;
  font-weight: 700;
}

.card-body {
  padding: 20px;
}

.form-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 16px;
  margin-bottom: 16px;
}

.form-row:has(.form-group:nth-child(3)) {
  grid-template-columns: 1fr 1fr 1fr;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.form-label {
  font-size: 0.85rem;
  font-weight: 600;
  color: #475569;
}

.form-input {
  padding: 10px 14px;
  border: 2px solid #e2e8f0;
  border-radius: 10px;
  font-size: 0.9rem;
  transition: all 0.3s;
}

.form-input:focus {
  outline: none;
  border-color: #667eea;
  box-shadow: 0 0 0 4px rgba(102, 126, 234, 0.1);
}

.form-actions {
  margin-top: 16px;
}

.btn {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  padding: 12px 20px;
  border-radius: 12px;
  font-size: 0.9rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s;
  border: none;
  box-shadow: 0 4px 12px rgba(0,0,0,0.1);
}

.btn-primary {
  background: linear-gradient(135deg, #667eea, #764ba2);
  color: white;
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

.btn-lg {
  padding: 14px 24px;
  font-size: 1rem;
}

.empty-message,
.alert-message {
  text-align: center;
  padding: 20px;
  color: #64748b;
  background: #f8fafc;
  border-radius: 12px;
}

.bloques-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.bloque-item {
  padding: 12px;
  background: #f8fafc;
  border-radius: 10px;
  border-left: 4px solid #667eea;
}

.bloque-info {
  display: flex;
  gap: 8px;
  margin-bottom: 8px;
  flex-wrap: wrap;
  align-items: center;
}

.bloque-curso {
  font-weight: 700;
  color: #1e293b;
}

.bloque-nombre {
  color: #64748b;
  font-size: 0.9rem;
}

.bloque-grupo {
  color: #64748b;
}

.bloque-tipo {
  padding: 2px 8px;
  border-radius: 6px;
  font-size: 0.75rem;
  font-weight: 600;
}

.bloque-tipo.tipo-teoria {
  background: #dbeafe;
  color: #2563eb;
}

.bloque-tipo.tipo-practica {
  background: #fef3c7;
  color: #d97706;
}

.bloque-detalles {
  display: flex;
  gap: 16px;
  font-size: 0.85rem;
  color: #64748b;
}

.loading-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 40px;
  color: #64748b;
}

.spinner-lg {
  width: 48px;
  height: 48px;
  border: 4px solid rgba(102, 126, 234, 0.2);
  border-top-color: #667eea;
  border-radius: 50%;
  animation: spin 1s linear infinite;
  margin-bottom: 16px;
}

@keyframes spin {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}

.success-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 40px;
  background: linear-gradient(135deg, #f0fdf4, #dcfce7);
  border-radius: 16px;
  color: #166534;
  text-align: center;
}

.success-state svg {
  color: #22c55e;
  margin-bottom: 16px;
}

.result-badge {
  padding: 6px 14px;
  border-radius: 20px;
  font-size: 0.8rem;
  font-weight: 600;
}

.result-badge.success {
  background: linear-gradient(135deg, #dcfce7, #bbf7d0);
  color: #166534;
}

.result-badge.has-issues {
  background: linear-gradient(135deg, #fef3c7, #fde68a);
  color: #92400e;
}

.findings-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.finding-item {
  padding: 14px;
  border-radius: 12px;
  border-left: 4px solid;
}

.finding-item.severity-blocker {
  background: #fef2f2;
  border-color: #ef4444;
}

.finding-item.severity-warning {
  background: #fffbeb;
  border-color: #f59e0b;
}

.finding-item.severity-info {
  background: #eff6ff;
  border-color: #3b82f6;
}

.finding-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 8px;
}

.finding-rule {
  font-weight: 700;
  font-family: monospace;
  font-size: 0.9rem;
  color: #1e293b;
}

.finding-severity {
  padding: 4px 10px;
  border-radius: 6px;
  font-size: 0.75rem;
  font-weight: 600;
}

.finding-severity.severity-blocker {
  background: #fee2e2;
  color: #dc2626;
}

.finding-severity.severity-warning {
  background: #fef3c7;
  color: #d97706;
}

.finding-severity.severity-info {
  background: #dbeafe;
  color: #2563eb;
}

.finding-message {
  margin: 0;
  font-size: 0.9rem;
  color: #475569;
}

.rules-list {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.rule-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 10px 12px;
  border-radius: 8px;
  font-size: 0.85rem;
}

.rule-item.blocker {
  background: #fef2f2;
}

.rule-item.warning {
  background: #fffbeb;
}

.rule-item.info {
  background: #eff6ff;
}

.rule-code {
  font-weight: 700;
  font-family: monospace;
  color: #667eea;
}

@media (max-width: 1024px) {
  .validation-grid {
    grid-template-columns: 1fr;
  }
}
</style>
