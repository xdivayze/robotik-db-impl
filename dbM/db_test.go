package dbM_test

import (
	"database/sql"
	_ "github.com/lib/pq"
	"robotik/dbM"
	"testing"
)

func TestSQLIntegrationT(t *testing.T) {
	var db *sql.DB
	var count uint8
	var res string

	if err := dbM.InitializeSQL(&db, &count); err != nil {
		t.Fatalf("err initializing: %s\n", err)
	}
	if count == 0 {
		t.Fatalf("count is nil\n")
	}

	if err := dbM.GetUserFromNumber(db, count, &res); res == "" {
		t.Fatalf("error getting user:%s\n", err)
	}

	if res == "" {
		t.Fatalf("got user is nil\n")
	}

	if err := dbM.InsertToTable(db, "ozangg", "05462507982", "11FF", "ROBO"); err != nil {
		t.Fatalf("err occured while inserting to table: %s\n", err)
	}
}
