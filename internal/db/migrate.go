package db

import (
	"database/sql"
	"fmt"
	"log"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source/file"
	"github.com/mohrezfadaei/ipresist/config"
)

func RunMigrations() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true",
		config.DB_USER, config.DB_PASSWORD, config.DB_HOST, config.DB_PORT, config.DB_NAME)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatalf("Failed to create migration driver: %v", err)
	}

	migrationsPath, err := filepath.Abs("internal/db/migrations")
	if err != nil {
		log.Fatalf("Failed to get absolute path of migrations: %v", err)
	}

	log.Printf("Migrations path: %s", migrationsPath)

	// Import file driver
	sourceDriver, err := (&file.File{}).Open("file://" + migrationsPath)
	if err != nil {
		log.Fatalf("Failed to create source driver: %v", err)
	}

	m, err := migrate.NewWithInstance(
		"file", sourceDriver, config.DB_NAME, driver)
	if err != nil {
		log.Fatalf("Failed to create migrate instance: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to apply migrations: %v", err)
	}

	log.Println("Database migrations applied successfully")
}
