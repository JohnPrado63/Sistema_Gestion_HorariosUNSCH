-- =============================================================================
-- DATOS DE PRUEBA E INICIALIZACIÓN (SEED DATA)
-- =============================================================================

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

-- 4. Usuarios Base (password: admin123)
INSERT INTO usuario (nombre, email, password_hash, rol) VALUES
('Administrador TI', 'admin@unsch.edu.pe', '$2a$10$rtqeT8saALppb511PgvD/u3iTH02gDEJN/qlqOslOUwgINa3jAs9O', 'ADMIN_TI'),
('Director de Escuela IS', 'director.sistemas@unsch.edu.pe', '$2a$10$rtqeT8saALppb511PgvD/u3iTH02gDEJN/qlqOslOUwgINa3jAs9O', 'DIRECTOR_ESCUELA'),
('Jefe DGA', 'dga@unsch.edu.pe', '$2a$10$rtqeT8saALppb511PgvD/u3iTH02gDEJN/qlqOslOUwgINa3jAs9O', 'DGA');