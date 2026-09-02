<template>
  <div class="catalogos-page">
    <header class="page-header">
      <div class="header-left">
        <h1 class="page-title">
          <svg width="28" height="28" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M4 19.5A2.5 2.5 0 0 1 6.5 17H20"/>
            <path d="M6.5 2H20v20H6.5A2.5 2.5 0 0 1 4 19.5v-15A2.5 2.5 0 0 1 6.5 2z"/>
          </svg>
          Catálogos Institucionales
        </h1>
        <p class="page-subtitle">Gestión de estructura académica y administrativa</p>
      </div>
      <div class="header-actions">
        <button class="btn btn-secondary" @click="refresh">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M23 4v6h-6"/>
            <path d="M1 20v-6h6"/>
            <path d="M3.51 9a9 9 0 0 1 14.85-3.36L23 10M1 14l4.64 4.36A9 9 0 0 0 20.49 15"/>
          </svg>
          Actualizar
        </button>
      </div>
    </header>

    <div class="page-content">
      <div class="tabs-container">
        <div class="tabs">
          <button class="tab" :class="{ active: activeTab === 'facultades' }" @click="activeTab = 'facultades'">
            <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M3 21h18"/>
              <path d="M5 21V7l8-4v18"/>
              <path d="M19 21V11l-6-4"/>
              <path d="M9 9v.01"/>
              <path d="M9 12v.01"/>
              <path d="M9 15v.01"/>
              <path d="M9 18v.01"/>
            </svg>
            Facultades
            <span class="tab-count">{{ catalogos.facultades.length }}</span>
          </button>
          <button class="tab" :class="{ active: activeTab === 'departamentos' }" @click="activeTab = 'departamentos'">
            <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M3 9l9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"/>
              <polyline points="9 22 9 12 15 12 15 22"/>
            </svg>
            Departamentos
            <span class="tab-count">{{ catalogos.departamentos.length }}</span>
          </button>
          <button class="tab" :class="{ active: activeTab === 'escuelas' }" @click="activeTab = 'escuelas'">
            <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M22 10v6M2 10l10-5 10 5-10 5z"/>
              <path d="M6 12v5c0 2 2 3 6 3s6-1 6-3v-5"/>
            </svg>
            Escuelas
            <span class="tab-count">{{ catalogos.escuelas.length }}</span>
          </button>
          <button class="tab" :class="{ active: activeTab === 'docentes' }" @click="activeTab = 'docentes'">
            <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"/>
              <circle cx="9" cy="7" r="4"/>
              <path d="M23 21v-2a4 4 0 0 0-3-3.87"/>
              <path d="M16 3.13a4 4 0 0 1 0 7.75"/>
            </svg>
            Docentes
            <span class="tab-count">{{ catalogos.docentes.length }}</span>
          </button>
          <button class="tab" :class="{ active: activeTab === 'cursos' }" @click="activeTab = 'cursos'">
            <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M4 19.5A2.5 2.5 0 0 1 6.5 17H20"/>
              <path d="M6.5 2H20v20H6.5A2.5 2.5 0 0 1 4 19.5v-15A2.5 2.5 0 0 1 6.5 2z"/>
            </svg>
            Cursos
            <span class="tab-count">{{ catalogos.cursos.length }}</span>
          </button>
          <button class="tab" :class="{ active: activeTab === 'aulas' }" @click="activeTab = 'aulas'">
            <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M3 9l9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"/>
              <polyline points="9 22 9 12 15 12 15 22"/>
            </svg>
            Aulas
            <span class="tab-count">{{ allAulas.length }}</span>
          </button>
          <button class="tab" :class="{ active: activeTab === 'periodos' }" @click="activeTab = 'periodos'">
            <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <rect x="3" y="4" width="18" height="18" rx="2" ry="2"/>
              <line x1="16" y1="2" x2="16" y2="6"/>
              <line x1="8" y1="2" x2="8" y2="6"/>
              <line x1="3" y1="10" x2="21" y2="10"/>
            </svg>
            Períodos
            <span class="tab-count">{{ catalogos.periodos.length }}</span>
          </button>
          <button class="tab" :class="{ active: activeTab === 'planes' }" @click="activeTab = 'planes'">
            <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M4 19.5A2.5 2.5 0 0 1 6.5 17H20"/>
              <path d="M6.5 2H20v20H6.5A2.5 2.5 0 0 1 4 19.5v-15A2.5 2.5 0 0 1 6.5 2z"/>
            </svg>
            Planes
            <span class="tab-count">{{ catalogos.planes.length }}</span>
          </button>
        </div>
      </div>

      <div v-if="loading" class="loading-overlay">
        <div class="spinner-lg"></div>
        <p>Cargando catálogos...</p>
      </div>

      <div v-else-if="error" class="error-banner">
        <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <circle cx="12" cy="12" r="10"/>
          <line x1="12" y1="8" x2="12" y2="12"/>
          <line x1="12" y1="16" x2="12.01" y2="16"/>
        </svg>
        <span>{{ error }}</span>
        <button @click="refresh" class="btn btn-sm">Reintentar</button>
      </div>

      <template v-else>
        <div v-show="activeTab === 'facultades'" class="tab-content">
          <div class="content-header">
            <div>
              <h2>Facultades</h2>
              <p>Listado de facultades registradas en el sistema</p>
            </div>
            <button class="btn btn-primary" @click="openNuevaFacultad">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <line x1="12" y1="5" x2="12" y2="19"/>
                <line x1="5" y1="12" x2="19" y2="12"/>
              </svg>
              Nueva Facultad
            </button>
          </div>
          <div class="table-container">
            <table class="data-table">
              <thead>
                <tr>
                  <th>ID</th>
                  <th>Nombre</th>
                  <th>Acciones</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="item in catalogos.facultades" :key="item.id_facultad">
                  <td class="id-cell">{{ item.id_facultad }}</td>
                  <td class="name-cell">
                    <div class="name-content">
                      <span class="entity-icon facultad-icon">
                        <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                          <path d="M3 21h18"/>
                          <path d="M5 21V7l8-4v18"/>
                        </svg>
                      </span>
                      {{ item.nombre }}
                    </div>
                  </td>
                  <td class="actions-cell">
                    <button class="btn btn-icon-sm" @click="openEditarFacultad(item)" title="Editar">
                      <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"/>
                        <path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z"/>
                      </svg>
                    </button>
                    <button class="btn btn-icon-sm btn-danger" @click="eliminarFacultad(item)" title="Eliminar">
                      <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <polyline points="3 6 5 6 21 6"/>
                        <path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"/>
                      </svg>
                    </button>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>

        <div v-show="activeTab === 'departamentos'" class="tab-content">
          <div class="content-header">
            <div>
              <h2>Departamentos Académicos</h2>
              <p>Departamentos organizados por facultad</p>
            </div>
            <button class="btn btn-primary" @click="openNuevoDepartamento">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <line x1="12" y1="5" x2="12" y2="19"/>
                <line x1="5" y1="12" x2="19" y2="12"/>
              </svg>
              Nuevo Departamento
            </button>
          </div>
          <div class="table-container">
            <table class="data-table">
              <thead>
                <tr>
                  <th>ID</th>
                  <th>Facultad</th>
                  <th>Nombre</th>
                  <th>Acciones</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="item in catalogos.departamentos" :key="item.id_departamento">
                  <td class="id-cell">{{ item.id_departamento }}</td>
                  <td>
                    <span class="facultad-badge">{{ getFacultadNombre(item.id_facultad) }}</span>
                  </td>
                  <td class="name-cell">
                    <div class="name-content">
                      <span class="entity-icon depto-icon">
                        <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                          <path d="M3 9l9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"/>
                        </svg>
                      </span>
                      {{ item.nombre }}
                    </div>
                  </td>
                  <td class="actions-cell">
                    <button class="btn btn-icon-sm" @click="openEditarDepartamento(item)" title="Editar">
                      <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"/>
                        <path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z"/>
                      </svg>
                    </button>
                    <button class="btn btn-icon-sm btn-danger" @click="eliminarDepartamento(item)" title="Eliminar">
                      <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <polyline points="3 6 5 6 21 6"/>
                        <path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"/>
                      </svg>
                    </button>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>

        <div v-show="activeTab === 'escuelas'" class="tab-content">
          <div class="content-header">
            <div>
              <h2>Escuelas Profesionales</h2>
              <p>Escuelas organizadas por departamento y facultad</p>
            </div>
            <button class="btn btn-primary" @click="openNuevaEscuela">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <line x1="12" y1="5" x2="12" y2="19"/>
                <line x1="5" y1="12" x2="19" y2="12"/>
              </svg>
              Nueva Escuela
            </button>
          </div>
          <div class="table-container">
            <table class="data-table">
              <thead>
                <tr>
                  <th>ID</th>
                  <th>Facultad</th>
                  <th>Departamento</th>
                  <th>Nombre</th>
                  <th>Acciones</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="item in catalogos.escuelas" :key="item.id_escuela">
                  <td class="id-cell">{{ item.id_escuela }}</td>
                  <td><span class="facultad-badge">{{ getFacultadNombre(item.id_facultad) }}</span></td>
                  <td><span class="depto-badge">{{ getDepartamentoNombre(item.id_departamento) }}</span></td>
                  <td class="name-cell">
                    <div class="name-content">
                      <span class="entity-icon escuela-icon">
                        <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                          <path d="M22 10v6M2 10l10-5 10 5z"/>
                          <path d="M6 12v5c0 2 2 3 6 3s6-1 6-3v-5"/>
                        </svg>
                      </span>
                      {{ item.nombre }}
                    </div>
                  </td>
                  <td class="actions-cell">
                    <button class="btn btn-icon-sm" @click="openEditarEscuela(item)" title="Editar">
                      <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"/>
                        <path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z"/>
                      </svg>
                    </button>
                    <button class="btn btn-icon-sm btn-danger" @click="eliminarEscuela(item)" title="Eliminar">
                      <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <polyline points="3 6 5 6 21 6"/>
                        <path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"/>
                      </svg>
                    </button>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>

        <div v-show="activeTab === 'docentes'" class="tab-content">
          <div class="content-header">
            <div>
              <h2>Docentes</h2>
              <p>Personal docente registrado en el sistema</p>
            </div>
            <button class="btn btn-primary" @click="openNuevoDocente">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <line x1="12" y1="5" x2="12" y2="19"/>
                <line x1="5" y1="12" x2="19" y2="12"/>
              </svg>
              Nuevo Docente
            </button>
          </div>
          <div class="table-container">
            <table class="data-table">
              <thead>
                <tr>
                  <th>ID</th>
                  <th>Departamento</th>
                  <th>Código Plaza</th>
                  <th>Nombres</th>
                  <th>Apellidos</th>
                  <th>Email</th>
                  <th>Acciones</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="item in catalogos.docentes" :key="item.id_docente">
                  <td class="id-cell">{{ item.id_docente }}</td>
                  <td><span class="depto-badge">{{ getDepartamentoNombre(item.id_departamento) }}</span></td>
                  <td><span class="codigo-badge">{{ item.codigo_plaza }}</span></td>
                  <td>{{ item.nombres }}</td>
                  <td class="fw-600">{{ item.apellidos }}</td>
                  <td class="email-cell">{{ item.email || '—' }}</td>
                  <td class="actions-cell">
                    <button class="btn btn-icon-sm" @click="openEditarDocente(item)" title="Editar">
                      <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"/>
                        <path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z"/>
                      </svg>
                    </button>
                    <button class="btn btn-icon-sm btn-danger" @click="eliminarDocente(item)" title="Eliminar">
                      <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <polyline points="3 6 5 6 21 6"/>
                        <path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"/>
                      </svg>
                    </button>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>

        <div v-show="activeTab === 'cursos'" class="tab-content">
          <div class="content-header">
            <div>
              <h2>Cursos</h2>
              <p>Cursos registrados en el sistema académico</p>
            </div>
            <button class="btn btn-primary" @click="openNuevoCurso">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <line x1="12" y1="5" x2="12" y2="19"/>
                <line x1="5" y1="12" x2="19" y2="12"/>
              </svg>
              Nuevo Curso
            </button>
          </div>
          <div class="table-container">
            <table class="data-table">
              <thead>
                <tr>
                  <th>ID</th>
                  <th>Código</th>
                  <th>Nombre</th>
                  <th>Escuela</th>
                  <th>Créditos</th>
                  <th>Hrs. Teoría</th>
                  <th>Hrs. Práctica</th>
                  <th>Acciones</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="item in catalogos.cursos" :key="item.id_curso">
                  <td class="id-cell">{{ item.id_curso }}</td>
                  <td><span class="codigo-badge curso-codigo">{{ item.codigo }}</span></td>
                  <td class="fw-600">{{ item.nombre }}</td>
                  <td><span class="escuela-badge">{{ item.escuela_nombre }}</span></td>
                  <td><span class="creditos-badge">{{ item.creditos }}</span></td>
                  <td>{{ item.horas_teoria }}</td>
                  <td>{{ item.horas_practica }}</td>
                  <td class="actions-cell">
                    <button class="btn btn-icon-sm" @click="openEditarCurso(item)" title="Editar">
                      <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"/>
                        <path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z"/>
                      </svg>
                    </button>
                    <button class="btn btn-icon-sm btn-danger" @click="eliminarCurso(item)" title="Eliminar">
                      <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <polyline points="3 6 5 6 21 6"/>
                        <path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"/>
                      </svg>
                    </button>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>

        <div v-show="activeTab === 'aulas'" class="tab-content">
          <div class="content-header">
            <div>
              <h2>Aulas</h2>
              <p>Gestionar disponibilidad de aulas en el sistema</p>
            </div>
            <button class="btn btn-primary" @click="openNuevaAula">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <line x1="12" y1="5" x2="12" y2="19"/>
                <line x1="5" y1="12" x2="19" y2="12"/>
              </svg>
              Nueva Aula
            </button>
          </div>
          <div class="table-container">
            <table class="data-table">
              <thead>
                <tr>
                  <th>ID</th>
                  <th>Código</th>
                  <th>Tipo</th>
                  <th>Pabellón</th>
                  <th>Aforo</th>
                  <th>Compartida</th>
                  <th>Estado</th>
                  <th>Acciones</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="item in allAulas" :key="item.id_aula" :class="{ 'row-inactive': !item.activo }">
                  <td class="id-cell">{{ item.id_aula }}</td>
                  <td><span class="codigo-badge">{{ item.codigo }}</span></td>
                  <td><span class="tipo-badge" :class="'tipo-' + item.tipo.toLowerCase()">{{ item.tipo }}</span></td>
                  <td>{{ getPabellonNombre(item.id_pabellon) }}</td>
                  <td>{{ item.aforo }}</td>
                  <td>
                    <span v-if="item.es_compartida" class="badge-compartida">Sí</span>
                    <span v-else class="badge-no">No</span>
                  </td>
                  <td>
                    <span v-if="item.activo" class="badge-activo">Activa</span>
                    <span v-else class="badge-inactivo">Inactiva</span>
                  </td>
                  <td class="actions-cell">
                    <button
                      class="btn btn-icon-sm"
                      @click="openEditarAula(item)"
                      title="Editar aula"
                    >
                      <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"/>
                        <path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z"/>
                      </svg>
                    </button>
                    <button
                      class="btn btn-icon-sm btn-danger"
                      @click="eliminarAula(item)"
                      title="Eliminar aula"
                    >
                      <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <polyline points="3 6 5 6 21 6"/>
                        <path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"/>
                      </svg>
                    </button>
                    <button
                      v-if="item.activo"
                      class="btn btn-icon-sm btn-warning"
                      @click="toggleAulaActivo(item)"
                      title="Desactivar aula"
                    >
                      <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <circle cx="12" cy="12" r="10"/>
                        <line x1="4.93" y1="4.93" x2="19.07" y2="19.07"/>
                      </svg>
                    </button>
                    <button
                      v-else
                      class="btn btn-icon-sm btn-success"
                      @click="toggleAulaActivo(item)"
                      title="Activar aula"
                    >
                      <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <circle cx="12" cy="12" r="10"/>
                        <polyline points="16 12 12 8 8 12"/>
                        <line x1="12" y1="16" x2="12" y2="8"/>
                      </svg>
                    </button>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>

        <div v-show="activeTab === 'periodos'" class="tab-content">
          <div class="content-header">
            <div>
              <h2>Períodos Académicos</h2>
              <p>Gestionar períodos lectivos del sistema</p>
            </div>
            <button class="btn btn-primary" @click="openNuevoPeriodo">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <line x1="12" y1="5" x2="12" y2="19"/>
                <line x1="5" y1="12" x2="19" y2="12"/>
              </svg>
              Nuevo Período
            </button>
          </div>
          <div class="table-container">
            <table class="data-table">
              <thead>
                <tr>
                  <th>ID</th>
                  <th>Código</th>
                  <th>Estado</th>
                  <th>Acciones</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="item in catalogos.periodos" :key="item.id_periodo">
                  <td class="id-cell">{{ item.id_periodo }}</td>
                  <td><span class="codigo-badge">{{ item.codigo }}</span></td>
                  <td>
                    <span v-if="item.activo" class="badge-activo">Activo</span>
                    <span v-else class="badge-inactivo">Inactivo</span>
                  </td>
                  <td class="actions-cell">
                    <button class="btn btn-icon-sm" @click="openEditarPeriodo(item)" title="Editar">
                      <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"/>
                        <path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z"/>
                      </svg>
                    </button>
                    <button class="btn btn-icon-sm btn-danger" @click="eliminarPeriodo(item)" title="Eliminar">
                      <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <polyline points="3 6 5 6 21 6"/>
                        <path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"/>
                      </svg>
                    </button>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>

        <div v-show="activeTab === 'planes'" class="tab-content">
          <div class="content-header">
            <div>
              <h2>Planes de Estudio</h2>
              <p>Planes de estudio asociados a escuelas profesionales</p>
            </div>
            <button class="btn btn-primary" @click="openNuevoPlan">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <line x1="12" y1="5" x2="12" y2="19"/>
                <line x1="5" y1="12" x2="19" y2="12"/>
              </svg>
              Nuevo Plan
            </button>
          </div>
          <div class="table-container">
            <table class="data-table">
              <thead>
                <tr>
                  <th>ID</th>
                  <th>Código</th>
                  <th>Nombre</th>
                  <th>Escuela</th>
                  <th>Acciones</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="item in catalogos.planes" :key="item.id_plan">
                  <td class="id-cell">{{ item.id_plan }}</td>
                  <td><span class="codigo-badge">{{ item.codigo_plan }}</span></td>
                  <td class="fw-600">{{ item.nombre }}</td>
                  <td><span class="escuela-badge">{{ getEscuelaNombre(item.id_escuela) }}</span></td>
                  <td class="actions-cell">
                    <button class="btn btn-icon-sm" @click="openEditarPlan(item)" title="Editar">
                      <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"/>
                        <path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z"/>
                      </svg>
                    </button>
                    <button class="btn btn-icon-sm btn-danger" @click="eliminarPlan(item)" title="Eliminar">
                      <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <polyline points="3 6 5 6 21 6"/>
                        <path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"/>
                      </svg>
                    </button>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </template>

      <!-- Modal Facultad -->
      <div v-if="showModalFacultad" class="modal-overlay" @click.self="showModalFacultad = false">
        <div class="modal modal-md">
          <div class="modal-header">
            <h2>
              <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M3 21h18"/>
                <path d="M5 21V7l8-4v18"/>
              </svg>
              {{ editingFacultad ? 'Editar' : 'Nueva' }} Facultad
            </h2>
            <button class="btn btn-icon" @click="showModalFacultad = false">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <line x1="18" y1="6" x2="6" y2="18"/>
                <line x1="6" y1="6" x2="18" y2="18"/>
              </svg>
            </button>
          </div>
          <div class="modal-body">
            <div class="form-group">
              <label class="form-label">Nombre de la Facultad</label>
              <input v-model="facultadForm.nombre" type="text" class="form-input" placeholder="Ej: Facultad de Ingeniería">
            </div>
          </div>
          <div class="modal-footer">
            <button class="btn btn-secondary" @click="showModalFacultad = false">Cancelar</button>
            <button class="btn btn-primary" @click="guardarFacultad" :disabled="saving">
              {{ saving ? 'Guardando...' : 'Guardar' }}
            </button>
          </div>
        </div>
      </div>

      <!-- Modal Departamento -->
      <div v-if="showModalDepartamento" class="modal-overlay" @click.self="showModalDepartamento = false">
        <div class="modal modal-md">
          <div class="modal-header">
            <h2>
              <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M3 9l9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"/>
              </svg>
              {{ editingDepartamento ? 'Editar' : 'Nuevo' }} Departamento
            </h2>
            <button class="btn btn-icon" @click="showModalDepartamento = false">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <line x1="18" y1="6" x2="6" y2="18"/>
                <line x1="6" y1="6" x2="18" y2="18"/>
              </svg>
            </button>
          </div>
          <div class="modal-body">
            <div class="form-group">
              <label class="form-label">Facultad</label>
              <select v-model="departamentoForm.id_facultad" class="form-input">
                <option value="">-- Seleccionar facultad --</option>
                <option v-for="f in catalogos.facultades" :key="f.id_facultad" :value="f.id_facultad">
                  {{ f.nombre }}
                </option>
              </select>
            </div>
            <div class="form-group">
              <label class="form-label">Nombre del Departamento</label>
              <input v-model="departamentoForm.nombre" type="text" class="form-input" placeholder="Ej: Departamento de Ingeniería de Sistemas">
            </div>
          </div>
          <div class="modal-footer">
            <button class="btn btn-secondary" @click="showModalDepartamento = false">Cancelar</button>
            <button class="btn btn-primary" @click="guardarDepartamento" :disabled="saving">
              {{ saving ? 'Guardando...' : 'Guardar' }}
            </button>
          </div>
        </div>
      </div>

      <!-- Modal Aula -->
      <div v-if="showModalAula" class="modal-overlay" @click.self="showModalAula = false">
        <div class="modal modal-md">
          <div class="modal-header">
            <h2>
              <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M3 9l9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"/>
                <polyline points="9 22 9 12 15 12 15 22"/>
              </svg>
              {{ editingAula ? 'Editar' : 'Nueva' }} Aula
            </h2>
            <button class="btn btn-icon" @click="showModalAula = false">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <line x1="18" y1="6" x2="6" y2="18"/>
                <line x1="6" y1="6" x2="18" y2="18"/>
              </svg>
            </button>
          </div>
          <div class="modal-body">
            <div class="form-group">
              <label class="form-label">Código del Aula *</label>
              <input v-model="aulaForm.codigo" type="text" class="form-input" placeholder="Ej: H-201">
            </div>
            <div class="form-group">
              <label class="form-label">Pabellón *</label>
              <select v-model="aulaForm.id_pabellon" class="form-input">
                <option value="">-- Seleccionar pabellón --</option>
                <option v-for="p in pabellones" :key="p.id_pabellon" :value="p.id_pabellon">
                  {{ p.codigo }} - {{ p.nombre }}
                </option>
              </select>
            </div>
            <div class="form-group">
              <label class="form-label">Tipo de Aula *</label>
              <select v-model="aulaForm.tipo" class="form-input">
                <option value="TEORIA">Teoría</option>
                <option value="PRACTICA">Práctica</option>
                <option value="LABORATORIO">Laboratorio</option>
              </select>
            </div>
            <div class="form-group">
              <label class="form-label">Aforo (número de estudiantes) *</label>
              <input v-model="aulaForm.aforo" type="number" min="1" class="form-input" placeholder="Ej: 40">
            </div>
            <div class="form-group">
              <label class="checkbox-label">
                <input v-model="aulaForm.es_compartida" type="checkbox">
                <span class="checkbox-custom"></span>
                Aula compartida entre escuelas
              </label>
            </div>
          </div>
          <div class="modal-footer">
            <button class="btn btn-secondary" @click="showModalAula = false">Cancelar</button>
            <button class="btn btn-primary" @click="guardarAula" :disabled="saving">
              {{ saving ? 'Guardando...' : 'Guardar' }}
            </button>
          </div>
        </div>
      </div>

      <!-- Modal Escuela -->
      <div v-if="showModalEscuela" class="modal-overlay" @click.self="showModalEscuela = false">
        <div class="modal modal-md">
          <div class="modal-header">
            <h2>
              <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M22 10v6M2 10l10-5 10 5z"/>
                <path d="M6 12v5c0 2 2 3 6 3s6-1 6-3v-5"/>
              </svg>
              {{ editingEscuela ? 'Editar' : 'Nueva' }} Escuela
            </h2>
            <button class="btn btn-icon" @click="showModalEscuela = false">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <line x1="18" y1="6" x2="6" y2="18"/>
                <line x1="6" y1="6" x2="18" y2="18"/>
              </svg>
            </button>
          </div>
          <div class="modal-body">
            <div class="form-group">
              <label class="form-label">Facultad *</label>
              <select v-model="escuelaForm.id_facultad" class="form-input">
                <option value="">-- Seleccionar facultad --</option>
                <option v-for="f in catalogos.facultades" :key="f.id_facultad" :value="f.id_facultad">
                  {{ f.nombre }}
                </option>
              </select>
            </div>
            <div class="form-group">
              <label class="form-label">Departamento *</label>
              <select v-model="escuelaForm.id_departamento" class="form-input">
                <option value="">-- Seleccionar departamento --</option>
                <option v-for="d in filteredDepartamentos" :key="d.id_departamento" :value="d.id_departamento">
                  {{ d.nombre }}
                </option>
              </select>
            </div>
            <div class="form-group">
              <label class="form-label">Nombre de la Escuela *</label>
              <input v-model="escuelaForm.nombre" type="text" class="form-input" placeholder="Ej: Escuela Profesional de Ingeniería de Sistemas">
            </div>
          </div>
          <div class="modal-footer">
            <button class="btn btn-secondary" @click="showModalEscuela = false">Cancelar</button>
            <button class="btn btn-primary" @click="guardarEscuela" :disabled="saving">
              {{ saving ? 'Guardando...' : 'Guardar' }}
            </button>
          </div>
        </div>
      </div>

      <!-- Modal Docente -->
      <div v-if="showModalDocente" class="modal-overlay" @click.self="showModalDocente = false">
        <div class="modal modal-md">
          <div class="modal-header">
            <h2>
              <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"/>
                <circle cx="9" cy="7" r="4"/>
                <path d="M23 21v-2a4 4 0 0 0-3-3.87"/>
                <path d="M16 3.13a4 4 0 0 1 0 7.75"/>
              </svg>
              {{ editingDocente ? 'Editar' : 'Nuevo' }} Docente
            </h2>
            <button class="btn btn-icon" @click="showModalDocente = false">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <line x1="18" y1="6" x2="6" y2="18"/>
                <line x1="6" y1="6" x2="18" y2="18"/>
              </svg>
            </button>
          </div>
          <div class="modal-body">
            <div class="form-group">
              <label class="form-label">Departamento *</label>
              <select v-model="docenteForm.id_departamento" class="form-input">
                <option value="">-- Seleccionar departamento --</option>
                <option v-for="d in catalogos.departamentos" :key="d.id_departamento" :value="d.id_departamento">
                  {{ d.nombre }}
                </option>
              </select>
            </div>
            <div class="form-group">
              <label class="form-label">Código de Plaza *</label>
              <input v-model="docenteForm.codigo_plaza" type="text" class="form-input" placeholder="Ej: 1TC04">
            </div>
            <div class="form-group">
              <label class="form-label">Nombres *</label>
              <input v-model="docenteForm.nombres" type="text" class="form-input" placeholder="Ej: Juan Carlos">
            </div>
            <div class="form-group">
              <label class="form-label">Apellidos *</label>
              <input v-model="docenteForm.apellidos" type="text" class="form-input" placeholder="Ej: Pérez García">
            </div>
            <div class="form-group">
              <label class="form-label">Email</label>
              <input v-model="docenteForm.email" type="email" class="form-input" placeholder="Ej: jperez@unsch.edu.pe">
            </div>
          </div>
          <div class="modal-footer">
            <button class="btn btn-secondary" @click="showModalDocente = false">Cancelar</button>
            <button class="btn btn-primary" @click="guardarDocente" :disabled="saving">
              {{ saving ? 'Guardando...' : 'Guardar' }}
            </button>
          </div>
        </div>
      </div>

      <!-- Modal Curso -->
      <div v-if="showModalCurso" class="modal-overlay" @click.self="showModalCurso = false">
        <div class="modal modal-md">
          <div class="modal-header">
            <h2>
              <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M4 19.5A2.5 2.5 0 0 1 6.5 17H20"/>
                <path d="M6.5 2H20v20H6.5A2.5 2.5 0 0 1 4 19.5v-15A2.5 2.5 0 0 1 6.5 2z"/>
              </svg>
              {{ editingCurso ? 'Editar' : 'Nuevo' }} Curso
            </h2>
            <button class="btn btn-icon" @click="showModalCurso = false">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <line x1="18" y1="6" x2="6" y2="18"/>
                <line x1="6" y1="6" x2="18" y2="18"/>
              </svg>
            </button>
          </div>
          <div class="modal-body">
            <div class="form-group">
              <label class="form-label">Serie (Ciclo) *</label>
              <select v-model="cursoForm.id_serie" class="form-input">
                <option value="">-- Seleccionar serie --</option>
                <option v-for="s in series" :key="s.id_serie" :value="s.id_serie">
                  {{ s.codigo_plan }} - Ciclo {{ s.numero_ciclo }}{{ s.subciclo > 1 ? ' (II)' : '' }}
                </option>
              </select>
            </div>
            <div class="form-group">
              <label class="form-label">Código del Curso *</label>
              <input v-model="cursoForm.codigo" type="text" class="form-input" placeholder="Ej: CS101">
            </div>
            <div class="form-group">
              <label class="form-label">Nombre del Curso *</label>
              <input v-model="cursoForm.nombre" type="text" class="form-input" placeholder="Ej: Introducción a la Programación">
            </div>
            <div class="form-row">
              <div class="form-group">
                <label class="form-label">Créditos *</label>
                <input v-model="cursoForm.creditos" type="number" min="1" class="form-input" placeholder="Ej: 4">
              </div>
              <div class="form-group">
                <label class="form-label">Hrs. Teoría *</label>
                <input v-model="cursoForm.horas_teoria" type="number" min="0" class="form-input" placeholder="Ej: 3">
              </div>
              <div class="form-group">
                <label class="form-label">Hrs. Práctica *</label>
                <input v-model="cursoForm.horas_practica" type="number" min="0" class="form-input" placeholder="Ej: 2">
              </div>
            </div>
          </div>
          <div class="modal-footer">
            <button class="btn btn-secondary" @click="showModalCurso = false">Cancelar</button>
            <button class="btn btn-primary" @click="guardarCurso" :disabled="saving">
              {{ saving ? 'Guardando...' : 'Guardar' }}
            </button>
          </div>
        </div>

        <!-- Modal Periodo -->
        <div v-if="showModalPeriodo" class="modal-overlay" @click.self="showModalPeriodo = false">
          <div class="modal modal-sm">
            <div class="modal-header">
              <h2>
                <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <rect x="3" y="4" width="18" height="18" rx="2" ry="2"/>
                  <line x1="16" y1="2" x2="16" y2="6"/>
                  <line x1="8" y1="2" x2="8" y2="6"/>
                  <line x1="3" y1="10" x2="21" y2="10"/>
                </svg>
                {{ editingPeriodo ? 'Editar' : 'Nuevo' }} Período
              </h2>
              <button class="btn btn-icon" @click="showModalPeriodo = false">
                <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <line x1="18" y1="6" x2="6" y2="18"/>
                  <line x1="6" y1="6" x2="18" y2="18"/>
                </svg>
              </button>
            </div>
            <div class="modal-body">
              <div class="form-group">
                <label class="form-label">Código del Período *</label>
                <input v-model="periodoForm.codigo" type="text" class="form-input" placeholder="Ej: 2025-I">
              </div>
              <div class="form-group">
                <label class="checkbox-label">
                  <input v-model="periodoForm.activo" type="checkbox">
                  <span class="checkbox-custom"></span>
                  Período activo
                </label>
              </div>
            </div>
            <div class="modal-footer">
              <button class="btn btn-secondary" @click="showModalPeriodo = false">Cancelar</button>
              <button class="btn btn-primary" @click="guardarPeriodo" :disabled="saving">
                {{ saving ? 'Guardando...' : 'Guardar' }}
              </button>
            </div>
          </div>
        </div>

        <!-- Modal Plan de Estudio -->
        <div v-if="showModalPlan" class="modal-overlay" @click.self="showModalPlan = false">
          <div class="modal modal-md">
            <div class="modal-header">
              <h2>
                <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <path d="M4 19.5A2.5 2.5 0 0 1 6.5 17H20"/>
                  <path d="M6.5 2H20v20H6.5A2.5 2.5 0 0 1 4 19.5v-15A2.5 2.5 0 0 1 6.5 2z"/>
                </svg>
                {{ editingPlan ? 'Editar' : 'Nuevo' }} Plan de Estudio
              </h2>
              <button class="btn btn-icon" @click="showModalPlan = false">
                <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <line x1="18" y1="6" x2="6" y2="18"/>
                  <line x1="6" y1="6" x2="18" y2="18"/>
                </svg>
              </button>
            </div>
            <div class="modal-body">
              <div class="form-group">
                <label class="form-label">Escuela *</label>
                <select v-model="planForm.id_escuela" class="form-input">
                  <option value="">-- Seleccionar escuela --</option>
                  <option v-for="e in catalogos.escuelas" :key="e.id_escuela" :value="e.id_escuela">
                    {{ e.nombre }}
                  </option>
                </select>
              </div>
              <div class="form-group">
                <label class="form-label">Código del Plan *</label>
                <input v-model="planForm.codigo" type="text" class="form-input" placeholder="Ej: 2020-I">
              </div>
              <div class="form-group">
                <label class="form-label">Nombre del Plan *</label>
                <input v-model="planForm.nombre" type="text" class="form-input" placeholder="Ej: Plan de Estudios 2020">
              </div>
            </div>
            <div class="modal-footer">
              <button class="btn btn-secondary" @click="showModalPlan = false">Cancelar</button>
              <button class="btn btn-primary" @click="guardarPlan" :disabled="saving">
                {{ saving ? 'Guardando...' : 'Guardar' }}
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import api from '../services/api'

