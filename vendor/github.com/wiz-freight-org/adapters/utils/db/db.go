package db

import (
	"context"
	"os"
	"strconv"
	"sync"

	"github.com/wiz-freight-org/adapters/internal/config"
	"github.com/wiz-freight-org/adapters/internal/config/globals"
	"github.com/wiz-freight-org/adapters/utils"
	"go.uber.org/zap"

	"github.com/jackc/pgconn"
	pgx "github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Pool of database connection
var ConnPool *pgxpool.Pool
var once sync.Once

// Init the connection to DB
func Init() error {
	if ConnPool == nil {
		once.Do(func() {
			l := globals.Logger
			poolConfig, err := pgxpool.ParseConfig(config.DatabaseURL)
			if err != nil {
				l.Fatal("Unable to parse db url", zap.Error(err))
				os.Exit(1)
				return
			}

			poolConfig.MaxConns = int32(config.MaxDBConn)

			ConnPool, err = pgxpool.ConnectConfig(context.Background(), poolConfig)
			if err != nil {
				l.Fatal("Unable to connect to Database", zap.Error(err))
				os.Exit(1)
				return
			}
		})
	}
	return nil
}

type DBConn struct {
	conn          *pgxpool.Pool
	tx            pgx.Tx
	isTransaction bool
	l             *utils.Logger
}

// Interface to abstract the queryer(dbconnection or transaction)
type Queryer interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}

func New(l *utils.Logger) *DBConn {
	return &DBConn{
		l:    l,
		conn: ConnPool,
	}
}

// Initialize the DB connection and assign the existing db connection
func (db *DBConn) Init(l *utils.Logger) {
	db.conn = ConnPool
	db.l = l
}

func (db *DBConn) GetQueryer() Queryer {
	if db.isTransaction {
		return db.tx
	} else {
		return db.conn
	}
}

// ExecuteInTransaction executes the given function in DB transaction, i.e. It commits
// only if there is not error otherwise it is rolledback.
func (db *DBConn) ExecuteInTransaction(f func() error) (err error) {
	ctx := context.Background()
	tx, err := db.conn.Begin(ctx)
	if err != nil {
		return err
	}

	db.tx = tx
	db.isTransaction = true

	defer func() {
		db.isTransaction = false
	}()

	err = f()
	if err != nil {
		db.rollbackTransaction(tx)
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		db.rollbackTransaction(tx)
		return err
	}
	return nil
}

func (db *DBConn) rollbackTransaction(tx pgx.Tx) {
	err := tx.Rollback(context.Background())
	if err != nil {
		db.l.Error("Error While rollback", zap.Error(err))
	}
}

type Args struct {
	cnt  int
	vals []interface{}
}

func (a *Args) Append(v interface{}) string {
	a.cnt++
	a.vals = append(a.vals, v)
	return "$" + strconv.Itoa(a.cnt)
}

func (a *Args) Vals() []interface{} {
	return a.vals
}
