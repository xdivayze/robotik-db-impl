package main

import (
	"database/sql"
	"testing"
)

func Test(t *testing.T) {
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
		t.Errorf("error getting user:%s\n", err)
	}

	if res == "" {
		t.Fatalf("got user is nil\n")
	}

}