const activeTab = ref('facultades')
const loading = ref(true)
const error = ref('')

const showModalFacultad = ref(false)
const editingFacultad = ref(null)
const saving = ref(false)
const facultadForm = ref({
  nombre: ''
})

const showModalDepartamento = ref(false)
const editingDepartamento = ref(null)
const departamentoForm = ref({
  id_facultad: '',
  nombre: ''
})

const catalogos = ref({
  facultades: [],
  departamentos: [],
  escuelas: [],
  docentes: [],
  cursos: [],
  periodos: [],
  planes: []
})

const allAulas = ref([])
const pabellones = ref([])

const showModalAula = ref(false)
const editingAula = ref(null)
const aulaForm = ref({
  codigo: '',
  id_pabellon: '',
  tipo: 'TEORIA',
  aforo: '',
  es_compartida: false
})

const showModalEscuela = ref(false)
const editingEscuela = ref(null)
const escuelaForm = ref({
  id_facultad: '',
  id_departamento: '',
  nombre: ''
})

const showModalDocente = ref(false)
const editingDocente = ref(null)
const docenteForm = ref({
  id_departamento: '',
  codigo_plaza: '',
  nombres: '',
  apellidos: '',
  email: ''
})

const showModalCurso = ref(false)
const editingCurso = ref(null)
const cursoForm = ref({
  id_serie: '',
  codigo: '',
  nombre: '',
  creditos: '',
  horas_teoria: '',
  horas_practica: ''
})

