package main

import (
	"log"
	"os"

	"github.com/stephencdaly/stephens-openbanking-test/http"
	"github.com/stephencdaly/stephens-openbanking-test/database"
)

func Main() error {
	db, err := database.NewDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	if err := db.Init(); err != nil {
		return err
	}

	http.Start(http.Config{
		DB: db})

	return nil
}

func main() {
	if err := Main(); err != nil {
		log.Fatal(err)
	}
}
