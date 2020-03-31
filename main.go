package main

import(

	"github.com/stephencdaly/stephens-open-banking-test/api"
)

func main() { 
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

	api.Start()
}