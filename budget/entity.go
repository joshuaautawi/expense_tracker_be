package budget

type Budget struct {
	ID         uint64 `gorm:"primaryKey" json:"id"`
	Amount     uint64
	Year       uint
	Month      uint   `json: "id"`
	CategoryId uint64 `json:"category_id"`
}
