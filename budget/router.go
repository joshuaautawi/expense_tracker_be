package budget

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routes(route *gin.Engine, db *gorm.DB) {
	repository := NewRepository(db)
	service := NewService(repository)
	handler := NewHandler(service)
	budget := route.Group("/budgets")
	budget.GET("/", handler.GetBudgets)
}
