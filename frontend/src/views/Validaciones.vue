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
        <p class="page-subtitle">Prueba las reglas de negocio del sistema</p>
      </div>
      <div class="header-actions">
        <span class="rules-badge">RV-01 a RV-09</span>
      </div>
    </header>

    <div class="page-content">
      <div class="stats-grid">
        <div class="stat-card stat-gradient-purple">
          <div class="stat-icon-wrapper">
            <div class="stat-icon white">
              <svg width="28" height="28" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M22 11.08V12a10 10 0 1 1-5.93-9.14"/>
                <polyline points="22 4 12 14.01 9 11.01"/>
              </svg>
            </div>
          </div>
          <div class="stat-content">
            <span class="stat-value">9</span>
            <span class="stat-label">Reglas Activas</span>
          </div>
        </div>

        <div class="stat-card stat-gradient-red">
          <div class="stat-icon-wrapper">
            <div class="stat-icon white">
              <svg width="28" height="28" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <circle cx="12" cy="12" r="10"/>
                <line x1="15" y1="9" x2="9" y2="15"/>
                <line x1="9" y1="9" x2="15" y2="15"/>
              </svg>
            </div>
          </div>
          <div class="stat-content">
            <span class="stat-value">5</span>
            <span class="stat-label">Bloqueantes</span>
          </div>
        </div>

        <div class="stat-card stat-gradient-amber">
          <div class="stat-icon-wrapper">
            <div class="stat-icon white">
              <svg width="28" height="28" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M10.29 3.86L1.82 18a2 2 0 0 0 1.71 3h16.94a2 2 0 0 0 1.71-3L13.71 3.86a2 2 0 0 0-3.42 0z"/>
                <line x1="12" y1="9" x2="12" y2="13"/>
                <line x1="12" y1="17" x2="12.01" y2="17"/>
              </svg>
            </div>
          </div>
          <div class="stat-content">
            <span class="stat-value">3</span>
            <span class="stat-label">Advertencias</span>
          </div>
        </div>

        <div class="stat-card stat-gradient-blue">
          <div class="stat-icon-wrapper">
            <div class="stat-icon white">
              <svg width="28" height="28" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <circle cx="12" cy="12" r="10"/>
                <line x1="12" y1="16" x2="12" y2="12"/>
                <line x1="12" y1="8" x2="12.01" y2="8"/>
              </svg>
            </div>
          </div>
          <div class="stat-content">
            <span class="stat-value">1</span>
            <span class="stat-label">Informativo</span>
          </div>
        </div>
      </div>

      <div class="main-grid">
        <div class="card test-card">
          <div class="card-header">
            <h2 class="card-title">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <polygon points="13 2 3 14 12 14 11 22 21 10 12 10 13 2"/>
              </svg>
              Probar Validaciones
            </h2>
          </div>
          <div class="card-body">
            <p class="test-description">
              Selecciona un escenario de prueba para verificar el comportamiento del motor de validaciones.
            </p>

            <div class="test-actions">
              <button class="btn btn-primary btn-lg" @click="ejecutarSimple">
                <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <polygon points="5 3 19 12 5 21 5 3"/>
                </svg>
                Escenario Simple
              </button>
              <button class="btn btn-secondary btn-lg" @click="ejecutarComplejo">
                <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <path d="M13 2L3 14h9l-1 8 10-12h-9l1-8z"/>
                </svg>
                Escenario Complejo
              </button>
            </div>

            <div v-if="loading" class="loading-state">
              <div class="spinner-lg"></div>
              <p>Ejecutando validación...</p>
            </div>

            <template v-else-if="resultado">
              <div class="result-header">
                <h3>Resultado de la Validación</h3>
                <span class="result-badge" :class="resultado.length === 0 ? 'success' : 'has-issues'">
                  {{ resultado.length === 0 ? 'Sin problemas' : resultado.length + ' hallazgo(s)' }}
                </span>
              </div>

              <div v-if="resultado.length === 0" class="success-state">
                <svg width="48" height="48" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <path d="M22 11.08V12a10 10 0 1 1-5.93-9.14"/>
                  <polyline points="22 4 12 14.01 9 11.01"/>
                </svg>
                <p>Sin observaciones - La asignación puede continuar</p>
              </div>

              <div v-else class="result-table">
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
                      <td><span class="rule-code">{{ finding.rule }}</span></td>
                      <td>
                        <span class="severity-badge" :class="getSeverityClass(finding.severity)">
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
                <svg width="64" height="64" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1">
                  <path d="M9 12l2 2 4-4"/>
                  <path d="M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>
                </svg>
                <p>Selecciona un escenario para ejecutar la validación</p>
              </div>
            </template>
          </div>
        </div>

        <div class="card rules-reference">
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
                <div class="rule-header">
                  <span class="rule-code">RV-01</span>
                  <span class="severity-pill blocker">Bloqueante</span>
                </div>
                <p>Conflicto de docente (mismo horario y docente)</p>
              </div>
              <div class="rule-item blocker">
                <div class="rule-header">
                  <span class="rule-code">RV-02</span>
                  <span class="severity-pill blocker">Bloqueante</span>
                </div>
                <p>Conflicto de aula (aula ocupada)</p>
              </div>
              <div class="rule-item blocker">
                <div class="rule-header">
                  <span class="rule-code">RV-03</span>
                  <span class="severity-pill blocker">Bloqueante</span>
                </div>
                <p>Conflicto con sesión de departamento</p>
              </div>
              <div class="rule-item blocker">
                <div class="rule-header">
                  <span class="rule-code">RV-04a</span>
                  <span class="severity-pill blocker">Bloqueante</span>
                </div>
                <p>Tiempo insuficiente para traslado entre pabellones</p>
              </div>
              <div class="rule-item warning">
                <div class="rule-header">
                  <span class="rule-code">RV-04b</span>
                  <span class="severity-pill warning">Advertencia</span>
                </div>
                <p>Intervalo libre menor al tiempo de desplazamiento</p>
              </div>
              <div class="rule-item warning">
                <div class="rule-header">
                  <span class="rule-code">RV-05</span>
                  <span class="severity-pill warning">Advertencia</span>
                </div>
                <p>Carga lectiva semanal supera 16 horas</p>
              </div>
              <div class="rule-item info">
                <div class="rule-header">
                  <span class="rule-code">RV-06</span>
                  <span class="severity-pill info">Informativo</span>
                </div>
                <p>Dos cursos de la misma serie en el mismo horario</p>
              </div>
              <div class="rule-item blocker">
                <div class="rule-header">
                  <span class="rule-code">RV-07</span>
                  <span class="severity-pill blocker">Bloqueante</span>
                </div>
                <p>Aula compartida ya reservada por otra escuela</p>
              </div>
              <div class="rule-item blocker">
                <div class="rule-header">
                  <span class="rule-code">RV-08</span>
                  <span class="severity-pill blocker">Bloqueante</span>
                </div>
                <p>Cambios en horario oficial requieren justificación</p>
              </div>
              <div class="rule-item warning">
                <div class="rule-header">
                  <span class="rule-code">RV-09</span>
                  <span class="severity-pill warning">Advertencia</span>
                </div>
                <p>Matrícula excede capacidad del aula</p>
              </div>
            </div>
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
    'BLOCKER': 'severity-blocker',
    'WARNING': 'severity-warning',
    'INFO': 'severity-info'
  }
  return map[severity] || 'severity-info'
}
</script>

