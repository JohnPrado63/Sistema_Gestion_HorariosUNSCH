package auth

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte(getJWTSecret())

func getJWTSecret() string {
	return "unsch-horarios-secret-2026"
}

type Claims struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	Rol    string `json:"rol"`
	JTI    string `json:"jti"`
	jwt.RegisteredClaims
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateToken(userID int, email, rol string) (string, string, error) {
	jti := uuid.New().String()

	claims := Claims{
		UserID: userID,
		Email:  email,
		Rol:    rol,
		JTI:    jti,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "unsch-horarios",
			ID:       jti,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	return tokenString, jti, err
}

func ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		tokenString := ""
		if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
			tokenString = authHeader[7:]
		} else {
			c.JSON(401, gin.H{"error": "Invalid authorization format"})
			c.Abort()
			return
		}

		claims, err := ValidateToken(tokenString)
		if err != nil {
			c.JSON(401, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		if IsRedisEnabled() {
			ctx := c.Request.Context()
			blacklisted, err := IsTokenBlacklisted(ctx, claims.JTI)
			if err != nil {
				c.JSON(500, gin.H{"error": "Error checking token status"})
				c.Abort()
				return
			}
			if blacklisted {
				c.JSON(401, gin.H{"error": "Token has been revoked"})
				c.Abort()
				return
			}
		}

		c.Set("user_id", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("rol", claims.Rol)
		c.Set("jti", claims.JTI)
		c.Set("claims", claims)
		c.Next()
	}
}

func RequireRole(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRol, exists := c.Get("rol")
		if !exists {
			c.JSON(403, gin.H{"error": "Rol no encontrado"})
			c.Abort()
			return
		}

		rol := userRol.(string)
		for _, r := range roles {
			if rol == r {
				c.Next()
				return
			}
		}

		c.JSON(403, gin.H{"error": "No tienes permisos para esta acción"})
		c.Abort()
	}
}

func GetCurrentUserID(c *gin.Context) int {
	userID, _ := c.Get("user_id")
	return userID.(int)
}

func GetCurrentUserEmail(c *gin.Context) string {
	email, _ := c.Get("email")
	return email.(string)
}

func GetCurrentUserRol(c *gin.Context) string {
	rol, _ := c.Get("rol")
	return rol.(string)
}

func GetTokenJTI(c *gin.Context) string {
	jti, exists := c.Get("jti")
	if !exists {
		return ""
	}
	return jti.(string)
}
