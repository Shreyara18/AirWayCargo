package handlers

import (
	"github.com/jackc/pgx/v4"
)

var dbConn *pgx.Conn

func InitiateDB(db *pgx.Conn) {
	dbConn = db
}