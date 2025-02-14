package openapi

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool

// InitDB инициализирует подключение к базе данных
func InitDB() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		// Если не задана переменная окружения, используем DSN по умолчанию
		dsn = "postgres://postgres:testpass@localhost:5432/postgres?sslmode=disable"
	}
	var err error
	db, err = pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	log.Println("Database connected")
}

// GetDB возвращает пул соединений с БД
func GetDB() *pgxpool.Pool {
	return db
}

// CloseDB закрывает соединение с БД
func CloseDB() {
	if db != nil {
		db.Close()
	}
}
