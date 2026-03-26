package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/skyia-jp/shiro-go/internal/bot"
	"github.com/skyia-jp/shiro-go/internal/config"
	"github.com/skyia-jp/shiro-go/internal/database"
	"github.com/skyia-jp/shiro-go/internal/logger"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("Warning: .env file not found, continuing with environment variables")
	}

	// Initialize logger
	log, err := logger.Initialize()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}
	defer log.Sync()

	// Initialize config
	cfg := config.LoadConfig()

	// Initialize database
	db, err := database.Initialize(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer func() {
		if err := database.Close(db); err != nil {
			log.Warnf("Failed to close database connection cleanly: %v", err)
		}
	}()

	// Initialize bot client
	botClient, err := bot.NewClient(cfg, db, log)
	if err != nil {
		log.Fatalf("Failed to create bot client: %v", err)
	}

	// Start bot
	if err := botClient.Start(); err != nil {
		log.Fatalf("Failed to start bot: %v", err)
	}

	log.Info("Bot started successfully")

	// Setup signal handling for graceful shutdown
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan struct{})

	go func() {
		<-sc
		log.Info("Shutting down...")
		botClient.Close()
		close(done)
	}()

	<-done
	log.Info("Bot stopped")
}
