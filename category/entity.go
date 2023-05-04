package category

type Category struct {
	ID   uint64 `gorm:"primaryKey" json:"id"`
	Name string
}