const series = ref([])

const showModalPeriodo = ref(false)
const editingPeriodo = ref(null)
const periodoForm = ref({
  codigo: '',
  activo: true
})

const showModalPlan = ref(false)
const editingPlan = ref(null)
const planForm = ref({
  id_escuela: '',
  codigo: '',
  nombre: ''
})

async function loadAll() {
  loading.value = true
  error.value = ''
  try {
    const [facultades, departamentos, escuelas, docentes, cursos, aulas, pabs, seriesData, periodos, planes] = await Promise.all([
      api.facultades.list(),
      api.departamentos.list(),
      api.escuelas.list(),
      api.docentes.list(),
      api.cursos.list(),
      api.aulas.listAll(),
      api.get('/pabellones'),
      api.series.list(),
      api.periodos.list(),
      api.planesEstudio.list()
    ])

    catalogos.value = { facultades, departamentos, escuelas, docentes, cursos, periodos, planes }
    allAulas.value = aulas
    pabellones.value = pabs
    series.value = seriesData
  } catch (e) {
    error.value = e.message
  } finally {
    loading.value = false
  }
}

function refresh() {
  loadAll()
}

function openNuevaFacultad() {
  editingFacultad.value = null
  facultadForm.value = { nombre: '' }
  showModalFacultad.value = true
}

