package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	var db *sql.DB
	var count uint8
	if err := initializeSQL(&db, &count); err != nil {
		fmt.Fprintf(os.Stderr, "error initialization: %s\n", err)
	}

	defer func() {
		if err := dropTable(db); err != nil {
			fmt.Fprintf(os.Stderr, "error dropping table: %s\n", err)
			os.Exit(-1)
		}
	}()
}
