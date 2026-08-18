-- =============================================================================
-- PARCHE: Agregar datos de prueba SIN afectar datos existentes
-- =============================================================================
-- Ejecutar ESTE archivo DESPUÉS del schema.sql principal
-- Solo inserta datos faltantes, no modifica ni elimina nada existente
-- =============================================================================

-- 1. DOCENTES (solo si la tabla está vacía)
-- Verificar si ya hay docentes
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM docente LIMIT 1) THEN
        INSERT INTO docente (id_departamento, codigo_plaza, nombres, apellidos, email) VALUES
        (1, 'DNI-001', 'María', 'López Rodríguez', 'maria.lopez@unsch.edu.pe'),
        (1, 'DNI-002', 'Carlos', 'González Villa', 'carlos.gonzalez@unsch.edu.pe'),
        (1, 'DNI-003', 'Juan', 'Pérez Torres', 'juan.perez@unsch.edu.pe'),
        (1, 'DNI-004', 'Ana', 'García Ruiz', 'ana.garcia@unsch.edu.pe'),
        (1, 'DNI-005', 'Luis', 'Ramírez Santos', 'luis.ramirez@unsch.edu.pe'),
        (1, 'DNI-006', 'Elena', 'Castro Morales', 'elena.castro@unsch.edu.pe');
        RAISE NOTICE 'Docentes insertados';
    ELSE
        RAISE NOTICE 'Docentes ya existen, omitidos';
    END IF;
END $$;

-- 2. PLAN DE ESTUDIO (solo si no existe para esta escuela)
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM plan_estudio WHERE id_escuela = 1 AND codigo_plan = '2024') THEN
        INSERT INTO plan_estudio (id_escuela, codigo_plan, nombre) VALUES (1, '2024', 'Plan de Estudios 2024 - IS');
        RAISE NOTICE 'Plan de estudio insertado';
    ELSE
        RAISE NOTICE 'Plan de estudio ya existe, omitido';
    END IF;
END $$;

-- 3. SERIES (ciclos) - insertar solo las que no existan
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM serie s JOIN plan_estudio p ON p.id_plan = s.id_plan WHERE p.codigo_plan = '2024' AND s.numero_ciclo = 100) THEN
        INSERT INTO serie (id_plan, numero_ciclo)
        SELECT p.id_plan, v.numero_ciclo
        FROM plan_estudio p
        CROSS JOIN (VALUES (100), (200), (300), (400), (500)) AS v(numero_ciclo)
        WHERE p.codigo_plan = '2024';
        RAISE NOTICE 'Series insertadas';
    ELSE
        RAISE NOTICE 'Series ya existen, omitidas';
    END IF;
END $$;

-- 4. CURSOS DE EJEMPLO (solo si NO existen cursos con esos códigos)
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM curso WHERE codigo = 'CS101') THEN
        INSERT INTO curso (id_serie, codigo, nombre, creditos, horas_teoria, horas_practica)
        SELECT s.id_serie, v.codigo, v.nombre, v.creditos, v.ht, v.hp
        FROM serie s
        JOIN plan_estudio p ON p.id_plan = s.id_plan
        CROSS JOIN (VALUES
            (100, 'CS101', 'Introducción a la Programación', 4, 2, 2),
            (100, 'MAT101', 'Matemática Básica', 3, 2, 1),
            (200, 'CS201', 'Programación Orientada a Objetos', 4, 2, 2),
            (200, 'MAT201', 'Cálculo I', 4, 3, 1)
        ) AS v(serie, codigo, nombre, creditos, ht, hp)
        WHERE p.codigo_plan = '2024' AND s.numero_ciclo = v.serie;
        RAISE NOTICE 'Cursos de ejemplo insertados';
    ELSE
        RAISE NOTICE 'Cursos ya existen, omitidos';
    END IF;
END $$;

-- 5. SESIÓN DE DEPARTAMENTO (solo si no existe para este dept y periodo)
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM sesion_departamento WHERE id_departamento = 1 AND id_periodo = 1) THEN
        INSERT INTO sesion_departamento (id_departamento, id_periodo, dia_semana, hora_inicio, hora_fin) VALUES
        (1, 1, 5, '12:00', '14:00');
        RAISE NOTICE 'Sesión de departamento insertada';
    ELSE
        RAISE NOTICE 'Sesión de departamento ya existe, omitida';
    END IF;
END $$;

-- =============================================================================
-- CONSULTA PARA VERIFICAR: Muestra docentes disponibles
-- =============================================================================
-- SELECT d.*, dep.nombre as departamento
-- FROM docente d
-- JOIN departamento_academico dep ON dep.id_departamento = d.id_departamento;

-- =============================================================================
-- NOTA: Las cargas académicas y grupos de ejemplo NO se insertan automáticamente
-- porque dependen de los IDs de cursos específicos que ya tienes.
-- Si quieres crear cargas de prueba, hazlo manualmente desde la interfaz.
-- =============================================================================
