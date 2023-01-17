package config

import (
	"Air_Way_Cargo/internal/handlers"
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
)

const (
    DB_USER     = "postgres"
    DB_PASSWORD = ""
    DB_NAME     = "Air_Way_Cargo"
    DB_ADDRESS  = "localhost:5432"
    DB_URL      = "postgres://" + DB_USER + ":" + DB_PASSWORD + "@" + DB_ADDRESS + "/" + DB_NAME
)

func Connect() *pgx.Conn {
    conn, err := pgx.Connect(context.Background(), DB_URL)
    pingErr := conn.Ping(context.Background())

    if err != nil || pingErr != nil {
        log.Printf("Failed to connect to DB %v\n", err)
        os.Exit(100)
    } else {
        log.Printf("Connected to DB")
    }
    handlers.InitiateDB(conn)

    return conn
}

