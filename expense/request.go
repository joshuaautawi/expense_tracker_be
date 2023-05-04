package expense

type ExpenseRequest struct {
	CategoryId  uint64 `json:"category_id" binding:"required"`
	Amount      uint64 `json:"amount" binding:"required"`
	Description string `json:"description"`
	Date        uint   `json:"date"`
}
