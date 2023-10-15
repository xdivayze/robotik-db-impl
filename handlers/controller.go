package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"robotik/dbM"

	"github.com/gin-gonic/gin"
)

func InitializeHandlers(r *gin.Engine, db *sql.DB) {
	r.GET("/api/get-random-user", func(c *gin.Context) {
		var res uint8
		dbM.GetUserCount(db, &res)
		getRandomUser(c, db, res)
	})

	r.POST("/api/insert", func(c *gin.Context) {
		insertUser(c, db)
	})
}

func insertUser(c *gin.Context, db *sql.DB) {
	type UserData struct {
		Name  string `json:"name"`
		Gsmno string `json:"gsmno"`
		Class string `json:"class"`
		Ekip  string `json:"ekip"`
	}
	var user UserData

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := dbM.InsertToTable(db, user.Name, user.Gsmno, user.Class, user.Ekip); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	c.Status(202)

}

func getRandomUser(c *gin.Context, db *sql.DB, count uint8) {
	var res string
	if err := dbM.GetUserFromNumber(db, count, &res); res == "" {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": fmt.Sprintf("error occured when trying to fetch user: %v\n", err),
		})
		return
	}
	if res == "" {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "response is null",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"name": res,
	})
}
