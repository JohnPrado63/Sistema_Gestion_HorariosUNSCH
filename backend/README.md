# Backend Go - Sistema de Horarios UNSCH

API inicial en Go para conectarse a PostgreSQL y exponer los primeros catalogos del sistema.

## Ejecutar

```powershell
cd backend
go mod tidy
go run ./cmd/api
```

La API escucha por defecto en:

```text
http://localhost:8081
```

## Endpoints iniciales

```text
GET /health
GET /ready
GET /api/v1/facultades
GET /api/v1/departamentos
GET /api/v1/escuelas
GET /api/v1/aulas
GET /api/v1/usuarios
GET /api/v1/planes
GET /api/v1/series
GET /api/v1/cursos
GET /api/v1/docentes
GET /api/v1/periodos
GET /api/v1/sesiones-departamento
GET /api/v1/locales
GET /api/v1/pabellones
GET /api/v1/matriz-distancias
GET /api/v1/cargas-academicas
GET /api/v1/grupos
GET /api/v1/horarios
GET /api/v1/bloques
GET /api/v1/bitacora
GET /validaciones
GET /api/v1/validaciones/escenarios
POST /api/v1/validaciones/placement
POST /api/v1/validaciones/audit
POST /api/v1/validaciones/carga
```

## Validaciones de horario

Ejemplo de validacion de ubicacion de bloque:

```powershell
curl.exe -X POST http://localhost:8081/api/v1/validaciones/placement `
  -H "Content-Type: application/json" `
  -d '{"proposed":{"teacher_id":10,"day":1,"start_slot":3,"end_slot":4,"room_id":5,"school_id":1},"existing":[{"teacher_id":10,"day":1,"start_slot":4,"end_slot":5,"room_id":6,"school_id":1}],"state":"BORRADOR"}'
```

Ejemplo de validacion de justificacion de horario:

```powershell
curl.exe -X POST http://localhost:8081/api/v1/validaciones/audit `
  -H "Content-Type: application/json" `
  -d '{"state":"OFICIAL","justification":"Ajuste urgente"}'
```

Ejemplo de validacion de carga lectiva:

```powershell
curl.exe -X POST http://localhost:8081/api/v1/validaciones/carga `
  -H "Content-Type: application/json" `
  -d '{"teacher_id":10,"weekly_hours":18,"confirmed":false}'
```

## Conexion

Por defecto usa:

```text
postgres://postgres:sulcaprado@localhost:5433/unsch_horarios?sslmode=disable
```

## Pruebas de reglas de negocio

Ejecutar todas las pruebas:

```powershell
go test ./...
```

Ejecutar solo las validaciones de horarios:

```powershell
go test ./internal/schedule/validation -v
```

Reglas cubiertas inicialmente:

```text
RV-01  Cruce de docente                         Bloqueo
RV-02  Cruce de aula                            Bloqueo
RV-03  Sesion de departamento                   Bloqueo
RV-04a Traslado sin tiempo                      Bloqueo
RV-04b Traslado con tiempo insuficiente         Advertencia
RV-05  Tope de carga lectiva mayor a 16h        Advertencia DGA
RV-06  Cruce en la misma serie                  Informativo
RV-07  Aula compartida reservada por otra escuela Bloqueo
RV-08  Justificacion en horario oficial/reajustado Bloqueo
RV-09  Aforo excedido en PRELIMINAR/EN_REAJUSTE Info/Advertencia
```

Estas pruebas usan logica pura en Go. El siguiente paso es conectar este validador al endpoint que cree o modifique bloques de horario.
