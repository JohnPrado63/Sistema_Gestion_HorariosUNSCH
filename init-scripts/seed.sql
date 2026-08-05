-- =============================================================================
-- DATOS DE PRUEBA E INICIALIZACIÓN (SEED DATA)
-- =============================================================================

-- Password para todos los usuarios: admin123
-- Hash bcrypt generado: $2a$10$nAAC17PkVLbcigHr47UWe.fogi7U9SshX5VyZnhBVn3bgca/aeeK.

-- 1. Periodo Académico
INSERT INTO periodo_academico (codigo, activo) VALUES ('2026-I', true);

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
('Coordinador de Turno', 'coordinador@unsch.edu.pe', '$2a$10$nAAC17PkVLbcigHr47UWe.fogi7U9SshX5VyZnhBVn3bgca/aeeK.', 'COORDINADOR'),

-- Docentes
('María López Rodríguez', 'maria.lopez@unsch.edu.pe', '$2a$10$nAAC17PkVLbcigHr47UWe.fogi7U9SshX5VyZnhBVn3bgca/aeeK.', 'DOCENTE'),
('Carlos González Villa', 'carlos.gonzalez@unsch.edu.pe', '$2a$10$nAAC17PkVLbcigHr47UWe.fogi7U9SshX5VyZnhBVn3bgca/aeeK.', 'DOCENTE'),
('Juan Pérez Torres', 'juan.perez@unsch.edu.pe', '$2a$10$nAAC17PkVLbcigHr47UWe.fogi7U9SshX5VyZnhBVn3bgca/aeeK.', 'DOCENTE'),
('Ana García Ruiz', 'ana.garcia@unsch.edu.pe', '$2a$10$nAAC17PkVLbcigHr47UWe.fogi7U9SshX5VyZnhBVn3bgca/aeeK.', 'DOCENTE'),

-- Estudiantes (representativos)
('Pedro Huamaní Quispe', 'pedro.huamani@unsch.edu.pe', '$2a$10$nAAC17PkVLbcigHr47UWe.fogi7U9SshX5VyZnhBVn3bgca/aeeK.', 'ESTUDIANTE'),
('Sofia Curo Quispe', 'sofia.curo@unsch.edu.pe', '$2a$10$nAAC17PkVLbcigHr47UWe.fogi7U9SshX5VyZnhBVn3bgca/aeeK.', 'ESTUDIANTE');
