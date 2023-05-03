package expense

type Expense struct {
	ID          uint64 `gorm:"primaryKey" json:"id"`
	CategoryId  uint64 `json:"category_id"`
	Amount      uint64
	Description string
	Date        string
	createdAt   string `json:"created_at"`
	updatedAt   string `json:"updated_at"`
}
