package expense

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Expense, error)
	Create(expense Expense) (Expense, error)
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

func (r *repository) Create(expense Expense) (Expense, error) {
	err := r.db.Create(&expense).Error
	return expense, err
}
