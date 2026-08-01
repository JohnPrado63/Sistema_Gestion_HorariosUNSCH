package auth

import (
	"net/http"

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
	ID    int    `json:"id_usuario"`
	Nombre string `json:"nombre"`
	Email  string `json:"email"`
	Rol    string `json:"rol"`
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

		token, err := GenerateToken(user.ID, user.Email, user.Rol)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al generar token"})
			return
		}

		c.JSON(http.StatusOK, LoginResponse{
			Token: token,
			User:  user,
		})
	}
}
