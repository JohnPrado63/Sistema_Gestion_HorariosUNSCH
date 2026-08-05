<template>
  <div class="dashboard-page">
    <header class="page-header">
      <div class="header-left">
        <h1 class="page-title">
          <svg width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M3 9l9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"/>
            <polyline points="9 22 9 12 15 12 15 22"/>
          </svg>
          Panel de Control
        </h1>
        <p class="page-subtitle">Sistema de Gestión de Horarios UNSCH</p>
      </div>
      <div class="header-right">
        <div class="status-badge status-active">
          <span class="status-dot"></span>
          Sistema Activo
        </div>
      </div>
    </header>

    <div class="page-content">
      <div class="stats-grid">
        <div class="stat-card stat-card-purple">
          <div class="stat-icon-wrapper">
            <div class="stat-icon">
              <svg width="28" height="28" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M4 19.5A2.5 2.5 0 0 1 6.5 17H20"/>
                <path d="M6.5 2H20v20H6.5A2.5 2.5 0 0 1 4 19.5v-15A2.5 2.5 0 0 1 6.5 2z"/>
              </svg>
            </div>
          </div>
          <div class="stat-content">
            <span class="stat-value">{{ stats.facultades }}</span>
            <span class="stat-label">Facultades</span>
          </div>
          <div class="stat-decoration"></div>
        </div>

        <div class="stat-card stat-card-teal">
          <div class="stat-icon-wrapper">
            <div class="stat-icon">
              <svg width="28" height="28" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"/>
                <circle cx="9" cy="7" r="4"/>
                <path d="M23 21v-2a4 4 0 0 0-3-3.87"/>
                <path d="M16 3.13a4 4 0 0 1 0 7.75"/>
              </svg>
            </div>
          </div>
          <div class="stat-content">
            <span class="stat-value">{{ stats.docentes }}</span>
            <span class="stat-label">Docentes</span>
          </div>
          <div class="stat-decoration"></div>
        </div>

        <div class="stat-card stat-card-amber">
          <div class="stat-icon-wrapper">
            <div class="stat-icon">
              <svg width="28" height="28" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M2 3h6a4 4 0 0 1 4 4v14a3 3 0 0 0-3-3H2z"/>
                <path d="M22 3h-6a4 4 0 0 0-4 4v14a3 3 0 0 1 3-3h7z"/>
              </svg>
            </div>
          </div>
          <div class="stat-content">
            <span class="stat-value">{{ stats.cursos }}</span>
            <span class="stat-label">Cursos</span>
          </div>
          <div class="stat-decoration"></div>
        </div>

        <div class="stat-card stat-card-rose">
          <div class="stat-icon-wrapper">
            <div class="stat-icon">
              <svg width="28" height="28" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <rect x="3" y="4" width="18" height="18" rx="2" ry="2"/>
                <line x1="16" y1="2" x2="16" y2="6"/>
                <line x1="8" y1="2" x2="8" y2="6"/>
                <line x1="3" y1="10" x2="21" y2="10"/>
              </svg>
            </div>
          </div>
          <div class="stat-content">
            <span class="stat-value">{{ stats.horarios }}</span>
            <span class="stat-label">Horarios Creados</span>
          </div>
          <div class="stat-decoration"></div>
        </div>
      </div>

      <div class="dashboard-grid">
        <div class="card info-card">
          <div class="card-header">
            <h2 class="card-title">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <circle cx="12" cy="12" r="10"/>
                <line x1="12" y1="16" x2="12" y2="12"/>
                <line x1="12" y1="8" x2="12.01" y2="8"/>
              </svg>
              Estado del Sistema
            </h2>
          </div>
          <div class="card-body">
            <div class="info-grid">
              <div class="info-item">
                <span class="info-label">Versión API</span>
                <span class="info-value version-badge">1.0.0</span>
              </div>
              <div class="info-item">
                <span class="info-label">Reglas de Validación</span>
                <span class="info-value">RV-01 a RV-09</span>
              </div>
              <div class="info-item">
                <span class="info-label">Período Activo</span>
                <span class="info-value periodo-badge">{{ stats.periodoActivo || 'No configurado' }}</span>
              </div>
              <div class="info-item">
                <span class="info-label">Última Actualización</span>
                <span class="info-value">{{ new Date().toLocaleDateString('es-PE', { weekday: 'long', year: 'numeric', month: 'long', day: 'numeric' }) }}</span>
              </div>
            </div>
          </div>
        </div>

        <div class="card actions-card">
          <div class="card-header">
            <h2 class="card-title">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <polygon points="13 2 3 14 12 14 11 22 21 10 12 10 13 2"/>
              </svg>
              Accesos Rápidos
            </h2>
          </div>
          <div class="card-body">
            <div class="quick-actions">
              <router-link to="/catalogos" class="action-btn action-btn-purple">
                <div class="action-icon">
                  <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <path d="M4 19.5A2.5 2.5 0 0 1 6.5 17H20"/>
                    <path d="M6.5 2H20v20H6.5A2.5 2.5 0 0 1 4 19.5v-15A2.5 2.5 0 0 1 6.5 2z"/>
                  </svg>
                </div>
                <span class="action-label">Catálogos</span>
                <span class="action-desc">Gestionar facultades, departamentos, escuelas</span>
              </router-link>

              <router-link to="/horarios" class="action-btn action-btn-teal">
                <div class="action-icon">
                  <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <rect x="3" y="4" width="18" height="18" rx="2" ry="2"/>
                    <line x1="16" y1="2" x2="16" y2="6"/>
                    <line x1="8" y1="2" x2="8" y2="6"/>
                    <line x1="3" y1="10" x2="21" y2="10"/>
                  </svg>
                </div>
                <span class="action-label">Horarios</span>
                <span class="action-desc">Crear y gestionar horarios académicos</span>
              </router-link>

              <router-link to="/carga-academica" class="action-btn action-btn-amber">
                <div class="action-icon">
                  <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/>
                    <polyline points="14 2 14 8 20 8"/>
                    <line x1="12" y1="18" x2="12" y2="12"/>
                    <line x1="9" y1="15" x2="15" y2="15"/>
                  </svg>
                </div>
                <span class="action-label">Carga Académica</span>
                <span class="action-desc">Asignar cursos y grupos a docentes</span>
              </router-link>

              <router-link to="/validaciones" class="action-btn action-btn-rose">
                <div class="action-icon">
                  <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <polyline points="9 11 12 14 22 4"/>
                    <path d="M21 12v7a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h11"/>
                  </svg>
                </div>
                <span class="action-label">Validaciones</span>
                <span class="action-desc">Probar reglas de negocio</span>
              </router-link>
            </div>
          </div>
        </div>

        <div class="card rules-card">
          <div class="card-header">
            <h2 class="card-title">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z"/>
              </svg>
              Reglas de Validación
            </h2>
          </div>
          <div class="card-body">
            <div class="rules-list">
              <div class="rule-item rule-blocker">
                <span class="rule-indicator"></span>
                <div class="rule-content">
                  <span class="rule-code">RV-01</span>
                  <span class="rule-desc">Conflicto de horario docente</span>
                </div>
                <span class="rule-severity severity-blocker">Bloqueante</span>
              </div>
              <div class="rule-item rule-blocker">
                <span class="rule-indicator"></span>
                <div class="rule-content">
                  <span class="rule-code">RV-02</span>
                  <span class="rule-desc">Conflicto de aula ocupada</span>
                </div>
                <span class="rule-severity severity-blocker">Bloqueante</span>
              </div>
              <div class="rule-item rule-blocker">
                <span class="rule-indicator"></span>
                <div class="rule-content">
                  <span class="rule-code">RV-03</span>
                  <span class="rule-desc">Superposición de sesiones del depto</span>
                </div>
                <span class="rule-severity severity-blocker">Bloqueante</span>
              </div>
              <div class="rule-item rule-warning">
                <span class="rule-indicator"></span>
                <div class="rule-content">
                  <span class="rule-code">RV-04b</span>
                  <span class="rule-desc">Transferencia sin tiempo suficiente</span>
                </div>
                <span class="rule-severity severity-warning">Advertencia</span>
              </div>
              <div class="rule-item rule-info">
                <span class="rule-indicator"></span>
                <div class="rule-content">
                  <span class="rule-code">RV-06</span>
                  <span class="rule-desc">Dos cursos de misma serie cruzan</span>
                </div>
                <span class="rule-severity severity-info">Info</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import api from '../services/api'

