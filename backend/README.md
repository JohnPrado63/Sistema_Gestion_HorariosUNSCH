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
