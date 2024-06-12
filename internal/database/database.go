package database

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // postgres driver
)

// Base struct containing db instances
type Base struct {
	Slave  *sqlx.DB
	Master *sqlx.DB
	Driver string
}

// Config struct containing database config
type Config struct {
	Slave  string `json:"slave"`
	Master string `json:"master"`
}

// db wrapper struct for new connection to db
type db struct {
	driver     string
	connection string

	maxConnection int
	maxIdle       int
}

// New create new connection to db
func New(cfg Config, driver string) (*Base, error) {
	slaveConn := &db{
		driver:        driver,
		connection:    cfg.Slave,
		maxIdle:       25,
		maxConnection: 100,
	}
	slave, err := slaveConn.connect()
	if err != nil {
		return nil, err
	}

	masterConn := &db{
		driver:        driver,
		connection:    cfg.Slave,
		maxIdle:       25,
		maxConnection: 100,
	}
	master, err := masterConn.connect()
	if err != nil {
		return nil, err
	}
	return &Base{Slave: slave, Master: master, Driver: driver}, nil
}

// connect connect to db
func (d *db) connect() (*sqlx.DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := sqlx.ConnectContext(ctx, d.driver, d.connection)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(d.maxIdle)
	db.SetMaxOpenConns(d.maxConnection)
	db.SetConnMaxIdleTime(10 * time.Minute)
	return db, err
}
