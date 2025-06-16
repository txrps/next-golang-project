package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/txrps/next-golang-project/internal/handlers"
)

func SetupAuthRoutes(r *gin.Engine, handler *handlers.Handler) {
	authen := r.Group("/authen")
	{
		authen.POST("/register", handler.RegisterHandler())
		authen.POST("/login", handler.LoginHandler())
	}
}