const stats = ref({
  facultades: '-',
  docentes: '-',
  cursos: '-',
  horarios: '-',
  periodoActivo: ''
})

async function loadStats() {
  try {
    const [facultades, docentes, cursos, horarios, periodos] = await Promise.all([
      api.facultades.list(),
      api.docentes.list(),
      api.cursos.list(),
      api.horarios.list(),
      api.periodos.list()
    ])

    stats.value = {
      facultades: facultades.length,
      docentes: docentes.length,
      cursos: cursos.length,
      horarios: horarios.length,
      periodoActivo: periodos.find(p => p.activo)?.codigo || 'No hay período activo'
    }
  } catch (error) {
    console.error('Error loading stats:', error)
  }
}

onMounted(loadStats)
</script>

<style scoped>
.dashboard-page {
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

.status-badge {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 10px 18px;
  border-radius: 30px;
  font-size: 0.9rem;
  font-weight: 600;
  background: rgba(255,255,255,0.15);
  border: 2px solid rgba(255,255,255,0.3);
}

.status-dot {
  width: 10px;
  height: 10px;
  border-radius: 50%;
  animation: pulse 2s infinite;
}

.status-active .status-dot {
  background: #4ade80;
  box-shadow: 0 0 8px rgba(74, 222, 128, 0.6);
}

@keyframes pulse {
  0%, 100% { opacity: 1; transform: scale(1); }
  50% { opacity: 0.7; transform: scale(1.2); }
}

.page-content {
  padding: 28px;
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 24px;
  margin-bottom: 28px;
}

.stat-card {
  background: white;
  border-radius: 20px;
  padding: 24px;
  display: flex;
  align-items: center;
  gap: 20px;
  position: relative;
  overflow: hidden;
  box-shadow: 0 8px 32px rgba(0,0,0,0.08);
  border: 1px solid rgba(0,0,0,0.05);
  transition: all 0.3s;
}

.stat-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 12px 40px rgba(0,0,0,0.12);
}

