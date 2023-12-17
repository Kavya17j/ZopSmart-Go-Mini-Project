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

	// Prepare statement for updating data
	updateStmt, err := con.Prepare("UPDATE Movies SET Title = ?, Director = ? WHERE ID = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer updateStmt.Close()

	// Execute the statement
	_, err = updateStmt.Exec("Updated Movie Title", "Updated Director", 1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Record updated successfully")
}
