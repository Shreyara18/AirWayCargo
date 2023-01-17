package orm

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/wiz-freight-org/adapters/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Config struct {
	URL               string
	MaxDBConn         int
	MigrationFilePath string
}

// Init the connection to DB
func Init(config *Config) error {
	if DB == nil {
		sqlDB, err := sql.Open("postgres", config.URL)
		if err != nil {
			log.Println("Unable to open postges connection. Err:", err)
			return err
		}

		sqlDB.SetMaxIdleConns(config.MaxDBConn)
		sqlDB.SetMaxOpenConns(config.MaxDBConn)
		sqlDB.SetConnMaxLifetime(time.Hour)

		DB, err = gorm.Open(postgres.New(postgres.Config{
			Conn: sqlDB,
		}), &gorm.Config{})
		if err != nil {
			log.Println("Unable to open postges gorm connection. Err:", err)
			return err
		}

		log.Println("Successfully established database connection")
	}

	return nil
}

func Migrate(models ...interface{}) {
	DB.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)

	for _, model := range models {
		DB.AutoMigrate(model)
	}
}

type DBConn struct {
	DB *gorm.DB
	l  *utils.Logger
}

func New(l *utils.Logger) *DBConn {
	return &DBConn{
		l:  l,
		DB: DB,
	}
}

type Args struct {
	cnt  int
	vals []interface{}
}

func (a *Args) Append(v interface{}) string {
	a.cnt++
	a.vals = append(a.vals, v)
	return "?"
}

func (a *Args) Vals() []interface{} {
	return a.vals
}