function openEditarFacultad(fac) {
  editingFacultad.value = fac
  facultadForm.value = { nombre: fac.nombre }
  showModalFacultad.value = true
}

async function guardarFacultad() {
  if (!facultadForm.value.nombre) {
    alert('Ingresa el nombre de la facultad')
    return
  }
  saving.value = true
  try {
    if (editingFacultad.value) {
      await api.facultades.update(editingFacultad.value.id_facultad, facultadForm.value)
    } else {
      await api.facultades.create(facultadForm.value)
    }
    showModalFacultad.value = false
    loadAll()
  } catch (e) {
    alert('Error: ' + e.message)
  } finally {
    saving.value = false
  }
}

async function eliminarFacultad(fac) {
  if (!confirm(`¿Eliminar la facultad "${fac.nombre}"?`)) return
  try {
    await api.facultades.delete(fac.id_facultad)
    loadAll()
  } catch (e) {
    alert('Error: ' + e.message)
  }
}

function openNuevoDepartamento() {
  editingDepartamento.value = null
  departamentoForm.value = {
    id_facultad: '',
    nombre: ''
  }
  showModalDepartamento.value = true
}

function openEditarDepartamento(depto) {
  editingDepartamento.value = depto
  departamentoForm.value = {
    id_facultad: depto.id_facultad,
    nombre: depto.nombre
  }
  showModalDepartamento.value = true
}

