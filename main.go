package main

import (
	"database/sql"
	"fmt"
	"os"
	"robotik/dbM"
	"robotik/handlers"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var router *gin.Engine = gin.Default()

func main() {
	var db *sql.DB
	var count uint8
	if err := dbM.InitializeSQL(&db, &count); err != nil {
		fmt.Fprintf(os.Stderr, "error initialization: %s\n", err)
	}

	defer func() {
		if err := dbM.DropTable(db); err != nil {
			fmt.Fprintf(os.Stderr, "error dropping table: %s\n", err)
			os.Exit(-1)
		}
	}()

	handlers.InitializeHandlers(router, db)
	router.Run("127.0.0.1:6471")
}
