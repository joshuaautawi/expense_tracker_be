package budget

type Budget struct {
	ID         uint64 `gorm:"primaryKey" json:"id"`
	Amount     uint64
	Year       uint
	CategoryId uint64 `json:"category_id"`
	Date       string `gorm:"varchar(50)"`
}