async function guardarDepartamento() {
  if (!departamentoForm.value.id_facultad || !departamentoForm.value.nombre) {
    alert('Completa todos los campos')
    return
  }
  saving.value = true
  try {
    if (editingDepartamento.value) {
      await api.departamentos.update(editingDepartamento.value.id_departamento, departamentoForm.value)
    } else {
      await api.departamentos.create(departamentoForm.value)
    }
    showModalDepartamento.value = false
    loadAll()
  } catch (e) {
    alert('Error: ' + e.message)
  } finally {
    saving.value = false
  }
}

async function eliminarDepartamento(depto) {
  if (!confirm(`¿Eliminar el departamento "${depto.nombre}"?`)) return
  try {
    await api.departamentos.delete(depto.id_departamento)
    loadAll()
  } catch (e) {
    alert('Error: ' + e.message)
  }
}

function getFacultadNombre(id) {
  const f = catalogos.value.facultades.find(x => x.id_facultad === id)
  return f ? f.nombre : `Facultad #${id}`
}

function getDepartamentoNombre(id) {
  const d = catalogos.value.departamentos.find(x => x.id_departamento === id)
  return d ? d.nombre : `Depto #${id}`
}

const filteredDepartamentos = computed(() => {
  if (!escuelaForm.value.id_facultad) return catalogos.value.departamentos
  return catalogos.value.departamentos.filter(d => d.id_facultad === parseInt(escuelaForm.value.id_facultad))
})

