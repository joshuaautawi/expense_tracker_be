package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joshua/expense_tracker/db"
	"github.com/joshua/expense_tracker/expense"
)

func main() {
	
	r := gin.Default()
	DB := db.InitDb()
	expense.Routes(r, DB)
	r.GET("/api/", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"message": "Hello"}) })
	r.Run()

}
