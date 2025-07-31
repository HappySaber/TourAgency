package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/pressly/goose/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func buildDBConfig() *DBConfig {
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalf("Invalid DB_PORT: %v", err)
	}

	return &DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     port,
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
	}
}

func (config *DBConfig) dsn() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sadslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DBName,
	)
}

func Init() *gorm.DB {
	cfg := buildDBConfig()

	// Открытие подключения через GORM
	var err error
	DB, err = gorm.Open(postgres.Open(cfg.dsn()), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Connected to the database (GORM)")

	// Получение *sql.DB для goose
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to get sql.DB from gorm.DB: %v", err)
	}

	// Применение миграций через goose
	if err := runMigrations(sqlDB); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	return DB
}

func runMigrations(db *sql.DB) error {
	migrationsDir := "internal/migrations"

	log.Println("Running migrations with Goose...")
	if err := goose.Up(db, migrationsDir); err != nil {
		return fmt.Errorf("goose migration failed: %w", err)
	}

	log.Println("Goose migrations applied successfully")
	return nil
}
