package expense

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{service}
}

func (h *handler) GetExpenses(c *gin.Context) {
	expenses, err := h.service.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": expenses,
	})

}

func (h *handler) Create(c *gin.Context) {
	var expenseRequest ExpenseRequest

	err := c.ShouldBindJSON(&expenseRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	expense, err := h.service.Create(expenseRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"data": expense,
	})
}
