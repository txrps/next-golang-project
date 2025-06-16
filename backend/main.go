// backend/main.go
package main

import (
	"context"
	_ "embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/txrps/next-golang-project/config"
	"github.com/txrps/next-golang-project/database"
	"github.com/txrps/next-golang-project/internal/handlers"
	"github.com/txrps/next-golang-project/internal/routes"
)

const (
	gracefulShutdownDuration = 10 * time.Second
	serverReadHeaderTimeout  = 5 * time.Second
	serverReadTimeout        = 5 * time.Second
	serverWriteTimeout       = 10 * time.Second
	handlerTimeout           = serverWriteTimeout - (time.Millisecond * 100)
	errLoadConfigMessage     = "Failed to load config %v"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf(errLoadConfigMessage, err)
	}

	db := database.ConnectDB(config.DatabaseURL)
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get underlying DB: %v", err)
	}
	defer sqlDB.Close()

	r := router()

	handler := handlers.NewHandler(db)

	routes.SetUpRoutes(r, handler)

	serverAddr := fmt.Sprintf(":%s", config.ServerPort)
	server := &http.Server{
		Addr:              serverAddr,
		Handler:           r,
		ReadHeaderTimeout: serverReadHeaderTimeout,
		ReadTimeout:       serverReadTimeout,
		WriteTimeout:      serverWriteTimeout,
		MaxHeaderBytes:    1 << 20,
	}

	fmt.Printf("Server is up and running on PORT %s\n", serverAddr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed %v", err)
	}

}

func router() *gin.Engine {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf(errLoadConfigMessage, err)
	}

	r := gin.New()
	r.Use(gin.Recovery(), gin.Logger())

	r.NoMethod(func(c *gin.Context) {
		c.JSON(405, gin.H{
			"error": "Method Not Allowed",
		})
	})

	if config.Environment == "development" {
		r.Use(gin.Logger())
	}

	{
		r.GET("/liveness", liveness())
	}

	r.Use(
		accessControl,
		handlerTimeoutMiddleware,
	)

	r.GET("/api/test", testHandler)

	r.HandleMethodNotAllowed = true

	return r
}

func testHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello from Go backend",
		"time":    time.Now().Format(time.RFC3339),
	})
}

func handlerTimeoutMiddleware(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), handlerTimeout)
	defer cancel()
	c.Request = c.Request.WithContext(ctx)
	c.Next()
}

var headers = []string{
	"Content-Type",
	"Content-Length",
	"Accept-Encoding",
	"X-CSRF-Token",
	"Authorization",
	"accept",
	"origin",
	"Cache-Control",
	"X-Requested-With",
}

func accessControl(c *gin.Context) {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf(errLoadConfigMessage, err)
	}
	c.Writer.Header().Set("Access-Control-Allow-Origin", config.AllowOrigins)
	c.Writer.Header().Set("Access-Control-Allow-Method", "POST, GET, PUT, OPTIONS")
	c.Writer.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}
	c.Next()
}

// go build -ldflags "-X main.commit=123456"
var commit string

//go:embed VERSION
var version string

func liveness() func(c *gin.Context) {
	h, err := os.Hostname()
	if err != nil {
		h = fmt.Sprintf("unknown host err: %s", err.Error())
	}
	return func(c *gin.Context) {
		// the liveness probe is only this API itself probe
		// others service healthy not responsibility of this API
		// however, if you need it please follow these steps yourself
		// - check db connection, redis connection, etc
		// - implement help check your service dependencies
		// - implement help check for Postgres, MongoDB, Redis, etc
		//   e.g. MongoDB database.IsMongoReady()
		//   e.g. Redis database.IsRedisReady()
		//   e.g. Kafka database.IsKafkaReady()

		// e.g. check if Postgres is ready
		// if !database.IsPostgresReady() {
		// 	c.Status(http.StatusInternalServerError)
		// 	return
		// }

		c.JSON(http.StatusOK, gin.H{
			"hostname": h,
			"version":  strings.ReplaceAll(version, "\n", ""),
			"commit":   commit,
		})
	}
}
