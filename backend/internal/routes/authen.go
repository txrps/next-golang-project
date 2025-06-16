package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/txrps/next-golang-project/internal/handlers"
)

func SetupAuthRoutes(r *gin.Engine, handler *handlers.Handler) {
	authen := r.Group("/api/authen")
	{
		authen.POST("/register", handler.RegisterHandler())
		authen.POST("/login", handler.LoginHandler())
		authen.POST("/protected", handler.ProtectedHandler())
		authen.POST("/logout", handler.LogoutHandler())
	}
}
