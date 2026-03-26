package database

import (
	"fmt"
	"strings"
	"time"

	"github.com/skyia-jp/shiro-go/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func resolveDialector(databaseURL string) gorm.Dialector {
	if strings.HasPrefix(databaseURL, "postgres://") || strings.HasPrefix(databaseURL, "postgresql://") {
		return postgres.Open(databaseURL)
	}
	return mysql.Open(databaseURL)
}

// Initialize creates a database connection
func Initialize(cfg *config.Config) (*gorm.DB, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	// Configure GORM logger
	gormLogger := logger.Default.LogMode(logger.Info)
	if cfg.LogLevel == "debug" {
		gormLogger = logger.Default.LogMode(logger.Info)
	} else {
		gormLogger = logger.Default.LogMode(logger.Silent)
	}

	// Connect to database (PostgreSQL or MySQL)
	db, err := gorm.Open(resolveDialector(cfg.DatabaseURL), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			NoLowerCase:   true,
		},
		Logger: gormLogger,
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql.DB from gorm: %w", err)
	}

	// Pool settings tuned for long-running bot process.
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}

// Close closes the database connection
func Close(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
