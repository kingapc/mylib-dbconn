package connection

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "library"
)

func GetConnection() (*sql.DB, bool) {

	var response bool
	psCredentials := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psCredentials)

	if err == nil {
		err = db.Ping()
		if err == nil {
			response = false
		} else {
			response = true
		}
	} else {
		response = true
	}

	return db, response
}

/*
func main() {

	db, err := getConnection()

	if err {
		fmt.Printf("\nUnable to connect to the data base!")
	} else {
		fmt.Printf("\nSuccessfully connected to database!")
		fmt.Print(db)
	}
}
*/
