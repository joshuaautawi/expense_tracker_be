package expense

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Expense, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Expense, error) {
	var expenses []Expense
	err := r.db.Find(&expenses).Error
	return expenses, err
}
