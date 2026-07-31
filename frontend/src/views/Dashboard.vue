<template>
  <div>
    <header class="topbar">
      <h1 class="page-title">Dashboard</h1>
      <div class="page-actions">
        <span class="badge badge-success">Sistema Activo</span>
      </div>
    </header>
    
    <div class="page-content">
      <div class="stats-grid">
        <div class="stat-card">
          <div class="stat-icon blue">
            <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M4 19.5A2.5 2.5 0 0 1 6.5 17H20"/>
              <path d="M6.5 2H20v20H6.5A2.5 2.5 0 0 1 4 19.5v-15A2.5 2.5 0 0 1 6.5 2z"/>
            </svg>
          </div>
          <div class="stat-value">{{ stats.facultades }}</div>
          <div class="stat-label">Facultades</div>
        </div>
        
        <div class="stat-card">
          <div class="stat-icon green">
            <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"/>
              <circle cx="9" cy="7" r="4"/>
              <path d="M23 21v-2a4 4 0 0 0-3-3.87"/>
              <path d="M16 3.13a4 4 0 0 1 0 7.75"/>
            </svg>
          </div>
          <div class="stat-value">{{ stats.docentes }}</div>
          <div class="stat-label">Docentes</div>
        </div>
        
        <div class="stat-card">
          <div class="stat-icon amber">
            <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M2 3h6a4 4 0 0 1 4 4v14a3 3 0 0 0-3-3H2z"/>
              <path d="M22 3h-6a4 4 0 0 0-4 4v14a3 3 0 0 1 3-3h7z"/>
            </svg>
          </div>
          <div class="stat-value">{{ stats.cursos }}</div>
          <div class="stat-label">Cursos</div>
        </div>
        
        <div class="stat-card">
          <div class="stat-icon purple">
            <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <rect x="3" y="4" width="18" height="18" rx="2" ry="2"/>
              <line x1="16" y1="2" x2="16" y2="6"/>
              <line x1="8" y1="2" x2="8" y2="6"/>
              <line x1="3" y1="10" x2="21" y2="10"/>
            </svg>
          </div>
          <div class="stat-value">{{ stats.horarios }}</div>
          <div class="stat-label">Horarios Creados</div>
        </div>
      </div>
      
      <div class="card">
        <div class="card-header">
          <h2 class="card-title">Estado del Sistema</h2>
        </div>
        <div class="card-body">
          <div class="detail-grid">
            <div class="detail-item">
              <label>Versión API</label>
              <span>1.0.0</span>
            </div>
            <div class="detail-item">
              <label>Reglas de Validación</label>
              <span>RV-01 a RV-09</span>
            </div>
            <div class="detail-item">
              <label>Período Activo</label>
              <span>{{ stats.periodoActivo || 'No configurado' }}</span>
            </div>
            <div class="detail-item">
              <label>Última Actualización</label>
              <span>{{ new Date().toLocaleDateString('es-PE') }}</span>
            </div>
          </div>
          
          <h3 style="margin-bottom: 16px; font-size: 1rem; color: var(--gray-700);">Accesos Rápidos</h3>
          <div style="display: flex; gap: 12px; flex-wrap: wrap;">
            <router-link to="/catalogos" class="btn btn-primary">
              Ver Catálogos
            </router-link>
            <router-link to="/horarios" class="btn btn-secondary">
              Gestionar Horarios
            </router-link>
            <router-link to="/validaciones" class="btn btn-secondary">
              Motor de Validaciones
            </router-link>
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
