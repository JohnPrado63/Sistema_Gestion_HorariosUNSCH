# Sistema_Gestion_HorariosUNSCH

Sistema de Gestión de Horarios y Carga Académica de la Universidad Nacional de San Cristóbal de Huamanga (UNSCH).

## Requisitos previos

- Go 1.21+
- Node.js 18+
- Docker y Docker Compose

## Configuración

1. Copiar el archivo de entorno de ejemplo:
   ```powershell
   cp .env.example .env
   ```

2. Editar `.env` y configurar las contraseñas:
   ```
   DB_PASSWORD=tu_contraseña_postgres
   REDIS_PASSWORD=tu_contraseña_redis
   ```

## Ejecución

- Levantar la base de datos local con Docker Compose:
  ```powershell
  docker compose up -d
  ```

- Iniciar el backend:
  ```powershell
  cd backend
  go mod tidy
  go run ./cmd/api
  ```

- Iniciar el frontend:
  ```powershell
  cd frontend
  npm install
  npm run dev
  ```

- Abrir la interfaz de usuario en el navegador en:
  ```text
  http://localhost:8081/app/
  ```

## Tecnologías

- **Backend**: Go, Gin, PostgreSQL, Redis
- **Frontend**: Vue 3, Vite

