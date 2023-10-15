package dbM

import "fmt"

var TableName string = "users"

var connStr string = "postgresql://sniffyjoe:564123@127.0.0.1:5432/resulzade_robotik_kullanici?sslmode=disable"

var CreateTableQuery string = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s
  (
  name VARCHAR(255), 
  gsmno BIGINT, 
  class VARCHAR(255),
  ekip VARCHAR(255)
  )`, TableName)

var DropTableQuery string = fmt.Sprintf("DROP TABLE %s", TableName)

var GetTableUserCount string = fmt.Sprintf("SELECT COUNT(*) FROM %s", TableName)

var GetFromCount string = fmt.Sprintf("SELECT name FROM %s OFFSET $1 LIMIT 1", TableName)

var InsertIntoTable string = fmt.Sprintf("INSERT INTO %s (name, gsmno, class, ekip) VALUES ($1, $2, $3, $4)", TableName)
