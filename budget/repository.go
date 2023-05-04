package budget

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Budget, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Budget, error) {
	var budgets []Budget
	err := r.db.Find(&budgets).Error
	return budgets, err
}
