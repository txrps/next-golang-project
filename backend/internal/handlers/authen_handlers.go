package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/txrps/next-golang-project/internal/interfaces"
	"github.com/txrps/next-golang-project/internal/utils"
	"github.com/txrps/next-golang-project/models"
)

type AuthHandler struct {
	svc interfaces.AuthService
}

func NewAuthHandler(s interfaces.AuthService) *AuthHandler {
	return &AuthHandler{svc: s}
}

// RegisterHandler godoc
// @Summary      Register new user
// @Description  Create a new user account with a unique username and password
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        user  body      models.User  true  "User credentials"
// @Success      201   {object}  models.ResultAPI  		 "Created"
// @Failure      406   {object}  models.ResultAPI        "Validation error"
// @Failure      409   {object}  models.ResultAPI        "Username conflict"
// @Failure      400   {object}  models.ResultAPI        "Bad request"
// @Failure      500   {object}  models.ResultAPI        "Internal error"
// @Router       /authen/register [post]
func (h *AuthHandler) RegisterHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, models.ResultAPI{
				StatusCode: http.StatusBadRequest,
				Message:    err.Error(),
			})
			return
		}

		newUser, status, msg := h.svc.Register(user)
		if status != http.StatusCreated {
			c.JSON(status, models.ResultAPI{
				StatusCode: status,
				Message:    msg,
			})
			return
		}

		c.JSON(http.StatusCreated, models.ResultAPI{
			StatusCode: status,
			Message:    msg,
			Result:     newUser,
		})
	}
}

// LoginHandler godoc
// @Summary      Login user
// @Description  Authenticates user with username and password, and sets cookies for JWT, session, and CSRF
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        credentials  body      models.LoginRequest  true  "User login credentials"
// @Success      200  {object}  models.ResultAPI         "Login successful"
// @Failure      400  {object}  models.ResultAPI         "Invalid request payload"
// @Failure      401  {object}  models.ResultAPI         "Invalid username or password"
// @Failure      500  {object}  models.ResultAPI         "Internal server/database error"
// @Router       /authen/login [post]
func (h *AuthHandler) LoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, models.ResultAPI{
				StatusCode: http.StatusBadRequest,
				Message:    err.Error(),
			})
			return
		}

		user, jwtToken, sessionToken, csrfToken, status, msg := h.svc.Login(req)
		if status != http.StatusOK {
			c.JSON(status, models.ResultAPI{
				StatusCode: status,
				Message:    msg,
			})
			return
		}

		c.SetCookie("jwt_token", jwtToken, 86400, "/", "", false, true)
		c.SetCookie("session_token", sessionToken, 86400, "/", "", false, true)
		c.SetCookie("csrf_token", csrfToken, 86400, "/", "", false, false)

		c.JSON(http.StatusOK, models.ResultAPI{
			StatusCode: http.StatusOK,
			Message:    msg,
			Result:     user,
		})
	}
}

// Post api/auth/login
func (h *Handler) ProtectedHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := utils.ExtractUserFromToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, models.ResultAPI{
				StatusCode: http.StatusUnauthorized,
				Message:    "Invalid token",
			})
			return
		}

		if err := utils.Authorize(c, h.DB, user); err != nil {
			c.JSON(http.StatusUnauthorized, models.ResultAPI{
				StatusCode: http.StatusUnauthorized,
				Message:    "Unauthorized",
			})
			return
		}

		c.JSON(http.StatusOK, models.ResultAPI{
			StatusCode: http.StatusOK,
			Message:    "CSRF validation successfully!",
		})
	}
}

// LogoutHandler godoc
// @Summary      Logout user
// @Description  Clears authentication cookies and session from database for the current user
// @Tags         Auth
// @Security     ApiKeyAuth
// @Produce      json
// @Success      200 {object} models.ResultAPI "Logout successful"
// @Failure      401 {object} models.ResultAPI "Invalid or missing token"
// @Failure      500 {object} models.ResultAPI "Failed to clear session"
// @Router       /authen/logout [post]
func (h *AuthHandler) LogoutHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, err := utils.ExtractUserFromToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, models.ResultAPI{
				StatusCode: http.StatusUnauthorized,
				Message:    "Invalid token",
			})
			return
		}

		status, msg := h.svc.Logout(username, c)
		if status != http.StatusOK {
			c.JSON(status, models.ResultAPI{
				StatusCode: status,
				Message:    msg,
			})
			return
		}

		// ลบ cookie
		c.SetCookie("jwt_token", "", -1, "/", "", false, true)
		c.SetCookie("session_token", "", -1, "/", "", false, true)
		c.SetCookie("csrf_token", "", -1, "/", "", false, false)

		c.JSON(http.StatusOK, models.ResultAPI{
			StatusCode: http.StatusOK,
			Message:    msg,
		})
	}
}
