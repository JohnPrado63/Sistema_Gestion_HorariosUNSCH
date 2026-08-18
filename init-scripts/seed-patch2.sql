-- =============================================================================
-- PARCHE 2: Crear cargas académicas para cursos nuevos y grupos de ejemplo
-- =============================================================================
-- Ejecutar después de seed-patch.sql
-- Solo inserta si no existen las cargas
-- =============================================================================

-- 1. Crear cargas académicas para los 4 cursos nuevos (en BORRADOR para poder probar)
INSERT INTO carga_academica (id_curso, id_periodo, id_escuela, estado)
SELECT id_curso, 1, 1, 'BORRADOR'
FROM curso
WHERE codigo IN ('CS101', 'MAT101', 'CS201', 'MAT201')
AND NOT EXISTS (
    SELECT 1 FROM carga_academica ca
    WHERE ca.id_curso = curso.id_curso AND ca.id_periodo = 1 AND ca.id_escuela = 1
);

-- 2. Crear grupos de ejemplo para poder probar asignación de docentes
-- CS101 Grupo A - Teoría (docente: María López - 4h)
INSERT INTO grupo (id_carga, id_docente, codigo_grupo, tipo_componente, es_nueva_necesidad, matriculados_proyectados)
SELECT ca.id_carga, 1, 'Grupo A', 'TEORIA', false, 35
FROM carga_academica ca
JOIN curso c ON c.id_curso = ca.id_curso
WHERE c.codigo = 'CS101' AND ca.id_periodo = 1 AND ca.id_escuela = 1
AND NOT EXISTS (
    SELECT 1 FROM grupo g WHERE g.id_carga = ca.id_carga AND g.codigo_grupo = 'Grupo A' AND g.tipo_componente = 'TEORIA'
);

-- CS101 Grupo B - Teoría (docente: Carlos González - 4h)
INSERT INTO grupo (id_carga, id_docente, codigo_grupo, tipo_componente, es_nueva_necesidad, matriculados_proyectados)
SELECT ca.id_carga, 2, 'Grupo B', 'TEORIA', false, 38
FROM carga_academica ca
JOIN curso c ON c.id_curso = ca.id_curso
WHERE c.codigo = 'CS101' AND ca.id_periodo = 1 AND ca.id_escuela = 1
AND NOT EXISTS (
    SELECT 1 FROM grupo g WHERE g.id_carga = ca.id_carga AND g.codigo_grupo = 'Grupo B' AND g.tipo_componente = 'TEORIA'
);

-- CS201 Grupo A - Teoría (docente: Juan Pérez - 4h)
INSERT INTO grupo (id_carga, id_docente, codigo_grupo, tipo_componente, es_nueva_necesidad, matriculados_proyectados)
SELECT ca.id_carga, 3, 'Grupo A', 'TEORIA', false, 30
FROM carga_academica ca
JOIN curso c ON c.id_curso = ca.id_curso
WHERE c.codigo = 'CS201' AND ca.id_periodo = 1 AND ca.id_escuela = 1
AND NOT EXISTS (
    SELECT 1 FROM grupo g WHERE g.id_carga = ca.id_carga AND g.codigo_grupo = 'Grupo A' AND g.tipo_componente = 'TEORIA'
);

-- CS201 Grupo A - Práctica (docente: Juan Pérez - 3h, linked)
INSERT INTO grupo (id_carga, id_docente, codigo_grupo, tipo_componente, es_nueva_necesidad, matriculados_proyectados, id_grupo_teoria_ref)
SELECT ca.id_carga, 3, 'Grupo A', 'PRACTICA', false, 30,
       (SELECT g.id_grupo FROM grupo g WHERE g.id_carga = ca.id_carga AND g.codigo_grupo = 'Grupo A' AND g.tipo_componente = 'TEORIA' LIMIT 1)
FROM carga_academica ca
JOIN curso c ON c.id_curso = ca.id_curso
WHERE c.codigo = 'CS201' AND ca.id_periodo = 1 AND ca.id_escuela = 1
AND NOT EXISTS (
    SELECT 1 FROM grupo g WHERE g.id_carga = ca.id_carga AND g.codigo_grupo = 'Grupo A' AND g.tipo_componente = 'PRACTICA'
);

-- MAT101 Grupo A - Teoría (docente: Ana García - 3h)
INSERT INTO grupo (id_carga, id_docente, codigo_grupo, tipo_componente, es_nueva_necesidad, matriculados_proyectados)
SELECT ca.id_carga, 4, 'Grupo A', 'TEORIA', false, 40
FROM carga_academica ca
JOIN curso c ON c.id_curso = ca.id_curso
WHERE c.codigo = 'MAT101' AND ca.id_periodo = 1 AND ca.id_escuela = 1
AND NOT EXISTS (
    SELECT 1 FROM grupo g WHERE g.id_carga = ca.id_carga AND g.codigo_grupo = 'Grupo A' AND g.tipo_componente = 'TEORIA'
);

-- 3. UNA NUEVA NECESIDAD de ejemplo (grupo sin docente - plaza vacante)
-- IS-182 Grupo B - sin docente asignado (Nueva Necesidad)
INSERT INTO grupo (id_carga, id_docente, codigo_grupo, tipo_componente, es_nueva_necesidad, matriculados_proyectados)
SELECT ca.id_carga, NULL, 'Grupo B', 'TEORIA', true, 35
FROM carga_academica ca
JOIN curso c ON c.id_curso = ca.id_curso
WHERE c.codigo = 'IS-182' AND ca.id_periodo = 1 AND ca.id_escuela = 1
AND NOT EXISTS (
    SELECT 1 FROM grupo g WHERE g.id_carga = ca.id_carga AND g.codigo_grupo = 'Grupo B' AND g.es_nueva_necesidad = true
);

-- =============================================================================
-- VERIFICAR: Consultas para ver lo insertado
-- =============================================================================
-- SELECT c.codigo, c.nombre, ca.id_carga, ca.estado,
--        g.id_grupo, g.codigo_grupo, g.tipo_componente, g.es_nueva_necesidad,
--        d.nombres || ' ' || d.apellidos as docente
-- FROM carga_academica ca
-- JOIN curso c ON c.id_curso = ca.id_curso
-- LEFT JOIN grupo g ON g.id_carga = ca.id_carga
-- LEFT JOIN docente d ON d.id_docente = g.id_docente
-- WHERE ca.id_periodo = 1 AND ca.id_escuela = 1
--   AND c.codigo IN ('CS101', 'MAT101', 'CS201', 'MAT101')
-- ORDER BY c.codigo;
