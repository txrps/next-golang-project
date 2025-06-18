package interfaces

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/txrps/next-golang-project/internal/utils"
	"github.com/txrps/next-golang-project/models"
	"gorm.io/gorm"
)

type AuthService interface {
	Register(user models.User) (interface{}, int, string)
	Login(req models.LoginRequest) (interface{}, string, string, string, int, string)
	Logout(username string, c *gin.Context) (int, string)
}

type authService struct {
	DB *gorm.DB
}

func NewAuthService(db *gorm.DB) AuthService {
	return &authService{DB: db}
}

func (s *authService) Register(user models.User) (interface{}, int, string) {
	if len(user.Username) < 5 || len(user.Password) < 5 {
		return nil, http.StatusNotAcceptable, "Username and password must be longer than 5 characters."
	}

	if exists, err := s.GetByUsername(user.Username); err == nil && exists != nil {
		return nil, http.StatusConflict, "Username already exists."
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, http.StatusInternalServerError, "Database error."
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, http.StatusInternalServerError, "Failed to hash password."
	}
	user.Password = hashedPassword

	if err := s.DB.Create(&user).Error; err != nil {
		return nil, http.StatusInternalServerError, "Failed to create user."
	}

	return &user, http.StatusCreated, "User created successfully."
}

func (s *authService) Login(req models.LoginRequest) (interface{}, string, string, string, int, string) {
	user, err := s.GetByUsername(req.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, "", "", "", http.StatusUnauthorized, "Invalid username."
		}
		return nil, "", "", "", http.StatusInternalServerError, "Database error."
	}

	if !utils.ComparePassword(user.Password, req.Password) {
		return nil, "", "", "", http.StatusUnauthorized, "Invalid password."
	}

	sessionToken := utils.GenerateToken(32)
	csrfToken := utils.GenerateToken(32)

	jwtToken, err := utils.GenerateJWT(int64(user.ID), user.Username)
	if err != nil {
		return nil, "", "", "", http.StatusInternalServerError, "Could not generate token."
	}

	user.SessionToken = sessionToken
	user.CRSFToken = csrfToken

	if err := s.DB.Save(user).Error; err != nil {
		return nil, "", "", "", http.StatusInternalServerError, "Failed to save session token."
	}

	return user, jwtToken, sessionToken, csrfToken, http.StatusOK, "Login successful"
}

func (s *authService) Logout(username string, c *gin.Context) (int, string) {

	if err := utils.Authorize(c, s.DB, username); err != nil {
		return http.StatusUnauthorized, "Unauthorized"
	}
	var user models.User
	if err := s.DB.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return http.StatusUnauthorized, "Unauthorized"
		}
		return http.StatusInternalServerError, "Database error"
	}

	user.SessionToken = ""
	user.CRSFToken = ""

	if err := s.DB.Save(&user).Error; err != nil {
		return http.StatusInternalServerError, "Failed to clear session tokens"
	}

	return http.StatusOK, "Logout successful"
}

func (s *authService) GetByUsername(username string) (*models.User, error) {
	var user models.User
	if err := s.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
