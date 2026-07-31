<template>
  <div>
    <header class="topbar">
      <h1 class="page-title">Motor de Validaciones</h1>
      <div class="page-actions">
        <span class="badge badge-primary">RV-01 a RV-09</span>
      </div>
    </header>
    
    <div class="page-content">
      <div class="stats-grid" style="margin-bottom: 32px;">
        <div class="stat-card">
          <div class="stat-icon blue">
            <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M22 11.08V12a10 10 0 1 1-5.93-9.14"/>
              <polyline points="22 4 12 14.01 9 11.01"/>
            </svg>
          </div>
          <div class="stat-value">9</div>
          <div class="stat-label">Reglas Activas</div>
        </div>
        
        <div class="stat-card">
          <div class="stat-icon green">
            <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <circle cx="12" cy="12" r="10"/>
              <line x1="12" y1="8" x2="12" y2="12"/>
              <line x1="12" y1="16" x2="12.01" y2="16"/>
            </svg>
          </div>
          <div class="stat-value">5</div>
          <div class="stat-label">Bloqueantes</div>
        </div>
        
        <div class="stat-card">
          <div class="stat-icon amber">
            <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M10.29 3.86L1.82 18a2 2 0 0 0 1.71 3h16.94a2 2 0 0 0 1.71-3L13.71 3.86a2 2 0 0 0-3.42 0z"/>
              <line x1="12" y1="9" x2="12" y2="13"/>
              <line x1="12" y1="17" x2="12.01" y2="17"/>
            </svg>
          </div>
          <div class="stat-value">3</div>
          <div class="stat-label">Advertencias</div>
        </div>
        
        <div class="stat-card">
          <div class="stat-icon purple">
            <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <circle cx="12" cy="12" r="10"/>
              <line x1="12" y1="16" x2="12" y2="12"/>
              <line x1="12" y1="8" x2="12.01" y2="8"/>
            </svg>
          </div>
          <div class="stat-value">1</div>
          <div class="stat-label">Informativo</div>
        </div>
      </div>
      
      <div class="card">
        <div class="card-header">
          <h2 class="card-title">Probar Validaciones</h2>
        </div>
        <div class="card-body">
          <p style="color: var(--gray-500); margin-bottom: 24px;">
            Selecciona un escenario de prueba para verificar el comportamiento del motor de validaciones.
          </p>
          
          <div style="display: flex; gap: 12px; margin-bottom: 24px;">
            <button class="btn btn-primary" @click="ejecutarSimple">
              <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <polygon points="5 3 19 12 5 21 5 3"/>
              </svg>
              Escenario Simple
            </button>
            <button class="btn btn-secondary" @click="ejecutarComplejo">
              Escenario Complejo
            </button>
          </div>
          
          <div v-if="loading" class="loading">
            <div class="spinner"></div>
          </div>
          
          <template v-else-if="resultado">
            <h3 style="font-size: 1rem; margin-bottom: 12px;">Resultado de la Validación</h3>
            
            <div v-if="resultado.length === 0" class="badge badge-success" style="margin-bottom: 16px; font-size: 0.875rem; padding: 8px 16px;">
              Sin observaciones - La asignación puede continuar
            </div>
            
            <div v-else class="table-wrapper">
              <table>
                <thead>
                  <tr>
                    <th>Regla</th>
                    <th>Severidad</th>
                    <th>Mensaje</th>
                  </tr>
                </thead>
                <tbody>
                  <tr v-for="(finding, idx) in resultado" :key="idx">
                    <td><span class="badge badge-gray">{{ finding.rule }}</span></td>
                    <td>
                      <span class="badge" :class="getSeverityClass(finding.severity)">
                        {{ finding.severity }}
                      </span>
                    </td>
                    <td>{{ finding.message }}</td>
                  </tr>
                </tbody>
              </table>
            </div>
          </template>
          
          <template v-else>
            <div class="empty-state">
              <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5">
                <path d="M9 12l2 2 4-4"/>
                <path d="M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>
              </svg>
              <p>Selecciona un escenario para ejecutar la validación</p>
            </div>
          </template>
        </div>
      </div>
      
      <div class="card" style="margin-top: 24px;">
        <div class="card-header">
          <h2 class="card-title">Reglas Disponibles</h2>
        </div>
        <div class="card-body">
          <div class="table-wrapper">
            <table>
              <thead>
                <tr>
                  <th>Código</th>
                  <th>Descripción</th>
                  <th>Severidad</th>
                </tr>
              </thead>
              <tbody>
                <tr><td>RV-01</td><td>Conflicto de docente (mismo horario y docente)</td><td><span class="badge badge-danger">Bloqueante</span></td></tr>
                <tr><td>RV-02</td><td>Conflicto de aula (aula ocupada)</td><td><span class="badge badge-danger">Bloqueante</span></td></tr>
                <tr><td>RV-03</td><td>Conflicto con sesión de departamento</td><td><span class="badge badge-danger">Bloqueante</span></td></tr>
                <tr><td>RV-04a</td><td>Tiempo insuficiente para traslado entre pabellones</td><td><span class="badge badge-danger">Bloqueante</span></td></tr>
                <tr><td>RV-04b</td><td>Intervalo libre menor al tiempo de desplazamiento</td><td><span class="badge badge-warning">Advertencia</span></td></tr>
                <tr><td>RV-05</td><td>Carga lectiva semanal supera 16 horas</td><td><span class="badge badge-warning">Advertencia</span></td></tr>
                <tr><td>RV-06</td><td>Dos cursos de la misma serie en el mismo horario</td><td><span class="badge badge-gray">Informativo</span></td></tr>
                <tr><td>RV-07</td><td>Aula compartida ya reservada por otra escuela</td><td><span class="badge badge-danger">Bloqueante</span></td></tr>
                <tr><td>RV-08</td><td>Cambios en horario oficial requieren justificación</td><td><span class="badge badge-danger">Bloqueante</span></td></tr>
                <tr><td>RV-09</td><td>Matrícula excede capacidad del aula</td><td><span class="badge badge-warning">Advertencia</span></td></tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import api from '../services/api'

