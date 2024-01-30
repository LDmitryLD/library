package db

import (
	"database/sql"
	"fmt"
	"projects/LDmitryLD/library/app/config"
	"projects/LDmitryLD/library/app/internal/db/adapter"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewSqlDB(dbCond config.DB) (*sqlx.DB, *adapter.SQLAdapter, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbCond.Host, dbCond.Port, dbCond.User, dbCond.Password, dbCond.Name)
	var dbRaw *sql.DB

	dbRaw, err := sql.Open(dbCond.Driver, dsn)
	if err != nil {
		return nil, nil, err
	}
	err = dbRaw.Ping()
	if err != nil {
		return nil, nil, err
	}

	db := sqlx.NewDb(dbRaw, dbCond.Driver)
	sqlAdapter := adapter.NewSQLAdapter(db)
	return db, sqlAdapter, nil
}
