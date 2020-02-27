package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

/*Create mysql connection*/
func CreateCon() *sql.DB {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "12345"
	dbName := "goblog"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)

	// defer db.Close()

	if err != nil {
		fmt.Println(err)
	}

	// make sure connection is available
	err = db.Ping()
	fmt.Println(err)

	if err != nil {
		fmt.Println("db is not connected")
	}

	return db
}
