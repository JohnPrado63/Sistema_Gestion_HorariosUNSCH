package auth

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type User struct {
	ID     int    `json:"id_usuario"`
	Nombre string `json:"nombre"`
	Email  string `json:"email"`
	Rol    string `json:"rol"`
}

type CreateUserRequest struct {
	Nombre   string `json:"nombre" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Rol      string `json:"rol" binding:"required"`
}

type UpdateUserRequest struct {
	Nombre   string `json:"nombre"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Rol      string `json:"rol"`
}

type LogoutResponse struct {
	Message string `json:"message"`
}

func LoginHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email y password son requeridos"})
			return
		}

		var user User
		var passwordHash string
		err := db.QueryRow(c.Request.Context(),
			"SELECT id_usuario, nombre, email, rol, COALESCE(password_hash, '') FROM usuario WHERE email = $1",
			req.Email,
		).Scan(&user.ID, &user.Nombre, &user.Email, &user.Rol, &passwordHash)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales inválidas"})
			return
		}

		if !CheckPassword(req.Password, passwordHash) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales inválidas"})
			return
		}

		token, jti, err := GenerateToken(user.ID, user.Email, user.Rol)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al generar token"})
			return
		}

		_ = jti

		c.JSON(http.StatusOK, LoginResponse{
			Token: token,
			User:  user,
		})
	}
}

func LogoutHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		jti := GetTokenJTI(c)
		if jti == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No token found"})
			return
		}

		ctx := c.Request.Context()

		err := BlacklistToken(ctx, jti, 24*time.Hour)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al cerrar sesión"})
			return
		}

		c.JSON(http.StatusOK, LogoutResponse{
			Message: "Sesión cerrada exitosamente",
		})
	}
}

func MeHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := GetCurrentUserID(c)

		var user User
		err := db.QueryRow(c.Request.Context(),
			"SELECT id_usuario, nombre, email, rol FROM usuario WHERE id_usuario = $1",
			userID,
		).Scan(&user.ID, &user.Nombre, &user.Email, &user.Rol)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

func ListUsersHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query(c.Request.Context(),
			"SELECT id_usuario, nombre, email, rol FROM usuario ORDER BY id_usuario")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		users := []User{}
		for rows.Next() {
			var u User
			if err := rows.Scan(&u.ID, &u.Nombre, &u.Email, &u.Rol); err != nil {
				continue
			}
			users = append(users, u)
		}

		c.JSON(http.StatusOK, users)
	}
}

func CreateUserHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateUserRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		hashedPassword, err := HashPassword(req.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al cifrar contraseña"})
			return
		}

		var userID int
		err = db.QueryRow(c.Request.Context(),
			`INSERT INTO usuario (nombre, email, password_hash, rol)
			 VALUES ($1, $2, $3, $4)
			 RETURNING id_usuario`,
			req.Nombre, req.Email, hashedPassword, req.Rol,
		).Scan(&userID)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear usuario: " + err.Error()})
			return
		}

		c.JSON(http.StatusCreated, User{
			ID:     userID,
			Nombre: req.Nombre,
			Email:  req.Email,
			Rol:    req.Rol,
		})
	}
}

func UpdateUserHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}

		var req UpdateUserRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if req.Password != "" {
			hashedPassword, err := HashPassword(req.Password)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al cifrar contraseña"})
				return
			}

			_, err = db.Exec(context.Background(),
				`UPDATE usuario SET nombre = $1, email = $2, password_hash = $3, rol = $4 WHERE id_usuario = $5`,
				req.Nombre, req.Email, hashedPassword, req.Rol, id,
			)
		} else {
			_, err = db.Exec(context.Background(),
				`UPDATE usuario SET nombre = $1, email = $2, rol = $3 WHERE id_usuario = $4`,
				req.Nombre, req.Email, req.Rol, id,
			)
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar usuario"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Usuario actualizado"})
	}
}

func DeleteUserHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}

		currentUserID := GetCurrentUserID(c)
		if currentUserID == id {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No puedes eliminarte a ti mismo"})
			return
		}

		result, err := db.Exec(context.Background(),
			"DELETE FROM usuario WHERE id_usuario = $1",
			id,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar usuario"})
			return
		}

		if result.RowsAffected() == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Usuario eliminado"})
	}
}
