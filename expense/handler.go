package expense

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

func (h *handler) GetExpenseById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}
	expense, err := h.service.FindById(uint64(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": expense,
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

func (h *handler) UpdateExpense(c *gin.Context) {

	var expenseRequest ExpenseRequest
	err := c.ShouldBindJSON(&expenseRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)

	expense, err := h.service.Update(uint64(id), expenseRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": expense,
	})

}

func (h *handler) DeleteExpense(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err,
			"message": "Error while convert",
		})
		return
	}

	expense, err := h.service.Delete(uint64(id))
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
