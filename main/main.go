package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/klb0000/retailshop/server"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "/Users/vikram/go/src/projects/retailshop/data/data.db")
	if err != nil {
		log.Fatal(err)
	}
	addr := "localhost:8080"
	fmt.Printf("starting server at %s\n", addr)
	server.ServeDB(addr, db)

}