function getPabellonNombre(id) {
  const p = pabellones.value.find(x => x.id_pabellon === id)
  return p ? p.codigo : `Pabellón #${id}`
}

function getEscuelaNombre(id) {
  const e = catalogos.value.escuelas.find(x => x.id_escuela === id)
  return e ? e.nombre : `Escuela #${id}`
}

async function toggleAulaActivo(aula) {
  const accion = aula.activo ? 'desactivar' : 'activar'
  if (!confirm(`¿${aula.activo ? 'Desactivar' : 'Activar'} el aula "${aula.codigo}"?`)) return
  try {
    await api.aulas.setActivo(aula.id_aula, !aula.activo)
    loadAll()
  } catch (e) {
    alert('Error: ' + e.message)
  }
}

function openNuevaAula() {
  editingAula.value = null
  aulaForm.value = {
    codigo: '',
    id_pabellon: '',
    tipo: 'TEORIA',
    afecto: '',
    es_compartida: false
  }
  showModalAula.value = true
}

function openEditarAula(aula) {
  editingAula.value = aula
  aulaForm.value = {
    codigo: aula.codigo,
    id_pabellon: aula.id_pabellon,
    tipo: aula.tipo,
    afecto: aula.aforo,
    es_compartida: aula.es_compartida
  }
  showModalAula.value = true
}

async function guardarAula() {
  if (!aulaForm.value.codigo || !aulaForm.value.id_pabellon || !aulaForm.value.aforo) {
    alert('Completa todos los campos requeridos')
    return
  }
  saving.value = true
  try {
    const data = {
      codigo: aulaForm.value.codigo,
      id_pabellon: parseInt(aulaForm.value.id_pabellon),
      tipo: aulaForm.value.tipo,
      afecto: parseInt(aulaForm.value.aforo),
      es_compartida: aulaForm.value.es_compartida
    }
    if (editingAula.value) {
      await api.aulas.update(editingAula.value.id_aula, data)
    } else {
      await api.aulas.create(data)
    }
    showModalAula.value = false
    loadAll()
  } catch (e) {
    alert('Error: ' + e.message)
  } finally {
    saving.value = false
  }
}

async function eliminarAula(aula) {
  if (!confirm(`¿Eliminar el aula "${aula.codigo}"? Esta accion no se puede deshacer.`)) return
  try {
    await api.aulas.delete(aula.id_aula)
    loadAll()
  } catch (e) {
    alert('Error: ' + e.message)
  }
}

// Periodo functions
function openNuevoPeriodo() {
  editingPeriodo.value = null
  periodoForm.value = {
    codigo: '',
    activo: true
  }
  showModalPeriodo.value = true
}

function openEditarPeriodo(periodo) {
  editingPeriodo.value = periodo
  periodoForm.value = {
    codigo: periodo.codigo,
    activo: periodo.activo
  }
  showModalPeriodo.value = true
}

async function guardarPeriodo() {
  if (!periodoForm.value.codigo) {
    alert('Ingresa el código del período')
    return
  }
  saving.value = true
  try {
    const data = {
      codigo: periodoForm.value.codigo,
      activo: periodoForm.value.activo
    }
    if (editingPeriodo.value) {
      await api.periodos.update(editingPeriodo.value.id_periodo, data)
    } else {
      await api.periodos.create(data)
    }
    showModalPeriodo.value = false
    loadAll()
  } catch (e) {
    alert('Error: ' + e.message)
  } finally {
    saving.value = false
  }
}

async function eliminarPeriodo(periodo) {
  if (!confirm(`¿Eliminar el período "${periodo.codigo}"?`)) return
  try {
    await api.periodos.delete(periodo.id_periodo)
    loadAll()
  } catch (e) {
    alert('Error: ' + e.message)
  }
}

// Plan functions
function openNuevoPlan() {
  editingPlan.value = null
  planForm.value = {
    id_escuela: '',
    codigo: '',
    nombre: ''
  }
  showModalPlan.value = true
}

function openEditarPlan(plan) {
  editingPlan.value = plan
  planForm.value = {
    id_escuela: plan.id_escuela,
    codigo: plan.codigo_plan,
    nombre: plan.nombre
  }
  showModalPlan.value = true
}

async function guardarPlan() {
  if (!planForm.value.id_escuela || !planForm.value.codigo || !planForm.value.nombre) {
    alert('Completa todos los campos requeridos')
    return
  }
  saving.value = true
  try {
    const data = {
      id_escuela: parseInt(planForm.value.id_escuela),
      codigo_plan: planForm.value.codigo,
      nombre: planForm.value.nombre
    }
    if (editingPlan.value) {
      await api.planesEstudio.update(editingPlan.value.id_plan, data)
    } else {
      await api.planesEstudio.create(data)
    }
    showModalPlan.value = false
    loadAll()
  } catch (e) {
    alert('Error: ' + e.message)
  } finally {
    saving.value = false
  }
}

async function eliminarPlan(plan) {
  if (!confirm(`¿Eliminar el plan "${plan.nombre}"?`)) return
  try {
    await api.planesEstudio.delete(plan.id_plan)
    loadAll()
  } catch (e) {
    alert('Error: ' + e.message)
  }
}

