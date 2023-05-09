package budget

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Budget, error)
	FindById(ID uint64) (Budget, error)
	Create(budget Budget) (Budget, error)
	Update(budget Budget) (Budget, error)
	Delete(budget Budget) (Budget, error)
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

func (r *repository) FindById(ID uint64) (Budget, error) {
	var budget Budget
	err := r.db.Find(&budget, ID).Error
	return budget, err
}

func (r *repository) Create(budget Budget) (Budget, error) {
	err := r.db.Create(&budget).Error
	return budget, err
}

func (r *repository) Update(budget Budget) (Budget, error) {
	err := r.db.Save(&budget).Error
	return budget, err
}

func (r *repository) Delete(budget Budget) (Budget, error) {
	err := r.db.Delete(&budget).Error
	return budget, err
}
