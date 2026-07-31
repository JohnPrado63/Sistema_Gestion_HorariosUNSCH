# Sistema_Gestion_HorariosUNSCH
Sistema de Gestión de Horarios y Carga Académica de la Universidad Nacional de San Cristóbal de Huamanga (UNSCH).

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
- Abrir la interfaz de usuario en el navegador en:
  ```text
  http://localhost:8081/app/
  ```

