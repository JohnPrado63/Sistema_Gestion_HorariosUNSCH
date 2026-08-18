-- =============================================================================
-- PARCHE 3: Corregir grupos usando los docentes reales del sistema
-- =============================================================================

-- Eliminar los grupos mal asignados del parche anterior
DELETE FROM grupo WHERE id_carga IN (
    SELECT ca.id_carga FROM carga_academica ca
    JOIN curso c ON c.id_curso = ca.id_curso
    WHERE c.codigo IN ('CS101', 'MAT101', 'CS201')
    AND ca.id_periodo = 1
);

-- CS101 Grupo A - Teoría (docente: María López - ID 2, 4h)
INSERT INTO grupo (id_carga, id_docente, codigo_grupo, tipo_componente, es_nueva_necesidad, matriculados_proyectados)
SELECT ca.id_carga, 2, 'Grupo A', 'TEORIA', false, 35
FROM carga_academica ca
JOIN curso c ON c.id_curso = ca.id_curso
WHERE c.codigo = 'CS101' AND ca.id_periodo = 1 AND ca.id_escuela = 1;

-- CS101 Grupo B - Teoría (docente: Carlos González - ID 3, 4h)
INSERT INTO grupo (id_carga, id_docente, codigo_grupo, tipo_componente, es_nueva_necesidad, matriculados_proyectados)
SELECT ca.id_carga, 3, 'Grupo B', 'TEORIA', false, 38
FROM carga_academica ca
JOIN curso c ON c.id_curso = ca.id_curso
WHERE c.codigo = 'CS101' AND ca.id_periodo = 1 AND ca.id_escuela = 1;

-- CS201 Grupo A - Teoría (docente: Juan Pérez - ID 1, 4h)
INSERT INTO grupo (id_carga, id_docente, codigo_grupo, tipo_componente, es_nueva_necesidad, matriculados_proyectados)
SELECT ca.id_carga, 1, 'Grupo A', 'TEORIA', false, 30
FROM carga_academica ca
JOIN curso c ON c.id_curso = ca.id_curso
WHERE c.codigo = 'CS201' AND ca.id_periodo = 1 AND ca.id_escuela = 1;

-- CS201 Grupo A - Práctica (docente: Juan Pérez - ID 1, 3h)
INSERT INTO grupo (id_carga, id_docente, codigo_grupo, tipo_componente, es_nueva_necesidad, matriculados_proyectados, id_grupo_teoria_ref)
SELECT ca.id_carga, 1, 'Grupo A', 'PRACTICA', false, 30,
       (SELECT g.id_grupo FROM grupo g WHERE g.id_carga = ca.id_carga AND g.codigo_grupo = 'Grupo A' AND g.tipo_componente = 'TEORIA' LIMIT 1)
FROM carga_academica ca
JOIN curso c ON c.id_curso = ca.id_curso
WHERE c.codigo = 'CS201' AND ca.id_periodo = 1 AND ca.id_escuela = 1;

-- MAT101 Grupo A - Teoría (docente: María López - ID 2, 3h)
INSERT INTO grupo (id_carga, id_docente, codigo_grupo, tipo_componente, es_nueva_necesidad, matriculados_proyectados)
SELECT ca.id_carga, 2, 'Grupo A', 'TEORIA', false, 40
FROM carga_academica ca
JOIN curso c ON c.id_curso = ca.id_curso
WHERE c.codigo = 'MAT101' AND ca.id_periodo = 1 AND ca.id_escuela = 1;

-- UNA NUEVA NECESIDAD (grupo sin docente para probar sustitución)
-- IS-182 Grupo B - sin docente (Nueva Necesidad)
INSERT INTO grupo (id_carga, id_docente, codigo_grupo, tipo_componente, es_nueva_necesidad, matriculados_proyectados)
SELECT ca.id_carga, NULL, 'Grupo B', 'TEORIA', true, 35
FROM carga_academica ca
JOIN curso c ON c.id_curso = ca.id_curso
WHERE c.codigo = 'IS-182' AND ca.id_periodo = 1 AND ca.id_escuela = 1
AND NOT EXISTS (
    SELECT 1 FROM grupo g WHERE g.id_carga = ca.id_carga AND g.codigo_grupo = 'Grupo B' AND g.es_nueva_necesidad = true
);

-- VERIFICAR
-- SELECT c.codigo, g.codigo_grupo, g.tipo_componente, g.es_nueva_necesidad,
--        d.nombres || ' ' || d.apellidos as docente
-- FROM carga_academica ca
-- JOIN curso c ON c.id_curso = ca.id_curso
-- LEFT JOIN grupo g ON g.id_carga = ca.id_carga
-- LEFT JOIN docente d ON d.id_docente = g.id_docente
-- WHERE ca.id_periodo = 1 AND ca.id_escuela = 1
--   AND c.codigo IN ('CS101', 'CS201', 'MAT101')
-- ORDER BY c.codigo, g.codigo_grupo;
