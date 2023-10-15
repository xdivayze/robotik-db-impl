package main

import (
	"database/sql"
	"testing"
)

func TestSQLIntegrationT(t *testing.T) {
	var db *sql.DB
	var count uint8
	var res string

	if err := initializeSQL(&db, &count); err != nil {
		t.Fatalf("err initializing: %s\n", err)
	}
	if count == 0 {
		t.Fatalf("count is nil\n")
	}

	if err := getUserFromNumber(db, count, &res); res == "" {
		t.Fatalf("error getting user:%s\n", err)
	}

	if res == "" {
		t.Fatalf("got user is nil\n")
	}

	if err := insertToTable(db, "ozangg", "05462507982", "11FF", "ROBO"); err != nil {
		t.Fatalf("err occured while inserting to table: %s\n", err)
	}
}
