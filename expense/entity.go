package expense

import (
	"github.com/joshua/expense_tracker/category"
)

type Expense struct {
	ID          uint64 `gorm:"primaryKey" json:"id"`
	CategoryId  uint64 `json:"category_id"`
	Category    category.Category
	Amount      uint64
	Description string
	Date        uint64
	CreatedAt   uint64 `json:"created_at"`
	UpdatedAt   uint64 `json:"updated_at"`
}
