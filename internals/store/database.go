package store

import (
	"database/sql"
	"fmt"
	"io/fs"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
)

func Open() (*sql.DB, error) {
	err := godotenv.Load()
    if err != nil {
        panic(`No .env file found, run "mv .env.example .env"`)
    }

    db_host := os.Getenv("DB_HOST")
	if db_host == "" {
		db_host = "localhost"
	}
	
	db, err := sql.Open("pgx","host="+ db_host +" user=postgres password=postgres dbname=postgres port=5432 sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("db: open %w", err)
	}

	fmt.Println("Connected to Database ...")
	return db, nil
}

func MigrateFS(db *sql.DB, migrationsFS fs.FS, dir string) error {
	goose.SetBaseFS(migrationsFS)
	defer func(){
		goose.SetBaseFS(nil)
	}()
	return Migrate(db, dir)
}

func Migrate(db *sql.DB, dir string) error {
	err := goose.SetDialect("postgres")
	if err != nil {
		return fmt.Errorf("migrate: %w",err)
	}
	err = goose.Up(db, dir)
	if err != nil {
		return fmt.Errorf("goose up: %w",err)
	}
	return nil
}