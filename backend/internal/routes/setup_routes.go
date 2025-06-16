package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/txrps/next-golang-project/internal/handlers"
)

func SetUpRoutes(r *gin.Engine, handler *handlers.Handler) {
	SetupAuthRoutes(r, handler)
}
