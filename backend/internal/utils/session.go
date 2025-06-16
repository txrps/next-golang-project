package utils

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"log"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

type AuthRequest struct {
	Username string `json:"username" binding:"required"`
}

type User struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	Username     string `json:"username" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
	Password     string `json:"password" binding:"required"`
	SessionToken string `json:"session_token"`
	CRSFToken    string `json:"crsf_token"`
}

func Authorize(c *gin.Context, DB *gorm.DB, username string) error {
	fmt.Printf("username is %s\n", username)
	var user User
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