// Escuela functions
function openNuevaEscuela() {
  editingEscuela.value = null
  escuelaForm.value = {
    id_facultad: '',
    id_departamento: '',
    nombre: ''
  }
  showModalEscuela.value = true
}

function openEditarEscuela(escuela) {
  editingEscuela.value = escuela
  escuelaForm.value = {
    id_facultad: escuela.id_facultad,
    id_departamento: escuela.id_departamento,
    nombre: escuela.nombre
  }
  showModalEscuela.value = true
}

async function guardarEscuela() {
  if (!escuelaForm.value.id_facultad || !escuelaForm.value.id_departamento || !escuelaForm.value.nombre) {
    alert('Completa todos los campos requeridos')
    return
  }
  saving.value = true
  try {
    const data = {
      id_facultad: parseInt(escuelaForm.value.id_facultad),
      id_departamento: parseInt(escuelaForm.value.id_departamento),
      nombre: escuelaForm.value.nombre
    }
    if (editingEscuela.value) {
      await api.escuelas.update(editingEscuela.value.id_escuela, data)
    } else {
      await api.escuelas.create(data)
    }
    showModalEscuela.value = false
    loadAll()
  } catch (e) {
    alert('Error: ' + e.message)
  } finally {
    saving.value = false
  }
}

async function eliminarEscuela(escuela) {
  if (!confirm(`¿Eliminar la escuela "${escuela.nombre}"?`)) return
  try {
    await api.escuelas.delete(escuela.id_escuela)
    loadAll()
  } catch (e) {
    alert('Error: ' + e.message)
  }
}

// Docente functions
function openNuevoDocente() {
  editingDocente.value = null
  docenteForm.value = {
    id_departamento: '',
    codigo_plaza: '',
    nombres: '',
    apellidos: '',
    email: ''
  }
  showModalDocente.value = true
}

function openEditarDocente(docente) {
  editingDocente.value = docente
  docenteForm.value = {
    id_departamento: docente.id_departamento,
    codigo_plaza: docente.codigo_plaza,
    nombres: docente.nombres,
    apellidos: docente.apellidos,
    email: docente.email || ''
  }
  showModalDocente.value = true
}

async function guardarDocente() {
  if (!docenteForm.value.id_departamento || !docenteForm.value.codigo_plaza || !docenteForm.value.nombres || !docenteForm.value.apellidos) {
    alert('Completa todos los campos requeridos')
    return
  }
  saving.value = true
  try {
    const data = {
      id_departamento: parseInt(docenteForm.value.id_departamento),
      codigo_plaza: docenteForm.value.codigo_plaza,
      nombres: docenteForm.value.nombres,
      apellidos: docenteForm.value.apellidos,
      email: docenteForm.value.email
    }
    if (editingDocente.value) {
      await api.docentes.update(editingDocente.value.id_docente, data)
    } else {
      await api.docentes.create(data)
    }
    showModalDocente.value = false
    loadAll()
  } catch (e) {
    alert('Error: ' + e.message)
  } finally {
    saving.value = false
  }
}

async function eliminarDocente(docente) {
  if (!confirm(`¿Eliminar el docente "${docente.nombres} ${docente.apellidos}"?`)) return
  try {
    await api.docentes.delete(docente.id_docente)
    loadAll()
  } catch (e) {
    alert('Error: ' + e.message)
  }
}

// Curso functions
function openNuevoCurso() {
  editingCurso.value = null
  cursoForm.value = {
    id_serie: '',
    codigo: '',
    nombre: '',
    creditos: '',
    horas_teoria: '',
    horas_practica: ''
  }
  showModalCurso.value = true
}

function openEditarCurso(curso) {
  editingCurso.value = curso
  cursoForm.value = {
    id_serie: curso.id_serie,
    codigo: curso.codigo,
    nombre: curso.nombre,
    creditos: curso.creditos,
    horas_teoria: curso.horas_teoria,
    horas_practica: curso.horas_practica
  }
  showModalCurso.value = true
}

async function guardarCurso() {
  if (!cursoForm.value.id_serie || !cursoForm.value.codigo || !cursoForm.value.nombre || !cursoForm.value.creditos) {
    alert('Completa todos los campos requeridos')
    return
  }
  saving.value = true
  try {
    const data = {
      id_serie: parseInt(cursoForm.value.id_serie),
      codigo: cursoForm.value.codigo,
      nombre: cursoForm.value.nombre,
      creditos: parseInt(cursoForm.value.creditos),
      horas_teoria: parseInt(cursoForm.value.horas_teoria) || 0,
      horas_practica: parseInt(cursoForm.value.horas_practica) || 0
    }
    if (editingCurso.value) {
      await api.cursos.update(editingCurso.value.id_curso, data)
    } else {
      await api.cursos.create(data)
    }
    showModalCurso.value = false
    loadAll()
  } catch (e) {
    alert('Error: ' + e.message)
  } finally {
    saving.value = false
  }
}

async function eliminarCurso(curso) {
  if (!confirm(`¿Eliminar el curso "${curso.nombre}"?`)) return
  try {
    await api.cursos.delete(curso.id_curso)
    loadAll()
  } catch (e) {
    alert('Error: ' + e.message)
  }
}

onMounted(loadAll)
</script>

<style scoped>
.catalogos-page {
  min-height: 100vh;
  background: var(--bg-primary);
}

