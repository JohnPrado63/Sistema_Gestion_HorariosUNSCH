-- SISTEMA DE GESTIÓN DE HORARIOS UNSCH - ESQUEMA DE BASE DE DATOS (PostgreSQL)
-- =============================================================================

-- Habilitar extensión para UUIDs si se requiere más adelante
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- -----------------------------------------------------------------------------
-- 1. TIPOS ENUMERADOS (ENUMS)
-- -----------------------------------------------------------------------------

CREATE TYPE tipo_aula_enum AS ENUM ('TEORIA', 'PRACTICA', 'COMPARTIDA');
CREATE TYPE estado_carga_enum AS ENUM ('BORRADOR', 'AUTORIZADO');
CREATE TYPE tipo_componente_enum AS ENUM ('TEORIA', 'PRACTICA');
CREATE TYPE estado_horario_enum AS ENUM ('BORRADOR', 'PRELIMINAR', 'EN_REAJUSTE', 'OFICIAL', 'REAJUSTADO');
CREATE TYPE rol_usuario_enum AS ENUM (
    'ADMIN_TI',
    'JEFE_DEPTO',
    'DGA',
    'DIRECTOR_ESCUELA',
    'COORDINADOR',
    'DOCENTE',
    'ESTUDIANTE'
);
-- -----------------------------------------------------------------------------
-- 2. DOMINIO ESTRUCTURA INSTITUCIONAL Y ACADÉMICA
-- -----------------------------------------------------------------------------

CREATE TABLE facultad (
    id_facultad SERIAL PRIMARY KEY,
    nombre VARCHAR(150) NOT NULL UNIQUE
);

CREATE TABLE departamento_academico (
    id_departamento SERIAL PRIMARY KEY,
    id_facultad INT NOT NULL REFERENCES facultad(id_facultad) ON DELETE RESTRICT,
    nombre VARCHAR(150) NOT NULL
);

CREATE TABLE escuela_profesional (
    id_escuela SERIAL PRIMARY KEY,
    id_facultad INT NOT NULL REFERENCES facultad(id_facultad) ON DELETE RESTRICT,
    id_departamento INT NOT NULL REFERENCES departamento_academico(id_departamento) ON DELETE RESTRICT,
    nombre VARCHAR(150) NOT NULL
);

CREATE TABLE plan_estudio (
    id_plan SERIAL PRIMARY KEY,
    id_escuela INT NOT NULL REFERENCES escuela_profesional(id_escuela) ON DELETE RESTRICT,
    codigo_plan VARCHAR(20) NOT NULL, -- Ej: '2024'
    nombre VARCHAR(150) NOT NULL,
    CONSTRAINT uk_plan_escuela UNIQUE (id_escuela, codigo_plan)
);

CREATE TABLE serie (
    id_serie SERIAL PRIMARY KEY,
    id_plan INT NOT NULL REFERENCES plan_estudio(id_plan) ON DELETE CASCADE,
    numero_ciclo INT NOT NULL CHECK (numero_ciclo >= 100 AND numero_ciclo <= 1000), -- Ej: 100, 200...
    CONSTRAINT uk_serie_plan UNIQUE (id_plan, numero_ciclo)
);

CREATE TABLE curso (
    id_curso SERIAL PRIMARY KEY,
    id_serie INT NOT NULL REFERENCES serie(id_serie) ON DELETE RESTRICT,
    codigo VARCHAR(20) NOT NULL UNIQUE,
    nombre VARCHAR(150) NOT NULL,
    creditos INT NOT NULL CHECK (creditos > 0),
    horas_teoria INT NOT NULL DEFAULT 0 CHECK (horas_teoria >= 0),
    horas_practica INT NOT NULL DEFAULT 0 CHECK (horas_practica >= 0),
    CONSTRAINT chk_horas_minimas CHECK (horas_teoria + horas_practica > 0)
);

CREATE TABLE docente (
    id_docente SERIAL PRIMARY KEY,
    id_departamento INT NOT NULL REFERENCES departamento_academico(id_departamento) ON DELETE RESTRICT,
    codigo_plaza VARCHAR(30) NOT NULL UNIQUE, -- DNI o Código de Plaza
    nombres VARCHAR(100) NOT NULL,
    apellidos VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE
);

