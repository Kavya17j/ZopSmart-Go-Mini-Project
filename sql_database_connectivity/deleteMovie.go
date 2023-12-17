package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Open connection
	con, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/MoviesDB")
	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()

	// Prepare statement for deleting data
	deleteStmt, err := con.Prepare("DELETE FROM Movies WHERE ID = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer deleteStmt.Close()

	// Execute the statement
	_, err = deleteStmt.Exec(2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Record deleted successfully")
}
