package main

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	envFile, _ := godotenv.Read("../../.env")
	databaseUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		envFile["DB_USER"],
		envFile["DB_PASSWORD"],
		envFile["DB_HOST"],
		envFile["DB_PORT"],
		envFile["DB_NAME"],
	)

	m, err := migrate.New(
		"file://../migrations",
		databaseUrl,
	)

	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}
