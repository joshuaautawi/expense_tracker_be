package category

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Category, error)
	FindById(ID uint64) (Category, error)
	Create(category Category) (Category, error)
	Update(category Category) (Category, error)
	Delete(category Category) (Category, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Category, error) {
	var categories []Category
	err := r.db.Find(&categories).Error
	return categories, err
}

func (r *repository) FindById(ID uint64) (Category, error) {
	var category Category
	err := r.db.Find(&category, ID).Error
	return category, err
}

func (r *repository) Create(category Category) (Category, error) {
	err := r.db.Create(&category).Error
	return category, err
}

func (r *repository) Update(category Category) (Category, error) {
	err := r.db.Save(&category).Error
	return category, err
}
func (r *repository) Delete(category Category) (Category, error) {
	err := r.db.Delete(&category).Error
	return category, err
}
