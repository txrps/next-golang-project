package main

import (
	_ "embed"
	"log"
	"time"

	_ "github.com/txrps/next-golang-project/docs"

	"github.com/txrps/next-golang-project/config"
	"github.com/txrps/next-golang-project/database"
	"gorm.io/gen"
	"gorm.io/gorm"
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
	generateDB(db)
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get underlying DB: %v", err)
	}
	defer sqlDB.Close()
}

func generateDB(DB *gorm.DB) {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./models",
		Mode:    gen.WithoutContext,
	})

	g.UseDB(DB)
	g.ApplyBasic(
		g.GenerateAllTable()...,
	)
	g.Execute()
}
