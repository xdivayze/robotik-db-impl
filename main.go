package main

import (
	"database/sql"
	"fmt"
	"os"
	"robotik/dbM"
	"robotik/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var router *gin.Engine = gin.Default()

func main() {
	c := cors.DefaultConfig()
	c.AllowAllOrigins = true
	router.Use(cors.New(c))

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
	router.RunTLS(":6471","/etc/letsencrypt/live/meral.club/fullchain.pem","/etc/letsencrypt/live/meral.club/privkey.pem")
}