<style scoped>
.validaciones-page {
  min-height: 100vh;
  background: linear-gradient(135deg, #f0f4ff 0%, #fdf2f8 50%, #f0fdf4 100%);
}

.page-header {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  padding: 24px 32px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  box-shadow: 0 8px 32px rgba(102, 126, 234, 0.3);
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

.rules-badge {
  background: rgba(255,255,255,0.2);
  color: white;
  padding: 8px 16px;
  border-radius: 20px;
  font-size: 0.85rem;
  font-weight: 600;
  border: 2px solid rgba(255,255,255,0.3);
}

.page-content {
  padding: 28px;
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 20px;
  margin-bottom: 28px;
}

.stat-card {
  background: white;
  border-radius: 20px;
  padding: 24px;
  display: flex;
  align-items: center;
  gap: 16px;
  box-shadow: 0 8px 32px rgba(0,0,0,0.08);
  transition: all 0.3s;
}

.stat-card:hover {
  transform: translateY(-4px);
}

.stat-gradient-purple { background: linear-gradient(135deg, #667eea, #764ba2); color: white; }
.stat-gradient-red { background: linear-gradient(135deg, #f56565, #e53e3e); color: white; }
.stat-gradient-amber { background: linear-gradient(135deg, #f6ad55, #dd6b20); color: white; }
.stat-gradient-blue { background: linear-gradient(135deg, #4299e1, #3182ce); color: white; }

.stat-icon {
  width: 52px;
  height: 52px;
  border-radius: 14px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: rgba(255,255,255,0.2);
}

.stat-content .stat-value {
  display: block;
  font-size: 2rem;
  font-weight: 700;
  line-height: 1;
}

.stat-content .stat-label {
  font-size: 0.85rem;
  opacity: 0.9;
}

.main-grid {
  display: grid;
  grid-template-columns: 2fr 1fr;
  gap: 24px;
}

.card {
  background: white;
  border-radius: 20px;
  box-shadow: 0 8px 32px rgba(0,0,0,0.08);
  overflow: hidden;
}

.card-header {
  padding: 20px 24px;
  border-bottom: 1px solid #f1f5f9;
  background: linear-gradient(135deg, #f8fafc 0%, #f0f4ff 100%);
}

.card-title {
  display: flex;
  align-items: center;
  gap: 10px;
  margin: 0;
  font-size: 1.1rem;
  color: #1e293b;
  font-weight: 700;
}

.card-body {
  padding: 24px;
}

.test-description {
  color: #64748b;
  margin: 0 0 24px;
}

.test-actions {
  display: flex;
  gap: 12px;
  margin-bottom: 24px;
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

.btn-lg {
  padding: 14px 24px;
  font-size: 1rem;
}

.btn-primary {
  background: linear-gradient(135deg, #667eea, #764ba2);
  color: white;
}

.btn-primary:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(102, 126, 234, 0.4);
}

.btn-secondary {
  background: linear-gradient(135deg, #11998e, #38ef7d);
  color: white;
}

.btn-secondary:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(17, 153, 142, 0.4);
}

.loading-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 60px 20px;
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

.result-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
}

.result-header h3 {
  margin: 0;
  font-size: 1rem;
  color: #1e293b;
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

.success-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 40px 20px;
  background: linear-gradient(135deg, #f0fdf4, #dcfce7);
  border-radius: 16px;
  color: #166534;
  text-align: center;
}

.success-state p {
  margin: 12px 0 0;
  font-weight: 500;
}

.result-table {
  background: #f8fafc;
  border-radius: 16px;
  overflow: hidden;
}

.result-table table {
  width: 100%;
  border-collapse: collapse;
}

.result-table th {
  text-align: left;
  padding: 14px 16px;
  font-size: 0.75rem;
  font-weight: 700;
  text-transform: uppercase;
  color: #64748b;
  background: white;
  border-bottom: 2px solid #e2e8f0;
}

.result-table td {
  padding: 14px 16px;
  border-bottom: 1px solid #f1f5f9;
  font-size: 0.9rem;
}

.result-table tr:last-child td {
  border-bottom: none;
}

.rule-code {
  font-weight: 700;
  font-family: monospace;
  color: #667eea;
}

.severity-badge {
  padding: 4px 10px;
  border-radius: 6px;
  font-size: 0.75rem;
  font-weight: 600;
}

.severity-badge.severity-blocker {
  background: #fef2f2;
  color: #dc2626;
}

.severity-badge.severity-warning {
  background: #fffbeb;
  color: #d97706;
}

.severity-badge.severity-info {
  background: #eff6ff;
  color: #2563eb;
}

.empty-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 60px 20px;
  color: #94a3b8;
  text-align: center;
}

.empty-state svg {
  opacity: 0.5;
  margin-bottom: 16px;
}

.rules-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.rule-item {
  padding: 14px 16px;
  border-radius: 12px;
  border-left: 4px solid;
}

.rule-item.blocker {
  background: #fef2f2;
  border-color: #ef4444;
}

.rule-item.warning {
  background: #fffbeb;
  border-color: #f59e0b;
}

.rule-item.info {
  background: #eff6ff;
  border-color: #3b82f6;
}

.rule-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 6px;
}

.rule-item .rule-code {
  font-weight: 700;
  font-family: monospace;
  font-size: 0.85rem;
}

.rule-item p {
  margin: 0;
  font-size: 0.8rem;
  color: #64748b;
}

.severity-pill {
  padding: 2px 8px;
  border-radius: 10px;
  font-size: 0.65rem;
  font-weight: 600;
  text-transform: uppercase;
}

.severity-pill.blocker {
  background: #fee2e2;
  color: #dc2626;
}

.severity-pill.warning {
  background: #fef3c7;
  color: #d97706;
}

.severity-pill.info {
  background: #dbeafe;
  color: #2563eb;
}

@media (max-width: 1024px) {
  .stats-grid {
    grid-template-columns: repeat(2, 1fr);
  }
  .main-grid {
    grid-template-columns: 1fr;
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
  .stats-grid {
    grid-template-columns: 1fr;
  }
  .test-actions {
    flex-direction: column;
  }
}
</style>