const loading = ref(false)
const resultado = ref(null)

async function ejecutarSimple() {
  loading.value = true
  resultado.value = null
  
  try {
    const payload = {
      proposed: {
        id: 101,
        teacher_id: 10,
        school_id: 1,
        group_id: 1,
        room_id: 5,
        series_id: 1,
        day: 1,
        start_slot: 3,
        end_slot: 4,
        enrollment: 42,
        room_capacity: 40
      },
      existing: [
        {
          id: 100,
          teacher_id: 10,
          school_id: 1,
          group_id: 2,
          room_id: 6,
          series_id: 2,
          day: 1,
          start_slot: 4,
          end_slot: 5,
          enrollment: 20,
          room_capacity: 40
        }
      ],
      state: 'BORRADOR'
    }
    
    resultado.value = await api.validaciones.placement(payload)
  } catch (e) {
    alert('Error: ' + e.message)
  } finally {
    loading.value = false
  }
}

async function ejecutarComplejo() {
  loading.value = true
  resultado.value = null
  
  try {
    const payload = {
      proposed: {
        id: 201,
        teacher_id: 10,
        school_id: 1,
        group_id: 1,
        course_id: 101,
        series_id: 1,
        room_id: 5,
        pavilion_id: 2,
        day: 1,
        start_slot: 3,
        end_slot: 4,
        enrollment: 55,
        room_capacity: 40
      },
      existing: [
        {
          id: 200,
          teacher_id: 10,
          school_id: 1,
          group_id: 2,
          course_id: 102,
          series_id: 2,
          room_id: 6,
          pavilion_id: 1,
          day: 1,
          start_slot: 4,
          end_slot: 5,
          enrollment: 20,
          room_capacity: 40
        }
      ],
      distances: [
        {from_pavilion_id: 1, to_pavilion_id: 2, minutes: 90}
      ],
      department_sessions: [
        {department_id: 7, day: 1, start_slot: 4, end_slot: 5}
      ],
      state: 'EN_REAJUSTE'
    }
    
    resultado.value = await api.validaciones.placement(payload)
  } catch (e) {
    alert('Error: ' + e.message)
  } finally {
    loading.value = false
  }
}

function getSeverityClass(severity) {
  const map = {
    'BLOCKER': 'badge-danger',
    'WARNING': 'badge-warning',
    'INFO': 'badge-gray'
  }
  return map[severity] || 'badge-gray'
}
</script>
