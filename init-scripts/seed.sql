-- =============================================================================
-- DATOS DE PRUEBA E INICIALIZACIÓN (SEED DATA)
-- =============================================================================

-- Password para todos los usuarios: admin123
-- Hash bcrypt generado: $2a$10$nAAC17PkVLbcigHr47UWe.fogi7U9SshX5VyZnhBVn3bgca/aeeK.

-- 1. Periodo Académico
INSERT INTO periodo_academico (codigo, activo) VALUES ('2026-I', true);
INSERT INTO periodo_academico (codigo, activo) VALUES ('2026-II', false);

-- 2. Estructura Institucional
INSERT INTO facultad (nombre) VALUES ('Facultad de Ingeniería de Minas, Geología y Civil');

INSERT INTO departamento_academico (id_facultad, nombre)
VALUES (1, 'Departamento Académico de Ingeniería de Sistemas');

INSERT INTO escuela_profesional (id_facultad, id_departamento, nombre)
VALUES (1, 1, 'Escuela Profesional de Ingeniería de Sistemas');

-- 3. Infraestructura
INSERT INTO local (nombre) VALUES ('Ciudad Universitaria');

INSERT INTO pabellon (id_local, codigo, nombre) VALUES
(1, 'PAB-IS', 'Pabellón de Ingeniería de Sistemas'),
(1, 'PAB-GENERAL', 'Pabellón de General de Aulas');

INSERT INTO matriz_distancia (id_pabellon_origen, id_pabellon_destino, tiempo_minutos) VALUES
(1, 2, 5),
(2, 1, 5);

INSERT INTO aula (id_pabellon, id_escuela, codigo, tipo, aforo, es_compartida) VALUES
(1, 1, 'AULA-101', 'TEORIA', 40, false),
(1, 1, 'AULA-102', 'TEORIA', 45, false),
(1, 1, 'LAB-101', 'PRACTICA', 30, false),
(2, NULL, 'AULA-AUDITORIO', 'COMPARTIDA', 100, true);

-- 4. Usuarios del Sistema
INSERT INTO usuario (nombre, email, password_hash, rol) VALUES
-- Administradores
('Administrador TI', 'admin@unsch.edu.pe', '$2a$10$nAAC17PkVLbcigHr47UWe.fogi7U9SshX5VyZnhBVn3bgca/aeeK.', 'ADMIN_TI'),
('Director General Académico', 'dga@unsch.edu.pe', '$2a$10$nAAC17PkVLbcigHr47UWe.fogi7U9SshX5VyZnhBVn3bgca/aeeK.', 'DGA'),

-- Directores y Jefes
('Director de Escuela IS', 'director.sistemas@unsch.edu.pe', '$2a$10$nAAC17PkVLbcigHr47UWe.fogi7U9SshX5VyZnhBVn3bgca/aeeK.', 'DIRECTOR_ESCUELA'),
('Jefe Dept. Ingeniería Sistemas', 'jefe.depto@unsch.edu.pe', '$2a$10$nAAC17PkVLbcigHr47UWe.fogi7U9SshX5VyZnhBVn3bgca/aeeK.', 'JEFE_DEPTO'),

-- Coordinadores
('Coordinador de Turno', 'coordinador@unsch.edu.pe', '$2a$10$nAAC17PkVLbcigHr47UWe.fogi7U9SshX5VyZnhBVn3bgca/aeeK.', 'COORDINADOR');

-- 5. DOCENTES (tabla separada de usuarios)
INSERT INTO docente (id_departamento, codigo_plaza, nombres, apellidos, email) VALUES
(1, 'DNI-001', 'María', 'López Rodríguez', 'maria.lopez@unsch.edu.pe'),
(1, 'DNI-002', 'Carlos', 'González Villa', 'carlos.gonzalez@unsch.edu.pe'),
(1, 'DNI-003', 'Juan', 'Pérez Torres', 'juan.perez@unsch.edu.pe'),
(1, 'DNI-004', 'Ana', 'García Ruiz', 'ana.garcia@unsch.edu.pe'),
(1, 'DNI-005', 'Luis', 'Ramírez Santos', 'luis.ramirez@unsch.edu.pe'),
(1, 'DNI-006', 'Elena', 'Castro Morales', 'elena.castro@unsch.edu.pe');

