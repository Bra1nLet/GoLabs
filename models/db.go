package models

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
)

const migration = `
CREATE SCHEMA IF NOT EXISTS gowebapp;
CREATE TABLE gowebapp.users
(
    User_ID        bigserial NOT NULL,
    User_Name      text NOT NULL,
    Pass_Word_Hash text NOT NULL,
    Name           text NOT NULL,
    Config         jsonb NOT NULL DEFAULT '{}'::JSONB,
    Created_At     timestamp NOT NULL DEFAULT NOW(),
    Is_Enabled     boolean NOT NULL DEFAULT TRUE,
    CONSTRAINT PK_users PRIMARY KEY ( User_ID )
);`

func Migrate() {
	db, _ := ConnectDB()
	_, err := db.Exec(migration)
	if err != nil {
		return
	}

}

func ConnectDB() (*sql.DB, error) {
	dbURI := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		GetAsString("DB_USER", "postgres"),
		GetAsString("DB_PASSWORD", "postgres"),
		GetAsString("DB_HOST", "awesomeproject3-postgres-1"),
		GetAsInt("DB_PORT", 5432),
		GetAsString("DB_NAME", "postgres"),
	)

	log.Default()
	// Open the database
	db, err := sql.Open("postgres", dbURI)
	if err != nil {
		panic(err)
	}
	return db, err
}

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db: tx,
	}
}

// GetAsString reads an environment or returns a default value
func GetAsString(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// GetAsInt reads an environment variable into integer or returns a default value
func GetAsInt(name string, defaultValue int) int {
	valueStr := GetAsString(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}

type Queries struct {
	db DBTX
}
