<template>
  <div class="siige-page">
    <header class="page-header">
      <div class="header-left">
        <h1 class="page-title">
          <svg width="28" height="28" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"/>
            <polyline points="17 8 12 3 7 8"/>
            <line x1="12" y1="3" x2="12" y2="15"/>
          </svg>
          Importador SIIGE
        </h1>
        <p class="page-subtitle">Importación masiva de datos desde el sistema SIIGE</p>
      </div>
      <div class="header-actions">
        <span class="beta-badge">
          <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <polygon points="13 2 3 14 12 14 11 22 21 10 12 10 13 2"/>
          </svg>
          Beta
        </span>
      </div>
    </header>

    <div class="page-content">
      <div class="main-grid">
        <div class="card import-card">
          <div class="card-header">
            <h2 class="card-title">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"/>
                <polyline points="17 8 12 3 7 8"/>
                <line x1="12" y1="3" x2="12" y2="15"/>
              </svg>
              Importar Datos
            </h2>
          </div>
          <div class="card-body">
            <p class="import-description">
              Pega el JSON exportado desde el sistema SIIGE para importar los datos de manera masiva.
              Esta herramienta permite validar el formato antes de realizar la importación.
            </p>

            <div class="form-group">
              <label class="form-label">Datos SIIGE (JSON)</label>
              <textarea
                v-model="jsonInput"
                class="form-input json-textarea"
                placeholder='{
  "periodo": "2026-I",
  "facultades": [...],
  "docentes": [...],
  "cursos": [...]
}'></textarea>
            </div>

            <div class="action-buttons">
              <button class="btn btn-primary" @click="validarJson">
                <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <path d="M22 11.08V12a10 10 0 1 1-5.93-9.14"/>
                  <polyline points="22 4 12 14.01 9 11.01"/>
                </svg>
                Validar Formato
              </button>
              <button class="btn btn-secondary" @click="simularImport">
                <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <polygon points="5 3 19 12 5 21 5 3"/>
                </svg>
                Simular Importación
              </button>
              <button class="btn btn-outline" @click="limpiar">
                <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <path d="M3 6h18"/>
                  <path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"/>
                </svg>
                Limpiar
              </button>
            </div>

            <div v-if="loading" class="loading-state">
              <div class="spinner-lg"></div>
              <p>Procesando importación...</p>
            </div>

            <template v-else-if="resultado">
              <div class="result-card" :class="resultado.valido ? 'valid' : 'invalid'">
                <div class="result-header">
                  <div class="result-icon" :class="resultado.valido ? 'success' : 'error'">
                    <svg v-if="resultado.valido" width="28" height="28" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                      <path d="M22 11.08V12a10 10 0 1 1-5.93-9.14"/>
                      <polyline points="22 4 12 14.01 9 11.01"/>
                    </svg>
                    <svg v-else width="28" height="28" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                      <circle cx="12" cy="12" r="10"/>
                      <line x1="15" y1="9" x2="9" y2="15"/>
                      <line x1="9" y1="9" x2="15" y2="15"/>
                    </svg>
                  </div>
                  <div class="result-info">
                    <h3>{{ resultado.valido ? 'JSON Válido' : 'JSON Inválido' }}</h3>
                    <p v-if="resultado.valido">El formato es correcto, puede proceder con la importación</p>
                    <p v-else class="error-message">{{ resultado.error }}</p>
                  </div>
                </div>

                <div v-if="resultado.valido && resultado.stats" class="stats-grid-result">
                  <div class="stat-item">
                    <span class="stat-value">{{ resultado.stats.facultades }}</span>
                    <span class="stat-label">Facultades</span>
                  </div>
                  <div class="stat-item">
                    <span class="stat-value">{{ resultado.stats.departamentos }}</span>
                    <span class="stat-label">Departamentos</span>
                  </div>
                  <div class="stat-item">
                    <span class="stat-value">{{ resultado.stats.escuelas }}</span>
                    <span class="stat-label">Escuelas</span>
                  </div>
                  <div class="stat-item">
                    <span class="stat-value">{{ resultado.stats.docentes }}</span>
                    <span class="stat-label">Docentes</span>
                  </div>
                  <div class="stat-item">
                    <span class="stat-value">{{ resultado.stats.cursos }}</span>
                    <span class="stat-label">Cursos</span>
                  </div>
                </div>

                <div v-if="resultado.preview" class="preview-section">
                  <h4>Vista Previa</h4>
                  <pre class="json-preview">{{ JSON.stringify(resultado.preview, null, 2) }}</pre>
                </div>
              </div>
            </template>

            <template v-else-if="error">
              <div class="error-banner">
                <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <circle cx="12" cy="12" r="10"/>
                  <line x1="12" y1="8" x2="12" y2="12"/>
                  <line x1="12" y1="16" x2="12.01" y2="16"/>
                </svg>
                <span>{{ error }}</span>
              </div>
            </template>
          </div>
        </div>

        <div class="card format-card">
          <div class="card-header">
            <h2 class="card-title">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/>
                <polyline points="14 2 14 8 20 8"/>
                <line x1="16" y1="13" x2="8" y2="13"/>
                <line x1="16" y1="17" x2="8" y2="17"/>
              </svg>
              Formato Esperado
            </h2>
          </div>
          <div class="card-body">
            <p class="format-description">
              El importador SIIGE espera un JSON con la siguiente estructura:
            </p>
            <pre class="code-block">{{ formatoEjemplo }}</pre>
          </div>
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

