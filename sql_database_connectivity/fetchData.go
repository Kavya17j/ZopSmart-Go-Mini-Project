package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//open connection
	con, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/MoviesDB")
	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()

	stmt, err := con.Prepare("select ID,Isbn,Title,Director from Movies where ID=?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var movieID string

	err := stmt.QueryRow(2).Scan(&movieID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("MovieID: %d", movieID)

}
