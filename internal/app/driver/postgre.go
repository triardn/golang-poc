package driver

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/go-gorp/gorp/v3"
	_ "github.com/lib/pq" // defines postgreSQL driver used
)

// DBPostgreOption options for postgre connection
type DBPostgreOption struct {
	IsEnable        bool
	Host            string
	Port            int
	Username        string
	Password        string
	DBName          string
	MaxOpenConn     int
	MaxIdleConn     int
	ConnMaxLifetime time.Duration
}

// NewPostgreDatabase return gorp dbmap object with postgre options param
func NewPostgreDatabase(option DBPostgreOption) (*gorp.DbMap, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", option.Host, option.Port, option.Username, option.DBName, option.Password))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(option.MaxOpenConn)
	db.SetConnMaxLifetime(option.ConnMaxLifetime)
	db.SetMaxIdleConns(option.MaxIdleConn)
	gorp := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	return gorp, nil
}
