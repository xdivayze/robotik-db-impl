package dbM

import (
	"context"
	"database/sql"
	"fmt"
	"math/rand"
	"time"
)

func InitializeSQL(db **sql.DB, count *uint8) error {
	newDb, err := sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("error opening database %s\n", err)
	}
	*db = newDb

	if err = CreateUserTable(*db); err != nil {
		return fmt.Errorf("error creating users table %v\n ", err)
	}

	if err = GetUserCount(*db, count); err != nil {
		return fmt.Errorf("error getting count of rows in db: %s\n", err)
	}
	return nil
}

func GetUserFromNumber(db *sql.DB, count uint8, res *string) error {
	randm := rand.Intn(int(count))
	err := db.QueryRow(GetFromCount, randm).Scan(res)
	return fmt.Errorf("error retrieving %d th user: %v \n", randm+1, err)
}

func DropTable(db *sql.DB) error {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	_, err := db.ExecContext(ctx, DropTableQuery)
	return err
}

func CreateUserTable(db *sql.DB) error {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	_, err := db.ExecContext(ctx, CreateTableQuery)
	return err
}

func InsertToTable(db *sql.DB, name string, gsm string, class string, ekip string) error {
	_, err := db.Exec(InsertIntoTable, name, gsm, class, ekip)
	return err
}

func GetUserCount(db *sql.DB, resInt *uint8) error {
	err := db.QueryRow(GetTableUserCount).Scan(resInt)
	return err
}