.page-header {
  background: linear-gradient(135deg, #8B0000 0%, #A52A2A 100%);
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
  box-shadow: 0 4px 12px var(--shadow);
}

.btn-secondary {
  background: rgba(255,255,255,0.2);
  color: white;
  border: 2px solid rgba(255,255,255,0.3);
}

.btn-secondary:hover {
  background: rgba(255,255,255,0.3);
  transform: translateY(-2px);
}

.btn-sm {
  padding: 6px 12px;
  font-size: 0.8rem;
}

.btn-icon-sm {
  background: var(--hover-bg);
  border: none;
  padding: 8px;
  border-radius: 8px;
  cursor: pointer;
  color: var(--text-secondary);
  transition: all 0.2s;
}

.btn-icon-sm:hover {
  background: var(--border-color);
  color: var(--text-primary);
}

.page-content {
  padding: 24px;
}

.tabs-container {
  margin-bottom: 24px;
}

.tabs {
  display: flex;
  gap: 8px;
  background: var(--bg-secondary);
  padding: 8px;
  border-radius: 16px;
  box-shadow: 0 4px 16px var(--shadow);
  overflow-x: auto;
}

.tab {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 12px 20px;
  border: none;
  background: transparent;
  border-radius: 10px;
  font-size: 0.9rem;
  font-weight: 600;
  color: var(--text-secondary);
  cursor: pointer;
  transition: all 0.3s;
  white-space: nowrap;
}

.tab:hover {
  background: var(--hover-bg);
  color: #8B0000;
}

.tab.active {
  background: linear-gradient(135deg, #8B0000 0%, #A52A2A 100%);
  color: white;
  box-shadow: 0 4px 12px rgba(139, 0, 0, 0.3);
}

.tab-count {
  background: rgba(0,0,0,0.1);
  padding: 2px 8px;
  border-radius: 10px;
  font-size: 0.75rem;
}

.tab.active .tab-count {
  background: rgba(255,255,255,0.3);
}

.loading-overlay {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 80px 20px;
  color: var(--text-secondary);
}

.spinner-lg {
  width: 48px;
  height: 48px;
  border: 4px solid rgba(17, 153, 142, 0.2);
  border-top-color: #11998e;
  border-radius: 50%;
  animation: spin 1s linear infinite;
}

@keyframes spin {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}

.error-banner {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 16px 20px;
  background: var(--bg-secondary);
  border-radius: 16px;
  box-shadow: 0 4px 20px rgba(239, 68, 68, 0.15);
  border-left: 4px solid #ef4444;
  color: #dc2626;
  margin-bottom: 24px;
}

.tab-content {
  animation: fadeIn 0.3s;
}

@keyframes fadeIn {
  from { opacity: 0; transform: translateY(10px); }
  to { opacity: 1; transform: translateY(0); }
}

.table-container {
  background: var(--bg-secondary);
  border-radius: 20px;
  box-shadow: 0 8px 32px var(--shadow);
  overflow: hidden;
  border: 1px solid var(--border-color);
}

.data-table {
  width: 100%;
  border-collapse: collapse;
}

.data-table th {
  text-align: left;
  padding: 16px 20px;
  font-size: 0.75rem;
  font-weight: 700;
  text-transform: uppercase;
  letter-spacing: 0.5px;
  color: var(--text-secondary);
  background: var(--bg-tertiary);
  border-bottom: 2px solid var(--border-color);
}

.data-table td {
  padding: 14px 20px;
  border-bottom: 1px solid var(--border-color);
  font-size: 0.9rem;
  color: var(--text-primary);
}

.data-table tr:last-child td {
  border-bottom: none;
}

.data-table tr:hover td {
  background: var(--bg-tertiary);
}

.id-cell {
  font-weight: 600;
  color: var(--text-secondary);
  width: 60px;
}

.name-cell {
  font-weight: 500;
}

.name-content {
  display: flex;
  align-items: center;
  gap: 10px;
}

.entity-icon {
  width: 32px;
  height: 32px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
}

.facultad-icon { background: linear-gradient(135deg, #667eea, #764ba2); }
.depto-icon { background: linear-gradient(135deg, #11998e, #38ef7d); }
.escuela-icon { background: linear-gradient(135deg, #f093fb, #f5576c); }

.actions-cell {
  width: 100px;
}

.fw-600 {
  font-weight: 600;
}

.facultad-badge {
  background: linear-gradient(135deg, #667eea20, #764ba220);
  color: #667eea;
  padding: 4px 10px;
  border-radius: 8px;
  font-size: 0.8rem;
  font-weight: 500;
}

.depto-badge {
  background: linear-gradient(135deg, #11998e20, #38ef7d20);
  color: #11998e;
  padding: 4px 10px;
  border-radius: 8px;
  font-size: 0.8rem;
  font-weight: 500;
}

.escuela-badge {
  background: linear-gradient(135deg, #f093fb20, #f5576c20);
  color: #f5576c;
  padding: 4px 10px;
  border-radius: 8px;
  font-size: 0.8rem;
  font-weight: 500;
}

.codigo-badge {
  background: var(--bg-tertiary);
  color: var(--text-primary);
  padding: 4px 10px;
  border-radius: 8px;
  font-size: 0.8rem;
  font-weight: 600;
  font-family: monospace;
}

.curso-codigo {
  background: linear-gradient(135deg, #f093fb20, #f5576c20);
  color: #f5576c;
}

.creditos-badge {
  background: linear-gradient(135deg, #fbbf24, #f59e0b);
  color: white;
  padding: 4px 10px;
  border-radius: 8px;
  font-size: 0.8rem;
  font-weight: 700;
}

.email-cell {
  color: var(--text-secondary);
  font-style: italic;
}

.btn-danger {
  background: #ef4444 !important;
  color: white !important;
}

.btn-danger:hover {
  background: #dc2626 !important;
}

.btn-success {
  background: #22c55e !important;
  color: white !important;
}

.btn-success:hover {
  background: #16a34a !important;
}

.btn-warning {
  background: #f59e0b !important;
  color: white !important;
}

.btn-warning:hover {
  background: #d97706 !important;
}

.content-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 20px;
}

.modal-overlay {
  position: fixed;
  inset: 0;
  background: rgba(0,0,0,0.4);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
  backdrop-filter: blur(8px);
}

.modal {
  background: var(--bg-secondary);
  border-radius: 24px;
  width: 100%;
  max-width: 560px;
  max-height: 90vh;
  overflow: hidden;
  box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.25);
  animation: modalIn 0.2s ease-out;
}

@keyframes modalIn {
  from { opacity: 0; transform: scale(0.95); }
  to { opacity: 1; transform: scale(1); }
}

.modal-md {
  max-width: 480px;
}

.modal-header {
  padding: 24px;
  border-bottom: 1px solid var(--border-color);
  display: flex;
  justify-content: space-between;
  align-items: center;
  background: var(--bg-tertiary);
}

.modal-header h2 {
  display: flex;
  align-items: center;
  gap: 12px;
  margin: 0;
  font-size: 1.25rem;
  color: var(--text-primary);
}

.modal-body {
  padding: 24px;
  max-height: calc(90vh - 180px);
  overflow-y: auto;
}

.modal-footer {
  padding: 20px 24px;
  border-top: 1px solid var(--border-color);
  display: flex;
  justify-content: flex-end;
  gap: 12px;
  background: var(--bg-tertiary);
}

.form-group {
  margin-bottom: 16px;
}

.form-row {
  display: flex;
  gap: 12px;
}

.form-row .form-group {
  flex: 1;
}

.form-label {
  display: block;
  font-size: 0.875rem;
  font-weight: 600;
  color: var(--text-primary);
  margin-bottom: 6px;
}

.form-input {
  width: 100%;
  padding: 12px 16px;
  border: 2px solid var(--border-color);
  border-radius: 12px;
  font-size: 0.95rem;
  transition: all 0.3s;
  box-sizing: border-box;
  background: var(--bg-secondary);
  color: var(--text-primary);
}

.form-input:focus {
  outline: none;
  border-color: #11998e;
  box-shadow: 0 0 0 4px rgba(17, 153, 142, 0.1);
}

@media (max-width: 1024px) {
  .tabs {
    flex-wrap: wrap;
  }
  .table-container {
    overflow-x: auto;
  }
  .data-table {
    min-width: 800px;
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
}

.tipo-badge {
  padding: 4px 10px;
  border-radius: 8px;
  font-size: 0.8rem;
  font-weight: 600;
  text-transform: uppercase;
}

.tipo-teoria {
  background: linear-gradient(135deg, #3b82f620, #8b5cf620);
  color: #3b82f6;
}

.tipo-practica {
  background: linear-gradient(135deg, #f59e0b20, #f9731620);
  color: #f59e0b;
}

.tipo-laboratorio {
  background: linear-gradient(135deg, #8b5cf620, #a855f720);
  color: #8b5cf6;
}

.badge-compartida {
  background: linear-gradient(135deg, #f59e0b20, #f9731620);
  color: #f59e0b;
  padding: 4px 10px;
  border-radius: 8px;
  font-size: 0.8rem;
  font-weight: 500;
}

.badge-no {
  background: var(--bg-tertiary);
  color: var(--text-secondary);
  padding: 4px 10px;
  border-radius: 8px;
  font-size: 0.8rem;
}

.badge-activo {
  background: linear-gradient(135deg, #22c55e20, #16a34a20);
  color: #22c55e;
  padding: 4px 10px;
  border-radius: 8px;
  font-size: 0.8rem;
  font-weight: 600;
}

.badge-inactivo {
  background: linear-gradient(135deg, #ef444420, #dc262620);
  color: #ef4444;
  padding: 4px 10px;
  border-radius: 8px;
  font-size: 0.8rem;
  font-weight: 600;
}

.row-inactive {
  opacity: 0.6;
  background: var(--bg-tertiary);
}

.row-inactive:hover td {
  background: var(--bg-tertiary) !important;
}
</style>
