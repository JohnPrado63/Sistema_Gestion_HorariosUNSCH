<template>
  <div>
    <header class="topbar">
      <h1 class="page-title">Gestión de Horarios</h1>
      <div class="page-actions">
        <button class="btn btn-primary">
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
            <button class="btn btn-secondary btn-sm" @click="loadHorarios">
              Actualizar
            </button>
          </div>
          
          <div v-if="horarios.length === 0" class="empty-state">
            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5">
              <rect x="3" y="4" width="18" height="18" rx="2" ry="2"/>
              <line x1="16" y1="2" x2="16" y2="6"/>
              <line x1="8" y1="2" x2="8" y2="6"/>
              <line x1="3" y1="10" x2="21" y2="10"/>
            </svg>
            <p>No hay horarios creados todavía</p>
            <button class="btn btn-primary" style="margin-top: 16px;">
              Crear Primer Horario
            </button>
          </div>
          
          <div v-else class="table-wrapper">
            <table>
              <thead>
                <tr>
                  <th>ID</th>
                  <th>Escuela</th>
                  <th>Período</th>
                  <th>Estado</th>
                  <th>Versión</th>
                  <th>Última Actualización</th>
                  <th>Acciones</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="item in horarios" :key="item.id_horario">
                  <td>{{ item.id_horario }}</td>
                  <td>{{ item.id_escuela }}</td>
                  <td>{{ item.id_periodo }}</td>
                  <td>
                    <span class="badge" :class="getEstadoClass(item.estado)">
                      {{ item.estado }}
                    </span>
                  </td>
                  <td>v{{ item.version_reajuste }}</td>
                  <td>{{ formatDate(item.fecha_actualizacion) }}</td>
                  <td>
                    <div style="display: flex; gap: 8px;">
                      <button class="btn btn-secondary btn-sm">Ver</button>
                      <button class="btn btn-secondary btn-sm">Editar</button>
                    </div>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
        
        <div class="card" style="margin-top: 24px;">
          <div class="card-header">
            <h2 class="card-title">Bloques Horario</h2>
          </div>
          
          <div v-if="bloques.length === 0" class="empty-state">
            <p>No hay bloques de horario registrados</p>
          </div>
          
          <div v-else class="table-wrapper">
            <table>
              <thead>
                <tr>
                  <th>ID</th>
                  <th>Horario</th>
                  <th>Grupo</th>
                  <th>Aula</th>
                  <th>Docente</th>
                  <th>Día</th>
                  <th>Hora Inicio</th>
                  <th>Hora Fin</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="item in bloques" :key="item.id_bloque">
                  <td>{{ item.id_bloque }}</td>
                  <td>{{ item.id_horario }}</td>
                  <td>{{ item.id_grupo }}</td>
                  <td>{{ item.id_aula }}</td>
                  <td>{{ item.id_docente || '—' }}</td>
                  <td>{{ getDiaNombre(item.dia_semana) }}</td>
                  <td>{{ getHora(item.slot_inicio) }}</td>
                  <td>{{ getHora(item.slot_fin) }}</td>
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

const loading = ref(true)
const error = ref('')
const horarios = ref([])
const bloques = ref([])

async function loadHorarios() {
  loading.value = true
  error.value = ''
  try {
    const [h, b] = await Promise.all([
      api.horarios.list(),
      api.bloques.list()
    ])
    horarios.value = h
    bloques.value = b
  } catch (e) {
    error.value = e.message
  } finally {
    loading.value = false
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

function getDiaNombre(dia) {
  const dias = ['', 'Lunes', 'Martes', 'Miércoles', 'Jueves', 'Viernes', 'Sábado']
  return dias[dia] || `Día ${dia}`
}

function getHora(slot) {
  if (!slot) return '—'
  const hora = 6 + slot
  return `${hora}:00`
}

function formatDate(dateStr) {
  if (!dateStr) return '—'
  return new Date(dateStr).toLocaleString('es-PE', {
    day: '2-digit',
    month: 'short',
    year: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  })
}

onMounted(loadHorarios)
</script>
