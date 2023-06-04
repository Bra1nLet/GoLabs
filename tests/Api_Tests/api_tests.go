package Api_Tests

import (
	"github.com/golang-migrate/migrate/v4"
	"log"
	"os"
	"testing"
)

import (
	"context"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var dbURI string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file, using system env variables")
	}

	dbURI = os.Getenv("DATABASE_URL")
}

func initDB() func() {
	m, err := migrate.New("file://../migrations", dbURI+"?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
	log.Println("db scheme is up to date")

	return func() {
		if err := m.Down(); err != nil {
			log.Fatal(err)
		}
	}
}

func TestAll(t *testing.T) {
	// Перед запуском тестів необхідно підключитись до БД
	// Міграцію можна виконувати за межами тестів, але в даному випадку вона запускається в тестах.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	resetDB := initDB()
	pool, err := pgxpool.New(ctx, dbURI)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer func() {
		resetDB()
		pool.Close()
	}()

}
