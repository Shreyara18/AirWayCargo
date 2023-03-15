package main

import (
	"database/sql"
	"embed"
	"encoding/json"
	"os"
	"strings"

	"flag"

	"AIRWAYCARGO/dtos"
	"AIRWAYCARGO/internal/config"
	"AIRWAYCARGO/internal/config/globals"
	"AIRWAYCARGO/internal/handlers"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"go.uber.org/zap"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {

	runMigration := flag.String("migration", "", "Flag to check if Migrations need to Run")

	flag.Parse()

	if runMigration != nil && strings.ToUpper(*runMigration) == "ON" {
		runMigrations()
		os.Exit(0)
	}

	initiateDB()
	// go jobs.a

	r := handlers.GetRouter()
	globals.Logger.Info("Listening to Port: " + config.Port)

	r.Run(":" + config.Port)

	allmodels()
}

func allmodels() {
	v := dtos.GetResp{}
	daostring, _ := json.Marshal(v)
	println(string(daostring))
}

func initiateDB() {
	orm.Init(&orm.Config{
		URL:       config.DatabaseURL,
		MaxDBConn: config.MaxDBConnections,
	})
}

func runMigrations() {
	migrationUserDbUrl := os.Getenv("MIGRATIONS_USER_DB_URL")

	if strings.TrimSpace(migrationUserDbUrl) == "" {
		globals.Logger.Fatal("MIGRATIONS_USER_DB_URL is not provided")
		os.Exit(1)
	}
	db, err := sql.Open("postgres", migrationUserDbUrl)
	if err != nil {
		globals.Logger.Fatal("PG DB Connection Failed", zap.Error(err))
		panic(err)
	}
	// setup database
	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("postgres"); err != nil {
		globals.Logger.Fatal("Setting Goose Postgres Dialect Failed", zap.Error(err))
		panic(err)
	}
	if err := goose.Up(db, "migrations"); err != nil {
		globals.Logger.Fatal("Goose Up Failed", zap.Error(err))
		panic(err)
	}
}