-- 6. Plan de Estudios, Series y Cursos (para Ingeniería de Sistemas)
INSERT INTO plan_estudio (id_escuela, codigo_plan, nombre) VALUES
(1, '2024', 'Plan de Estudios 2024 - IS');

-- Series (ciclos) - Serie 100, 200, 300...
INSERT INTO serie (id_plan, numero_ciclo) VALUES
(1, 100),  -- I ciclo
(1, 200),  -- II ciclo
(1, 300),  -- III ciclo
(1, 400),  -- IV ciclo
(1, 500);  -- V ciclo

-- Cursos del I-II ciclo (Serie 100 y 200)
INSERT INTO curso (id_serie, codigo, nombre, creditos, horas_teoria, horas_practica) VALUES
-- Serie 100 (I ciclo)
(100, 'CS101', 'Introducción a la Programación', 4, 2, 2),
(100, 'MAT101', 'Matemática Básica', 3, 2, 1),
(100, 'FIS101', 'Física General', 3, 2, 1),
(100, 'LEL101', 'Lenguaje y Comunicación', 2, 1, 1),

-- Serie 200 (II ciclo)
(200, 'CS201', 'Programación Orientada a Objetos', 4, 2, 2),
(200, 'MAT201', 'Cálculo I', 4, 3, 1),
(200, 'EDD201', 'Estructuras de Datos', 4, 2, 2),
(200, 'FIS201', 'Física Computacional', 3, 2, 1);

-- 7. CARGAS ACADÉMICAS DE EJEMPLO (Periodo 2026-I)
-- Estas están en estado BORRADOR para poder probar
INSERT INTO carga_academica (id_curso, id_periodo, id_escuela, estado) VALUES
-- CS101: 2 grupos (A y B teoría)
(1, 1, 1, 'BORRADOR'),
-- CS201: 1 grupo
(5, 1, 1, 'BORRADOR'),
-- MAT101: 1 grupo
(2, 1, 1, 'BORRADOR');

-- 8. GRUPOS DE EJEMPLO
-- CS101 Grupo A - Teoría (docente: María López)
INSERT INTO grupo (id_carga, id_docente, codigo_grupo, tipo_componente, es_nueva_necesidad, matriculados_proyectados)
VALUES (1, 1, 'Grupo A', 'TEORIA', false, 35);

-- CS101 Grupo B - Teoría (docente: Carlos González)
INSERT INTO grupo (id_carga, id_docente, codigo_grupo, tipo_componente, es_nueva_necesidad, matriculados_proyectados)
VALUES (1, 2, 'Grupo B', 'TEORIA', false, 38);

-- CS201 Grupo A - Teoría (docente: Juan Pérez)
INSERT INTO grupo (id_carga, id_docente, codigo_grupo, tipo_componente, es_nueva_necesidad, matriculados_proyectados)
VALUES (2, 3, 'Grupo A', 'TEORIA', false, 30);

-- CS201 Grupo A - Práctica (docente: Juan Pérez,linked to teoria ref)
INSERT INTO grupo (id_carga, id_docente, codigo_grupo, tipo_componente, es_nueva_necesidad, matriculados_proyectados, id_grupo_teoria_ref)
VALUES (2, 3, 'Grupo A', 'PRACTICA', false, 30, 3);

-- MAT101 Grupo A - Teoría (docente: Ana García) - 3 horas
INSERT INTO grupo (id_carga, id_docente, codigo_grupo, tipo_componente, es_nueva_necesidad, matriculados_proyectados)
VALUES (3, 4, 'Grupo A', 'TEORIA', false, 40);

-- 9. SESIÓN DE DEPARTAMENTO (Jefe depto define franja semanal)
INSERT INTO sesion_departamento (id_departamento, id_periodo, dia_semana, hora_inicio, hora_fin) VALUES
(1, 1, 5, '12:00', '14:00');  -- Viernes 12:00-14:00