.stat-icon-wrapper {
  position: relative;
  z-index: 1;
}

.stat-icon {
  width: 56px;
  height: 56px;
  border-radius: 16px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
}

.stat-card-purple .stat-icon { background: linear-gradient(135deg, #667eea, #764ba2); }
.stat-card-teal .stat-icon { background: linear-gradient(135deg, #11998e, #38ef7d); }
.stat-card-amber .stat-icon { background: linear-gradient(135deg, #f093fb, #f5576c); }
.stat-card-rose .stat-icon { background: linear-gradient(135deg, #fc466b, #3f5efb); }

.stat-content {
  position: relative;
  z-index: 1;
}

.stat-value {
  display: block;
  font-size: 2rem;
  font-weight: 700;
  line-height: 1;
  margin-bottom: 4px;
}

.stat-card-purple .stat-value { color: #667eea; }
.stat-card-teal .stat-value { color: #11998e; }
.stat-card-amber .stat-value { color: #f5576c; }
.stat-card-rose .stat-value { color: #fc466b; }

.stat-label {
  font-size: 0.85rem;
  color: #64748b;
  font-weight: 500;
}

.stat-decoration {
  position: absolute;
  top: -20px;
  right: -20px;
  width: 100px;
  height: 100px;
  border-radius: 50%;
  opacity: 0.1;
}

.stat-card-purple .stat-decoration { background: #667eea; }
.stat-card-teal .stat-decoration { background: #11998e; }
.stat-card-amber .stat-decoration { background: #f5576c; }
.stat-card-rose .stat-decoration { background: #fc466b; }

.dashboard-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 24px;
}

.card {
  background: white;
  border-radius: 20px;
  box-shadow: 0 8px 32px rgba(0,0,0,0.08);
  border: 1px solid rgba(0,0,0,0.05);
  overflow: hidden;
}

.info-card {
  grid-column: span 1;
}

.actions-card {
  grid-column: span 2;
}

.rules-card {
  grid-column: span 3;
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

.info-grid {
  display: grid;
  gap: 16px;
}

.info-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 16px;
  background: linear-gradient(135deg, #f8fafc 0%, #f0f4ff 100%);
  border-radius: 12px;
}

.info-label {
  font-size: 0.85rem;
  color: #64748b;
}

.info-value {
  font-size: 0.9rem;
  font-weight: 600;
  color: #1e293b;
}

.version-badge {
  background: linear-gradient(135deg, #667eea20, #764ba220);
  color: #667eea;
  padding: 4px 12px;
  border-radius: 20px;
}

.periodo-badge {
  background: linear-gradient(135deg, #11998e20, #38ef7d20);
  color: #11998e;
  padding: 4px 12px;
  border-radius: 20px;
}

.quick-actions {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 16px;
}

.action-btn {
  display: flex;
  flex-direction: column;
  align-items: center;
  text-align: center;
  padding: 24px;
  border-radius: 16px;
  background: linear-gradient(135deg, #f8fafc 0%, #f0f4ff 100%);
  border: 2px solid transparent;
  transition: all 0.3s;
  text-decoration: none;
}

.action-btn:hover {
  transform: translateY(-4px);
  box-shadow: 0 8px 24px rgba(0,0,0,0.12);
}

.action-btn-purple:hover { border-color: #667eea; background: linear-gradient(135deg, #667eea10, #764ba210); }
.action-btn-teal:hover { border-color: #11998e; background: linear-gradient(135deg, #11998e10, #38ef7d10); }
.action-btn-amber:hover { border-color: #f5576c; background: linear-gradient(135deg, #f093fb10, #f5576c10); }
.action-btn-rose:hover { border-color: #fc466b; background: linear-gradient(135deg, #fc466b10, #3f5efb10); }

.action-icon {
  width: 52px;
  height: 52px;
  border-radius: 14px;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 12px;
  color: white;
}

.action-btn-purple .action-icon { background: linear-gradient(135deg, #667eea, #764ba2); }
.action-btn-teal .action-icon { background: linear-gradient(135deg, #11998e, #38ef7d); }
.action-btn-amber .action-icon { background: linear-gradient(135deg, #f093fb, #f5576c); }
.action-btn-rose .action-icon { background: linear-gradient(135deg, #fc466b, #3f5efb); }

.action-label {
  font-size: 1rem;
  font-weight: 700;
  color: #1e293b;
  margin-bottom: 4px;
}

.action-desc {
  font-size: 0.8rem;
  color: #64748b;
}

.rules-list {
  display: grid;
  grid-template-columns: repeat(5, 1fr);
  gap: 12px;
}

.rule-item {
  display: flex;
  flex-direction: column;
  padding: 16px;
  border-radius: 12px;
  background: linear-gradient(135deg, #f8fafc 0%, #f0f4ff 100%);
  border-left: 4px solid;
}

.rule-item.rule-blocker { border-color: #ef4444; }
.rule-item.rule-warning { border-color: #f59e0b; }
.rule-item.rule-info { border-color: #3b82f6; }

.rule-indicator {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  margin-bottom: 8px;
}

.rule-blocker .rule-indicator { background: #ef4444; }
.rule-warning .rule-indicator { background: #f59e0b; }
.rule-info .rule-indicator { background: #3b82f6; }

.rule-content {
  margin-bottom: 8px;
}

.rule-code {
  display: block;
  font-size: 0.85rem;
  font-weight: 700;
  color: #1e293b;
  margin-bottom: 2px;
}

.rule-desc {
  display: block;
  font-size: 0.75rem;
  color: #64748b;
  line-height: 1.3;
}

.rule-severity {
  font-size: 0.65rem;
  padding: 3px 8px;
  border-radius: 10px;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.severity-blocker { background: #fef2f2; color: #dc2626; }
.severity-warning { background: #fffbeb; color: #d97706; }
.severity-info { background: #eff6ff; color: #2563eb; }

@media (max-width: 1200px) {
  .stats-grid {
    grid-template-columns: repeat(2, 1fr);
  }
  .dashboard-grid {
    grid-template-columns: 1fr;
  }
  .actions-card, .rules-card {
    grid-column: span 1;
  }
  .rules-list {
    grid-template-columns: repeat(3, 1fr);
  }
}

@media (max-width: 768px) {
  .stats-grid {
    grid-template-columns: 1fr;
  }
  .quick-actions {
    grid-template-columns: 1fr;
  }
  .rules-list {
    grid-template-columns: repeat(2, 1fr);
  }
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
