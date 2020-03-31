package database

import (
	"database/sql"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
)

type DB struct {
	conn *sql.DB
}

func NewDB(connstr string) (*DB, error) {
	conn, err := sql.open("postgres", connstr)
	if err != nil {
		return nil, err
	}

	return &DB{conn: conn}, nil
}

func (db *DB) Ping() error {
	return db.conn.Ping()
}

func (db *DB) Init() error {
	driver, err := postgres.WithInstance(db.conn, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance("file://"+sqlDir(), "postgres", driver)
	if err != nil {
		return err
	}

	defer m.Close()
	if err := m.Up(); err != migrate.ErrNoChange {
		return err
	}

	return nil
}

func (db *DB) Close() error {
	return db.conn.Close()
}
