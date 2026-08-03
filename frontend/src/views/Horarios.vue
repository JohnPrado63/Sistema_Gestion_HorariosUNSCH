<template>
  <div>
    <header class="topbar">
      <h1 class="page-title">Gestión de Horarios</h1>
      <div class="page-actions">
        <select v-model="selectedPeriodo" class="form-input" style="width: 140px" @change="loadHorarios">
          <option value="">Periodo...</option>
          <option v-for="p in periodos" :key="p.id_periodo" :value="p.id_periodo">
            {{ p.codigo }}
          </option>
        </select>
        <select v-model="selectedEscuela" class="form-input" style="width: 200px" @change="loadHorarios">
          <option value="">Todas las escuelas</option>
          <option v-for="e in escuelas" :key="e.id_escuela" :value="e.id_escuela">
            {{ e.nombre }}
          </option>
        </select>
        <button class="btn btn-primary" @click="openCrearHorario">
          <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
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
                <p>{{ getPeriodoCodigo(h.id_periodo) }}</p>
              </div>
              <div class="horario-footer" @click.stop>
                <span class="horario-bloques">{{ getBloquesCount(h.id_horario) }} bloques</span>
                <button class="btn btn-primary btn-sm" @click.stop="openAgregarBloque(h, 1, 1)">+ Bloque</button>
              </div>
            </div>
          </div>
        </div>

        <div v-if="selectedHorario" class="card" style="margin-top: 24px;">
          <div class="card-header">
            <h2 class="card-title">Grilla Horaria - {{ getEscuelaNombre(selectedHorario.id_escuela) }}</h2>
            <div style="display: flex; gap: 8px; align-items: center;">
              <select v-model="filtroDia" class="form-input" style="width: 120px;">
                <option value="">Todos los días</option>
                <option v-for="(d, i) in dias" :key="i" :value="i">{{ d }}</option>
              </select>
              <button class="btn btn-secondary btn-sm" @click="verificarConflictos">Verificar Conflictos</button>
            </div>
          </div>

          <div class="grilla-wrapper">
            <table class="grilla-table">
              <thead>
                <tr>
                  <th class="hora-col"></th>
                  <th v-for="(dia, i) in dias" :key="i" class="dia-col">{{ dia }}</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="slot in slots" :key="slot">
                  <td class="hora-cell">{{ formatSlot(slot) }}</td>
                  <td v-for="(dia, diaIdx) in dias" :key="diaIdx"
                      class="dia-cell"
                      :class="{ 'bloque-preview': getBloqueAt(slot, diaIdx) }"
                      @click="diaIdx > 0 && !getBloqueAt(slot, diaIdx) && openAgregarBloque(selectedHorario, diaIdx, slot)">
                    <div v-if="getBloqueAt(slot, diaIdx)" class="bloque-preview-inner" :class="getBloqueAt(slot, diaIdx).tipo_componente">
                      <template v-if="getBloqueAt(slot, diaIdx).isStart">
                        <span class="bloque-curso">{{ getBloqueAt(slot, diaIdx).codigo_curso }}</span>
                        <span class="bloque-nombre">{{ getBloqueAt(slot, diaIdx).nombre_curso }}</span>
                        <span class="bloque-grupo">{{ getBloqueAt(slot, diaIdx).codigo_grupo }}</span>
                      </template>
                    </div>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </template>
    </div>

    <!-- Modal Crear Horario -->
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
        </div>
        <div class="modal-footer">
          <button class="btn btn-secondary" @click="showModalHorario = false">Cancelar</button>
          <button class="btn btn-primary" @click="crearHorario" :disabled="saving">
            {{ saving ? 'Creando...' : 'Crear Horario' }}
          </button>
        </div>
      </div>
    </div>

    <!-- Modal Agregar Bloque -->
    <div v-if="showModalBloque" class="modal-overlay" @click.self="showModalBloque = false">
      <div class="modal">
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
          <div v-if="conflictoError" class="error-alert" style="margin-bottom: 16px;">
            <strong>Conflictos encontrados:</strong>
            <ul style="margin: 8px 0 0 20px;">
              <li v-for="(c, i) in conflictoError" :key="i">{{ c.mensaje }}</li>
            </ul>
          </div>

          <div class="form-group">
            <label class="form-label">Grupo</label>
            <select v-model="bloqueForm.id_grupo" class="form-input">
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
                <option v-for="(d, i) in dias" :key="i" :value="i">{{ d }}</option>
              </select>
            </div>
          </div>

          <div class="form-row">
            <div class="form-group">
              <label class="form-label">Slot Inicio</label>
              <select v-model="bloqueForm.slot_inicio" class="form-input">
                <option v-for="s in slots" :key="s" :value="s">{{ formatSlot(s) }}</option>
              </select>
            </div>
            <div class="form-group">
              <label class="form-label">Slot Fin</label>
              <select v-model="bloqueForm.slot_fin" class="form-input">
                <option v-for="s in slots" :key="s" :value="s">{{ formatSlot(s) }}</option>
              </select>
            </div>
          </div>

          <div class="form-group">
            <label class="form-label">Aula</label>
            <select v-model="bloqueForm.id_aula" class="form-input">
              <option value="">-- Seleccionar --</option>
              <option v-for="a in aulas" :key="a.id_aula" :value="a.id_aula">
                {{ a.codigo }} - {{ a.nombre }} ({{ a.tipo }})
              </option>
            </select>
          </div>

          <div class="form-group">
            <label class="form-label">
              <input v-model="bloqueForm.verificar_solo" type="checkbox">
              Solo verificar sin crear
            </label>
          </div>
        </div>
        <div class="modal-footer">
          <button class="btn btn-secondary" @click="showModalBloque = false">Cancelar</button>
          <button class="btn btn-secondary" @click="verificarBloque">Verificar</button>
          <button class="btn btn-primary" @click="crearBloque" :disabled="saving">
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
const aulas = ref([])
const gruposDisponibles = ref([])

