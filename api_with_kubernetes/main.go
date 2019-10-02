package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {

	dbService := NewDatabaseServiceImp(HOST, PORT, DBNAME, COLLECTIONNAME)

	r := gin.Default()

	r.GET("/empid/:empid", func(context *gin.Context) {
		id := context.Param("empid")
		idasint, err := strconv.Atoi(id)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"status": "Internal error" + err.Error()})
			return
		}
		var userinfo *UserInfo
		if userinfo, err = dbService.Retrieve(idasint); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"status": "Internal error " + err.Error()})
			return
		}
		context.JSON(http.StatusOK, gin.H{"status": "ok", "user": userinfo})
	})

	r.POST("/emp", func(context *gin.Context) {
		var user UserInfo

		if context.Bind(&user) == nil {
			if err := dbService.Save(&user); err != nil {
				context.JSON(http.StatusInternalServerError, gin.H{"status": "Internal Server Error " + err.Error()})
				return
			}
		}
		context.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.Run(":8080")
}