<style scoped>
.siige-page {
  min-height: 100vh;
  background: linear-gradient(135deg, #f0f4ff 0%, #fdf2f8 50%, #f0fdf4 100%);
}

.page-header {
  background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%);
  color: white;
  padding: 24px 32px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  box-shadow: 0 8px 32px rgba(240, 147, 251, 0.3);
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

.beta-badge {
  display: flex;
  align-items: center;
  gap: 6px;
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

.main-grid {
  display: grid;
  grid-template-columns: 1.5fr 1fr;
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

.import-description {
  color: #64748b;
  margin: 0 0 24px;
  line-height: 1.6;
}

.form-group {
  margin-bottom: 20px;
}

.form-label {
  display: block;
  font-size: 0.875rem;
  font-weight: 600;
  color: #374151;
  margin-bottom: 8px;
}

.form-input {
  width: 100%;
  padding: 14px 16px;
  border: 2px solid #e2e8f0;
  border-radius: 12px;
  font-size: 0.9rem;
  transition: all 0.3s;
  box-sizing: border-box;
}

.form-input:focus {
  outline: none;
  border-color: #f5576c;
  box-shadow: 0 0 0 4px rgba(245, 87, 108, 0.1);
}

.json-textarea {
  min-height: 200px;
  font-family: 'Monaco', 'Menlo', monospace;
  font-size: 0.85rem;
  resize: vertical;
  background: #f8fafc;
}

.action-buttons {
  display: flex;
  gap: 12px;
  margin-bottom: 24px;
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
  background: linear-gradient(135deg, #f093fb, #f5576c);
  color: white;
}

.btn-primary:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(245, 87, 108, 0.4);
}

.btn-secondary {
  background: linear-gradient(135deg, #11998e, #38ef7d);
  color: white;
}

.btn-secondary:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(17, 153, 142, 0.4);
}

.btn-outline {
  background: white;
  color: #64748b;
  border: 2px solid #e2e8f0;
}

.btn-outline:hover {
  border-color: #f5576c;
  color: #f5576c;
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
  border: 4px solid rgba(245, 87, 108, 0.2);
  border-top-color: #f5576c;
  border-radius: 50%;
  animation: spin 1s linear infinite;
  margin-bottom: 16px;
}

@keyframes spin {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}

.result-card {
  border-radius: 16px;
  padding: 20px;
  animation: fadeIn 0.3s;
}

@keyframes fadeIn {
  from { opacity: 0; transform: translateY(10px); }
  to { opacity: 1; transform: translateY(0); }
}

.result-card.valid {
  background: linear-gradient(135deg, #f0fdf4, #dcfce7);
  border: 2px solid #86efac;
}

.result-card.invalid {
  background: linear-gradient(135deg, #fef2f2, #fecaca);
  border: 2px solid #fca5a5;
}

.result-header {
  display: flex;
  align-items: center;
  gap: 16px;
  margin-bottom: 20px;
}

.result-icon {
  width: 56px;
  height: 56px;
  border-radius: 14px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.result-icon.success {
  background: linear-gradient(135deg, #11998e, #38ef7d);
  color: white;
}

.result-icon.error {
  background: linear-gradient(135deg, #f56565, #e53e3e);
  color: white;
}

.result-info h3 {
  margin: 0 0 4px;
  font-size: 1.1rem;
  color: #1e293b;
}

.result-info p {
  margin: 0;
  font-size: 0.9rem;
  color: #64748b;
}

.result-info .error-message {
  color: #dc2626;
}

.stats-grid-result {
  display: grid;
  grid-template-columns: repeat(5, 1fr);
  gap: 12px;
  margin-bottom: 20px;
}

.stat-item {
  text-align: center;
  padding: 12px;
  background: white;
  border-radius: 10px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.05);
}

.stat-item .stat-value {
  display: block;
  font-size: 1.5rem;
  font-weight: 700;
  background: linear-gradient(135deg, #f093fb, #f5576c);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}

.stat-item .stat-label {
  font-size: 0.7rem;
  color: #64748b;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.preview-section {
  background: white;
  border-radius: 12px;
  padding: 16px;
}

.preview-section h4 {
  margin: 0 0 12px;
  font-size: 0.85rem;
  color: #64748b;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.json-preview {
  background: #1e293b;
  color: #e2e8f0;
  padding: 16px;
  border-radius: 10px;
  font-size: 0.8rem;
  overflow-x: auto;
  max-height: 300px;
  margin: 0;
  line-height: 1.6;
}

.error-banner {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 16px 20px;
  background: #fef2f2;
  border-radius: 12px;
  border: 2px solid #fca5a5;
  color: #dc2626;
}

.format-description {
  color: #64748b;
  margin: 0 0 16px;
}

.code-block {
  background: #1e293b;
  color: #e2e8f0;
  padding: 20px;
  border-radius: 12px;
  font-size: 0.8rem;
  overflow-x: auto;
  line-height: 1.8;
  margin: 0;
}

@media (max-width: 1024px) {
  .main-grid {
    grid-template-columns: 1fr;
  }
  .stats-grid-result {
    grid-template-columns: repeat(3, 1fr);
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
  .action-buttons {
    flex-direction: column;
  }
  .stats-grid-result {
    grid-template-columns: repeat(2, 1fr);
  }
}
</style>
