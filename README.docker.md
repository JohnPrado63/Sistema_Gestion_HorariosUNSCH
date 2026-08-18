# Uso de Docker con el Sistema de Horarios UNSCH

## Requisitos previos

- Docker y Docker Compose v2+
- (Opcional) Make para simplificar comandos

## Archivos generados

```
├── backend/
│   ├── Dockerfile          # Producción (multi-stage, imagen pequeña)
│   └── Dockerfile.dev       # Desarrollo (hot-reload)
├── frontend/
│   ├── Dockerfile          # Producción (multi-stage + Nginx)
│   ├── Dockerfile.dev      # Desarrollo (Vite hot-reload)
│   └── nginx.conf          # Configuración Nginx para SPA
├── docker-compose.yml      # Servicios: postgres, redis, backend, frontend
├── docker-compose.override.yml  # Desarrollo con volúmenes (aplicado automáticamente)
└── .env.docker             # Variables para desarrollo
```

## Modo Desarrollo (default)

```powershell
# Usar archivo .env.docker
Copy-Item .env.docker .env

# Levantar todo
docker compose up -d

# Ver logs
docker compose logs -f

# Ver estado
docker compose ps
```

- **Backend**: http://localhost:8080
- **Frontend**: http://localhost:8081/app/
- **PostgreSQL**: localhost:5433
- **Redis**: localhost:6379

Los volúmenes montan el código fuente, así que los cambios se reflejan automáticamente (hot-reload).

## Modo Producción

```powershell
# Build sin usar override
docker compose -f docker-compose.yml build

# Solo producción (sin hot-reload)
docker compose -f docker-compose.yml up -d --build
```

## Comandos útiles

```powershell
# Reiniciar un servicio
docker compose restart backend

# Ver logs de un servicio
docker compose logs -f backend

# Entrar al contenedor
docker compose exec backend sh

# Reconstruir desde cero
docker compose down -v --rmi local
docker compose build --no-cache
docker compose up -d

# Ver uso de recursos
docker stats
```

## Estructura de imágenes

### Backend (Go)
- **Builder stage**: Compila el binary con `CGO_ENABLED=0`
- **Runner stage**: Alpine Linux (~20MB final)
- Puerto: 8080
- Healthcheck: `/health`

### Frontend (Vue)
- **Builder stage**: `npm run build` genera los estáticos
- **Runner stage**: Nginx sirve los estáticos en puerto 8081
- Puerto: 8081
- Healthcheck: `/health.html`

## Variables de entorno

| Variable | Default | Descripción |
|----------|---------|-------------|
| DB_HOST | localhost | Host de PostgreSQL |
| DB_PORT | 5433 | Puerto de PostgreSQL |
| DB_NAME | unsch_horarios | Nombre de la base |
| DB_USER | postgres | Usuario de PostgreSQL |
| DB_PASSWORD | (requerido) | Contraseña |
| REDIS_HOST | localhost | Host de Redis |
| REDIS_PORT | 6379 | Puerto de Redis |
| REDIS_PASSWORD | (requerido) | Contraseña Redis |
| APP_PORT | 8080 | Puerto del backend |
| FRONTEND_PORT | 8081 | Puerto del frontend |