CREATE TABLE periodo_academico (
    id_periodo SERIAL PRIMARY KEY,
    codigo VARCHAR(15) NOT NULL UNIQUE, -- Ej: '2026-I'
    activo BOOLEAN NOT NULL DEFAULT FALSE
);
CREATE TABLE sesion_departamento (
    id_sesion SERIAL PRIMARY KEY,
    id_departamento INT NOT NULL REFERENCES departamento_academico(id_departamento) ON DELETE CASCADE,
    id_periodo INT NOT NULL REFERENCES periodo_academico(id_periodo) ON DELETE CASCADE,
    dia_semana INT NOT NULL CHECK (dia_semana BETWEEN 1 AND 6), -- 1: Lunes, 6: Sábado
    hora_inicio TIME NOT NULL,
    hora_fin TIME NOT NULL,
    CONSTRAINT chk_rango_horas_sesion CHECK (hora_fin > hora_inicio)
);

-- -----------------------------------------------------------------------------
-- 3. DOMINIO INFRAESTRUCTURA
-- -----------------------------------------------------------------------------

CREATE TABLE local (
    id_local SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL UNIQUE
);

CREATE TABLE pabellon (
    id_pabellon SERIAL PRIMARY KEY,
    id_local INT NOT NULL REFERENCES local(id_local) ON DELETE RESTRICT,
    codigo VARCHAR(20) NOT NULL UNIQUE,
    nombre VARCHAR(100) NOT NULL
);

CREATE TABLE matriz_distancia (
    id_pabellon_origen INT NOT NULL REFERENCES pabellon(id_pabellon) ON DELETE CASCADE,
    id_pabellon_destino INT NOT NULL REFERENCES pabellon(id_pabellon) ON DELETE CASCADE,
    tiempo_minutos INT NOT NULL CHECK (tiempo_minutos >= 0),
    PRIMARY KEY (id_pabellon_origen, id_pabellon_destino),
    CONSTRAINT chk_origen_destino_distinto CHECK (id_pabellon_origen <> id_pabellon_destino)
);

CREATE TABLE aula (
    id_aula SERIAL PRIMARY KEY,
    id_pabellon INT NOT NULL REFERENCES pabellon(id_pabellon) ON DELETE RESTRICT,
    id_escuela INT REFERENCES escuela_profesional(id_escuela) ON DELETE SET NULL, -- NULL = Uso Común
    codigo VARCHAR(20) NOT NULL UNIQUE,
    tipo tipo_aula_enum NOT NULL DEFAULT 'TEORIA',
    aforo INT NOT NULL CHECK (aforo > 0),
    es_compartida BOOLEAN NOT NULL DEFAULT FALSE
);

-- -----------------------------------------------------------------------------
-- 4. DOMINIO PRE-SISTEMA Y CARGA ACADÉMICA
-- -----------------------------------------------------------------------------

CREATE TABLE carga_academica (
    id_carga SERIAL PRIMARY KEY,
    id_curso INT NOT NULL REFERENCES curso(id_curso) ON DELETE RESTRICT,
    id_periodo INT NOT NULL REFERENCES periodo_academico(id_periodo) ON DELETE RESTRICT,
    id_escuela INT NOT NULL REFERENCES escuela_profesional(id_escuela) ON DELETE RESTRICT,
    estado estado_carga_enum NOT NULL DEFAULT 'BORRADOR',
    fecha_aprobacion TIMESTAMP,
    CONSTRAINT uk_carga_curso_periodo_escuela UNIQUE (id_curso, id_periodo, id_escuela)
);

CREATE TABLE grupo (
    id_grupo SERIAL PRIMARY KEY,
    id_carga INT NOT NULL REFERENCES carga_academica(id_carga) ON DELETE CASCADE,
    id_docente INT REFERENCES docente(id_docente) ON DELETE SET NULL,
    id_grupo_teoria_ref INT REFERENCES grupo(id_grupo) ON DELETE SET NULL, -- Relación 1:1 Práctica -> Teoría
    codigo_grupo VARCHAR(10) NOT NULL, -- Ej: 'Grupo A'
    tipo_componente tipo_componente_enum NOT NULL,
    es_nueva_necesidad BOOLEAN NOT NULL DEFAULT FALSE,
    matriculados_proyectados INT NOT NULL DEFAULT 0 CHECK (matriculados_proyectados >= 0),
    matriculados_reales INT NOT NULL DEFAULT 0 CHECK (matriculados_reales >= 0)
);
-- -----------------------------------------------------------------------------
-- 5. DOMINIO MOTOR DE HORARIOS
-- -----------------------------------------------------------------------------