const selectedPeriodo = ref('')
const selectedEscuela = ref('')
const selectedHorario = ref(null)
const filtroDia = ref('')

const showModalHorario = ref(false)
const showModalBloque = ref(false)
const saving = ref(false)
const conflictoError = ref(null)

const horarioForm = ref({
  id_escuela: '',
  id_periodo: ''
})

const bloqueForm = ref({
  id_grupo: '',
  dia_semana: 1,
  slot_inicio: 1,
  slot_fin: 2,
  id_aula: '',
  verificar_solo: false
})

const dias = ['', 'Lunes', 'Martes', 'Miércoles', 'Jueves', 'Viernes', 'Sábado']
const slots = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12]

async function loadCatalogos() {
  try {
    const [p, e, a] = await Promise.all([
      api.periodos.list(),
      api.escuelas.list(),
      api.aulas.list()
    ])
    periodos.value = p
    escuelas.value = e
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

    const b = await api.bloques.list()
    bloquesRaw.value = b
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
    bloquesRaw.value = b
  } catch (e) {
    console.error('Error loading bloques:', e)
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

function getBloquesCount(idHorario) {
  return bloquesRaw.value.filter(x => x.id_horario === idHorario).length
}

function getBloqueAt(slot, diaIdx) {
  if (!selectedHorario.value || diaIdx === 0) return null

  const bloques = bloquesRaw.value.filter(b =>
    b.id_horario === selectedHorario.value.id_horario &&
    b.dia_semana === diaIdx &&
    b.slot_inicio <= slot &&
    b.slot_fin > slot
  )

  if (bloques.length === 0) return null

  const bloque = bloques[0]
  const span = bloque.slot_fin - bloque.slot_inicio
  return {
    ...bloque,
    isStart: bloque.slot_inicio === slot,
    span: span
  }
}

function formatSlot(slot) {
  const hora = 6 + slot
  return `${hora.toString().padStart(2, '0')}:00`
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
  horarioForm.value = { id_escuela: selectedEscuela.value || '', id_periodo: selectedPeriodo.value }
  showModalHorario.value = true
}

async function crearHorario() {
  if (!horarioForm.value.id_escuela || !horarioForm.value.id_periodo) {
    alert('Selecciona escuela y periodo')
    return
  }

  saving.value = true
  try {
    await api.horarios.create(horarioForm.value)
    showModalHorario.value = false
    loadHorarios()
  } catch (e) {
    alert('Error: ' + e.message)
  } finally {
    saving.value = false
  }
}

async function openAgregarBloque(horario, diaIdx = 1, slotIdx = 1) {
  selectedHorario.value = horario
  bloqueForm.value = {
    id_grupo: '',
    dia_semana: diaIdx,
    slot_inicio: slotIdx,
    slot_fin: slotIdx + 1,
    id_aula: '',
    verificar_solo: false
  }
  conflictoError.value = null
  showModalBloque.value = true

  try {
    gruposDisponibles.value = await api.gruposHorario.list(horario.id_escuela, horario.id_periodo)
  } catch (e) {
    gruposDisponibles.value = []
  }
}

async function verificarBloque() {
  if (!bloqueForm.value.id_grupo || !bloqueForm.value.id_aula) {
    alert('Selecciona grupo y aula')
    return
  }

  try {
    const result = await api.bloques.verificar({
      id_horario: selectedHorario.value.id_horario,
      id_grupo: bloqueForm.value.id_grupo,
      id_aula: parseInt(bloqueForm.value.id_aula),
      dia_semana: bloqueForm.value.dia_semana,
      slot_inicio: bloqueForm.value.slot_inicio,
      slot_fin: bloqueForm.value.slot_fin
    })

    if (result.tiene_conflicto) {
      conflictoError.value = result.conflictos
    } else {
      conflictoError.value = null
      alert('No hay conflictos. El bloque puede crearse.')
    }
  } catch (e) {
    alert('Error: ' + e.message)
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
    await api.bloques.create({
      id_horario: selectedHorario.value.id_horario,
      id_grupo: parseInt(bloqueForm.value.id_grupo),
      id_aula: parseInt(bloqueForm.value.id_aula),
      dia_semana: bloqueForm.value.dia_semana,
      slot_inicio: bloqueForm.value.slot_inicio,
      slot_fin: bloqueForm.value.slot_fin
    })

    showModalBloque.value = false
    loadBloquesHorario(selectedHorario.value.id_horario)
  } catch (e) {
    const err = e.message
    if (err.includes('conflictos')) {
      try {
        const result = JSON.parse(err.replace('Error: ', ''))
        conflictoError.value = result.conflictos || []
      } catch {
        conflictoError.value = [{ mensaje: err }]
      }
    } else {
      alert('Error: ' + err)
    }
  } finally {
    saving.value = false
  }
}

async function verificarConflictos() {
  if (!selectedHorario.value) return
  alert('Verificación de conflictos iniciada para el horario ' + selectedHorario.value.id_horario)
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
.horarios-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: 16px;
  padding: 20px;
}

.horario-card {
  background: #f8fafc;
  border: 2px solid #e2e8f0;
  border-radius: 12px;
  padding: 16px;
  cursor: pointer;
  transition: all 0.2s;
}

.horario-card:hover {
  border-color: #3b82f6;
  transform: translateY(-2px);
}

.horario-card.selected {
  border-color: #1e40af;
  background: #eff6ff;
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
}

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

.grilla-wrapper {
  overflow-x: auto;
  padding: 20px;
}

.grilla-table {
  width: 100%;
  border-collapse: collapse;
  min-width: 700px;
}

.hora-col {
  width: 60px;
  background: #f8fafc;
}

.dia-col {
  width: 120px;
  background: #f8fafc;
  padding: 12px 8px;
  font-weight: 600;
  color: #475569;
  text-align: center;
  border: 1px solid #e2e8f0;
}

.hora-cell {
  background: #f8fafc;
  padding: 8px 4px;
  font-size: 0.7rem;
  color: #64748b;
  text-align: center;
  border: 1px solid #e2e8f0;
  white-space: nowrap;
}

.dia-cell {
  background: white;
  height: 50px;
  cursor: pointer;
  transition: background 0.2s;
  border: 1px solid #e2e8f0;
  padding: 2px;
}

.dia-cell:hover {
  background: #f0f9ff;
}

.bloque-preview {
  padding: 0;
  background: transparent;
}

.bloque-preview-inner {
  height: 100%;
  padding: 4px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  border-radius: 4px;
  font-size: 0.65rem;
  overflow: hidden;
}

.bloque-preview-inner.TEORIA {
  background: #dbeafe;
  border: 2px solid #3b82f6;
}

.bloque-preview-inner.PRACTICA {
  background: #dcfce7;
  border: 2px solid #22c55e;
}

.bloque-curso {
  font-weight: 700;
  color: #1e40af;
  font-size: 0.7rem;
}

.bloque-nombre {
  color: #1e3a8a;
  font-size: 0.6rem;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  max-width: 100%;
  text-align: center;
}

.bloque-grupo {
  color: #3b82f6;
  font-size: 0.6rem;
}

.form-row {
  display: flex;
  gap: 16px;
}

.form-row .form-group {
  flex: 1;
}

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

.form-group {
  margin-bottom: 16px;
}

.form-label {
  display: block;
  font-size: 0.875rem;
  font-weight: 500;
  color: #374151;
  margin-bottom: 4px;
}

.form-input {
  width: 100%;
  padding: 8px 12px;
  border: 1px solid #d1d5db;
  border-radius: 6px;
  font-size: 0.875rem;
}
</style>
