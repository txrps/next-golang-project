package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/txrps/next-golang-project/internal/handlers"
	"github.com/txrps/next-golang-project/internal/interfaces"
)

func SetupAuthRoutes(r *gin.Engine, handler *handlers.Handler) {
	authService := interfaces.NewAuthService(handler.DB)
	authHandler := handlers.NewAuthHandler(authService)
	authen := r.Group("/api/authen")
	{
		authen.POST("/register", authHandler.RegisterHandler())
		authen.POST("/login", authHandler.LoginHandler())
		authen.POST("/protected", handler.ProtectedHandler())
		authen.POST("/logout", authHandler.LogoutHandler())
	}
}