CREATE TABLE horario (
    id_horario SERIAL PRIMARY KEY,
    id_escuela INT NOT NULL REFERENCES escuela_profesional(id_escuela) ON DELETE RESTRICT,
    id_periodo INT NOT NULL REFERENCES periodo_academico(id_periodo) ON DELETE RESTRICT,
    estado estado_horario_enum NOT NULL DEFAULT 'BORRADOR',
    version_reajuste INT NOT NULL DEFAULT 0 CHECK (version_reajuste >= 0),
    fecha_actualizacion TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT uk_horario_escuela_periodo UNIQUE (id_escuela, id_periodo)
);

CREATE TABLE bloque_horario (
    id_bloque SERIAL PRIMARY KEY,
    id_horario INT NOT NULL REFERENCES horario(id_horario) ON DELETE CASCADE,
    id_grupo INT NOT NULL REFERENCES grupo(id_grupo) ON DELETE RESTRICT,
    id_aula INT NOT NULL REFERENCES aula(id_aula) ON DELETE RESTRICT,
    id_docente INT REFERENCES docente(id_docente) ON DELETE SET NULL, -- Asignación por sesión
    dia_semana INT NOT NULL CHECK (dia_semana BETWEEN 1 AND 6), -- 1: Lunes, 6: Sábado
    slot_inicio INT NOT NULL CHECK (slot_inicio BETWEEN 1 AND 14), -- Slots rígidos de 60 min (1 = 07:00-08:00)
    slot_fin INT NOT NULL CHECK (slot_fin BETWEEN 1 AND 14),
    CONSTRAINT chk_slot_valido CHECK (slot_fin >= slot_inicio)
);

-- -----------------------------------------------------------------------------
-- 6. DOMINIO USUARIOS Y CONTROL (AUDITORÍA)
-- -----------------------------------------------------------------------------

CREATE TABLE usuario (
    id_usuario SERIAL PRIMARY KEY,
    nombre VARCHAR(150) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    rol rol_usuario_enum NOT NULL
);

CREATE TABLE bitacora_auditoria (
    id_log SERIAL PRIMARY KEY,
    id_horario INT NOT NULL REFERENCES horario(id_horario) ON DELETE CASCADE,
    id_usuario INT NOT NULL REFERENCES usuario(id_usuario) ON DELETE RESTRICT,
    accion VARCHAR(100) NOT NULL,
    motivo_justificacion TEXT NOT NULL,
    version_resultante INT NOT NULL,
    fecha_hora TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- =============================================================================
-- 7. ÍNDICES DE RENDIMIENTO (OPTIMIZACIÓN DEL MOTOR DE VALIDACIONES)
-- =============================================================================

-- RV-01: Búsqueda rápida de cruce de docente
CREATE INDEX idx_bloque_docente_dia_slot 
ON bloque_horario (id_docente, dia_semana, slot_inicio, slot_fin) 
WHERE id_docente IS NOT NULL;

-- RV-02 & RV-07: Búsqueda rápida de ocupación de aula
CREATE INDEX idx_bloque_aula_dia_slot 
ON bloque_horario (id_aula, dia_semana, slot_inicio, slot_fin);

-- Consulta acelerada de bloques por horario
CREATE INDEX idx_bloque_horario_id 
ON bloque_horario (id_horario);

-- Búsqueda de grupos por carga académica
CREATE INDEX idx_grupo_carga 
ON grupo (id_carga);

-- RV-03: Búsqueda de sesiones de departamento por periodo
CREATE INDEX idx_sesion_depto_periodo 
ON sesion_departamento (id_departamento, id_periodo);

-- RV-06: Búsqueda de asignaturas por serie para cruces en la misma serie
CREATE INDEX idx_curso_serie 
ON curso (id_serie);