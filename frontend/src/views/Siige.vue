<template>
  <div>
    <header class="topbar">
      <h1 class="page-title">Importador SIIGE</h1>
      <div class="page-actions">
        <span class="badge badge-amber">Beta</span>
      </div>
    </header>
    
    <div class="page-content">
      <div class="card">
        <div class="card-header">
          <h2 class="card-title">Importar Datos desde SIIGE</h2>
        </div>
        <div class="card-body">
          <p style="color: var(--gray-500); margin-bottom: 24px;">
            Pega el JSON exportado desde el sistema SIIGE para importar los datos de manera masiva.
            Esta herramienta permite validar el formato antes de realizar la importación.
          </p>
          
          <div class="form-group">
            <label class="form-label">Datos SIIGE (JSON)</label>
            <textarea 
              v-model="jsonInput" 
              class="form-input" 
              style="min-height: 200px; font-family: monospace;"
              placeholder='{
  "periodo": "2026-I",
  "facultades": [...],
  "docentes": [...],
  "cursos": [...]
}'
            ></textarea>
          </div>
          
          <div style="display: flex; gap: 12px; margin-bottom: 24px;">
            <button class="btn btn-primary" @click="validarJson">
              <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M22 11.08V12a10 10 0 1 1-5.93-9.14"/>
                <polyline points="22 4 12 14.01 9 11.01"/>
              </svg>
              Validar Formato
            </button>
            <button class="btn btn-secondary" @click="simularImport">
              Simular Importación
            </button>
            <button class="btn btn-secondary" @click="limpiar">
              Limpiar
            </button>
          </div>
          
          <div v-if="loading" class="loading">
            <div class="spinner"></div>
          </div>
          
          <template v-else-if="resultado">
            <div class="card" style="background: var(--gray-50); border: 1px solid var(--gray-200);">
              <div class="card-body">
                <h3 style="font-size: 1rem; margin-bottom: 16px; display: flex; align-items: center; gap: 8px;">
                  <svg v-if="resultado.valido" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="#10b981" stroke-width="2">
                    <path d="M22 11.08V12a10 10 0 1 1-5.93-9.14"/>
                    <polyline points="22 4 12 14.01 9 11.01"/>
                  </svg>
                  <svg v-else width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="#ef4444" stroke-width="2">
                    <circle cx="12" cy="12" r="10"/>
                    <line x1="15" y1="9" x2="9" y2="15"/>
                    <line x1="9" y1="9" x2="15" y2="15"/>
                  </svg>
                  Resultado de Validación
                </h3>
                
                <div v-if="resultado.valido" class="detail-grid">
                  <div class="detail-item">
                    <label>Estado</label>
                    <span style="color: var(--success);">✓ Válido</span>
                  </div>
                  <div class="detail-item">
                    <label>Registros</label>
                    <span>{{ resultado.stats?.total || 0 }}</span>
                  </div>
                </div>
                
                <div v-else class="error-alert" style="margin-bottom: 0;">
                  {{ resultado.error }}
                </div>
                
                <div v-if="resultado.preview" style="margin-top: 16px;">
                  <h4 style="font-size: 0.875rem; color: var(--gray-600); margin-bottom: 8px;">Vista Previa:</h4>
                  <pre style="background: white; padding: 16px; border-radius: var(--radius); overflow-x: auto; font-size: 0.8125rem;">{{ JSON.stringify(resultado.preview, null, 2) }}</pre>
                </div>
              </div>
            </div>
          </template>
          
          <template v-else-if="error">
            <div class="error-alert">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <circle cx="12" cy="12" r="10"/>
                <line x1="12" y1="8" x2="12" y2="12"/>
                <line x1="12" y1="16" x2="12.01" y2="16"/>
              </svg>
              {{ error }}
            </div>
          </template>
        </div>
      </div>
      
      <div class="card" style="margin-top: 24px;">
        <div class="card-header">
          <h2 class="card-title">Formatos Soportados</h2>
        </div>
        <div class="card-body">
          <p style="color: var(--gray-500); margin-bottom: 16px;">
            El importador SIIGE espera un JSON con la siguiente estructura:
          </p>
          <pre style="background: var(--gray-800); color: var(--gray-100); padding: 20px; border-radius: var(--radius); overflow-x: auto; font-size: 0.8125rem; line-height: 1.8;">{{ formatoEjemplo }}</pre>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const jsonInput = ref('')
const resultado = ref(null)
const error = ref('')
const loading = ref(false)

const formatoEjemplo = `{
  "periodo": "2026-I",
  "facultades": [
    { "codigo": "FI", "nombre": "Facultad de Ingeniería" }
  ],
  "departamentos": [
    { "codigo": "COMP", "nombre": "Ciencias de la Computación", "facultad": "FI" }
  ],
  "escuelas": [
    { "codigo": "ING-SIST", "nombre": "Ingeniería de Sistemas", "departamento": "COMP" }
  ],
  "docentes": [
    { "codigo": "DOC001", "nombres": "Juan", "apellidos": "Pérez", "email": "jperez@unsch.edu", "departamento": "COMP" }
  ],
  "cursos": [
    { "codigo": "CS101", "nombre": "Introducción a la Programación", "creditos": 4, "escuela": "ING-SIST" }
  ]
}`

function validarJson() {
  error.value = ''
  resultado.value = null
  
  if (!jsonInput.value.trim()) {
    error.value = 'Por favor ingresa datos JSON para validar'
    return
  }
  
  try {
    const data = JSON.parse(jsonInput.value)
    
    const stats = {
      facultades: Array.isArray(data.facultades) ? data.facultades.length : 0,
      departamentos: Array.isArray(data.departamentos) ? data.departamentos.length : 0,
      escuelas: Array.isArray(data.escuelas) ? data.escuelas.length : 0,
      docentes: Array.isArray(data.docentes) ? data.docentes.length : 0,
      cursos: Array.isArray(data.cursos) ? data.cursos.length : 0
    }
    
    resultado.value = {
      valido: true,
      stats: {
        total: stats.facultades + stats.departamentos + stats.escuelas + stats.docentes + stats.cursos,
        ...stats
      },
      preview: data
    }
  } catch (e) {
    resultado.value = {
      valido: false,
      error: 'JSON inválido: ' + e.message
    }
  }
}

function simularImport() {
  validarJson()
  
  if (resultado.value?.valido) {
    loading.value = true
    setTimeout(() => {
      loading.value = false
      resultado.value.stats.simulado = true
      alert('Simulación completada. En producción, esto importaría los datos a la base de datos.')
    }, 1000)
  }
}

function limpiar() {
  jsonInput.value = ''
  resultado.value = null
  error.value = ''
}
</script>
