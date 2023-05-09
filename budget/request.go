package budget

type BudgetRequest struct {
	Amount     uint64 `json:"amount"`
	Year       uint   `json:"year"`
	Month      uint   `json:"month"`
	CategoryId uint64 `json:"category_id"`
}
