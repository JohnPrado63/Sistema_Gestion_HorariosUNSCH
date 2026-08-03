# Sistema_Gestion_HorariosUNSCH

Sistema de Gestión de Horarios y Carga Académica de la Universidad Nacional de San Cristóbal de Huamanga (UNSCH).

## Requisitos previos

- Go 1.21+
- Node.js 18+
- Docker y Docker Compose

## Configuración

1. Copiar el archivo de entorno de ejemplo:
   ```powershell
   cp .env.example .env
   ```

2. Editar `.env` y configurar las contraseñas:
   ```
   DB_PASSWORD=tu_contraseña_postgres
   REDIS_PASSWORD=tu_contraseña_redis
   APP_PORT=8080
   ```

**Nota**: El backend usa puerto `8080` y el frontend puerto `8081`. Asegúrate de que Docker tenga los puertos `5433` y `6379` disponibles.

## Ejecución

- Levantar la base de datos local con Docker Compose:
  ```powershell
  docker compose up -d
  ```

- Iniciar el backend:
  ```powershell
  cd backend
  go mod tidy
  go run ./cmd/api
  ```

- Iniciar el frontend:
  ```powershell
  cd frontend
  npm install
  npm run dev
  ```

- Abrir la interfaz de usuario en el navegador en:
  ```text
  http://localhost:8081/app/
  ```

## Tecnologías

- **Backend**: Go, Gin, PostgreSQL, Redis
- **Frontend**: Vue 3, Vite

## Requisitos Funcionales

### RF-01. Configuración
Gestión de facultades, departamentos académicos, escuelas profesionales y series por plan (Series 100, 200, 300…).

### RF-01.1. Configuración – Registro de Sesión de Departamento
El Jefe de Departamento o el Administrador define la franja horaria semanal reservada para las sesiones de su departamento académico. Bloquea automáticamente la agenda de todos sus docentes en todas las escuelas.

### RF-02. Infraestructura
Gestión de locales, pabellones y aulas (tipos Teoría y Práctica). Registra aforo, código y pertenencia (Aula de Escuela vs. Aula Compartida/Uso Común). *Fuera de alcance: laboratorios especiales.*

### RF-03. Infraestructura – Matriz de Distancias
Configuración de la matriz de distancias/tiempos estimados de desplazamiento entre pabellones y locales.

### RF-04. Integración con SIIGE
Importación unidireccional desde SIIGE de planes de estudio, asignaturas, créditos, horas lectivas (teoría/práctica) y estudiantes matriculados. SIIGE es la fuente única de verdad.

### RF-04.1. Periodos
Gestión de semestres académicos (ej. 2026-I, 2026-II). Función de Clonación de Plantilla: copia días/horas/aulas de un semestre homólogo anterior hacia un nuevo horario en estado BORRADOR.

### RF-05. Carga Académica (Pre-sistema)
Módulo de asignación docente-curso-grupo. Permite registrar entradas Nueva Necesidad cuando la plaza está vacante.

### RF-05.1. Disponibilidad – Bloqueo Distribuido Inter-Escuela
Vector centralizado (Redis) con la disponibilidad de cada docente en tiempo real, para impedir choques entre escuelas distintas.

### RF-05.2. Proyección de Grupos
Sugiere o exige la apertura de múltiples grupos (Grupo A, Grupo B…) cuando la matrícula proyectada supera el aforo promedio de aula.

### RF-05.3. Sustitución de Nueva Necesidad
Cuando se contrata al docente titular, la DGA asocia su identidad real. El sistema valida disponibilidad contra la agenda global del docente.

### RF-06. Aprobación de Carga Académica
La DGA aprueba la Carga Académica (Borrador → Autorizado), incluyendo la validación del tope de 16 horas lectivas semanales por docente (RV-05).

### RF-07. Formulación
Permite al Director de Escuela diagramar la grilla horaria asignando día, hora y aula para cada grupo de Teoría y su correspondiente grupo de Práctica (relación 1:1).

### RF-08. Reserva de Aulas
Selección y reserva de aulas o pabellones de uso compartido para escuelas sin infraestructura propia suficiente.

### RF-09. Ciclo de Vida del Horario
Administración de la máquina de estados: BORRADOR → PRELIMINAR → EN_REAJUSTE → OFICIAL → REAJUSTADO_vX.

### RF-10a. Reajuste Masivo
Reajuste Post-Matrícula: permite al Director ajustar masivamente aulas y apertura/cierre de grupos tras la consolidación de inscritos en SIIGE (PRELIMINAR → EN_REAJUSTE → OFICIAL).

### RF-10b. Reajuste Individual
Reajuste Extemporáneo: permite modificar la asignación de un curso/docente sobre un horario OFICIAL o REAJUSTADO_vX vigente, generando una nueva versión con justificación obligatoria.

### RF-11. Consultas
Consulta e impresión de grilla semanal interactiva, filtrable por Escuela, Serie/Ciclo, Docente, Aula y Local/Pabellón.

### RF-12. Exportación
Exportación del horario en formatos PDF y Excel, por escuela, docente o aula.

### RF-13. Auditoría
Registro de bitácora (logs) de todas las creaciones, ediciones, reajustes y aprobaciones, indicando usuario, rol, fecha, hora y motivo.

## Reglas de Negocio

### RV-01. Cruce de Docente
Un docente no puede figurar en dos escuelas/grupos en el mismo día y rango horario.
**Acción**: Bloqueo estricto.

### RV-02. Cruce de Aula
Un aula no puede tener más de una clase asignada en el mismo horario.
**Acción**: Bloqueo estricto.

### RV-03. Sesión de Departamento
Ningún docente puede dictar clases durante la franja semanal reservada para la sesión de su departamento.
**Acción**: Bloqueo estricto.

### RV-04a. Traslado Sin Tiempo
Dos clases consecutivas de un mismo docente en pabellones distantes, sin intervalo libre intermedio.
**Acción**: Bloqueo estricto.

### RV-04b. Traslado Ajustado
Intervalo libre entre pabellones distantes, pero menor al tiempo estimado en la matriz de distancias.
**Acción**: Advertencia con justificación.

### RV-05. Topes de Carga Lectiva (16h)
La suma de horas lectivas asignadas a un docente supera las 16 horas semanales.
**Acción**: Advertencia con confirmación de la DGA. Se valida exclusivamente al aprobar la Carga Académica.

### RV-06. Cruce en Misma Serie
Dos asignaturas obligatorias de la misma Serie/Ciclo coinciden en el mismo horario.
**Acción**: Alerta informativa al Director; no bloquea.

### RV-07. Aulas Compartidas
Intento de asignar un aula/local compartido que ya fue reservado por otra escuela.
**Acción**: Bloqueo estricto si está ocupada.

### RV-08. Control de Auditoría
Toda modificación sobre un horario OFICIAL o REAJUSTADO_vX exige registrar la justificación del cambio.
**Acción**: Registro obligatorio en la bitácora.

### RV-09. Reajuste por Aforo
La matrícula excede la capacidad del aula asignada.
**Acción**: En PRELIMINAR, alerta leve; en EN_REAJUSTE, alerta de reajuste que sugiere cambiar de aula o abrir Grupo B.
