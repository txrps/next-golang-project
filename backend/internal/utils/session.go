package utils

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"strings"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/txrps/next-golang-project/models"
	"gorm.io/gorm"
)

type Claims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

var secretKey = []byte("secret-key")
var ErrAuth = errors.New("unauthorized")
var ErrBadRequest = errors.New("badrequest")

func GenerateToken(length int) string {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		log.Fatalf("Failed to generate token: %v", err)
	}

	return base64.URLEncoding.EncodeToString(bytes)
}

func GenerateJWT(userID int64, username string) (string, error) {
	claims := Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			Issuer:    "https://softthaiapp.com",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func Authorize(c *gin.Context, DB *gorm.DB, username string) error {
	fmt.Printf("username is %s\n", username)
	var user models.User
	if err := DB.Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username"})
			return ErrAuth
		}
	}

	sessionToken, err := c.Cookie("session_token")
	fmt.Printf("sessionToken is %s\n", sessionToken)
	if err != nil || sessionToken == "" || sessionToken != user.SessionToken {
		return ErrAuth
	}

	crsf := c.GetHeader("X-CRSF-Token")
	if crsf == "" || crsf != user.CRSFToken {
		return ErrAuth
	}

	return nil
}

func ExtractUserFromToken(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return "", errors.New("missing authorization header")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("invalid authorization header format")
	}

	tokenStr := parts[1]
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("cannot parse claims")
	}

	username, ok := claims["username"].(string)
	if !ok {
		return "", errors.New("username not found in token")
	}

	return username, nil
}
