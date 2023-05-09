package category

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routes(route *gin.Engine, db *gorm.DB) {
	repository := NewRepository(db)
	service := NewService(repository)
	handler := NewHandler(service)
	category := route.Group("/category")
	category.GET("/", handler.GetCategories)
	category.GET("/:id", handler.GetCategoryById)
	category.POST("/", handler.Create)
	category.PUT("/:id", handler.UpdateCategory)
	category.DELETE("/:id", handler.DeleteCategory)
}
