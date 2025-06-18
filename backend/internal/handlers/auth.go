package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/txrps/next-golang-project/internal/utils"
	"github.com/txrps/next-golang-project/models"
	"gorm.io/gorm"
)

const (
	usernameCondition = "username = ?"
)

// POST api/auth/register
func (h *Handler) RegisterHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		username := user.Username
		password := user.Password
		if len(username) < 5 || len(password) < 5 {
			c.JSON(http.StatusNotAcceptable, gin.H{"invalid data": "Please ensure that the username and password are longer than 8 characters."})
			return
		}

		var existingUser models.User
		if err := h.DB.Where(usernameCondition, user.Username).First(&existingUser).Error; err == nil {
			c.JSON(http.StatusConflict, gin.H{"error": "Username already exists."})
			return
		} else if err != gorm.ErrRecordNotFound {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error."})
			return
		}

		hashedPassword, err := utils.HashPassword(user.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password."})
			return
		}
		user.Password = hashedPassword

		if err := h.DB.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user."})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "User created successfully",
			"user":    user,
		})
	}
}

// POST api/auth/login
func (h *Handler) LoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var user models.User
		if err := h.DB.Where(usernameCondition, req.Username).First(&user).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
			return
		}

		isPassword := utils.ComparePassword(user.Password, req.Password)
		if !isPassword {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
			return
		}

		sessionToken := utils.GenerateToken(32)
		csrfToken := utils.GenerateToken(32)
		jwtToken, err := utils.GenerateJWT(int64(user.ID), user.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
			return
		}

		c.SetCookie(
			"jwt_token", // name
			jwtToken,    // value
			86400,       // maxAge in seconds
			"/",         // path
			"",          // domain (optional)
			false,       // secure
			true,        // httpOnly
		)

		c.SetCookie(
			"session_token",
			sessionToken,
			86400,
			"/",
			"",
			false,
			true,
		)

		c.SetCookie(
			"csrf_token",
			csrfToken,
			86400,
			"/",
			"",
			false,
			false,
		)
		user.SessionToken = sessionToken
		user.CRSFToken = csrfToken

		if err := h.DB.Save(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":  "Login successful",
			"username": user.Username,
			"email":    user.Email,
		})
	}
}

// Post api/auth/login
func (h *Handler) ProtectedHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := utils.ExtractUserFromToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		if err := utils.Authorize(c, h.DB, user); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "CSRF validation successfully!",
		})
	}
}

// Post api/auth/logout
func (h *Handler) LogoutHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, err := utils.ExtractUserFromToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		if err := utils.Authorize(c, h.DB, username); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		c.SetCookie(
			"jwt_token",
			"",
			-1,
			"/",
			"",
			false,
			true,
		)

		c.SetCookie(
			"session_token",
			"",
			-1,
			"/",
			"",
			false,
			true,
		)

		c.SetCookie(
			"csrf_token",
			"",
			-1,
			"/",
			"",
			false,
			false,
		)

		var user models.User
		if err := h.DB.Where(usernameCondition, username).First(&user).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				return
			}
		}

		user.SessionToken = ""
		user.CRSFToken = ""
		if err := h.DB.Save(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Logout successful",
		})
	}
}
