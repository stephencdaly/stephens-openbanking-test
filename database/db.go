package database

import (
	"database/sql"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
)

type Charge struct {
	ExternalId  string `json:"payment_id"`
	Amount      int    `json:"amount"`
	Reference   string `json:"reference"`
	Description string `json:"description"`
	ReturnUrl   string `json:"return_url"`
	Status      string `json:"status"`
}

type DB struct {
	conn *sql.DB
}

func sqlDir() string {
	root := os.Getenv("APP_ROOT")
	if root == "" {
		root = os.Getenv("PWD")
	}
	if root == "" {
		root, _ = os.Getwd()
	}
	return filepath.Join(root, "database", "sql")
}

func NewDB(connstr string) (*DB, error) {
	conn, err := sql.Open("postgres", connstr)
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

	// defer m.Close()
	if err := m.Up(); err != migrate.ErrNoChange {
		return err
	}

	return nil
}

func (db *DB) Close() error {
	return db.conn.Close()
}

func (db *DB) InsertCharge(charge Charge) error {
	_, err := db.conn.Exec(`INSERT INTO CHARGES (
		external_id,
		amount,
	  reference,
		description, 
		return_url,
		status) 
	VALUES ($1, $2, $3, $4, $5, $6)`,
		charge.ExternalId,
		charge.Amount,
		charge.Reference,
		charge.Description,
		charge.ReturnUrl,
		charge.Status)

	return err
}
