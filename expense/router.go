package expense

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routes(route *gin.Engine, db *gorm.DB) {
	repository := NewRepository(db)
	service := NewService(repository)
	handler := NewHandler(service)
	expense := route.Group("/expenses")
	expense.GET("/", handler.GetExpenses)
}
