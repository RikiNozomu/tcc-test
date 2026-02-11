package sqlite

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SQLiteAdapter struct {
	db *gorm.DB
}

// New creates a new SQLite database adapter
func New(dbPath string) (*SQLiteAdapter, error) {
	fmt.Println("Hello, World!")
	fmt.Println("Hello, World!")
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return &SQLiteAdapter{db: db}, nil
}

// GetDB returns the underlying GORM database instance
func (a *SQLiteAdapter) Get() *gorm.DB {
	return a.db
}

// Close closes the database connection
func (a *SQLiteAdapter) Close() error {
	sqlDB, err := a.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

// Migrate runs migrations for the given models
func (a *SQLiteAdapter) Migrate(models ...interface{}) error {
	return a.db.AutoMigrate(models...)
}

// Health checks the database connection
func (a *SQLiteAdapter) Health() error {
	sqlDB, err := a.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Ping()
}
